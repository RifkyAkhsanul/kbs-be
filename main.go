package main

import "github.com/RifkyAkhsanul/kbs-be/routes"

func main() {
	e := routes.NewRoutes()

	e.Logger.Fatal(e.Start("8080"))
}
