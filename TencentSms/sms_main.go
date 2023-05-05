package main

import (
	"TencentSms/sms/tencent"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	engine := gin.Default()

	engine.GET("/sendMsg", func(context *gin.Context) {
		sms := tencent.Sms{}
		phones := []string{"18603293432"}
		_, err := sms.Send(phones, []string{"122322", "4"})
		context.JSONP(http.StatusOK, gin.H{
			"msg": "success",
			"err": err,
		})
	})
	engine.Run()

}
