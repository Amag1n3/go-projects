package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type Recipe struct {
	ID              int    `json: "id"`
	Title           string `json:"title"`
	Image           string `json:"image"`
	UsedIngredients []struct {
		Name string `json:"name"`
	} `json:"usedIngredients"`
}

func main() {
	var ingredients string
	fmt.Println("Enter your ingredients separated by commas (", "):")
	fmt.Scanln(&ingredients)
	fields := strings.Split(ingredients, ",")

	recipes := getRecipe(fields)
	if len(recipes) > 0 {
		displayRecipes(recipes)
	} else {
		fmt.Println("No recipes found for the given ingredients.")
	}
}

func getRecipe(fields []string) []Recipe {
	apiKey := "aa8ec3b7e7fe44f69efccd78ff2298ca"
	ingredientList := strings.Join(fields, ",")
	url := fmt.Sprintf("https://api.spoonacular.com/recipes/findByIngredients?ingredients=%s&number=5&apiKey=%s", ingredientList, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching recipes: ", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading data: ", err)
	}

	var recipes []Recipe
	err = json.Unmarshal(body, &recipes)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	return recipes
}

func displayRecipes(recipes []Recipe) {
	table := tablewriter.NewWriter(os.Stdout)

	// Define column widths based on your terminal size
	const (
		nameColumnWidth        = 30
		ingredientsColumnWidth = 25
		imageColumnWidth       = 30
		linkColumnWidth        = 50
	)

	table.SetHeader([]string{"S.No", "Name", "Ingredients Used", "Final Product", "Link to Recipe"})

	for i, recipe := range recipes {
		ingredients := make([]string, len(recipe.UsedIngredients))
		for j, ing := range recipe.UsedIngredients {
			ingredients[j] = ing.Name
		}
		recipeURL := fmt.Sprintf("https://api.spoonacular.com/recipes/%d", recipe.ID)
		// Truncate text if it exceeds column width
		truncate := func(text string, width int) string {
			if len(text) > width {
				return text[:width-3] + "..."
			}
			return text
		}

		table.Append([]string{
			fmt.Sprintf("%d", i+1),
			truncate(recipe.Title, nameColumnWidth),
			truncate(strings.Join(ingredients, ", "), ingredientsColumnWidth),
			truncate(recipe.Image, imageColumnWidth),
			truncate(recipeURL, linkColumnWidth),
		})
	}

	table.Render()
}
