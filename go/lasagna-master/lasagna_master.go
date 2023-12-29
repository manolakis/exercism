package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}

	return len(layers) * time
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}

	return
}

func AddSecretIngredient(friendIngredients, myIngredients []string) {
	myIngredients[len(myIngredients)-1] = friendIngredients[len(friendIngredients)-1]
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	result := make([]float64, len(amounts))

	for i, amount := range amounts {
		result[i] = amount / 2 * float64(portions)
	}

	return result
}
