package main

import (
	"fmt"
	"github.com/jessehorne/wrenchat/engine"
)

func main() {
	serv := engine.NewServer("wrenchat official v0.0.1", "the first version of wrenchat", "3000")
	err := serv.Start("3000")
	if err != nil {
		fmt.Println("Server Error:", err)
	}
}
