package main

import (
	"fmt"

	"os"
	"strconv"

	"github.com/Fapiko/fccalc/fcrecipes"
)

const fcdata = "/media/steam/SteamLibrary/steamapps/common/FortressCraft/Default/Data"

func main() {
	recipes := fcrecipes.NewRecipes(fcrecipes.LoadAllRecipes(fcdata).CraftData)

	itemName := os.Args[1]

	var numItems int

	if len(os.Args) == 2 {
		numItems = 1
	} else {
		var err error

		numItems, err = strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(itemName)
	recipes.GetRecursiveCosts(0, itemName, numItems)
}
