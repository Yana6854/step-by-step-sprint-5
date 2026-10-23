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

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает строку с данными о прогулке и записывает их в структуру.
// Формат строки: "678,0h50m"
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")

	if len(parts) != 2 {
		return errors.New("неверный формат данных")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше 0")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше 0")
	}
	ds.Duration = duration

	return nil
}

// ActionInfo формирует строку с информацией о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	// Вычисляем дистанцию с помощью функции из пакета spentenergy
	dist := spentenergy.Distance(ds.Steps, ds.Height)

	// Считаем сожжённые калории (для прогулки использую WalkingSpentCalories)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	// Формирую итоговую строку
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, dist, calories)

	return result, nil
}
