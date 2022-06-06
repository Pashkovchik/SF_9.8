package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func main() {
	people := map[string]Man{
		"Антон":  {"Антон", "Петров", 37, "M", 1},
		"Артем":  {"Артем", "Иванов", 23, "M", 0},
		"Петр":   {"Петр", "Габиев", 21, "M", 2},
		"Иван":   {"Иван", "Жуков", 18, "M", 4},
		"Трофим": {"Трофим", "Трофимов", 30, "M", 3},
		"Анна":   {"Анна", "Маслякова", 41, "W", 3},
		"Алена":  {"Алена", "Сигидаева", 23, "W", 0},
	}
	MaxCrimes := -1
	suspects := make([]string, 0, 0)
	//вычисляем людей с наибольшим количеством преступлений
	MostDangerousPeople := make([]string, 0, 0)
	for name, max := range people {
		suspects = append(suspects, name)
		if max.Crimes > MaxCrimes {
			MaxCrimes = max.Crimes
		}
	}
	if len(people) == 0 {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым")
	} else {
		for name, crime := range people {
			if crime.Crimes == MaxCrimes {
				MostDangerousPeople = append(MostDangerousPeople, name)
			}
		}
	}
	fmt.Println("Весь список подозреваемых:")
	fmt.Println(suspects)
	fmt.Println("Самые опасные преступники:")
	fmt.Println(MostDangerousPeople)
}
