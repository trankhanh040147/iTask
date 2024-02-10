package main

import (
	"log"
	"net/http"
	cmdworker "paradise-booking/cmd/worker"
	"paradise-booking/config"
	"paradise-booking/constant"
	accounthandler "paradise-booking/modules/account/handler"
	accountstorage "paradise-booking/modules/account/storage"
	accountusecase "paradise-booking/modules/account/usecase"
	amenityhandler "paradise-booking/modules/amenity/handler"
	amenitystorage "paradise-booking/modules/amenity/storage"
	amenityusecase "paradise-booking/modules/amenity/usecase"
	bookinghandler "paradise-booking/modules/booking/handler"
	bookingstorage "paradise-booking/modules/booking/storage"
	bookingusecase "paradise-booking/modules/booking/usecase"
	bookingdetailstorage "paradise-booking/modules/booking_detail/storage"
	bookingratinghandler "paradise-booking/modules/booking_rating/handler"
	bookingratingstorage "paradise-booking/modules/booking_rating/storage"
	bookingratingusecase "paradise-booking/modules/booking_rating/usecase"
	"paradise-booking/modules/middleware"
	paymenthandler "paradise-booking/modules/payment/handler"
	paymentstorage "paradise-booking/modules/payment/store"
	paymentusecase "paradise-booking/modules/payment/usecase"
	placehandler "paradise-booking/modules/place/handler"
	placestorage "paradise-booking/modules/place/storage"
	placeusecase "paradise-booking/modules/place/usecase"
	placewishlisthandler "paradise-booking/modules/place_wishlist/handler"
	placewishliststorage "paradise-booking/modules/place_wishlist/storage"
	placewishlistusecase "paradise-booking/modules/place_wishlist/usecase"
	policieshandler "paradise-booking/modules/policy/handler"
	policiesstorage "paradise-booking/modules/policy/storage"
	policiesusecase "paradise-booking/modules/policy/usecase"
	uploadhandler "paradise-booking/modules/upload/handler"
	uploadusecase "paradise-booking/modules/upload/usecase"
	verifyemailshanlder "paradise-booking/modules/verify_emails/handler"
	verifyemailsstorage "paradise-booking/modules/verify_emails/storage"
	verifyemailsusecase "paradise-booking/modules/verify_emails/usecase"
	wishlisthandler "paradise-booking/modules/wishlist/handler"
	wishliststorage "paradise-booking/modules/wishlist/storage"
	wishlistusecase "paradise-booking/modules/wishlist/usecase"
	"paradise-booking/provider/cache"
	googlemapprovider "paradise-booking/provider/googlemap"
	momoprovider "paradise-booking/provider/momo"
	mysqlprovider "paradise-booking/provider/mysql"
	redisprovider "paradise-booking/provider/redis"
	s3provider "paradise-booking/provider/s3"
	"paradise-booking/utils"
	"paradise-booking/worker"
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

	utils.RunDBMigration(cfg)

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
	googleMap := googlemapprovider.NewGoogleMap(cfg)

	// momo
	momo := momoprovider.NewMomo(cfg)

	// declare dependencies account
	accountSto := accountstorage.NewAccountStorage(db)
	accountCache := cache.NewAuthUserCache(accountSto, cacheRedis)

	// declare verify email usecase
	verifyEmailsSto := verifyemailsstorage.NewVerifyEmailsStorage(db)
	verifyEmailsUseCase := verifyemailsusecase.NewVerifyEmailsUseCase(verifyEmailsSto, accountSto)
	verifyEmailsHdl := verifyemailshanlder.NewVerifyEmailsHandler(verifyEmailsUseCase)

	accountUseCase := accountusecase.NewUserUseCase(cfg, accountSto, verifyEmailsUseCase, taskDistributor)
	accountHdl := accounthandler.NewAccountHandler(cfg, accountUseCase)

	// prepare for placewishlist storeage
	placeWishListSto := placewishliststorage.NewPlaceWishListStorage(db)
	// declare cache for place_wishlist
	placeWishListCache := cache.NewPlaceWishListCache(placeWishListSto, cacheRedis)

	// declare dependencies
	// prepare for wish list
	wishListSto := wishliststorage.NewWishListStorage(db)
	wishListUseCase := wishlistusecase.NewWishListUseCase(wishListSto, placeWishListSto, cacheRedis)
	wishListHdl := wishlisthandler.NewWishListHandler(wishListUseCase)

	// prepare for payment
	paymentSto := paymentstorage.NewPaymentStorage(db)
	paymentUC := paymentusecase.NewPaymentUseCase(paymentSto)
	paymentHdl := paymenthandler.NewPaymentHandler(paymentUC)
	// prepare for place
	bookingSto := bookingstorage.NewBookingStorage(db)

	placeSto := placestorage.NewPlaceStorage(db)
	placeCache := cache.NewPlaceStoCache(placeSto, cacheRedis)
	placeUseCase := placeusecase.NewPlaceUseCase(cfg, placeSto, accountCache, googleMap, placeWishListCache, placeCache, bookingSto)
	placeHdl := placehandler.NewPlaceHandler(placeUseCase)

	// prepare for booking detail
	bookingDetailSto := bookingdetailstorage.NewBookingDetailStorage(db)

	// prepare for booking
	bookingUseCase := bookingusecase.NewBookingUseCase(bookingSto, bookingDetailSto, cfg, taskDistributor, accountSto, placeSto, momo, paymentSto)
	bookingHdl := bookinghandler.NewBookingHandler(bookingUseCase)

	// prepare place wish list
	placeWishListUseCase := placewishlistusecase.NewPlaceWishListUseCase(placeWishListSto, placeSto, cacheRedis)
	placeWishListHdl := placewishlisthandler.NewPlaceWishListHandler(placeWishListUseCase)

	// prepare for place rating
	bookingRatingSto := bookingratingstorage.Newbookingratingstorage(db)
	bookingRatingUC := bookingratingusecase.Newbookingratingusecase(bookingRatingSto, accountSto, placeSto, cacheRedis)
	bookingRatingHdl := bookingratinghandler.Newbookingratinghandler(bookingRatingUC)

	// prepare for amenities
	amenitySto := amenitystorage.NewAmenityStorage(db)
	amenityUC := amenityusecase.NewAmenityUseCase(amenitySto, cfg)
	amenityHdl := amenityhandler.NewAmenityHandler(amenityUC)

	// upload file to s3
	s3Provider := s3provider.NewS3Provider(cfg)
	uploadUC := uploadusecase.NewUploadUseCase(cfg, s3Provider)
	uploadHdl := uploadhandler.NewUploadHandler(cfg, uploadUC)

	// prepare for policy
	policySto := policiesstorage.NewPolicyStorage(db)
	policyUC := policiesusecase.NewPolicyUseCase(policySto)
	policyHdl := policieshandler.NewPolicyHandler(policyUC)

	// run task processor
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		cmdworker.RunTaskProcessor(&redisOpt, accountSto, cfg, verifyEmailsUseCase, bookingSto, bookingUseCase)
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
	v1.GET("/accounts", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.AdminRole), accountHdl.GetAllAccountUserAndVendor())
	v1.PATCH("/account/role/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.AdminRole), accountHdl.UpdateAccountRoleByID())
	v1.POST("/change/password", middlewares.RequiredAuth(), accountHdl.ChangePassword())
	v1.POST("/change/status", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.AdminRole), accountHdl.ChangeStatusAccount())
	v1.POST("/forgot/password", accountHdl.ForgotPassword())
	v1.POST("/reset/password", accountHdl.ResetPassword())

	// // Place
	// v1.POST("/places", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), placeHdl.CreatePlace())
	// v1.PUT("/places", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), placeHdl.UpdatePlace())
	// v1.GET("/places/:id", placeHdl.GetPlaceByID())
	// v1.GET("/places/owner", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), placeHdl.ListPlaceByVendor())
	// v1.GET("/places/owner/:vendor_id", placeHdl.ListPlaceByVendorID())
	// v1.DELETE("/places", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), placeHdl.DeletePlaceByID())
	// v1.POST("/places/list", placeHdl.ListAllPlace())
	// v1.GET("/places/dates_booked", placeHdl.GetDatesBookedPlace())
	// v1.GET("/places/check_date_available", placeHdl.CheckDateBookingAvailable())
	// v1.GET("/places/status_booking", placeHdl.GetStatusPlaceToBook())

	// // booking
	// v1.POST("/bookings", bookingHdl.CreateBooking())
	// v1.GET("/confirm_booking", bookingHdl.UpdateStatusBooking())
	// v1.POST("/booking_list", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), bookingHdl.ListBooking())
	// v1.GET("/bookings/:id", bookingHdl.GetBookingByID())
	// v1.GET("/bookings", middlewares.RequiredAuth(), bookingHdl.GetBookingByPlaceID())
	// v1.GET("/bookings_list/manage_reservation", middlewares.RequiredAuth(), bookingHdl.ListBookingNotReservation())
	// v1.DELETE("/bookings/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), bookingHdl.DeleteBookingByID())
	// v1.POST("/cancel_booking", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), bookingHdl.CancelBookingByID())

	// // wish list
	// v1.POST("/wish_lists", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), wishListHdl.CreateWishList())
	// v1.GET("/wish_lists/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), wishListHdl.GetWishListByID())
	// v1.GET("/wish_lists/user/:user_id", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), wishListHdl.GetWishListByUserID())
	// v1.PUT("/wish_lists/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), wishListHdl.UpdateWishListByID())
	// v1.DELETE("/wish_lists/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), wishListHdl.DeleteWishListByID())

	// // place wish list
	// v1.POST("/place_wish_lists", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), placeWishListHdl.CreatePlaceWishList())
	// v1.DELETE("/place_wish_lists/:place_id/:wishlist_id", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), placeWishListHdl.DeletePlaceWishList())
	// v1.GET("/place_wish_lists/place", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), placeWishListHdl.ListPlaceByWishListID())

	// // booking rating
	// v1.POST("/booking_ratings", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), bookingRatingHdl.MakeComment())
	// v1.GET("/booking_ratings/places/:id", bookingRatingHdl.GetCommentByPlaceID())
	// v1.GET("/booking_ratings/bookings/:id", bookingRatingHdl.GetCommentByBookingID())
	// v1.GET("/booking_ratings/users/:id", bookingRatingHdl.GetCommentByUserID())
	// v1.GET("/booking_ratings/vendors/:id", bookingRatingHdl.GetCommentByVendorID())
	// v1.GET("/booking_ratings/statistics/:place_id", bookingRatingHdl.GetStatisTicsByPlaceId())

	// verify email
	v1.GET("/verify_email", verifyEmailsHdl.CheckVerifyCodeIsMatching())

	// verify reset code password
	v1.GET("/verify_reset_password", verifyEmailsHdl.CheckResetCodePasswordIsMatching())

	// upload file to s3
	v1.POST("/upload", middlewares.RequiredAuth(), uploadHdl.UploadFile())

	// amenities
	v1.POST("/amenities", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), amenityHdl.CreateAmenity())
	v1.DELETE("/amenities/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), amenityHdl.DeleteAmenityByID())
	v1.GET("/amenities/config", amenityHdl.GetAllConfigAmenity())
	v1.GET("/amenities/place/:place_id", amenityHdl.ListAmenityByPlaceID())
	v1.POST("/amenities/place/remove", amenityHdl.DeleteAmenityByListID())

	// policies
	v1.POST("/policies", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), policyHdl.UpsertPolicy())
	v1.GET("/policies/:place_id", policyHdl.GetPolicyByPlaceId())

	// payment
	v1.POST("/payments/list_by_vendor", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), paymentHdl.ListPaymentByVendorID())

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
