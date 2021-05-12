package main

import (
	"fmt"

	routes "github.com/crisley/simulator/application/router"
)

func main() {
	route := routes.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringJson, _ := route.ExportJsonPositions()
	fmt.Println(stringJson[0])
}
