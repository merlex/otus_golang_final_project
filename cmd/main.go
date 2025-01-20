package main

import (
	"context"
	"flag"
	"io"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/merlex/otus-image-previewer/internal/config"
	"github.com/merlex/otus-image-previewer/internal/http"
	"github.com/merlex/otus-image-previewer/internal/logger"
	srv "github.com/merlex/otus-image-previewer/internal/service"
)

var (
	configFile string
	wg         sync.WaitGroup
)

func init() {
	flag.StringVar(&configFile, "configs", "./config.yaml",
		"Path to configuration file")
}

func main() {
	flag.Parse()

	cfg := config.ReadConfig(configFile)
	var w io.Writer

	if cfg.Logger.LogToFile {
		file, err := os.OpenFile(cfg.Logger.File, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
		if err != nil {
			log.Println("error opening logfile: " + err.Error())
			return
		}
		defer file.Close()

		if cfg.Logger.LogToConsole {
			w = file
		} else {
			w = io.MultiWriter(file, os.Stdout)
		}
	} else {
		w = os.Stdout
	}
	applogger := logger.New(cfg.Logger.Level, w)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	service := srv.NewImageProcessingService(applogger, cfg.Cache)

	httpServer := http.NewServer(ctx, applogger, cfg.HTTP, service)

	wg.Add(1)
	go func() {
		defer wg.Done()
		httpServer.Start(ctx)
	}()

	wg.Wait()
}
