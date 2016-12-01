package main

import (
	"strings"

	"fmt"

	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	"github.com/fapiko/fccalc/fcrecipes"
)

const fcdata = "/media/steam/SteamLibrary/steamapps/common/FortressCraft/Default/Data"

func main() {
	recipes := fcrecipes.NewRecipes(fcrecipes.LoadAllRecipes(fcdata).CraftData)

	recipeNames := make([]string, 0)

	for recipeName, _ := range recipes {
		recipeNames = append(recipeNames, recipeName)
	}

	autocompletionScript := fmt.Sprintf(`_fccalc()
{
	local cur opts
	COMPREPLY=()
	cur="${COMP_WORDS[COMP_CWORD]}"
	opts="%s"

	if [ $COMP_CWORD -eq 1 ] ; then
		COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
	fi
}
complete -F _fccalc fccalc`, strings.Join(recipeNames, " "))

	err := ioutil.WriteFile("/etc/bash_completion.d/fccalc", []byte(autocompletionScript), 0644)
	if err != nil {
		log.Error(err)
	}
}
