package main

import (
	"fmt"
	"os"
	"rpn/lib"
)

func main() {
	line := "STC(CODE(5), CAN_PREVENT, NAME(\"foo\"), FUNCTION(add, add_funcs, TYPE(I64, I64), TYPE(I8, I8)))"

	// Parse STC line
	e, err := lib.ParseStcHeader(line)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	e.Code = 6
	//conf := config.CreateTOMLConfig("config.debug.toml")
	//Compile(conf, "app.rpn")

}
