package main

import (
	"fmt"
	"net/http"
	transportHTTP "rsathishtechit/go-rest-api/internal/transport/http"
)

//App - the struct which containst things like pointers to database connections
type App struct{}

//Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our App")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080",handler.Router); err != nil {
		fmt.Println("Failer to setup server")
		return err
	}
	
	return nil
}

func main() {
	fmt.Println("GO REST API services") 
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up out REST API")
		fmt.Println(err)
	}
}