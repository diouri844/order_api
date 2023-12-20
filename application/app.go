package application


import (
	"net/http"
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

// define the app structur :
type App struct {
	router http.Handler
	// add the redis clinet to the app type :
	rdb *sql.DB
}



// define the constructoor of the app :

func New() *App {
	// try to open a connection to the database :
	db, err := sql.Open("sqlite3", "../OrderDatabase.db")
	// check error :
	if err != nil {
		fmt.Println( err )
		fmt.Println("Faild to open database")
	}
	if err = createTables(db); err != nil {
		fmt.Println(err)
		panic("Failed to create tables")
	}
	// setup order table if not already present :
	app := &App{
		// setup app routs :
		router: loadRoutes(),
		// setup app reids client:
		rdb: db,
	}
	return app
}

// define a start method :
func (app *App) Start(ctx context.Context ) error {
	// test ping :
	server := &http.Server{
		// define the http server port :
		Addr: ":3000",
		// define the http server handler :
		Handler: app.router,
	}
	// run the server:
	fmt.Println(":: Staring Server on Dev mode ....")
	err := server.ListenAndServe()
	// check if there was an error :
	if err != nil {
		return fmt.Errorf("Failed to Runing Server: %w",err)
	}
	return nil
}


// define setup tables for the database :
// createTables creates necessary tables in the SQLite database
func createTables(db *sql.DB) error {
	// Create 'orders' table with UUID as the data type for the 'id' column
	_, err := db.Exec(`
   CREATE TABLE IF NOT EXISTS orders (
      id TEXT PRIMARY KEY,
      customer_id INTEGER NOT NULL,
      product_id INTEGER NOT NULL,
      quantity INTEGER NOT NULL,
      order_date DATETIME NOT NULL,
      )
	`)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
