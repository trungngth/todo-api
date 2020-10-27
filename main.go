package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"todo/model"
	"todo/transport"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	userName = "trungnt"
	password = "trungnt"
	dbName   = "notes"
)

func main() {
	//Init database using gorm
	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
			userName,
			password,
			dbName,
		))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&model.Note{}, &model.User{})

	//Init the default router
	r := gin.Default()
	gin.SetMode(gin.DebugMode)

	//Init the routes
	transport.InitRoutes(r, db)

	//log.Fatal(http.ListenAndServe(":8080", nil))
	srv := &http.Server{
		Addr:        ":" + "8080",
		Handler:     r,
		IdleTimeout: time.Duration(30) * time.Second,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, calcelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	calcelFunc()
	srv.Shutdown(ctx)

}
