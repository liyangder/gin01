package controller

import (
	"bubble/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetList(c *gin.Context) {
	type Param struct {
		Name     string
		Password string
		Age      int
	}
	var param = &Param{}
	param.Name = c.PostForm("username")
	param.Password = c.PostForm("password")
	param.Age, _ = strconv.Atoi(c.PostForm("age"))

	utils.Push(param)

	c.JSON(http.StatusOK, param)
}
