package main

import "fmt"

func main() {
	fmt.Println(Evaporator(10, 10, 5))
}

func Evaporator(content float64, evapPerDay int, threshold int) int {
	stop := content * float64(threshold) / 100 // условие остановки
	var days int                               // счетчик дней
	for cont := content; cont > stop; cont = cont - cont*float64(evapPerDay)/100 {
		days++
	}

	return days

}
