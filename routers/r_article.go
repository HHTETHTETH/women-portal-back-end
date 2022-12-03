package routers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"

	m "back/models"
	u "back/utils"

	"github.com/gin-gonic/gin"
)

func createArticle(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Article
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.InsertArticle(requestData)
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

func fetchArticles(context *gin.Context) {
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
		fmt.Println("Admin..")
		var tmp m.Admin
		mapstructure.Decode(claims, &tmp)
		fmt.Println("tmp : ", tmp)
		if tmp.Password != "" {
			_, err := u.FindAdminByPwd(tmp.Password)
			if err == nil {
				articles, err := u.FindAllArticles()
				if err != nil {
					context.JSON(http.StatusCreated, gin.H{
						"status":  http.StatusNonAuthoritativeInfo,
						"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
					})
					return
				}
				context.JSON(http.StatusOK, gin.H{
					"status":   http.StatusOK,
					"message":  "",
					"articles": articles,
				})
				return
			}
		} else {
			fmt.Println("User..")
			var tmp m.User
			mapstructure.Decode(claims, &tmp)
			fmt.Println("Here tmp ..", tmp)
			if tmp.Phone != "" {
				_, err := u.FindUserByPhone(tmp.Phone)
				if err == nil {
					fmt.Println(" Found User")
					articles, err := u.FindAllArticles()
					if err != nil {
						context.JSON(http.StatusCreated, gin.H{
							"status":  http.StatusNonAuthoritativeInfo,
							"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
						})
						return
					}
					context.JSON(http.StatusOK, gin.H{
						"status":   http.StatusOK,
						"message":  "",
						"articles": articles,
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

	// context.Header("Content-Type", "application/json")
	// var requestData m.Token
	// if err := context.ShouldBindJSON(&requestData); err != nil {
	// 	context.JSON(http.StatusOK, gin.H{
	// 		"status":  http.StatusBadGateway,
	// 		"message": "လုပ္ေဆာင္ခ်က္ မေအာင္ျမင္ပါ token မရှိပါ",
	// 	})
	// 	return
	// }

	// token, _ := jwt.Parse(requestData.Token, func(token *jwt.Token) (interface{}, error) {
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("There was an error")
	// 	}
	// 	return []byte("secret"), nil
	// })
	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	var tmp m.Admin
	// 	mapstructure.Decode(claims, &tmp)
	// 	if tmp.Password != "" {
	// 		//fmt.Println(" Found Phone ", tmp.Password)
	// 		//DB User Status
	// 		_, err := u.FindAdminByPwd(tmp.Password)
	// 		if err == nil {
	// 			//fmt.Println(" Found User")
	// 			articles, err := u.FindAllArticles()
	// 			if err != nil {
	// 				context.JSON(http.StatusCreated, gin.H{
	// 					"status":  http.StatusNonAuthoritativeInfo,
	// 					"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
	// 				})
	// 				return
	// 			}
	// 			context.JSON(http.StatusOK, gin.H{
	// 				"status":   http.StatusOK,
	// 				"message":  "",
	// 				"articles": articles,
	// 			})
	// 			return
	// 		}
	// 	}
	// } else {
	// 	context.JSON(http.StatusOK, gin.H{
	// 		"status":  http.StatusBadRequest,
	// 		"message": "လုပ္ေဆာင္ခ်က္ မေအာင္ျမင္ပါ",
	// 	})
	// }

	/*var requestData m.Token
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadGateway,
			"message": "လုပ္ေဆာင္ခ်က္ မေအာင္ျမင္ပါ token မရှိပါ",
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
		var tmp m.User
		mapstructure.Decode(claims, &tmp)
		if tmp.Phone != "" {
			fmt.Println(" Found Phone ", tmp.Phone)
			//DB User Status
			_, err := u.FindUserByPhone(tmp.Phone)
			if err == nil {
				fmt.Println(" Found User")
				articles, err := u.FindAllArticles()
				if err != nil {
					context.JSON(http.StatusCreated, gin.H{
						"status":  http.StatusNonAuthoritativeInfo,
						"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
					})
					return
				}
				context.JSON(http.StatusOK, gin.H{
					"status":   http.StatusOK,
					"message":  "",
					"articles": articles,
				})
				return
			}
		}
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ္ေဆာင္ခ်က္ မေအာင္ျမင္ပါ",
		})
	}*/

}

func updateArticle(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Article
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.UpdateArticle(requestData)
	if err == nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"article": requestData,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusInternalServerError,
		"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
	})
	return
}

func deleteArticle(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Article
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}

	err := u.DeleteArticle(requestData)
	if err == nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"article": requestData,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusInternalServerError,
		"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
	})
	return
}

func uploadImage(context *gin.Context) {
	//_, header, err := context.Request.FormFile("upload")
	// _, header, err := context.Request.FormFile("file")
	// if err == nil {
	// 	fmt.Println(header.Filename)
	// } else {
	// 	fmt.Println(err)
	// }

	_, header, err := context.Request.FormFile("file")
	if err != nil {
		context.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}

	err = context.SaveUploadedFile(header, "public/"+header.Filename)
	if err != nil {
		log.Fatal(err)
	}
	context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", header.Filename))
	// filename := header.Filename
	// out, err := os.Create("public/" + filename)
	// if err != nil {
	// 	log.Println("err : ", err)
	// }
	// defer out.Close()
	// ans, err := io.Copy(out, file)
	// if err != nil {
	// 	log.Println("err1 : ", err)
	// }
	// log.Println("ans :: ", ans)
	// filepath := "http://localhost:9008/public/" + filename
	// context.JSON(http.StatusOK, gin.H{"filepath": filepath})
}
