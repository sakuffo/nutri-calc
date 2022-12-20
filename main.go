package main

import "fmt"

func main() {
	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(),
		Sugars:              SugarGrams(),
		SaturatedFattyAcids: SaturatedFattyAcids(),
		Sodium:              SodiumMilligrams(),
		Fruits:              FruitsPercentage(),
		Fiber:               FiberGram(),
		Protein:             ProtienGram(),
	}, Food)

	fmt.Printf("Nutritional Score: %d\n", ns.Value)
}

func GetNutritionalScore(nd NutritionalData, st ScoreType) NutritionalScore {
	value, positive, negative := 0, 0, 0

	if st != Water {

		negative = nd.Energy.GetPoints(st) + nd.Sugars.GetPoints(st) + nd.SaturatedFattyAcids.GetPoints(st) + nd.Sodium.GetPoints(st)
		positive = nd.Fruits.GetPoints(st) + nd.Fiber.GetPoints(st) + nd.Protein.GetPoints(st)
	}

	return NutritionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: st,
	}
}

func EnergyFromKcal(kcal float64) EnergyKJ {
	return EnergyKJ(kcal * 4.184)
}

func SodiumFromSalt(saltMg float64) SodiumMilligram {
	return SodiumMilligram(saltMg * 2.5)
}
