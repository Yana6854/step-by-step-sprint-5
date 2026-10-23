package actioninfo

import (
	"log"
)

// DataParser — это интерфейс, в котором объявлена сигнатуры методов Parse() и ActionInfo().
type DataParser interface {
	// Parse разбирает строку с данными.
	Parse(datastring string) (err error)
	// ActionInfo создает строку с результатом.
	ActionInfo() (string, error)
}

// Info принимает слайс строк с данными и объект, реализующий интерфейс DataParser.
// В цикле парсит каждую строку и выводит результат на экран.
func Info(dataset []string, dp DataParser) {
		// Перебираем все строки в слайсе
    	for _, data := range dataset {
    		// Здесь пытаемся распарсить строку
    		err := dp.Parse(data)
    		if err != nil {
    			// Если ошибка парсинга — логируем и переходим к следующей строке
    			log.Println(err)
    			continue
    		}

    		// Формируем строку с информацией
    		info, err := dp.ActionInfo()
    		if err != nil {
    			// Если ошибка при формировании, то логируем и переходим к следующей
    			log.Println(err)
    			continue
    		}
    		log.Println(info)
    	}
}
