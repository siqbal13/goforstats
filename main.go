package main

import (
	"fmt"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Anscombe represents the Anscombe data set
type Anscombe struct {
	X1, X2, X3, X4 []float64
	Y1, Y2, Y3, Y4 []float64
}

// LoadAnscombe loads the Anscombe data set
func LoadAnscombe() Anscombe {
	return Anscombe{
		X1: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		X2: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		X3: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		X4: []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
		Y1: []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
		Y2: []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},
		Y3: []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
		Y4: []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},
	}
}

// RunAnalysis performs linear regression analysis on the Anscombe data set
func RunAnalysis(sets map[string][]float64) {
	for name, x := range sets {
		fmt.Println("Set:", name)
		y := sets["Y"+name[1:]]

		// Perform linear regression
		m, c := linreg(x, y)
		fmt.Printf("Linear Regression: y = %.2f * x + %.2f\n", m, c)

		// Create scatter plot
		p, err := plot.New()
		if err != nil {
			fmt.Println("Error creating plot:", err)
			return
		}
		p.Title.Text = "Set " + name
		p.X.Label.Text = "X" + name[1:]
		p.Y.Label.Text = "Y" + name[1:]
		s, err := plotter.NewScatter(plotter.XYs{})
		if err != nil {
			fmt.Println("Error creating scatter plot:", err)
			return
		}
		for i := range x {
			s.XYs = append(s.XYs, struct{ X, Y float64 }{x[i], y[i]})
		}
		p.Add(s)
		if err := p.Save(5*vg.Inch, 5*vg.Inch, "set_"+name+".png"); err != nil {
			fmt.Println("Error saving plot:", err)
		}
	}
}

// linreg performs linear regression on the given data
func linreg(x, y []float64) (float64, float64) {
	n := float64(len(x))

	// Calculate mean of x and y
	xMean, yMean := 0.0, 0.0
	for i := range x {
		xMean += x[i]
		yMean += y[i]
	}
	xMean /= n
	yMean /= n

	// Calculate slope and intercept
	numerator, denominator := 0.0, 0.0
	for i := range x {
		numerator += (x[i] - xMean) * (y[i] - yMean)
		denominator += (x[i] - xMean) * (x[i] - xMean)
	}
	m := numerator / denominator
	c := yMean - m*xMean
	return m, c
}

func main() {
	// Load Anscombe data set
	anscombe := LoadAnscombe()

	// Perform linear regression analysis
	sets := map[string][]float64{
		"Set I":   anscombe.X1,
		"Set II":  anscombe.X2,
		"Set III": anscombe.X3,
		"Set IV":  anscombe.X4,
	}
	RunAnalysis(sets)
}
