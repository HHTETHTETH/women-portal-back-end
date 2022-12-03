package routers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitConnection - Create Routers
func InitConnection(portNumber string) {
	router := gin.Default()
	gin.SetMode(gin.DebugMode) //
	//router.Use(gin.Logger(//))
	//router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://khintkabar.com", "http://13.251.20.26:3013/landing", "*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://khintkabar.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	//router.Use(cors.Default())

	v1 := router.Group("/api/v1/")
	{
		//MA Integration
		//v1.GET("/callback", maCallBack)
		//v1.POST("/login", maLogin)
		//v1.POST("/login-check", maLoginCheck)
		//v1.POST("/status", maStatus)
		//v1.POST("/cg", maCg)
		//router.GET("/welcome", maWelcome)
		//route.GET("/login", maLogin)
		//router.GET("/login", maLogin)

		router.GET("/he", heCheck)
		//router.GET("/", maCheck)
		router.POST("/clickMap", createClickMap)
		router.GET("/subscribe2", maForceMARedirect)
		router.POST("/subscribe", maSubscribe)
		router.POST("/unsubscribe", maUnsubscribe)
		router.POST("/login", maLogin)
		router.POST("/loginCheck", maOTPCheck)
		router.POST("/status", maStatus)
		router.POST("/statusNumber", maStatusByNumber)
		router.POST("/profile", maProfile)
		router.GET("/welcome", welcome)
		router.GET("/resource", resourceBlocker)

		v1.POST("/add/user", createUser)
		v1.GET("/users", fetchUsers)
		v1.POST("/get/user/phone", fetchUsersByPhone)

		//admin
		v1.GET("/admins", fetchAdmins)
		v1.POST("/get/admin/info", fetchAdminByInfo)
		v1.POST("/add/admin", createAdmin)
		v1.POST("/update/admin", updateAdmin)
		v1.POST("/delete/admin", deleteAdmin)

		v1.GET("/callback", handleCallbackAPI2)
		v1.GET("/callbacks", callbacks) //TestBed

		//Article CRUD
		v1.POST("/add/article", createArticle)
		v1.POST("/articles", fetchArticles)
		v1.POST("/update/article", updateArticle)
		v1.POST("/delete/article", deleteArticle)

		//Promotion CRUD
		v1.POST("/add/promotion", createPromotion)
		v1.POST("/promotions", fetchPromotion)
		v1.POST("/promotions/date", getPromotionByDate)

		//Cook CURD
		v1.POST("/cook/newcategory", createNewFoodCategory)
		v1.POST("/cook/getcategories", fetchCategories)
		v1.POST("/cook/deletecategory", deleteCategory) //fixed wrong function
		v1.POST("/cook/updatecategory", updateCateogry) //fixed wrong function
		v1.POST("/cook/recipe", fetchRecipes)
		v1.POST("/cook/add/recipe", createRecipes)
		v1.POST("/cook/update/recipe", updateRecipes) //fixed return wrong message
		v1.POST("/cook/delete/recipe", deleteRecipes) //fixed return wrong message
		v1.POST("/cook/filter", filterRecipes)

		//Fashion CURD
		v1.POST("/update/fashion", updateFashion)
		v1.POST("/delete/fashion", deleteFashion)
		v1.POST("/add/fashion", createFashion)
		v1.POST("/fashions", fetchFashion)

		//Song CURD
		v1.GET("/singers", fetchAllSingers)
		v1.POST("/add/singer", createSinger)
		v1.POST("/update/singer", updateSinger)
		v1.POST("/delete/singer", deleteSinger)

		v1.GET("/songs", fetchAllSongs)
		v1.POST("/add/songs", createSong)
		v1.POST("/update/songs", updateSong)
		v1.POST("/delete/songs", deleteSong)
		v1.POST("song/filter", filterSong)

		//Knowledge
		v1.POST("/books", fetchBooks)
		v1.POST("/add/book", createBook)
		//------------------------------
		v1.POST("/update/book", updateBook)
		v1.POST("/delete/book", deleteBook)

		v1.POST("/movies", fetchMovie)
		v1.POST("/add/movie", createMovie)    //fixed 'didn't catch error and return msg'
		v1.POST("/delete/movie", deleteMovie) //fixed 'didn't catch error and return msg'
		v1.POST("/update/movie", updateMovie) //fixed 'didn't catch error and return msg'

		v1.POST("/inspires", fetchInspires)
		v1.POST("/add/inspire", createInspire)
		v1.POST("/update/inspire", updateInspire)
		v1.POST("/delete/inspire", deleteInspire)

		//DIY
		v1.POST("/diys", fetchDIYs)
		v1.POST("/add/diy", createDIY)
		v1.POST("/update/diy", updateDIY)
		v1.POST("/delete/diy", deleteDIY)

		//Fun
		v1.GET("/horoscope", getHoroDate)
		v1.GET("/meme", getmeme)

		//Image Upload Test
		v1.POST("/upload", uploadImage)
	}
	portNumber = ":" + portNumber
	router.Run(portNumber)
}
