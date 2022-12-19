package main

import (
	"CSV_COnvert_Tool/route"
)

func main() {
	router := route.GetRouter()
	router.Run("127.0.0.1:8080")
}
