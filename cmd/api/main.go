package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/database/postgresql"
	authH "github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/auth/handler"
	userH "github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user/handler"
	userRepo "github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user/repository"
	userUC "github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user/usecase"
	myValidator "github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/validator"
)

func main() {
	godotenv.Load()
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8000", "https://qashir.ranggarifqi.com"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Validator = myValidator.NewMyValidator()

	db := postgresql.InitDB()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	v1Group := e.Group("/api/v1")

	// urlRepository := urlRepo.NewUrlRepository(db)
	// urlUsecase := urlUC.NewUrlUsecase(urlRepository)
	// urlH.NewUrlHandler(v1Group, urlUsecase)

	userRepository := userRepo.NewUserRepository(db)
	userUsecase := userUC.NewUserUsecase(userRepository)
	userH.NewUserHandler(v1Group, userUsecase)

	authH.NewAuthHandler(v1Group, userUsecase)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
