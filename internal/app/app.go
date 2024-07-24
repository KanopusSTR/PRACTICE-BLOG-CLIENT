package app

import (
	"client/internal/api"
	"client/internal/config"
)

type App struct {
	serviceProvider *serviceProvider
	client          api.Client
	url             string
}

func NewApp() (*App, error) {
	app := &App{}

	if err := app.initDeps(); err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) Run() error {
	return app.runHTTPClient()
}

func (app *App) initDeps() error {
	inits := []func() error{
		app.initConfig,
		app.initServiceProvider,
		app.initHTTPClient,
	}

	for _, f := range inits {
		err := f()
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *App) initConfig() error {
	err := config.LoadEnv()
	if err != nil {
		return err
	}
	url, err := config.GetURL()
	if err != nil {
		return err
	}
	app.url = url
	return nil
}

func (app *App) initServiceProvider() error {
	app.serviceProvider = newServiceProvider()
	return nil
}

func (app *App) initHTTPClient() error {
	app.client = api.NewClient(app.serviceProvider.handler(app.url))
	return nil
}

func (app *App) runHTTPClient() error {
	return app.client.Run()
}
