package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cecepsprd/foodstore-server/config"
	"github.com/cecepsprd/foodstore-server/handler"
	"github.com/cecepsprd/foodstore-server/repository"
	"github.com/cecepsprd/foodstore-server/service"
	"github.com/cecepsprd/foodstore-server/utils/logger"
	"github.com/cecepsprd/foodstore-server/utils/validate"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

func RunServer() {
	var cfg = config.NewConfig()

	db, err := cfg.MongoConnect()
	if err != nil {
		log.Fatal("error connecting to database: ", err.Error())
	}

	if err = logger.Init(cfg.App.LogLevel, cfg.App.LogTimeFormat); err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete,
		},
	}))

	customValidator := validate.NewValidator()
	en_translations.RegisterDefaultTranslations(customValidator.Validator, customValidator.Translator)
	e.Validator = customValidator

	timeoutContext := time.Duration(cfg.App.ContextTimeout) * time.Second

	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)
	tagRepository := repository.NewTagRepository(db)
	cartRepository := repository.NewCartRepository(db)

	userService := service.NewUserService(userRepository, timeoutContext)
	authService := service.NewAuthService(userService, cfg.App.JWTSecret)
	productService := service.NewProductService(productRepository, timeoutContext)
	categoryService := service.NewCategoryService(categoryRepository, timeoutContext)
	tagService := service.NewTagService(tagRepository, timeoutContext)
	cartService := service.NewCartService(cartRepository, timeoutContext)

	handler.NewAuthHandler(e, authService, userService)
	handler.NewProductHandler(e, productService)
	handler.NewCategoryHandler(e, categoryService)
	handler.NewTagHandler(e, tagService)
	handler.NewRegionHandler(e)
	handler.NewCartHandler(e, cartService)

	e.Static("/api/images", "images")

	// Starting server
	go func() {
		err := e.Start(cfg.App.HTTPPort)
		if err != nil {
			log.Fatal("error starting server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	// Block until a signal is received.
	<-quit

	log.Println("server shutdown of 5 second.")

	// gracefully shutdown the server, waiting max 5 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	e.Shutdown(ctx)
}
