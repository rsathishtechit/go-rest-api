package main

import (
	"fmt"
	"net/http"
	"rsathishtechit/go-rest-api/internal/comment"
	"rsathishtechit/go-rest-api/internal/database"
	transportHTTP "rsathishtechit/go-rest-api/internal/transport/http"
)

//App - the struct which containst things like pointers to database connections
type App struct{}

//Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our App")

	var err error
	db , err := database.NewDatabase()	
	
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return  err
	}
	commentService := comment.NewService(db)
	handler := transportHTTP.NewHandler(commentService)
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