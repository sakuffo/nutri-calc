package main

type ScoreType int

// ScoreType breaks down into types of food etc

const (
	Food ScoreType = iota
	Bevarage
	Water
	Cheese
)

type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

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

}

func (s SugarGram) GetPoints(st ScoreType) int {

}

func (sfa SaturatedFattyAcids) GetPoints(st ScoreType) int {

}

func (s SodiumMilligram) GetPoints(st ScoreType) int {

}

func (f FruitsPercentage) GetPoints(st ScoreType) int {

}

func (f FiberGram) GetPoints(st ScoreType) int {

}

func (p ProteinGram) GetPoints(st ScoreType) int {

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
