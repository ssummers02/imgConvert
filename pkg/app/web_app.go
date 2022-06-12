package app

import (
	"context"
	"imgConverter/pkg/handler"
	"imgConverter/pkg/service"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func RunWeb(path string) {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(path); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	err := os.MkdirAll(handler.ImgDirectory, os.ModePerm)
	if err != nil {
		logrus.Fatalf("error create directory: %s", err.Error())

		return
	}

	services := service.NewService()
	srv := handler.NewServer(viper.GetString("port"), services)

	go func() {
		if err := srv.Run(); err != nil {
			logrus.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()

	logrus.Print("ImgApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down: %s", err.Error())
	}
}

func initConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	return viper.ReadInConfig()
}
