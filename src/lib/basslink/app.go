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
	Recaptcha  *RecaptchaClient
	Mailgun    *MailgunClient

	EmailMsgChannel chan *EmailNotificationMesage
	SignalChannel   chan os.Signal
}

type Config struct {
	JwtKey                  string           `json:"jwt_key"`
	DB                      *DBConfig        `json:"db"`
	Http                    *HttpConfig      `json:"http"`
	Recaptcha               *RecaptchaConfig `json:"recaptcha"`
	Storage                 *StorageConfig   `json:"storage"`
	Mailgun                 *MailgunConfig   `json:"mailgun"`
	PaymentConfirmationLink string           `json:"payment_confirmation_link"`
}

func New(serviceName string, config *Config) *App {
	return &App{
		serviceName:     serviceName,
		Config:          config,
		EmailMsgChannel: make(chan *EmailNotificationMesage, 5),
		SignalChannel:   make(chan os.Signal),
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

func (app *App) CreateRecaptchaClient() {
	if app.Config.Recaptcha == nil {
		panic("recaptcha config is not set")
	}

	app.Recaptcha = NewRecaptchaClient(app.Config.Recaptcha)
}

func (app *App) CreateMailgunClient() {
	if app.Config.Mailgun == nil {
		panic("mailgun config is not set")
	}

	app.Mailgun = NewMailgunClient(app.Config.Mailgun)
}
