package note

import (
	"Gin/models"
	"fmt"
	"math"
	"math/rand"
)

type NoteRecommendService struct{}

func (this *NoteRecommendService) GetUserPostInfoService(args models.RecommendArgs) (resp *models.RecommendPosts) {
	// 先返回5个, 一个自己浏览的 和 一个朋友
	var inte map[int][]models.UserAndPostIntersection
	inte[args.ID] = make([]models.UserAndPostIntersection, 0)
	res := this.GetRecommendPostService(args.ID, inte)
	return res
}

func (this *NoteRecommendService) GetRecommendPostService(id int, inter map[int][]models.UserAndPostIntersection) *models.RecommendPosts {
	if len(inter) == 0 {
		return &models.RecommendPosts{Posts: make([]models.Note, 0)}
	}
	var recommendPosts = &models.RecommendPosts{
		Posts: make([]models.Note, 8),
	}

	fmt.Println(recommendPosts)

	var FriendIDs []int
	for _, user := range inter[id] {

		FriendIDs = append(FriendIDs, user.FriendID)

	}
	chooseFriendID := rand.Intn(len(FriendIDs))
	// 随机获得一位friend 将friend的Note信息拿到跟当前该用户的信息 做矩阵
	friendPosts := this.GetFriendPostService(chooseFriendID, inter[chooseFriendID])
	friend := inter[chooseFriendID]
	for i, value := range friendPosts {
		friend[i].PostID = value.ID

	}
	res := this.BecomeUserAndPostsMatrix(id, chooseFriendID, inter[id], friend)
	recommendPosts.Posts = res
	return recommendPosts
	//var recommendPosts
}
func (this *NoteRecommendService) BecomeUserAndPostsMatrix(ownerID, friendIdD int, owner, friend []models.UserAndPostIntersection) []models.Note {

	UserPostMaxtrix := make([][]float64, len(owner)+len(friend))

	// 构建用户-帖子矩阵
	for i := range UserPostMaxtrix {
		UserPostMaxtrix[i] = make([]float64, len(owner)+len(friend))
		for j := range UserPostMaxtrix[i] {
			UserPostMaxtrix[i][j] = 0
		}
	}

	// 填充-用户帖子矩阵
	// 填充 自己-自己帖子的矩阵范围
	for _, ownInter := range owner {
		ownInterIndex := this.findIndexOfPostInMatrix(UserPostMaxtrix, ownInter.UserID, ownInter.PostID)
		ownerID = ownInterIndex
		if ownInterIndex != -1 {
			UserPostMaxtrix[ownInterIndex][ownInter.PostID] = this.calculateWeightedScore(ownInter)
		}
	}

	// 填充-朋友和朋友帖子的矩阵范围
	for _, fri := range friend {
		friInterIndex := this.findIndexOfPostInMatrix(UserPostMaxtrix, fri.UserID, fri.PostID)
		friendIdD = friInterIndex
		if friInterIndex != -1 {
			UserPostMaxtrix[friInterIndex][fri.PostID] = this.calculateWeightedScore(fri)
		}
	}
	// 检查矩阵中是否有NA
	formatUserPostMaxtrix := this.ReplaceNA(UserPostMaxtrix)

	// 计算相似性
	o := formatUserPostMaxtrix[ownerID]
	f := formatUserPostMaxtrix[friendIdD]
	// 相似度在0-1之间
	similarityNum := this.JaccardSimilarity(o, f)
	// 推荐总数为 8
	// 如果相似度大于等于50,则将朋友的随机三篇文章推进行 + 自己热爱的标签ID下的系统库文章三篇 + 其他标签下系统库两篇
	var posts []models.Note
	if similarityNum >= 0.50 {
		friendPosts := this.getRandomNotesFromFriend(friendIdD, 3)
		ownLovePosts := this.getNotesByOwnerFavoriteTag(ownerID, 3)
		otherTagsPosts := this.getNotesFromOtherTags(ownerID, 2)
		posts = append(posts, friendPosts...)
		posts = append(posts, ownLovePosts...)
		posts = append(posts, otherTagsPosts...)

	}
	// 如果相似度大于等于25小于50 则将朋友的随机两篇文章推进来 + 自己热爱的标签ID下的系统库文章三篇 + 其他标签下系统库三篇
	if similarityNum >= 25 && similarityNum < 50 {
		friendPosts := this.getRandomNotesFromFriend(friendIdD, 2)
		ownLovePosts := this.getNotesByOwnerFavoriteTag(ownerID, 3)
		otherTagsPosts := this.getNotesFromOtherTags(ownerID, 3)
		posts = append(posts, friendPosts...)
		posts = append(posts, ownLovePosts...)
		posts = append(posts, otherTagsPosts...)
	}
	// 如果相似度大于10小于25 则将朋友的随机一篇文章推进来 + 自己热爱的标签下ID下的系统库文章四篇 + 其他标签下系统库三篇
	if similarityNum >= 10 && similarityNum < 25 {
		friendPosts := this.getRandomNotesFromFriend(friendIdD, 1)
		ownLovePosts := this.getNotesByOwnerFavoriteTag(ownerID, 4)
		otherTagsPosts := this.getNotesFromOtherTags(ownerID, 3)
		posts = append(posts, friendPosts...)
		posts = append(posts, ownLovePosts...)
		posts = append(posts, otherTagsPosts...)
	}
	// 如果相似度小于10 则将随机推系统库里所有文章
	if similarityNum < 10 {

		otherTagsPosts := this.getNotesFromOtherTags(ownerID, 8)

		posts = append(posts, otherTagsPosts...)
	}

	return posts
}

// getRandomNotesFromFriend 获取朋友的随机笔记
func (this *NoteRecommendService) getRandomNotesFromFriend(friendID int, count int) []models.Note {
	return nil
}

// getNotesByOwnerFavoriteTag 获取用户最喜爱标签下的文章
func (this *NoteRecommendService) getNotesByOwnerFavoriteTag(ownerID int, count int) []models.Note {
	// ... 获取用户最喜爱标签下的文章的代码 ...
	return nil
}

// getNotesFromOtherTags 获取其他标签下的文章
func (this *NoteRecommendService) getNotesFromOtherTags(ownerID int, count int) []models.Note {
	// ... 获取其他标签下的文章的代码 ...
	return nil
}

// recommendNotes 推荐文章
func (this *NoteRecommendService) recommendNotes(notes []models.Note) []models.Note {
	// ... 推荐文章的代码 ...
	return nil
}

func (this *NoteRecommendService) JaccardSimilarity(a, b []float64) float64 {
	intersection := 0
	union := 0

	// 交集/并集
	for id := range a {
		union++
		if id > 0 && this.existsInSlice(b, id) {
			intersection++
		}
	}

	for id := range b {
		if id > 0 && !this.existsInSlice(a, id) {
			union++
		}
	}
	if union == 0 {
		return 0
	}
	return float64(intersection) / float64(union)
}

func (this *NoteRecommendService) existsInSlice(slice []float64, value int) bool {
	for _, v := range slice {
		if v == float64(value) {
			return true
		}
	}
	return false
}

func (this *NoteRecommendService) ReplaceNA(prefs [][]float64) [][]float64 {
	var arr [][]float64
	for i := 0; i < len(prefs); i++ {
		arr = append(arr, make([]float64, len(prefs[0])))
		for j := 0; j < len(prefs[i]); j++ {
			if float64(prefs[i][j]) == math.NaN() { // 检查是否为 NaN
				arr[i][j] = float64(0)
			} else {
				arr[i][j] = prefs[i][j]
			}
		}
	}
	return arr
}
func (this *NoteRecommendService) calculateWeightedScore(intersection models.UserAndPostIntersection) float64 {
	// 这里可以根据实际情况来计算加权分数
	weightedScore := float64(intersection.Support)*0.4 + float64(intersection.Comment)*0.4 + float64(intersection.Like)*0.2
	return weightedScore
}

func (this *NoteRecommendService) findIndexOfPostInMatrix(matrix [][]float64, postID int, userID int) int {
	for i, row := range matrix {
		for j, _ := range row {
			if i == userID && j == postID {
				return i
			}
		}
	}
	return -1
}

func (this *NoteRecommendService) GetFriendPostService(friendID int, friends []models.UserAndPostIntersection) []models.Note {
	var notes []models.Note
	for _, friend := range friends {
		notes = append(notes, models.Note{
			ID: friend.PostID,
		})
	}

	return notes
}
