package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

type Bill struct {
	SubMchId   string
	OutTradeNo string
	CheckCode  string
}

func main() {
	r := gin.Default()
	//r.LoadHTMLGlob("templates/*")

	r.GET("/ticket", func(c *gin.Context) {
		bill := &Bill{
			SubMchId:   c.Query("sub_mch_id"),
			OutTradeNo: c.Query("out_trade_no"),
			CheckCode:  c.Query("check_code"),
		}
		t, err := template.ParseFiles("templates/index.tmpl")
		if err != nil {
			log.Println("Parse template failed, err%v", err)
			return
		}
		buf := new(bytes.Buffer)
		err = t.Execute(buf, bill)

		type H struct {
			ContentType string `json:"Content-Type"`
		}
		h := &H{
			ContentType: "text/html",
		}
		c.PureJSON(200, gin.H{
			"isBase64Encoded": false, // 是否使用base64编码，值为 true 或者 false
			"statusCode":      200,   // http请求状态码
			"headers":         h,     // Content-Type 只支持字符串，不支持数组
			"body":            buf.String(),
		})
	})

	r.StaticFile("/WXPAY_verify_1600662107.txt", "static/WXPAY_verify_1600662107.txt")

	// 推销的图片
	r.StaticFile("/coupon.jpg", "static/coupon.jpg")

	r.Run()
}
