package main

import (
	"fmt"
	"project-1-ar/domain"
)

func main() {
	dmytro := domain.Human{
		Id:   1,
		Name: "Dmytro",
		Age:  19,
	}

	var viktoria domain.Human
	viktoria.Id = 2
	viktoria.Name = "Viktoria"
	viktoria.Age = 19
	viktoria.Country = "Ukraine"

	fmt.Println(dmytro, viktoria)
}
