package main

import (
	"rpn/util/config"
)

func main() {

	// Parse STC line

	conf := config.CreateTOMLConfig("config.debug.toml")
	Compile(conf, "app.rpn")

}
