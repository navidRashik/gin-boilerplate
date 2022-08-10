package controllers

import (
	"github.com/gin-gonic/gin"
	"gin-boilerplate/models" 
	"net/http"
)

func GetData(ctx *gin.Context) {
	var example []models.Example
	models.GetAll(&example)
	ctx.JSON(http.StatusOK, &example)

}