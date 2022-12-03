package routers

import (
	"net/http"

	m "back/models"
	u "back/utils"

	"github.com/gin-gonic/gin"
)

// fetchInspires -
func fetchInspires(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	inspires, err := u.FindAllInspires()
	if err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"inspires": inspires,
	})
	return
}

func createInspire(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Inspire
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.InsertInspire(requestData)
	if err == nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "အောင်မြင်စွာထည့်ပြီးပါပြီ",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusInternalServerError,
		"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
	})
	return
}

func updateInspire(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Inspire
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.UpdateInspire(requestData)
	if err == nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "အောင်မြင်စွာထည့်ပြီးပါပြီ",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusInternalServerError,
		"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
	})
	return
}

func deleteInspire(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Inspire
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.DeleteInspire(requestData)
	if err == nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "အောင်မြင်စွာဖျက်ပြီးပါပြီ",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusInternalServerError,
		"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
	})
	return
}
