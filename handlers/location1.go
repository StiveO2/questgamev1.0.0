package handlers

import (
	"errors"
	"fmt"
)

func location1() {
	fmt.Println("ты находишся внутри первой локции")
	var location1 string
	for {
		fmt.Println("осмотрись вокруг и будь внимателен!")
		fmt.Scanf("%s\n", &location1)
		var waylc1, err = locationWay1(location1)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(waylc1)
	}
}
func locationWay1(waylc1 string) (string, error) {
	switch waylc1 {
	case "осмотреться":
		return "Ты осмотрелся вокруг и увидел руины разрушенного города, ты можешь выбрать направление движения", nil
	case "север":
		return "Ты направляешся на Север", nil
	case "юг":
		return "Ты направляешся на Юг", nil
	case "запад":
		return "Ты направляешся на Запад", nil
	case "восток":
		return "Ты направляешся на Восток ", nil
	}
	return "", errors.New("invalid input")
}
