package main

import (
	"fmt"
	"context"
	"github.com/diouri844/order_api/application"
)




func  main()  {
	// define my app :
	server := application.New()
	// start the server ( app ) with defined start method 
	// and check if there is any error :
	erro := server.Start(context.TODO())
	if erro != nil {
		fmt.Println("Failed to Runing Server: ", erro)
	}
}


