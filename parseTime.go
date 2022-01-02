package main

import (
	"fmt"
	"time"
	"os"
	"path/filepath"
)

func main() {
	var myTime string
	if len(os.Args) != 2 {
		fmt.Println("usage: %s string\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	myTime = os.Args[1]

	d, err := time.Parse("15:04", myTime)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Hour(), d.Minute())
	} else {
		fmt.Println(err)
	}
}