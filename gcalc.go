package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите значение для расчета")
	reader := bufio.NewReader(os.Stdin) //Читаем введенные даные через буфер Для создания потока ввода через буфер применяется функция bufio.NewReader():
	for {                               //запускаем цикл, в котором присваеваем переменной консоль значение строки, пока не встретим символ перевода строки \n
		console, _ := reader.ReadString('\n')
		s := strings.ReplaceAll(console, " ", "")   // присваиваем переменной s копию считанной строки, но без пробелов, ReplaceAll уберет все встречающиейся пробелы
		base(strings.ToUpper(strings.TrimSpace(s))) //передаем функции base, убрав пробелы по краям
		//fmt.Println("Проверим, что получилось после прочтения", s)
	}

} // Запускаем калькулятор

var roman = map[string]int{ //Привязали текстовые значения к цифрам замапить[строку]целому числу
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}
var convIntToRoman = [14]int{ //Создаем массив с типом int потом конвертируем целые числа в римские, чтобы выдвавть ответ в римских
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}
var a, b *int                          //Объявили переменные для двух чисел (операндов)
var operators = map[string]func() int{ //привязываем строковые значения к арифметическим функциям
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"/": func() int { return *a / *b },
	"*": func() int { return *a * *b },
}
var data []string //срез для хранения операндов

const (
	NONE        = "Ошибка. Строка не является математической операцией."
	MoreThenOne = "Ошибка. Только два операнда и один оператор (+, -, /, *)."
	OnlyOne     = "Ошибка. Используются одновременно разные системы счисления."
	NegNum      = "Ошибка. В римской системе нет отрицательных чисел."
	Zero        = "Ошибка. В римской системе нет числа 0."
	RANGE       = "Ошибка. Доступны олько с арабские целыми числа или римские цифры от 1 до 10 включительно"
)

func base(s string) {
	var operator string           //тут храним найденный оператор +-*/
	var stringsFound int          //а тут храним  римские цифры
	numbers := make([]int, 0)     //слайс для хранения арабских
	romans := make([]string, 0)   //для римских
	romansToInt := make([]int, 0) // для конвертированных
	for ops := range operators {  //перебор по маппингу операторов
		for _, val := range s { // перебираем по всем символам в строке и присваиваем значение переменной val
			if ops == string(val) { // если символ в строке равен оператору, конвертируем текущее значение в строку чтобы сравнить
				operator += ops                   // то добавляем его в оператор
				data = strings.Split(s, operator) //Делим строку по оператору и записываем в data
			}
		}
	}
	switch { //проверим сколько операторов в строке, по задаче должен быть только 1
	case len(operator) > 1: //Если больше одного то выдать паник
		panic(MoreThenOne)
	case len(operator) < 1: //или вообще не одного
		panic(NONE)
	}
	for _, elem := range data { //перебор каждого элемента  в data, _чтобы убрать индекс из выдачи
		num, err := strconv.Atoi(elem) //конвертируем  строку в целое число, если будет ошибка то err == ошибке
		if err != nil {                //не получилось, err не равен 0 - значит это римские
			stringsFound++                // добавляем в счетчик римских цифр stringsFound 0,1,2
			romans = append(romans, elem) //и в срез римских
		} else {
			numbers = append(numbers, num) //иначе вписываем в срез арабских
		}
	}

	switch stringsFound { //проверяем кол-во найденных римских

	case 0: //Если римских нет, то проверяем что диапазон до 10 и считаем арабские
		errCheck := numbers[0] > 0 && numbers[0] < 11 && //первый элемент в массиве >0&<11
			numbers[1] > 0 && numbers[1] < 11 // второй так же
		if val, ok := operators[operator]; ok && errCheck == true { //присваиваем к вал оператора найденного и к ОК тру/фелс из условий выше, если ок, то считаем
			a, b = &numbers[0], &numbers[1] //берем а и б из среза намберс 0,1
			fmt.Println(val())              //выводим функцию операции  вал, которая описана в операторс
		} else {
			panic(RANGE) //Если что-то из этого ложно то выдаем ошибку, а т.к. в этом кейсе stringsFound 0, значит дело в диапазоне
		}
	case 1:
		panic(OnlyOne) // stringsFound 1 - значит одна цифра римская, одна арабская
	case 2: //stringsFound две римских
		for _, elem := range romans { //перебор элементов в срезе романс
			if val, ok := roman[elem]; ok && val > 0 && val < 11 { //присваиваем к вал каждый элемиент, но сразу из маппинга получаем целочисленное значение и ,ОК тру/фелс, если найдена в мапе, то тру и  если найденная римская цифра в диапазоне то тру
				romansToInt = append(romansToInt, val) //Добавляем в срез, если условия выполнены кадлый элемент
			} else {
				panic(RANGE) //иначе нарушено правило до 10, т.к. это кейс в котором 2 римские цифры
			}
		}
		if val, ok := operators[operator]; ok { //Присвоили к вал оператораи проверили что оператор есть в мапе
			a, b = &romansToInt[0], &romansToInt[1] //прис указателям значения из найденных римских цифр, переведенных на арабские
			intToRoman(val())                       //выполнили функцию согласно найденному оператору и мапе, передали в intToRoman, в арабском виде
		}
	}
}
func intToRoman(romanResult int) { //получили результат из второго кейса
	var romanNum string   //куда будем записывать римские
	if romanResult == 0 { //Если результат вычислений 0, то ошибка
		panic(Zero)
	} else if romanResult < 0 { //меньше ноля тоже быть не может
		panic(NegNum)
	}
	for romanResult > 0 { //больше ноля
		for _, elem := range convIntToRoman { //перебор элементов массива convIntToRoman
			for i := elem; i <= romanResult; { //пока элемент массива  convIntToRoman меньше либо равен результату вычислений
				for index, value := range roman { //получаем числовой эквивалент для каждого элемента мапинга roman
					if value == elem { //сопоставляем элементы из value (roman) и elem(convIntToRoman)
						romanNum += index   //Добавляем римскую цифру (index) к строке romanNum.
						romanResult -= elem //Вычитаем числовое значение (elem) из romanResult.
					}
				}
			}
		}
	}
	fmt.Println(romanNum)
}
