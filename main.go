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
	storage2 "iTask/modules/project_member_invited/storage"
	"iTask/modules/project_members/storage"
	"iTask/modules/project_members/transport/gin"
	gintag "iTask/modules/tag/transport"
	gintask "iTask/modules/task/transport/gin"
	gintaskassignee "iTask/modules/task_assignees/transport/gin"
	gintaskcomments "iTask/modules/task_comments/transport"
	uploadhandler "iTask/modules/upload/handler"
	uploadusecase "iTask/modules/upload/usecase"
	verifyemailshanlder "iTask/modules/verify_emails/handler"
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

	// declare project member storage
	projectMemberStore := storage.NewSQLStore(db)

	projectMemberInvitedStore := storage2.NewSQLStore(db)

	// declare verify email usecase
	verifyEmailsSto := verifyemailsstorage.NewVerifyEmailsStorage(db)
	verifyEmailsUseCase := verifyemailsusecase.NewVerifyEmailsUseCase(verifyEmailsSto, accountSto, projectMemberInvitedStore, projectMemberStore)
	verifyEmailsHdl := verifyemailshanlder.NewVerifyEmailsHandler(verifyEmailsUseCase)

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
	v1.POST("/change/password", middlewares.RequiredAuth(), accountHdl.ChangePassword())
	v1.POST("/change/status", middlewares.RequiredAuth(), middlewares.RequiredRoles(entities.RoleAdmin), accountHdl.ChangeStatusAccount())
	v1.POST("/forgot/password", accountHdl.ForgotPassword())
	v1.POST("/reset/password", accountHdl.ResetPassword())

	// Project
	projects := v1.Group("/projects", middlewares.RequiredAuth())
	{
		projects.GET("/simple-list", ginproject.ListSimpleProjects(db))
		projects.GET("", ginproject.ListProject(db))
		projects.GET("/:id", ginproject.GetProject(db))
		projects.POST("", ginproject.CreateProject(db))
		projects.POST("/:id", ginproject.UpdateProject(db))
		projects.DELETE("/:id", ginproject.DeleteProject(db))
	}

	// ProjectMembers
	projectMembers := v1.Group("/projectMembers", middlewares.RequiredAuth())
	{
		projectMembers.GET("/:project_id", ginprojectmembers.ListMembersById(db))
		projectMembers.POST("/invitation", ginprojectmembers.InviteMember(db, taskDistributor))
		projectMembers.GET("/invited-member", ginprojectmembers.FindUninvitedMember(db))
		projectMembers.GET("/unassigned-members", ginprojectmembers.ListUnassignedMembers(db))
		projectMembers.DELETE("", ginprojectmembers.DeleteMember(db))
		projectMembers.POST("", ginprojectmembers.UpdateMemberRole(db))
	}

	// TaskAssignee
	taskAssignees := v1.Group("/taskAssignees", middlewares.RequiredAuth())
	{
		taskAssignees.GET("", gintaskassignee.ListAssignee(db))
		taskAssignees.POST("", gintaskassignee.CreateAssignee(db))
		taskAssignees.DELETE("", gintaskassignee.DeleteAssignee(db))
	}

	// Task
	tasks := v1.Group("/tasks", middlewares.RequiredAuth())
	{
		tasks.GET("", gintask.ListTask(db))
		tasks.GET("/:id", gintask.GetTask(db))
		tasks.POST("", gintask.CreateTask(db))
		tasks.POST("/:id", gintask.UpdateTask(db))
		tasks.DELETE("/:id", gintask.DeleteTask(db))
	}

	// Tag
	tags := v1.Group("/tags", middlewares.RequiredAuth())
	{
		tags.GET("", gintag.ListTag(db))
		//tags.GET("/:id", gintag.GetProject(db))
		//tags.POST("", gintag.CreateProject(db))
		//tags.POST("/:id", gintag.UpdateProject(db))
	}

	comments := v1.Group("/comments", middlewares.RequiredAuth())
	{
		comments.GET("", gintaskcomments.ListTaskCommentsByTaskId(db))
		comments.POST("", gintaskcomments.CreateTaskComment(db))
		comments.POST("/:id", gintaskcomments.UpdateComment(db))
	}

	// verify email
	v1.GET("/verify_email", verifyEmailsHdl.CheckVerifyCodeIsMatching())

	// verify reset code password
	v1.GET("/verify_reset_password", verifyEmailsHdl.CheckResetCodePasswordIsMatching())

	// project member invitation
	v1.GET("/project-invitation", verifyEmailsHdl.CheckProjectInvitation())

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
