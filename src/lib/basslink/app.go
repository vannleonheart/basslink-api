package basslink

import (
	"os"
	"time"

	"github.com/vannleonheart/goutil"
)

type App struct {
	serviceName string

	Config     *Config
	Tz         *time.Location
	DB         *DBClient
	HttpServer *HttpServer
	Storage    *StorageClient

	SignalChannel chan os.Signal
}

type Config struct {
	JwtKey  string         `json:"jwt_key"`
	DB      *DBConfig      `json:"db"`
	Http    *HttpConfig    `json:"http"`
	Storage *StorageConfig `json:"storage"`
}

func New(serviceName string, config *Config) *App {
	return &App{
		serviceName:   serviceName,
		Config:        config,
		SignalChannel: make(chan os.Signal),
	}
}

func (app *App) LoadLocation(timeZone string) {
	tz, err := time.LoadLocation(timeZone)

	if err != nil {
		panic(err)
	}

	app.Tz = tz
}

func (app *App) LoadConfigFromFile(configFile string) {
	if _, err := goutil.LoadJsonFile(configFile, &app.Config); err != nil {
		panic(err)
	}
}

func (app *App) ConnectToDatabase() {
	if app.Config.DB == nil {
		panic("database config is not set")
	}

	app.DB = NewDBClient(app.Config.DB)

	if err := app.DB.Connect(); err != nil {
		panic(err)
	}
}

func (app *App) CreateHttpService() {
	if app.Config.Http == nil {
		panic("http config is not set")
	}

	app.HttpServer = NewHttpServer(*app.Config.Http, app.serviceName)
}

func (app *App) CreateStorageClient() {
	if app.Config.Storage == nil {
		panic("storage config is not set")
	}

	cl, err := NewStorageClient(app.Config.Storage)
	if err != nil {
		panic(err)
	}

	app.Storage = cl
}
