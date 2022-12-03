package routers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	u "back/utils"
	"github.com/gin-gonic/gin"
)

func getHoroDate(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	requestData := context.Query("x-date")
	fmt.Println(" requested Date ", requestData)
	if requestData != "" {
		index := u.Horodate(requestData)
		context.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"index":  index,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"index":  0,
	})

}
func getmeme(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	requestData := context.Query("x-query")
	fmt.Println(" requested MEME ", requestData)
	switch requestData {
	case "how_much_skill_do_you_have_in_cooking":
		id := rand.Intn(4-1) + 1
		url := "http://khintkabar.com/resources/herportal/meme/cooking/" + strconv.Itoa(id) + ".png"
		time.Sleep(3 * time.Second)
		context.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"url":    url,
		})
	case "which_princess_is_look_like_you":
		id := rand.Intn(4-1) + 1
		url := "http://khintkabar.com/resources/herportal/meme/prince/" + strconv.Itoa(id) + ".png"
		time.Sleep(3 * time.Second)
		context.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"url":    url,
		})
	default:
		id := rand.Intn(3-1) + 1
		url := "http://khintkabar.com/resources/herportal/" + strconv.Itoa(id) + ".jpg"
		time.Sleep(3 * time.Second)
		context.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"url":    url,
		})
	}

}
