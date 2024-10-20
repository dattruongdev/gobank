package bootstrap

import "gorm.io/gorm"

type Application struct {
	Env *Env
	Db  *gorm.DB
}

func NewApp() *Application {
	env := NewEnv()
	db := NewPostgresqlDatabase(env)

	app := &Application {
		Env: env,
		Db: db,
	}

	return app
}