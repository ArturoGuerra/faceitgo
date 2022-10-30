package main

import (
	"fmt"
	"os"

	"github.com/arturoguerra/faceitgo"
)

func main() {
	cfg := &faceitgo.RESTConfig{
		Token:     os.Getenv("FACEIT_TOKEN"),      // Get your token from https://developers.faceit.com/faceit-api/overview
		PrivToken: os.Getenv("FACEIT_PRIV_TOKEN"), // Get your token from inspecting networking traffic in your browser while logged in to faceit.com
	}

	rest := faceitgo.New(cfg)

	// Get player info
	player, err := rest.GetPlayerByID("1d212b72-3d21-4f3d-95d7-d2953f4bc740-")
	if err != nil {
		panic(err)
	}

	fmt.Println(player)
}
