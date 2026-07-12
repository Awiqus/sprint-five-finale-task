package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

var (
	ErrorNotEnoughValues        = errors.New("trainings: not enough values")
	ErrorCannotConvertValue     = errors.New("trainings: cannot convert value")
	ErrorUnexpectedReturnValues = errors.New("trainings: fetched values are incorrect")
	ErrorIncorrectInputValues   = errors.New("trainings: incorrect input values")
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	slicedData := strings.Split(datastring, ",")
	if len(slicedData) != 3 {
		return fmt.Errorf("%w: datastring has wrong number of values (!= 3); Parse()", ErrorNotEnoughValues)
	}
	steps, err := strconv.Atoi(slicedData[0])
	if err != nil {
		return fmt.Errorf("%w: steps to int from string; Parse()", ErrorCannotConvertValue)
	}
	if steps <= 0 {
		return fmt.Errorf("%w: steps value; Parse()", ErrorIncorrectInputValues)
	}
	t.Steps = steps
	t.TrainingType = slicedData[1]
	duration, err := time.ParseDuration(slicedData[2])
	if err != nil {
		return fmt.Errorf("%w: strinf to time.Duration; Parse()", ErrorCannotConvertValue)
	}
	if duration <= 0 {
		return fmt.Errorf("%w: duration value; Parse()", ErrorIncorrectInputValues)
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distnce := spentenergy.Distance(t.Steps, t.Height)
	if distnce <= 0 {
		return "", fmt.Errorf("%w: fetched value of distance is equal/lower than zero; ActionInfo()", ErrorUnexpectedReturnValues)
	}
	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	if avgSpeed <= 0 {
		return "", fmt.Errorf("%w: fetched value of avgSpeed is equal/lower than zero; ActionInfo()", ErrorUnexpectedReturnValues)
	}

	var caloriesSpent float64
	var err error

	switch t.TrainingType {
	case "Бег":
		caloriesSpent, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("%w: incorrect value; ActionInfo", ErrorUnexpectedReturnValues)
		}
	case "Ходьба":
		caloriesSpent, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("%w: incorrect value; ActionInfo", ErrorUnexpectedReturnValues)
		}
	default:
		return "", fmt.Errorf("%w: неизвестный тип тренировки", ErrorNotEnoughValues)
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.02f ч.\nДистанция: %.02f км.\nСкорость: %.02f км/ч\nСожгли калорий: %.02f\n", t.TrainingType, t.Duration.Hours(), distnce, avgSpeed, caloriesSpent), nil
}
