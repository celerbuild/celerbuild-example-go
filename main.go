package main

import (
	handler "github.com/celerbuild/celerbuild-example-go/internel/handle"
)

func main() {
	r := handler.SetupRouter()
	r.Run(":8080")
}
