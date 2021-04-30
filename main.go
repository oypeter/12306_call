package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
)

type CallParam struct {
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func makeCall(c *gin.Context) {
	var json CallParam
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.Password != "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	clnt := ypclnt.New("")
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = ""
	param[ypclnt.CODE] = "111111"
	response := clnt.Voice().Send(param)
	if response.Code == 0 {
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": nil})
	} else {
		fmt.Println(time.Now(), "Error", response.Data, response.Detail, response.Msg)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": response.Msg})
	}
}

func main() {
	r := gin.Default()
	r.POST("/call", makeCall)
	r.Run(":8999")
}
