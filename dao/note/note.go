package note

import (
	"Gin/models"
	db2 "Gin/pkg/db"
	"Gin/pkg/log"
	r "Gin/redis"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	NoteTable                         = "note"
	KeyNoteInfoHashPrefix             = "gin:note:"
	KeyNoteTimeZSet                   = "gin:note:time"
	KeyNoteVotedZSetPrefix            = "gin:note:voted:"
	KeyNoteScoreSet                   = "gin:note:score"
	KeyCommunityNoteSetPrefix         = "gin:community:"
	OneWeekInSeconds                  = 7 * 24 * 3600
	VoteScore                 float64 = 432 // 每一票的值432分
	PostPerAge                        = 20
)

var ZSetKey = "mylist"

type NoteDaoGroup struct {
	NoteInfoDao
}
type NoteInfoDao struct {
}

func (NoteInfoDao *NoteInfoDao) VoteForNote(args *models.VoteDataArgs) (err error) {

	voteTime := r.Client.ZScore(KeyNoteTimeZSet, args.NoteID).Val()
	if float64(time.Now().Unix())-voteTime > OneWeekInSeconds {
		return errors.New("不允许投票了")
	}
	// 2.更新帖子的分数
	key := KeyNoteVotedZSetPrefix + args.NoteID
	ov := r.Client.ZScore(key, args.NoteID).Val()
	if args.Direction == ov {
		return errors.New("投票时间已经过去了")
	}
	var op float64
	if args.Direction > ov {
		op = 1
	} else {
		op = -1
	}
	diffAbs := math.Abs(ov - args.Direction) // 计算两次投票的差值
	pipeline := r.Client.TxPipeline()        // 事务操作
	_, err = pipeline.ZIncrBy(KeyNoteScoreSet, VoteScore*diffAbs*op, args.NoteID).Result()
	if err != nil {
		return err
	}

	if args.Direction == 0 {
		_, err = r.Client.ZRem(key, args.NoteID).Result()
	} else {
		pipeline.ZAdd(key, redis.Z{ // 记录已投票
			Score:  args.Direction, // 赞成票还是反对票
			Member: args.NoteID,
		})
	}
	_, err = pipeline.Exec()
	return err
}

func (noteInfoDao *NoteInfoDao) AddNote(args models.GetNoteInfoArgs) bool {
	err := db2.Db.Table(NoteTable).Create(&args).Error
	if err != nil {
		log.Error("[Dao] AddNote error : %v", err)
		return false
	}
	now := float64(time.Now().Unix())
	votedKey := KeyNoteVotedZSetPrefix + strconv.Itoa(int(args.ID))
	communityKey := KeyCommunityNoteSetPrefix + strconv.Itoa(int(args.Label)) // 1  2 3
	noteInfo := map[string]interface{}{
		"title":   args.Title,
		"content": args.Content,
		"desc":    args.Desc,
		"note:id": args.ID,
		"support": args.Support,
		"label":   args.Label,
		"newsID":  args.NewsID,
		"time":    now,
	}
	//事务操作
	support := cast.ToString(args.Support)
	pipeline := r.Client.TxPipeline()
	pipeline.ZAdd(votedKey, redis.Z{ //添加到分数
		Score:  cast.ToFloat64(support),
		Member: args.User,
	})
	pipeline.Expire(votedKey, time.Second*OneWeekInSeconds) // 一周时间

	pipeline.HMSet(KeyNoteInfoHashPrefix+strconv.Itoa(int(args.ID)), noteInfo)
	pipeline.ZAdd(KeyNoteScoreSet, redis.Z{
		Score:  cast.ToFloat64(support),
		Member: args.ID,
	})
	pipeline.ZAdd(KeyNoteTimeZSet, redis.Z{
		Score:  now,
		Member: args.ID,
	})
	pipeline.SAdd(communityKey, args.ID)
	_, err = pipeline.Exec()
	if err != nil {
		fmt.Println(err)
	}
	return true

}

func (noteInfoDao *NoteInfoDao) EditNote(args models.EditNoteInfoArgs) (res models.DelOrEditReply) {
	fmt.Println(args.Id)
	noteInfo := &models.GetNoteInfoArgs{
		Title:       args.Title,
		Content:     args.Content,
		Desc:        args.Desc,
		User:        args.User,
		Label:       args.Label,
		UpdatedTime: args.UpdatedTime,
	}
	/**
	c, err := conn.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		//TODO
	}
	defer c.Close()
	**/
	fmt.Println(noteInfo)
	i := db2.Db.Table(NoteTable).Where("id = ?", args.Id).Updates(noteInfo).RowsAffected
	if i <= 0 {
		res.Bool = false
		log.Warn("[Dao] Edit note have no influence in this data !!!")
		return res
	} else {
		res.Bool = true
		//reply, err := c.Do("HMSet", "RNode", "id", args.Id, "title", args.Title, "content", args.Content, "desc", args.Desc, "user", args.User)
		//_, _ = c.Do("expire", "id", -1)

	}

	return res
}

func (noteInfoDao *NoteInfoDao) DelNote(args models.DelNoteInfoArgs) (res models.DelOrEditReply) {

	i := db2.Db.Table(NoteTable).Where("id = ? ", args.Id).Delete(&models.GetNoteInfoArgs{}).RowsAffected
	if i > 0 {
		res.Bool = true
		return res
	} else {
		res.Bool = false
		return res
	}

}

func (noteInfoDao *NoteInfoDao) QueryNote(args models.QueryNoteArgs) (res []models.QueryNoteInfoReply) {
	var info []models.QueryNoteInfoReply
	//c, _ := conn.Dial("tcp", "localhost:6379")
	//defer c.Close()
	//	reply, err := c.Do("lrange", "mylist", 0, -1)

	//pkeys, _ := conn.Strings(reply, err)
	//fmt.Println(pkeys)

	db2.Db.Table(NoteTable).Find(&info).Limit(args.PageNum).Offset(args.PageSize)
	/**
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	_, err = conn.Do("zrevrange", ZSetKey, "0", "-1", "support")
	if err != nil {
		return
	}
	for _, v := range info {
		fmt.Println(v)
		conn.Do("zadd", ZSetKey, v)
	}
	do, _ := redis.ByteSlices(conn.Do("zrevrange", ZSetKey, 0, -1, "support"))
	for _, v := range do {
		fmt.Println(string(v))
	}
	**/
	return info
}

func (noteInfoDao *NoteInfoDao) QueryNoteByID(args models.QueryNoteByIDArgs) (res models.QueryNoteInfoReply) {
	var info models.QueryNoteInfoReply
	db2.Db.Table(NoteTable).Where("id = ?", args.Id).Find(&info)
	return info
}

/**

func (NoteInfoDao *NoteInfoDao) GetNoteHotSortingBang(args models.GetNoteHotSortingBangArgs) {

	// redis 链接
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer Rdb.Close()

	insertList := []redis.Z{
		{Score: float64(args.LoveNum), Member: args.Id},
	}
	num, err := Rdb.ZAdd(ZSetKey, insertList...).Result()
	if err != nil {
		fmt.Printf("添加失败, err:%v\n", err)
		return
	}
	fmt.Printf("添加成功 %d\n", num)

	ret, err := Rdb.ZRevRangeWithScores(ZSetKey, 0, -1).Result()
	if err != nil {
		fmt.Printf("获取所有失败,err:%v\n", err)

	}
	for k, v := range ret {
		fmt.Printf("%v,%v\n", k, v.Member)
	}

}
**/
/**
func (NoteInfoDao *NoteInfoDao) GetNoteSortingBangTopN(args models.GetNoteHotSortingBangArgs) (res []models.GetNoteHotSortingBangTopAllNReply) {
	var ids []string
	// redis 链接
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer Rdb.Close()
	ret, err := Rdb.ZRevRangeWithScores(ZSetKey, 0, -1).Result()
	if err != nil {
		fmt.Printf("获取所有失败,err:%v\n", err)

	}

	for k, v := range ret {
		ids = append(ids, v.Member.(string))
		res = NoteInfoDao.GetNoteHotInfoFromMysql(ids[k])

		fmt.Printf("返回值为: %v\n", res)

	}

	return res
}
**/

func (NoteInfoDao *NoteInfoDao) GetHotBangFromRedis(num int) (res []models.GetNoteInfoArgs) {
	ret, err := r.Client.ZRevRangeWithScores(KeyNoteScoreSet, 0, 2).Result()
	if err != nil {
		fmt.Println("zrevrange failed , err : %v")
	}
	for _,v := range ret {

		id := cast.ToInt(v.Member)

		req := NoteInfoDao.GetNoteInfoFromMysql(id)

       res = append(res,req)
	}
	return res
}

func (NoteInfoDao *NoteInfoDao) GetNoteInfoFromMysql(id int) (req models.GetNoteInfoArgs) {
	 res := new(models.GetNoteInfoArgs)
	err  := db2.Db.Table(NoteTable).Where("id = ?",id).Find(res).Error
	if  err != nil {
		fmt.Println(err)
	}
	req = *res
	return req
}

//  ID拿MYSQL文章详情信息
func (NoteInfoDao *NoteInfoDao) GetNoteHotInfoFromMysql(id string) (res []models.GetNoteHotSortingBangTopAllNReply) {
	var info []models.GetNoteHotSortingBangTopAllNReply
	db2.Db.Table(NoteTable).Where("id = ?", id).Find(&info)
	return info
}

// PayForYourLovePassageDao
func (NoteInfoDao *NoteInfoDao) PayForYourLovePassageDao(args models.PayForYourLovePassageArgs) (res models.PayForYourLovePassageReply) {
	var valueFromDB models.QueryNoteInfoReply
	var rep models.PayForYourLovePassageReply
	db2.Db.Table(NoteTable).Where("id = ?", args.Id).Find(&valueFromDB)

	j, _ := strconv.Atoi(args.Packet)
	fmt.Println("j", j)
	support := valueFromDB.Support + j
	fmt.Println("s", support)
	rep.Id = args.Id
	rep.Support = strconv.FormatInt(int64(support), 10)
	o := db2.Db.Table(NoteTable).Where("id = ?", args.Id).Update("support", support).RowsAffected

	if o > 0 {

		return rep
	} else {
		return models.PayForYourLovePassageReply{Id: 0, Support: ""}
	}
}
