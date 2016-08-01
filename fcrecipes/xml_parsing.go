package fcrecipes

import (
	"errors"
	"strings"
)

type ArrayOfCraftData struct {
	CraftData []*CraftData
}

type CraftData struct {
	Key           string
	Category      string
	Tier          int
	CraftedKey    string
	CraftedAmount int
	Costs         *Costs
	Description   string
	Hint          string
	MasterRecipe  string
	// ScanRequirements
	// ResearchRequirements
}

type Costs struct {
	CraftCost []*CraftCost
}

type CraftCost struct {
	Key    string
	Amount int
}

func (recipes *ArrayOfCraftData) GetRecipe(key string) (*CraftData, error) {
	for _, recipe := range recipes.CraftData {
		if recipe.Key == strings.ToLower(key) {
			return recipe, nil
		}
	}

	return nil, errors.New("No key found for item: " + key)
}
