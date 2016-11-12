package main

import (
	"upendo"
	_ "upendo_example/controller"
)

func main() {
	upendo.Start("sample-app")
}
