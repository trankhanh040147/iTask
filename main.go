package main

import (
	cmdworker "iTask/cmd/worker"
	"iTask/config"
	"iTask/entities"
	accounthandler "iTask/modules/account/handler"
	accountstorage "iTask/modules/account/storage"
	accountusecase "iTask/modules/account/usecase"
	"iTask/modules/middleware"
	ginproject "iTask/modules/project/transport/gin"
	gintag "iTask/modules/tag/transport"
	gintask "iTask/modules/task/transport/gin"
	uploadhandler "iTask/modules/upload/handler"
	uploadusecase "iTask/modules/upload/usecase"
	verifyemailsstorage "iTask/modules/verify_emails/storage"
	verifyemailsusecase "iTask/modules/verify_emails/usecase"
	"iTask/provider/cache"
	mysqlprovider "iTask/provider/mysql"
	redisprovider "iTask/provider/redis"
	s3provider "iTask/provider/s3"
	"log"
	"net/http"

	"iTask/worker"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Get config error", err)
		return
	}

	// Declare DB
	db, err := mysqlprovider.NewMySQL(cfg)
	if err != nil {
		log.Fatalln("Can not connect mysql: ", err)
	}

	//utils.RunDBMigration(cfg)

	// Declare redis
	redis, err := redisprovider.NewRedisClient(cfg)
	if err != nil {
		log.Fatalln("Can not connect redis: ", err)
	}

	// declare redis client for asynq
	redisOpt := asynq.RedisClientOpt{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
	}

	// cache residis
	cacheRedis := cache.NewRedisCache(redis)

	// declare task distributor
	taskDistributor := worker.NewRedisTaskDistributor(&redisOpt)

	// google map
	// googleMap := googlemapprovider.NewGoogleMap(cfg)

	// momo
	//momo := momoprovider.NewMomo(cfg)

	// declare dependencies account
	accountSto := accountstorage.NewAccountStorage(db)
	accountCache := cache.NewAuthUserCache(accountSto, cacheRedis)

	// declare verify email usecase
	verifyEmailsSto := verifyemailsstorage.NewVerifyEmailsStorage(db)
	verifyEmailsUseCase := verifyemailsusecase.NewVerifyEmailsUseCase(verifyEmailsSto, accountSto)

	accountUseCase := accountusecase.NewUserUseCase(cfg, accountSto, verifyEmailsUseCase, taskDistributor)
	accountHdl := accounthandler.NewAccountHandler(cfg, accountUseCase)

	// upload file to s3
	s3Provider := s3provider.NewS3Provider(cfg)
	uploadUC := uploadusecase.NewUploadUseCase(cfg, s3Provider)
	uploadHdl := uploadhandler.NewUploadHandler(cfg, uploadUC)

	// run task processor
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		cmdworker.RunTaskProcessor(&redisOpt, accountSto, cfg, verifyEmailsUseCase)
		cmdworker.RunTaskProcessor(&redisOpt, accountSto, cfg, verifyEmailsUseCase)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cmdworker.RunTaskScheduler(&redisOpt, cfg)
	}()

	router := gin.Default()

	// config CORS
	configCORS := setupCors()
	router.Use(cors.New(configCORS))

	middlewares := middleware.NewMiddlewareManager(cfg, accountCache)
	router.Use(middlewares.Recover())

	v1 := router.Group("/api/v1")

	// health check
	v1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Hello": "World"})
	})
	v1.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	// User
	v1.POST("/register", accountHdl.RegisterAccount())
	v1.POST("/login", accountHdl.LoginAccount())
	v1.PATCH("/account/:id", accountHdl.UpdatePersonalInfoAccountById())
	v1.GET("/profile", accountHdl.GetAccountByEmail())
	v1.GET("/profile/:id", accountHdl.GetAccountByID())
	v1.GET("/accounts", middlewares.RequiredAuth(), middlewares.RequiredRoles(entities.RoleAdmin), accountHdl.GetAllAccountUserAndVendor())
	v1.PATCH("/account/role/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles(entities.RoleAdmin), accountHdl.UpdateAccountRoleByID())
	v1.POST("/change/status", middlewares.RequiredAuth(), middlewares.RequiredRoles(entities.RoleAdmin), accountHdl.ChangeStatusAccount())

	// Project
	projects := v1.Group("/projects", middlewares.RequiredAuth())
	{
		projects.GET("", ginproject.ListProject(db))
		projects.GET("/:id", ginproject.GetProject(db))
		projects.POST("", ginproject.CreateProject(db))
		projects.POST("/:id", ginproject.UpdateProject(db))
		projects.DELETE("/:id", ginproject.DeleteProject(db))
	}

	// Task
	tasks := v1.Group("/tasks", middlewares.RequiredAuth())
	{
		tasks.GET("", gintask.ListTask(db))
		//projects.GET("/:id", ginproject.GetTask(db))
		//projects.POST("", ginproject.CreateTask(db))
		tasks.POST("/:id", gintask.UpdateTask(db))
	}

	// Tag
	tags := v1.Group("/tags", middlewares.RequiredAuth())
	{
		tags.GET("", gintag.ListTag(db))
		//tags.GET("/:id", gintag.GetProject(db))
		//tags.POST("", gintag.CreateProject(db))
		//tags.POST("/:id", gintag.UpdateProject(db))
	}

	// upload file to s3
	v1.POST("/upload", middlewares.RequiredAuth(), uploadHdl.UploadFile())

	// google login
	//v1.GET("/google/login")
	router.Run(":" + cfg.App.Port)
	wg.Wait()

}

func setupCors() cors.Config {
	configCORS := cors.DefaultConfig()
	configCORS.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	configCORS.AllowHeaders = []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Cache-Control", "X-Requested-With", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials"}
	configCORS.AllowCredentials = true
	//configCORS.AllowOrigins = []string{"http://localhost:3000"}
	configCORS.AllowAllOrigins = true

	return configCORS
}
