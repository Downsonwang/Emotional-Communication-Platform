package register

import (
	"Gin/dao"
	"Gin/models"
	"fmt"

	email2 "github.com/jordan-wright/email"
	"log"
	"math/rand"
	"net/smtp"
	"strings"
	"time"

	"regexp"
)

type RegisterInfoService struct {
}

var authInfoDao = dao.DaoGroupInfo.RegDaoGroup.RegDaoInfo

func (regInfoService *RegisterInfoService) SendCodeInfo(email string) (fcode string, b bool, msg string) {
	if email == "" {
		return "", false, "邮箱信息为空,请重新填写!"
	}

	if VertifyEmailFormat(email) == false {
		return "", false, "邮箱格式不对,请重新确认格式!"
	}
	e := email2.NewEmail()
	e.From = "茄子交流 <downsonliteracy@foxmail.com>"
	e.To = []string{email}
	e.Subject = "你好,这里是星座情感交流中心."
	//设置文件发送的内容
	code := GenValidateCode(6)
	e.Text = []byte("您的平台注册验证码如下:" + code)

	//dao.GetEmailRegisterInfo(&models.RegUserInfoArgs{Email: args.Email,Password:args.Password ,Code: code})
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "downsonliteracy@foxmail.com", "kcdzxborkrbmddcj", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)

	}
	//authInfoDao.GetEmailRegisterInfo()
	//authInfoDao.GetEmailRegisterInfo(&models.RegUserInfoArgs{Email: email,Code: code,Password: ""})
	b, _ = authInfoDao.GetEmailCodeInfo(email, code)

	return code, true, "验证码确认中."
}
func (regInfoService *RegisterInfoService) GetEmailRegisterInfo(args *models.RegUserInfoArgs) bool {
	code := authInfoDao.GetSqlCodeInfo(args)
	if code == args.Code {
		_ = authInfoDao.UpdateEmailPasswordInfo(args)
		return true
	}
	return false
}

func VertifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
