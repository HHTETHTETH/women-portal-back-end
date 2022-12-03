package routers

import (
	"net/http"

	m "back/models"
	u "back/utils"
	"github.com/gin-gonic/gin"
)

func createClickMap(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.ClickMap
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
		})
		return
	}
	err := u.InsertClickMap(requestData)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}
