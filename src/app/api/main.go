package main

import (
	"CRM/src/lib/basslink"
	"flag"
	"os"
	"os/signal"
	"syscall"
)

const (
	appName  = "api"
	timeZone = "Asia/Jakarta"
)

var (
	configFile *string
	app        *basslink.App
)

func main() {
	readFlags()

	app = basslink.New(appName, nil)

	app.LoadConfigFromFile(*configFile)
	app.LoadLocation(timeZone)

	signal.Notify(app.SignalChannel,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGHUP,
	)

	app.CreateRecaptchaClient()
	app.ConnectToDatabase()
	app.CreateStorageClient()
	app.CreateHttpService()
	app.CreateMailgunClient()

	initRouter()

	go func() {
		if err := app.HttpServer.Start(); err != nil {
			panic(err)
		}
	}()

	for {
		select {
		case <-app.SignalChannel:
			go handleQuitSignal()
		case msg := <-app.EmailMsgChannel:
			go app.HandleEmailNotification(msg)
		}
	}
}

func readFlags() {
	configFile = flag.String("config", "", "configuration file path")

	flag.Parse()

	if configFile == nil || len(*configFile) <= 0 {
		panic("config file is not set")
	}
}

func handleQuitSignal() {
	defer func() {
		os.Exit(0)
	}()

	if app.DB != nil {
		if err := app.DB.Close(); err != nil {

		} else {

		}
	}

	if app.HttpServer != nil {
		if err := app.HttpServer.Stop(); err != nil {

		} else {

		}
	}
}
