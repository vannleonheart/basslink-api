package main

import (
	"CRM/src/lib/basslink"
	"flag"
	"os"
	"os/signal"
	"syscall"
)

const (
	appName        = "api"
	appVersion     = "1.0.0"
	timeZone       = "Asia/Jakarta"
	dateTimeFormat = "2006-01-02 15:04:05"
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

	app.ConnectToDatabase()
	app.CreateHttpService()

	initRouter()

	go func() {
		if err := app.HttpServer.Start(); err != nil {
			panic(err)
		}
	}()

	for {
		select {
		case sig := <-app.SignalChannel:
			go handleQuitSignal(sig)
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

func handlePanicRecovery(source string, data interface{}) {
	if r := recover(); r != nil {

	}
}

func handleQuitSignal(sig os.Signal) {
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
