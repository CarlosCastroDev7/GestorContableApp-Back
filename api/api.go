package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ExecuteAPI() {
	servername := viper.GetString("Microservice.name")
	fecha := time.Now()
	logName := fmt.Sprintf("%s.%d.%d.%d.log", servername, fecha.Day(), fecha.Month(), fecha.Year())

	file, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error al leer el archivo de logs ", err)
		return
	}
	mw := io.MultiWriter(file)
	log.SetOutput(mw)
	logrus.SetOutput(mw)

	gin.DefaultWriter = io.MultiWriter(file)
	gin.DisableConsoleColor()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
	}))

	// group := router.Group("/api")

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", ConfigPort()),
		Handler:        router,
		ReadTimeout:    2 * time.Minute,
		WriteTimeout:   2 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		logrus.Infof("Serving api at http://127.0.0.1:%d", ConfigPort())
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			logrus.Error(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logrus.Info("signal caught. shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown error:", err)
	}

	select {
	case <-ctx.Done():
		logrus.Info("Server down.")
	}
}
