package db

import (
	"kredit-plus/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	EngineSQL *xorm.Engine
)

func InitoDatabase() (err error) {
	env := config.New()

	dataSourceName := env.Get("DB_USER") + ":" + env.Get("DB_PASSWORD") + "@tcp(" + env.Get("DB_HOST") + ":" + env.Get("DB_PORT") + ")/" + env.Get("DB_NAME") + "?parseTime=true"
	EngineSQL, err = xorm.NewEngine(env.Get("DB_DRIVER"), dataSourceName)
	if err != nil {
		return err
	}

	EngineSQL.SetConnMaxLifetime(1 * time.Minute)

	err = EngineSQL.Ping()
	if err != nil {
		return err
	}
	return
}
