package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/thinhpq0112/soa-backend/config"
	_ "github.com/thinhpq0112/soa-backend/docs"
	"github.com/thinhpq0112/soa-backend/internal/middleware"
	"github.com/thinhpq0112/soa-backend/internal/repository"
	"github.com/thinhpq0112/soa-backend/internal/service"
	"github.com/thinhpq0112/soa-backend/internal/transport"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	loadConfig()

	db, err := config.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	productRepo := repository.NewProductRepo(db)
	categoryRepo := repository.NewCategoryRepo(db)
	supplierRepo := repository.NewSupplierRepo(db)

	productService := service.NewProductService(productRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	supplierService := service.NewSupplierService(supplierRepo)

	//router := gin.Default()
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.LogMiddleWare())

	api := router.Group("/api/")
	api.Use(middleware.AuthMiddleware())
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	productHandler := transport.NewProductHandler(productService)
	productHandler.RegisterRoutes(api)

	categoryHandler := transport.NewCategoryHandler(categoryService)
	categoryHandler.RegisterRoutes(api)

	supplierHandler := transport.NewSupplierHandler(supplierService)
	supplierHandler.RegisterRoutes(api)

	distanceService := service.NewDistanceService()
	distanceHandler := transport.NewDistanceHandler(distanceService)
	distanceHandler.RegisterRoutes(api)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}

func loadConfig() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading .env file: %s", err)
	}
}
