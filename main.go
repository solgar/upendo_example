package main

import (
	"github.com/solgar/upendo"
	_ "upendo_example/controller"
)

func main() {
	upendo.Start("sample-app")
}
