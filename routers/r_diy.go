package routers

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"

	m "back/models"
	u "back/utils"

	"github.com/gin-gonic/gin"
)

func createDIY(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Diy
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.InsertDIY(requestData)
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

func fetchDIYs(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Token
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadGateway,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ token မရှိပါ",
		})
		return
	}

	token, _ := jwt.Parse(requestData.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var tmp m.Admin
		mapstructure.Decode(claims, &tmp)
		if tmp.Password != "" {
			_, err := u.FindAdminByPwd(tmp.Password)
			if err == nil {
				diys, err := u.FindAllDIYS()
				if err != nil {
					context.JSON(http.StatusCreated, gin.H{
						"status":  http.StatusNonAuthoritativeInfo,
						"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
					})
					return
				}
				context.JSON(http.StatusOK, gin.H{
					"status":  http.StatusOK,
					"message": "",
					"diys":    diys,
				})
				return
			}
		} else {
			var tmp m.User
			mapstructure.Decode(claims, &tmp)
			if tmp.Phone != "" {
				fmt.Println(" Found Phone ", tmp.Phone)
				//DB User Status
				_, err := u.FindUserByPhone(tmp.Phone)
				if err == nil {
					fmt.Println(" Found User")
					diys, err := u.FindAllDIYS()
					if err != nil {
						context.JSON(http.StatusCreated, gin.H{
							"status":  http.StatusNonAuthoritativeInfo,
							"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
						})
						return
					}
					context.JSON(http.StatusOK, gin.H{
						"status":  http.StatusOK,
						"message": "",
						"diys":    diys,
					})
					return
				}
			}
		}
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
	}
	return
}

func updateDIY(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Diy
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.UpdateDIY(requestData)
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

func deleteDIY(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Diy
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.DeleteDIY(requestData)
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
