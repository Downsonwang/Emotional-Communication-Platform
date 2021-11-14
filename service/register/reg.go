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

func GetEmailRegisterInfo(args *models.RegUserInfoArgs) (bool, string) {
	fmt.Println(args.Email,args.Password)
	if args.Email == "" {
		return false, "邮箱不能为空"
	}

	if VertifyEmailFormat(args.Email) == false {
		return false, "邮箱格式不对"
	}
	e := email2.NewEmail()
	e.From = "茄子交流 <downsonliteracy@foxmail.com>"
	e.To = []string{args.Email}
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
	_, err = dao.GetEmailRegisterInfo(&models.RegUserInfoArgs{Email: args.Email, Code: code,Password: args.Password})
	if err != nil {
		log.Fatal(err)
	}

	return true, ""

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
