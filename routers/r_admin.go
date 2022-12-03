package routers

import (
	m "back/models"
	u "back/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createAdmin(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Admin
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.InsertAdmin(requestData)
	if err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "အောင်မြင်စွာထည့်ပြီးပါပြီ",
	})
	return
}

func fetchAdmins(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	admins, err := u.FindAllAdmins()
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
		"users":   admins,
	})
}

func fetchAdminByInfo(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Admin
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	_, err := u.FindAdminByInfo(requestData.UserName, requestData.Password)
	if err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	} else {
		// requestData.Password = "09260517328"
		// token, err := u.GenerateToken(requestData.Password)
		// if err != nil {
		// 	context.JSON(http.StatusCreated, gin.H{
		// 		"status":  http.StatusInternalServerError,
		// 		"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		// 	})
		// 	return
		// }
		// context.JSON(http.StatusOK, gin.H{
		// 	"status":  http.StatusOK,
		// 	"message": "",
		// 	"token":   token,
		// })
		// return
		token, err := u.CreateToken(requestData.Password)
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
			"token":   token,
		})
		return
	}
}

func updateAdmin(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Admin
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.UpdateAdmin(requestData)
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

func deleteAdmin(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Admin
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.DeleteAdmin(requestData)
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
