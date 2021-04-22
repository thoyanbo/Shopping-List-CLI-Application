package main

import "fmt"

func (sl im) saveShoppingList() (result []im) {
	var n int = 0
	var retChoice int
	tempMap := make(map[string]itemInfo) //creation of  tempMap to copy over current shopping list
	fmt.Println("1. Save existing shopping list")
	fmt.Println("2. Retrieve previous shopping list")
	fmt.Println("3. Return to main menu")
	fmt.Println("Enter your selection.")
	//fmt.Println(len(itemMapSlice))
	fmt.Scanln(&choice)
	if choice == 1 {
		for n = 0; n < len(itemMapSlice); n++ {
			if n == len(itemMapSlice)-1 {
				for key, value := range itemMap {
					tempMap[key] = value
				}
			}
		}
		result = append(itemMapSlice, tempMap)
		fmt.Println("Shopping list has been saved at index", n) //Note: itemMapSlice[0] is just an element of empty map
		//fmt.Println(result)
		//fmt.Println(len(result))
	} else if choice == 2 {
		if len(itemMapSlice) == 1 {
			fmt.Println("There is no saved shopping list!")
		} else {
			fmt.Println("Which previously saved shopping list would you like to retrieve?")
			for n = 1; n < len(itemMapSlice); n++ {
				fmt.Println("Index:", n, "Shopping List:", itemMapSlice[n])
			}
			fmt.Scanln(&retChoice)
			if retChoice < len(itemMapSlice) {
				for key := range itemMap {
					delete(itemMap, key) // we completely delete itemMap to clear existing data
				}
				for key, value := range itemMapSlice[retChoice] {
					itemMap[key] = value // we copy stored element from shopping list index unto itemMap
				}
			} else {
				fmt.Println("Existing index not selected, returning to Main Menu")
			}
		}
		result = itemMapSlice
	} else if choice == 3 {
		fmt.Println("Returning back to Main Menu.")
	} else {
		fmt.Println("Selection not found, returning back to Main Menu.")
	}
	return result

}
