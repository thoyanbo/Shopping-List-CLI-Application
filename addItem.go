package main

import "fmt"

func (sl im) addItem() (string, itemInfo) {
	var newStruct itemInfo
	var newItem string
	var newCategory string
	var newCat int
	var newUnitCost float64
	var newQuantity int
	var n int = 0
	var exist bool
	fmt.Println("Add Item")
	fmt.Println("What is the name of your item?")
	fmt.Scanln(&newItem)
	fmt.Println("What category does it belong to?")
	fmt.Scanln(&newCategory)
	_, exist = checkExistingCat(newCategory) // this check will ensure category entered is an existing category
	if exist == true {
		for n = 0; n < len(cat); n++ {
			if newCategory == cat[n] {
				newCat = n
			}
		}
	}
	if exist == false {
		newCat = mustEnterExistingCat(newCategory)
	}

	fmt.Println("How many units are there?")
	fmt.Scanln(&newQuantity)
	fmt.Println("How much does it cost per unit?")
	fmt.Scanln(&newUnitCost)
	for n = 0; n < len(cat); n++ {
		if newCategory == cat[n] {
			newCat = n
			break
		}
	}

	newStruct = itemInfo{
		category: newCat,
		quantity: newQuantity,
		unitCost: newUnitCost,
	}
	itemMap[newItem] = newStruct
	//fmt.Println(newItem, newStruct) <- this is to check on details of item added
	fmt.Println("======================================================================================================================")
	return newItem, newStruct
}
