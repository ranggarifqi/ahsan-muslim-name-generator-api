package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	// db := postgresql.InitDB()

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World")
	// })

	// v1Group := e.Group("/api/v1")

	// urlRepository := urlRepo.NewUrlRepository(db)
	// urlUsecase := urlUC.NewUrlUsecase(urlRepository)
	// urlH.NewUrlHandler(v1Group, urlUsecase)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
