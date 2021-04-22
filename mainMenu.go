package main

import "fmt"

func (sl im) mainMenu() {
	var choice int
	for choice != 10 {
		fmt.Println("Shopping List Application")
		fmt.Println("==========================")
		fmt.Println("1. View entire shopping list.")
		fmt.Println("2. Generate Shopping List Report.")
		fmt.Println("3. Add Items.")
		fmt.Println("4. Modify Items.")
		fmt.Println("5. Delete Item.")
		fmt.Println("6. Print Current Data.")
		fmt.Println("7. Add New Category Name.")
		fmt.Println("8. Modify or delete Category")
		fmt.Println("9. Save shopping list.")
		fmt.Println("10. Terminate program.")
		fmt.Println("Select your choice:")
		fmt.Scanln(&choice)

		if choice == 1 {
			im(itemMap).printShoppingList()

		} else if choice == 2 {
			im(itemMap).generateReport()

		} else if choice == 3 {
			im(itemMap).addItem()

		} else if choice == 4 {
			im(itemMap).modify()

		} else if choice == 5 {
			im(itemMap).deleteItem()

		} else if choice == 6 {
			im(itemMap).currentDataFields()

		} else if choice == 7 {
			cat = addCategory(cat)
		} else if choice == 8 {
			im(itemMap).modifyCategory()
		} else if choice == 9 {
			itemMapSlice = im(itemMap).saveShoppingList()
		} else if choice == 10 {
			break
		} else {
			fmt.Println("Please select from existing Menu.")
		}

	}
}
