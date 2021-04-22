package main

import "fmt"

func checkExistingName(name1 string) (existName bool) {
	for n := range itemMap {
		if name1 == n {
			existName = true
			break
		} else {
			existName = false
		}
	}
	return existName
}

func mustEnterExistingName(name1 string) (existingName string) {
	var exist bool
	for exist != true {
		exist = checkExistingName(name1)
		if exist == true {
			break
		}
		fmt.Println("Name does not exist! Please select from below existing items!")
		for n := range itemMap {
			fmt.Println(n)
		}
		fmt.Scanln(&name1)
	}
	existingName = name1
	return existingName
}

func checkExistingCat(cat1 string) (n int, exist bool) {
	for n = 0; n < len(cat); n++ {
		if cat1 == cat[n] {
			exist = true
			break
		}
	}
	if n == len(cat) {
		exist = false
	}
	return n, exist
}

func mustEnterExistingCat(cat1 string) (n int) {
	var exist bool
	for exist != true {
		n, exist = checkExistingCat(cat1)
		if exist == true {
			break
		}
		fmt.Println("Category does not exist! Please select from", cat)
		fmt.Scanln(&cat1)
	}
	return n
}

func (sl im) modify() {
	var modItem string
	var modName string
	var modCat string
	var modQuan int
	var modUnitCost float64
	var modStruct itemInfo
	var mn string
	var mc int
	var mq int
	var muc float64
	var n int = 0
	var k int
	var exist bool
	fmt.Println("Modify Items.")
	fmt.Println("Which item would you wish to modify?")
	fmt.Scanln(&modItem)

	verifiedItem := mustEnterExistingName(modItem)

	for i := range sl {
		if verifiedItem == i {
			fmt.Println("Current item name is", i, "- Category is", cat[sl[i].category], "- Quantity is", sl[i].quantity, "- Unit Cost", sl[i].unitCost)
		} else {
			n++
		}
	}

	fmt.Println("Enter new Name. Enter for no change.")
	fmt.Scanln(&modName)

	fmt.Println("Enter new Category. Enter for no change.")
	fmt.Scanln(&modCat)

	if modCat != "" {
		_, exist = checkExistingCat(modCat)
		if exist == true {
			for n = 0; n < len(cat); n++ {
				if modCat == cat[n] {
					mc = n
				}
			}
		}
		if exist == false {
			mc = mustEnterExistingCat(modCat)
		}
	}
	fmt.Println("Enter new Quantity. Enter for no change.")
	fmt.Scanln(&modQuan)

	fmt.Println("Enter new Unit cost. Enter for no change.")
	fmt.Scanln(&modUnitCost)

	for k = 0; k < len(cat); k++ {
		if modCat == "" {
			fmt.Println("No changes to category made")
			mc = sl[modItem].category
			break
		}
		if modCat == cat[k] {
			mc = k
			break
		}
	}

	if modQuan == 0 {
		fmt.Println("No changes to quantity made")
		mq = sl[modItem].quantity
	} else {
		mq = modQuan
	}

	if modUnitCost == 0 {
		fmt.Println("No changes to unit cost made")
		muc = sl[modItem].unitCost
	} else {
		muc = modUnitCost
	}

	if modName == "" {
		fmt.Println("No Changes to item name made.")
		mn = modItem //mn is a temp name to store modItem if no modification to name
	} else {
		mn = modName //mn is assigned to modName
	}

	modStruct = itemInfo{
		category: mc,
		quantity: mq,
		unitCost: muc,
	}

	if modName != "" {
		delete(sl, modItem)
		itemMap[mn] = modStruct
	} else {
		itemMap[mn] = modStruct
	}
	//fmt.Println(itemMap) <- this is used to check itemMap
	fmt.Println("======================================================================================================================")
}
