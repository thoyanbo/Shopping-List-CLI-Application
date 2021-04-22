package main

import "fmt"

func (sl im) printShoppingList() {
	fmt.Println("Shopping List Contents:")
	for item := range sl {
		fmt.Println("Category:", cat[sl[item].category], " - Item:", item, "Quantity:", sl[item].quantity, "Unit Cost:", sl[item].unitCost)
	}
	fmt.Println("======================================================================================================================")
}
