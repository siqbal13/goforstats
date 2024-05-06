package main

import (
	"os"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

func TestLoadAnscombe(t *testing.T) {
	anscombe := LoadAnscombe()
	if len(anscombe.X1) != 11 || len(anscombe.Y1) != 11 {
		t.Error("LoadAnscombe: Incorrect length of data")
	}
}

func TestLinreg(t *testing.T) {
	x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}
	m, c := linreg(x, y)
	if m != 0.5 || c != 3.0 {
		t.Error("linreg: Incorrect regression parameters")
	}
}

func TestRunAnalysis(t *testing.T) {
	sets := map[string][]float64{
		"Set I":   {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		"Set II":  {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		"Set III": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		"Set IV":  {8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
	}
	RunAnalysis(sets)

	// Check if plots are generated
	for name := range sets {
		if _, err := os.Stat("set_" + name + ".png"); os.IsNotExist(err) {
			t.Errorf("RunAnalysis: Plot file for %s does not exist", name)
		}
	}
}

func TestMain(m *testing.M) {
	// Call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}
