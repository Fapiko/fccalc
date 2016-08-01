package fcrecipes

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type Recipes map[string][]*CraftCost

func NewRecipes(craftDatums []*CraftData) Recipes {
	recipes := make(Recipes, len(craftDatums))

	for _, craftData := range craftDatums {
		if craftData.MasterRecipe != "" {
			continue
		}

		costs := make([]*CraftCost, 0)

		for _, cost := range craftData.Costs.CraftCost {
			costs = append(costs, cost)
		}
		recipes[craftData.CraftedKey] = costs
	}

	return recipes
}

func (recipes Recipes) GetRecursiveCosts(level int, name string, numItems int) {
	if costs, ok := recipes[name]; ok {
		for _, cost := range costs {
			fmt.Printf("%s%s: %d (%d)\n", strings.Repeat("-", level*2), cost.Key, cost.Amount*numItems, cost.Amount)
			recipes.GetRecursiveCosts(level+1, cost.Key, numItems*cost.Amount)
		}
	}
}

func LoadAllRecipes(basePath string) *ArrayOfCraftData {
	recipes := &ArrayOfCraftData{}

	loadFiles := []string{
		"BlastFurnaceRecipes.xml",
		"CoilerRecipes.xml",
		"ExtruderRecipes.xml",
		"ManufacturerRecipes.xml",
		"PCBAssemblerRecipes.xml",
		"PipeExtruderRecipes.xml",
		"RefineryRecipes.xml",
		"ResearchAssemblerRecipes.xml",
		"SmelterRecipes.xml",
		"StamperRecipes.xml",
	}

	for _, filename := range loadFiles {
		newRecipes, err := LoadRecipesFromFile(basePath, filename)
		if err != nil {
			log.Info(err)
			continue
		}

		recipes.CraftData = append(recipes.CraftData, newRecipes.CraftData...)
	}

	return recipes
}

func LoadRecipesFromFile(basePath string, filename string) (*ArrayOfCraftData, error) {
	recipes := &ArrayOfCraftData{}

	data, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", basePath, filename))
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(data, recipes)
	if err != nil {
		panic(err)
	}

	return recipes, nil
}
