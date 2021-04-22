package main

import "fmt"

func (sl im) currentDataFields() {
	fmt.Println("Print Current Data.")
	for i, ii := range sl {
		fmt.Println(i, "-", ii)
	}
	if len(sl) == 0 {
		fmt.Println("No data found!")
	}
	fmt.Println("======================================================================================================================")
}
