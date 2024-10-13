package main

import (
	"rpn/util/config"
)

func main() {
	conf := config.CreateTOMLConfig("config.debug.toml")
	Compile(conf, "app.rpn")

}
