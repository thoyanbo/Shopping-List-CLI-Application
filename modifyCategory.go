package main

import (
	"fmt"
)

func (sl im) modifyCategory() {
	var modCat string
	var postModCat string
	var i int = 0
	fmt.Println("Modify or delete category.")
	fmt.Println("1. Modify Category.")
	fmt.Println("2. Delete category.")
	fmt.Println("3. Main Menu")
	fmt.Println()
	fmt.Println("Make your selection:")
	fmt.Scanln(&choice)

	if choice == 1 {
		fmt.Println("Which category would you like to modify?")
		fmt.Scanln(&modCat)
		for i = 0; i < len(cat); i++ {
			if modCat == cat[i] {
				fmt.Println("What is the name of new category?")
				fmt.Scanln(&postModCat)
				cat[i] = postModCat
				break
			}
		}
		if i == len(cat) {
			fmt.Println("Category does not exist")
		}
	} else if choice == 2 {
		fmt.Println("Which category would you like to delete?")
		fmt.Scanln(&modCat)

		for i = 0; i < len(cat); i++ { // iterating through category string to find match with input
			if modCat == cat[i] {
				for k := range itemMap { //iterating through item map to get key value category index
					if sl[k].category == i {
						delete(itemMap, k) // if category index matches, we delete item from itemMap
					} else if sl[k].category > i {
						temp := sl[k]                      //we assign value struct into temp
						temp.category = sl[k].category - 1 //decrease those index higher than user input by 1
						sl[k] = temp                       // we store the temp struct back into original map
					}
				}
				break
			}

		}
		copy(cat[i:], cat[i+1:]) // Shift cat[i+1:] left one index
		cat[len(cat)-1] = ""     // Erase last element (it takes value zero)
		cat = cat[:len(cat)-1]   // Truncate slice
	} else if choice == 3 {
		fmt.Println("Returning back to Main Menu.")
		//fmt.Println(cat)
		//fmt.Println(itemMap)
	} else {
		fmt.Println("Selection not found, returning back to Main Menu.")
	}
}
