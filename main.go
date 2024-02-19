package main

import "fmt"

func main() {
	names := []string{
		"Vladislav",
		"Artem",
		"Maksym",
		"Stanislav",
		"Anastasia",
		"Dmytro",
		"Bohdan",
	}

	names = append(names, "Vitalii")
	names = append(names, "Viktoria")

	fmt.Println(names)

	var (
		parni   []int
		neparni []int
	)

	for i := 0; i <= 20; i++ {
		if i%2 != 0 {
			neparni = append(neparni, i)
		}
		if i%2 == 0 {
			parni = append(parni, i)
		}
	}

	fmt.Println(parni)
	fmt.Println(neparni)
}
