package main

import (
	"fmt"
	"number-guessing-game/app/controllers"
	"number-guessing-game/app/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
