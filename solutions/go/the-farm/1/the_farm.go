package thefarm

import (
    "fmt"
    "errors"
)

// TODO: define the 'DivideFood' function
func DivideFood(f FodderCalculator, i int)(float64, error) {
	fodder, err := f.FodderAmount(i)
    if err != nil {
        return fodder, err
    }
    factor, err := f.FatteningFactor()
    if err != nil {
        return factor, err
    }
    
    totalFood := fodder * factor
	foodPerCow := totalFood / float64(i)
    return foodPerCow, err
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(f FodderCalculator, i int)(float64, error) {
    if i > 0 {
        foodPerCow, err := DivideFood(f, i)
        return foodPerCow, err
    } else {
        return 0, errors.New("invalid number of cows")
    }
}

type InvalidCowsError struct {
	numberOfCows int
	message      string
}

// Implement the error interface using fmt.Sprintf
func (e InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numberOfCows, e.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(i int) error {
	var customMessage string

	if i < 0 {
		customMessage = "there are no negative cows"
	} else if i == 0 {
		customMessage = "no cows don't need food"
	} else {
		// Validation successful
		return nil
	}

	// Validation failed, return a new InvalidCowsError
	return &InvalidCowsError{
		numberOfCows: i,
		message:      customMessage,
	}
}

