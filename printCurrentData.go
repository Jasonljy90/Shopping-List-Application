package main

import (
	"fmt"
)

func printCurrentData(currentItemsList map[string]itemInfo) {
	fmt.Println("Print Current Data.")
	if len(currentItemsList) == 0 {
		fmt.Println("No data found!")
	} else {
		for itemName, itemInfo := range currentItemsList {
			fmt.Printf("%s - %v \n", itemName, itemInfo)
		}
	}
}
