package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	engine "github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/config/gin"
	"github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/config/postgres"
	authrepo "github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/repository/auth"
	commentrepo "github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/repository/comment"
	photorepo "github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/repository/photo"
	socialmediarepo "github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/repository/socialmedia"
	userrepo "github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/repository/user"
	authhandler "github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/server/http/handler/auth"
	commenthandler "github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/server/http/handler/comment"
	photohandler "github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/server/http/handler/photo"
	socialmediahandler "github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/server/http/handler/socialmedia"
	userhandler "github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/server/http/handler/user"
	"github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/server/http/middleware"
	router "github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/server/http/router/v1"
	authusecase "github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/usecase/auth"
	commentusecase "github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/usecase/comment"
	photousecase "github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/usecase/photo"
	socialmediausecase "github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/usecase/socialmedia"
	userusecase "github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/usecase/user"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	postgresHost := os.Getenv("MY_GRAM_POSTGRES_HOST")
	postgresPort := os.Getenv("MY_GRAM_POSTGRES_PORT")
	postgresDatabase := os.Getenv("MY_GRAM_POSTGRES_DATABASE")
	postgresUsername := os.Getenv("MY_GRAM_POSTGRES_USERNAME")
	postgresPassword := os.Getenv("MY_GRAM_POSTGRES_PASSWORD")

	postgresCln := postgres.NewPostgresConnection(postgres.Config{
		Host:         postgresHost,
		Port:         postgresPort,
		User:         postgresUsername,
		Password:     postgresPassword,
		DatabaseName: postgresDatabase,
	})

	ginEngine := engine.NewGinHttp(engine.Config{
		Port: ":8080",
	})

	ginEngine.GetGin().Use(
		gin.Recovery(),
		gin.Logger(),
	)

	startTime := time.Now()
	ginEngine.GetGin().GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message":    "server up and running",
			"start_time": startTime,
		})
	})

	userRepo := userrepo.NewUserRepo(postgresCln)
	userUsecase := userusecase.NewUserUsecase(userRepo)
	userHandler := userhandler.NewUserHandler(userUsecase)

	authRepo := authrepo.NewAuthRepo(postgresCln)
	authUsecase := authusecase.NewAuthUsecase(authRepo, userUsecase)
	authHandler := authhandler.NewAuthHandler(authUsecase)

	photoRepo := photorepo.NewPhotoRepo(postgresCln)
	photoUsecase := photousecase.NewPhotoUsecase(photoRepo, userUsecase)
	photoHandler := photohandler.NewPhotoHandler(photoUsecase)

	commentRepo := commentrepo.NewCommentRepo(postgresCln)
	commentUsecase := commentusecase.NewCommentUsecase(commentRepo, photoUsecase)
	commentHandler := commenthandler.NewCommentHandler(commentUsecase)

	socialMediaRepo := socialmediarepo.NewSocialMediaRepo(postgresCln)
	socialMediaUsecase := socialmediausecase.NewSocialMediaUsecase(socialMediaRepo)
	socialMediaHandler := socialmediahandler.NewSocialMediaHandler(socialMediaUsecase)

	authMiddleware := middleware.NewAuthMiddleware(userUsecase)

	router.NewUserRouter(ginEngine, userHandler, authMiddleware).Routers()
	router.NewAuthRouter(ginEngine, authHandler, authMiddleware).Routers()
	router.NewPhotoRouter(ginEngine, photoHandler, authMiddleware).Routers()
	router.NewCommentRouter(ginEngine, commentHandler, authMiddleware).Routers()
	router.NewSocialMediaRouter(ginEngine, socialMediaHandler, authMiddleware).Routers()

	ginEngine.Serve()
}
