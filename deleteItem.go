package main

import "fmt"

func (sl im) deleteItem() {
	var del string
	fmt.Println("Delete Item.")
	fmt.Println("Enter item name to delete:")
	fmt.Scanln(&del)
	_, ok := itemMap[del]
	if ok == false {
		fmt.Println("Item not found, Nothing to delete!")
	}
	for i := range itemMap {
		if i == del {
			delete(itemMap, del)
		}

	}
	fmt.Println("======================================================================================================================")
}
