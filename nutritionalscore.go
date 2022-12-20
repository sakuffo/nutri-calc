package main

type ScoreType int

// ScoreType breaks down into types of food etc

const (
	Food ScoreType = iota
	Bevarage
	Water
	Cheese
)

// *Stephen-Guess I am betting that GetNutritionalScore has
// a return type of this struct, NutritionalScore.

type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

var scoreToLetter = []string{"A", "B", "C", "D", "E"}

type EnergyKJ float64

type SugarGram float64

type SaturatedFattyAcids float64

type SodiumMilligram float64

type FruitsPercentage float64

type FiberGram float64

type ProteinGram float64

type NutritionalData struct {
	Energy              EnergyKJ
	Sugars              SugarGram
	SaturatedFattyAcids SaturatedFattyAcids
	Sodium              SodiumMilligram
	Fruits              FruitsPercentage
	Fiber               FiberGram
	Protein             ProteinGram
	IsWater             bool
}

// It looks like we are populating these slices wit magic numbers
// numbers come from a pdf he showed on screen
var energyLevels = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1340, 1005, 670, 335, 0}
var sugarsLevels = []float64{45, 60, 36, 31, 27, 22.5, 18, 13.5, 9, 4.5}
var saturatedFattyAcidsLevels = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var sodiumLevels = []float64{900, 810, 720, 630, 540, 450, 360, 270, 180, 90}
var fiberLevels = []float64{4.7, 3.7, 2.8, 1.9, 0.9}
var proteinLevels = []float64{8, 6.4, 4.8, 3.2, 1.6}

var energyLevelsBevarage = []float64{270, 240, 210, 180, 150, 120, 90, 60, 30, 0}
var sugarsLevelsBevarage = []float64{13.5, 12, 10.5, 9, 7.5, 6, 4.5, 3, 1.5, 0}

func (e EnergyKJ) GetPoints(st ScoreType) int {
	if st == Bevarage {
		return getPointsFromRange(float64(e), energyLevelsBevarage)
	}

	return getPointsFromRange(float64(e), energyLevels)
}

func (s SugarGram) GetPoints(st ScoreType) int {
	if st == Bevarage {
		return getPointsFromRange(float64(s), sugarsLevelsBevarage)
	}
	return getPointsFromRange(float64(s), sugarsLevels)
}

func (sfa SaturatedFattyAcids) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(sfa), saturatedFattyAcidsLevels)
}

func (s SodiumMilligram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(s), sodiumLevels)
}

func (f FruitsPercentage) GetPoints(st ScoreType) int {
	if st == Bevarage {
		if f > 80 {
			return 10
		} else if f > 60 {
			return 4
		} else if f > 40 {
			return 2
		}
		return 0
	}

	if f > 80 {
		return 5
	} else if f > 60 {
		return 2
	} else if f > 40 {
		return 1
	}
	return 0
}

func (f FiberGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(f), fiberLevels)
}

func (p ProteinGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(p), proteinLevels)
}

func EnergyFromKcal(kcal float64) EnergyKJ {
	return EnergyKJ(kcal * 4.184)
}

func SodiumFromSalt(saltMg float64) SodiumMilligram {
	return SodiumMilligram(saltMg * 2.5)
}

func GetNutritionalScore(nd NutritionalData, st ScoreType) NutritionalScore {
	value, positive, negative := 0, 0, 0

	if st != Water {
		fruitPoints := nd.Fruits.GetPoints(st)
		// fiberPoints := nd.Fiber.GetPoints(st)

		negative = nd.Energy.GetPoints(st) + nd.Sugars.GetPoints(st) + nd.SaturatedFattyAcids.GetPoints(st) + nd.Sodium.GetPoints(st)
		positive = fruitPoints + nd.Fiber.GetPoints(st) + nd.Protein.GetPoints(st)

		if st == Cheese {
			value = negative - positive
		} else {
			if negative >= 11 && fruitPoints < 5 {
				value = negative - positive - fruitPoints
			} else {
				value = negative - positive
			}
		}
	}

	return NutritionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: st,
	}
}

func (ns NutritionalScore) GetNutriScore() string {
	if ns.ScoreType == Food {
		return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{18, 10, 2, -1})]
	}
	if ns.ScoreType == Water {
		return scoreToLetter[0]
	}
	return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{9, 5, 1, -2})]
}

// *Stephen-Guess
// I think this works by getting the value the user enters and compares
// it to the slices of magic numbers we setup for the various milestone/brackets

func getPointsFromRange(v float64, steps []float64) int {
	lenSteps := len(steps)

	for i, l := range steps {
		if v > l {
			return lenSteps - i
		}
	}
	return 0
}

// Are we going to need a function that collects and "maps"(adds up) the points
