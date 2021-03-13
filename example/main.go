package main

import (
	"fmt"
	"time"

	tinydate "github.com/wagslane/go-tinydate"
)

func main() {
	td, err := tinydate.New(2020, 04, 3)
	if err != nil {
		fmt.Println(err.Error())
	}

	td = td.Add(time.Hour * 48)
	fmt.Println(td)
	// prints "2020-04-05"
}
