package main

import (
	"flag"
	"fmt"

	"github.com/aInstaller/utils/steam"
)

var (
	steamMode bool
)

func init() {
	flag.BoolVar(&steamMode, "steam-location", false, "")
}

func main() {
	flag.Parse()

	if steamMode {
		libs, err := steam.FindSteamLibrary()
		if err != nil {
			panic(err)
		}

		fmt.Println(libs[0])
	}
}
