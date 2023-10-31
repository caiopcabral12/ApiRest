package main

import (
	db "APIRest/database"
	md "APIRest/models"
	rt "APIRest/routes"
	"fmt"
)

func main() {

	md.President = []md.Presidents{
		{Id: 1, Name: "Caio", Biography: "Maluco eh brabo", Birth: "01/02/2002"},
		{Id: 2, Name: "Bya", Biography: "Menos braba q o Caio", Birth: "07/09/1996"},
	}
	fmt.Println("OI")
	db.DbConnect()

	rt.HandleRequest()
}
