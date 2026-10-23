package spentenergy

// Я в этом пакете не стала расписывать подробные комментарии,
// т.к. код с такой же логикой как и в 4 спринте.
// Ниже указано краткое описание каждой функции.
import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// WalkingSpentCalories рассчитывает количество калорий, потраченных при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные входные параметры")
	}
	speed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * speed * durationInMinutes) / minInH

	calories *= walkingCaloriesCoefficient

	return calories, nil
}

// RunningSpentCalories рассчитывает количество калорий, потраченных при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные входные параметры")
	}
	speed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * speed * durationInMinutes) / minInH
	return calories, nil
}

// MeanSpeed рассчитывает среднюю скорость, принимает шаги, длительснть и рост.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	return dist / duration.Hours()
}

// Distance рассчитывает дистанцию в километрах, принимает шаги и рост пользователя.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return (float64(steps) * stepLength) / mInKm
}
