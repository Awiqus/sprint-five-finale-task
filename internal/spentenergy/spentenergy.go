package spentenergy

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrorIncorrectInputValues = errors.New("spentcalories: incorrect values")
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("%w: steps/weight/height/duration is less than 0; RunningSpentCalories()", ErrorIncorrectInputValues)
	}
	avgSpeed := MeanSpeed(steps, height, duration)
	calories := weight * avgSpeed * duration.Minutes()
	calories /= minInH
	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("%w: steps/weight/height/duration is less than 0; RunningSpentCalories()", ErrorIncorrectInputValues)
	}
	avgSpeed := MeanSpeed(steps, height, duration)
	return (weight * avgSpeed * duration.Minutes()) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}
	stepLength := height * stepLengthCoefficient
	distance := stepLength * float64(steps)
	return distance / mInKm
}
