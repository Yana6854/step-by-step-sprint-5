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

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает строку с данными о тренировке и записывает их в структуру.
// Формат строки: "3456,Бег,3h00m".
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		return errors.New("неверный формат данных")
	}

	// Преобразуем количество шагов в число
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше 0")
	}
	t.Steps = steps

	// Сохраняем тип тренировки
	t.TrainingType = parts[1]

	// Преобразуем строку с длительностью в формат времени
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше 0")
	}
	t.Duration = duration
	return nil
}

// ActionInfo формирует строку с информацией о тренировке.
func (t Training) ActionInfo() (string, error) {
	// Вычисляем дистанцию с помощью функции из пакета spentenergy
	dist := spentenergy.Distance(t.Steps, t.Height)

	// Вычисляем среднюю скорость
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	// В зависимости от типа тренировки считаем калории
	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	// Формирование итоговой строки
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), dist, speed, calories)

	return result, nil
}
