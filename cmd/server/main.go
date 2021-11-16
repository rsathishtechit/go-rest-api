package main

import "fmt"

//App - the struct which containst things like pointers to database connections
type App struct{}

//Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our App")
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