package routers

import (
	"net/http"

	m "back/models"
	u "back/utils"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.User
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.InsertUser(requestData)
	if err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "အောင်မြင်စွာထည့်ပြီးပါပြီ",
	})
	return
}

func fetchUsers(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	users, err := u.FindAllUsers()
	if err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "",
		"users":   users,
	})
}

func fetchUsersByPhone(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.User
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	users, err := u.FindUserByPhone(requestData.Phone)
	if err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "",
		"users":   users,
	})
	return
}
