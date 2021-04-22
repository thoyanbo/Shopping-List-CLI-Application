package main

import (
	"fmt"
)

func addCategory(c []string) (result []string) {
	var newCat string
	var i int
	fmt.Println("Add New Category Name")
	fmt.Println("What is the New Category Name to add?")
	fmt.Scanln(&newCat)
	for i = 0; i < len(c); i++ {
		if newCat == "" {
			fmt.Println("No Input Found!")
			result = c
			break
		}
		if newCat == c[i] {
			fmt.Println(newCat, "already exist at index", i, "!")
			result = c
			break
		}
	}
	if i == len(c) {
		result = append(c, newCat)
		fmt.Println("New category:", newCat, "added at index", len(result)-1)
	}

	//fmt.Println(result)
	fmt.Println("======================================================================================================================")
	return result
}
