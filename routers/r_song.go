package routers

import (
	"net/http"

	m "back/models"
	u "back/utils"

	"github.com/gin-gonic/gin"
)

func fetchAllSingers(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	singers, err := u.FindAllSingers()
	if err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"singers": singers,
	})
	return
}

func createSinger(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Singer
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.InsertSinger(requestData)
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

func updateSinger(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Singer
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.UpdateSinger(requestData)
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

func deleteSinger(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Singer
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.DeleteSinger(requestData)
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

func fetchAllSongs(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	songs, err := u.FindAllSongs()
	if err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"songs":  songs,
	})
	return
}

func createSong(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Song
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.InsertSong(requestData)
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

func updateSong(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Song
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.UpdateSong(requestData)
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

func deleteSong(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Song
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.DeleteSong(requestData)
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

func filterSong(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Singer
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	songs, err := u.FindSongsBySinger(requestData.Name)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadGateway,
			"message": "error",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"songs":  songs,
	})
	return
}
