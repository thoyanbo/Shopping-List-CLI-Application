package main

import "fmt"

var choice int

var totalCost []float64

func (sl im) generateReport() {
	totalCost = make([]float64, len(cat))
	fmt.Println("Generate Report")
	fmt.Println("1. Total Cost of each category")
	fmt.Println("2. List of item by category.")
	fmt.Println("3. Main Menu")
	fmt.Println()
	fmt.Println("Choose your report:")
	fmt.Scanln(&choice)

	if choice == 1 {
		fmt.Println("Total cost by Category.")
		for i := 0; i < len(totalCost); i++ {
			for _, ii := range sl {
				if ii.category == i {
					itemCost := float64(ii.quantity) * ii.unitCost
					totalCost[i] = totalCost[i] + itemCost
				}
			}
			fmt.Println(cat[i], "cost :", totalCost[i])
		}

	} else if choice == 2 {
		fmt.Println("List by Category.")
		for i := 0; i < len(totalCost); i++ {
			for item, ii := range sl {
				if ii.category == i {
					fmt.Println("Category:", cat[sl[item].category], " - Item:", item, "Quantity:", sl[item].quantity, "Unit Cost:", sl[item].unitCost)
				}
			}
		}

	} else if choice == 3 {
		fmt.Println("Returning to main Menu")
	} else {
		fmt.Println("Selection not found, returning back to Main Menu.")
	}
	fmt.Println("======================================================================================================================")
}
