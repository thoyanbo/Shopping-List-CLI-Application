package main

var (
	cat          []string
	itemMap      map[string]itemInfo
	itemMapSlice []im
)

type (
	itemInfo struct { //item info are stored in a struct
		category int
		quantity int
		unitCost float64
	}
	im map[string]itemInfo
)

func main() {
	cat = []string{"Household", "Food", "Drinks"}
	itemMap = make(map[string]itemInfo)
	itemMapSlice = make([]im, 1) // this is the slice of maps created for store shopping list
	itemMap["Cups"] = itemInfo{category: 0, quantity: 5, unitCost: 3}
	itemMap["Cake"] = itemInfo{category: 1, quantity: 3, unitCost: 1}
	itemMap["Sprite"] = itemInfo{category: 2, quantity: 5, unitCost: 2}
	itemMap["Fork"] = itemInfo{category: 0, quantity: 4, unitCost: 3}
	itemMap["Bread"] = itemInfo{category: 1, quantity: 2, unitCost: 2}
	itemMap["Plates"] = itemInfo{category: 0, quantity: 4, unitCost: 3}
	itemMap["Coke"] = itemInfo{category: 2, quantity: 5, unitCost: 2}

	im(itemMap).mainMenu()

}
