package main

import (
	"fmt"

	"github.com/smartwalle/alipay"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

var (
	publicKey  = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAgJzxeiRX1+Ie1kpmAxx4czcNqLbFgM2rosvt0SljiLS9NfdEE0GDzCBzUBPdTpmwnStPyz/IHBd2AjxZDD1fVCEN5PeXlvmZ1tPsT602YGUYZF/Q9hda7C1bMVRr/czp6sAUbBDUUxdHaHVlaYPC8RC9B20AH0ZYohRaD8KSugcXCl6aPhkpBpqCK5gCjUidsRJC+X0xWYKfzA2Ucsfk++9EYadpU/DmcCMy1b82ljTiI/EgpjhIli0KvMw3rS1SCa2LIEzDJCmDqAhEh/hTgwXQ3dN4vEpMYAH8I7gDCDRiHq+b1yTekyshXwcrXrRCX/Ut1PbUkL3gKy39s4cwvQIDAQAB"
	privateKey = "MIIEowIBAAKCAQEAgJzxeiRX1+Ie1kpmAxx4czcNqLbFgM2rosvt0SljiLS9NfdEE0GDzCBzUBPdTpmwnStPyz/IHBd2AjxZDD1fVCEN5PeXlvmZ1tPsT602YGUYZF/Q9hda7C1bMVRr/czp6sAUbBDUUxdHaHVlaYPC8RC9B20AH0ZYohRaD8KSugcXCl6aPhkpBpqCK5gCjUidsRJC+X0xWYKfzA2Ucsfk++9EYadpU/DmcCMy1b82ljTiI/EgpjhIli0KvMw3rS1SCa2LIEzDJCmDqAhEh/hTgwXQ3dN4vEpMYAH8I7gDCDRiHq+b1yTekyshXwcrXrRCX/Ut1PbUkL3gKy39s4cwvQIDAQABAoIBAEZWDlWvBH9jAVxOKcXvzwurgwPOmOqdEA2jGzc1PGLp/URDKu6g+LQs1wNnKbOvp3/8zNGp2wVJ61bDrCtecQDwZsnegf+mF3T+RxE3+DH1d8aFBCFhmm7pSyMKOVj+tLqWO1TySzv50iqVcVWoInd1oPsrqFJyRUy2dp6B2X50VGsY1mcFND9XtoO60qafr49whxhrr6YQa6frTBJ43YVd1JiAS16R4I4uXr9AAGApkoWdcXgbKNoD3hRCFtp1MWMH8IKcatbo/C8OEAE1t19CEQHjXAyZOuBOkpUqc0omtWnKHzL7PDOxc4rpa42FTYOPAHGCwZ6h+KsBDH3ZZEECgYEAy1HO8/BuHh0T4Uq3NfVlvYw+f1MAA9mJ8My/57TRT907+RrpPOWSN3MtY7aGd4Y75Y5FwdAyarBRllsunx1/LCtGLPICrlI7geFvQAm9F8EcQrE6T+A1STYkTCT+NT7jMTa5bF8sBp4NrWr0M+qZJ6BG7WeDsn8ErkGpz+zMLc0CgYEAoe/Zwh4vwikLu54Mi5W0j65QyuwEZRfNalUdmtDMgVGeLGkZjsIHsTWF/5hBfAcCSeU7Uq60inEa+T2xqKvSbwlbfjC2FQ2sEog7wWSc10DZN9zXOtEPGEfRrnQ+Y8clUu4fOFbOOqtl7YTT1E/kxSF1wNPS826aMZVzAu9znrECgYBW+L2wN47DLukMHCvW1wwYUt+BE34UJ4AME2mbsgs/QIGhA1P2CrXLJmeqGa/XfQIuJffM+kQ2wwmDaam9wp5dxH0WSMfAzchKvKYcHI5YlPCDztAEzwG8OX65W7GqESaaLLCQnP12LzHmlrOzaYzuGG9+qiL5vb/A0rHTEZYdmQKBgQCgIC/PR1IqR7QWXX5COFJHETNKtH7UB+lTD78qavUNmssxVuuywiIhFK/fwYpfPf8BWu6CH8bebmEhu+OEygJXUYVrDne98bsIyDVuUjLUloc0OeeJfVfTxpCMVN2shCtgCZ5Nc2iRkxJC/2kHM8kfNKrfdZXOqLtH9GO5r1VAwQKBgD50qc4yDe75IGLpGYRFPhlYtA8hsSrv2nb6L1gHl75FWbB4+ddW2iwyEo8Qf88t1VT3vg9wTJlOf6dl4Xh0TBUDUFm2ULrB6WFHq2RQZF1ravBg6+aKhru3ShDFejJ3nnqDmx0Q+CX8nVKnDct+S63akjBniNbrlVWpEyEBGP45"
	appId      = "2021000119603478"
	//沙箱环境
	client = alipay.New(appId, publicKey, privateKey, false)
)

func init() {
	//client.Load
	//client.LoadAliPayPublicCert("应用公钥证书")
	//client.LoadAppPublicCert("支付宝公钥证书")
	//client.LoadAliPayRootCert("支付宝根证书")
}

func WebPageAlipay() {
	var p = alipay.AliPayTradePagePay{}
	p.NotifyURL = "http://www.pibigstar/alipay"
	p.ReturnURL = "http://localhost:8088/return"
	p.Subject = "BLOG_COMMENT_PAY"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	p.OutTradeNo = time.Now().String()
	p.TotalAmount = "3.00"
	url, err := client.TradePagePay(p)
	if err != nil {
		fmt.Println(err)
	}
	payURL := url.String()
	fmt.Println(payURL)
	payURL = strings.Replace(payURL, "&", "^&", -1)
	exec.Command("cmd", "/c", "start", payURL).Start()

}

/**
func WapAlipay() {
	pay := alipay.TradeWapPay{}

	pay.ReturnURL = "http://localhost:8088/return"
	pay.Subject = "支付宝测试"
	pay.OutTradeNo = time.Now().String()
	pay.TotalAmount = "3.00"

	url,err := client.TradeWapPay(pay)
	if err != nil {
		fmt.Println(err)
	}
	payURL := url.String()
	fmt.Println(payURL)
	payURL = strings.Replace(payURL,"&","^&",-1)
	exec.Command("cmd", "/c", "start",payURL).Start()
}

*/
func main() {

	//WebPageAlipay()
	WebPageAlipay()
	http.HandleFunc("/return", func(resp http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		ok, err := client.VerifySign(request.Form)
		if err == nil && ok {
			resp.Write([]byte("支付成功"))
		}

	})
	http.HandleFunc("/alipay", func(rep http.ResponseWriter, req *http.Request) {
		var noti, _ = client.GetTradeNotification(req)
		if noti != nil {
			fmt.Println("支付成功")
			//修改订单状态。。。。
		} else {
			fmt.Println("支付失败")
		}
		alipay.AckNotification(rep) // 确认收到通知消息
	})
	fmt.Println("server start")
	http.ListenAndServe(":8088", nil)
}
