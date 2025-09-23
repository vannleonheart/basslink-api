package basslink

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBClient struct {
	Config     *DBConfig
	Connection *gorm.DB
}

type DBConfig struct {
	Host                  string  `json:"host"`
	Port                  *string `json:"port"`
	User                  string  `json:"user"`
	Password              string  `json:"password"`
	DatabaseName          string  `json:"database_name"`
	MaxIdleConnection     *int    `json:"max_idle_connection"`
	MaxOpenConnection     *int    `json:"max_open_connection"`
	MaxConnectionIdleTime *int    `json:"max_connection_idle_time"`
	MaxConnectionLifeTime *int    `json:"max_connection_life_time"`
	Debug                 *struct {
		Enable bool   `json:"enable"`
		Level  string `json:"level"`
	} `json:"debug"`
}

const (
	DefaultPort                  = "5432"
	DefaultMaxIdleConnection     = 5
	DefaultMaxOpenConnection     = 25
	DefaultMaxConnectionIdleTime = 300
	DefaultMaxConnectionLifeTime = 0
)

func NewDBClient(config *DBConfig) *DBClient {
	return &DBClient{
		Config: config,
	}
}

func (c *DBClient) Connect() error {
	port := DefaultPort

	if c.Config.Port != nil {
		port = *c.Config.Port
	}

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		c.Config.Host,
		c.Config.User,
		c.Config.Password,
		c.Config.DatabaseName,
		port,
	)

	config := gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	}

	if c.Config.Debug != nil {
		if c.Config.Debug.Enable {
			logLevel := logger.Silent

			switch c.Config.Debug.Level {
			case "silent":
				logLevel = logger.Silent
			case "error":
				logLevel = logger.Error
			case "warn":
				logLevel = logger.Warn
			case "info":
				logLevel = logger.Info
			}

			config.Logger = logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logLevel,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      false,
				Colorful:                  true,
			})
		}
	}

	connection, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true,
	}), &config)

	if err != nil {
		return err
	}

	db, err := connection.DB()
	if err != nil {
		return err
	}

	maxIdleConn := DefaultMaxIdleConnection
	if c.Config.MaxIdleConnection != nil {
		maxIdleConn = *c.Config.MaxIdleConnection
	}
	db.SetMaxIdleConns(maxIdleConn)

	maxOpenConn := DefaultMaxOpenConnection
	if c.Config.MaxOpenConnection != nil {
		maxOpenConn = *c.Config.MaxOpenConnection
	}
	db.SetMaxOpenConns(maxOpenConn)

	maxConnIdleTime := DefaultMaxConnectionIdleTime
	if c.Config.MaxConnectionIdleTime != nil {
		maxConnIdleTime = *c.Config.MaxConnectionIdleTime
	}
	db.SetConnMaxIdleTime(time.Duration(maxConnIdleTime) * time.Second)

	maxConnLifeTime := DefaultMaxConnectionLifeTime
	if c.Config.MaxConnectionLifeTime != nil {
		maxConnLifeTime = *c.Config.MaxConnectionLifeTime
	}
	db.SetConnMaxLifetime(time.Duration(maxConnLifeTime) * time.Second)

	c.Connection = connection

	return nil
}

func (c *DBClient) Close() error {
	if c.Connection != nil {
		db, err := c.Connection.DB()
		if err != nil {
			return err
		}

		return db.Close()
	}

	return nil
}
