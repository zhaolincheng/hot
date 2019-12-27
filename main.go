package main

import (
	_ "hot/config"
	"hot/controllers"
)

func main() {
	controllers.Router()
}
