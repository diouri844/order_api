package application


import (
	"net/http"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

// define the app structur :
type App struct {
	router http.Handler
	// add the redis clinet to the app type :
	rdb *redis.Client
}



// define the constructoor of the app :


func New() *App {
	app := &App{
		// setup app routs :
		router: loadRoutes(),
		// setup app reids client:
		rdb: redis.NewClient(&redis.Options{}),
	}
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
	// ping to the redis running instance :
	err := app.rdb.Ping(ctx).Err()
	// check if the connection with redis is failed:
	if err != nil {
		return fmt.Errorf("Failed to Connect redis: %w", err)
	}
	// run the server:
	fmt.Println(":: Staring Server on Dev mode ....")
	err = server.ListenAndServe()
	// check if there was an error :
	if err != nil {
		return fmt.Errorf("Failed to Runing Server: %w",err)
	}
	return nil
}

