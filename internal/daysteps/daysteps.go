package daysteps

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
	ErrorNotEnoughValues        = errors.New("daysteps: not enough values")
	ErrorCannotConvertValue     = errors.New("daysteps: cannot convert value")
	ErrorUnexpectedReturnValues = errors.New("daysteps: fetched values are incorrect")
	ErrorIncorrectInputValues   = errors.New("daysteps: incorrect input values")
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	slicedData := strings.Split(datastring, ",")
	if len(slicedData) != 2 {
		return fmt.Errorf("%w: datastring has wrong number of values (!= 3); Parse()", ErrorNotEnoughValues)
	}
	steps, err := strconv.Atoi(slicedData[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return fmt.Errorf("%w: steps value; Parse()", ErrorIncorrectInputValues)
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(slicedData[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return fmt.Errorf("%w: duration value; Parse()", ErrorIncorrectInputValues)
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distnce := spentenergy.Distance(ds.Steps, ds.Height)
	if distnce <= 0 {
		return "", fmt.Errorf("%w: fetched value of distance is equal/lower than zero; ActionInfo()", ErrorUnexpectedReturnValues)
	}
	avgSpeed := spentenergy.MeanSpeed(ds.Steps, ds.Height, ds.Duration)
	if avgSpeed <= 0 {
		return "", fmt.Errorf("%w: fetched value of avgSpeed is equal/lower than zero; ActionInfo()", ErrorUnexpectedReturnValues)
	}

	caloriesSpent, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("%w: incorrect value; ActionInfo", ErrorUnexpectedReturnValues)
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.02f км.\nВы сожгли %.02f ккал.\n", ds.Steps, distnce, caloriesSpent), nil
}
