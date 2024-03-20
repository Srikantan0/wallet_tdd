package square

import "errors"

func CreateNewSquare(a float64) (*float64, error) {
	if a <= 0 {
		return nil, errors.New("square cant have negative or zero dimensions")
	}
	return &a, nil
}

func AreaOfSquare(a float64) (*float64, error) {
	temp, err := CreateNewSquare(a)
	b := *temp
	b = b * b
	return &b, err
}

func PerimetereOfSquare(a float64) (*float64, error) {
	temp, err := CreateNewSquare(a)
	b := *temp
	b = 4 * b
	return &b, err
}
