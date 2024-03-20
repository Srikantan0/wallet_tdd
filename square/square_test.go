package square

import (
	"fmt"
	"testing"
	//"github.com/stretchr/testify"
)

func TestAreaOfSquare(t *testing.T) {
	var tab = []struct {
		sideLength, areaOfSquare float64
	}{
		{4, 16},
		{5, 25},
	}
	for _, tt := range tab {
		tes := fmt.Sprintf("%f", tt.sideLength)
		t.Run(tes, func(t *testing.T) {
			ans, _ := AreaOfSquare(tt.sideLength)
			if *ans != tt.areaOfSquare {
				t.Errorf("got %f, want %f", *ans, tt.areaOfSquare)
			}
		})
	}
}

func TestIfSquareCanExist(t *testing.T) {
	var tab = []struct {
		sideLength, areaOfSquare float64
	}{
		{6, -1},
		{5, -1},
	}
	for _, tt := range tab {
		tes := fmt.Sprintf("%f", tt.sideLength)
		t.Run(tes, func(t *testing.T) {
			got, _ := CreateNewSquare(tt.sideLength)
			if got == nil {
				t.Errorf("Empty or negative dimensions. Cant create square")
			}
		})
	}
}
