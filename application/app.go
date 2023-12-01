package application


import (
	"net/http"
	"context"
	"fmt"
)

// define the app structur :
type App struct {
	router http.Handler
}



// define the constructoor of the app :


func New() *App {
	app := &App{router: loadRoutes()}
	return app
}

// define a start method :
func (app *App) Start(ctx context.Context ) error {
	server := &http.Server{
		// define the http server port :
		Addr: ":3000",
		// define the http server handler :
		Handler: app.router,
	}
	// run the server:
	err := server.ListenAndServe()
	// check if there was an error :
	if err != nil {
		return fmt.Errorf("Failed to Runing Server: %w", err)
	}
	return nil
}

