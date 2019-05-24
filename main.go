package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./model"
	"./transport"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//Init database using gorm
	db, err := gorm.Open("mysql", "root:vtnt@/notes?charset=utf8&parseTime=True&loc=Local")
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
