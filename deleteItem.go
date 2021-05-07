package main

import (
	"fmt"
	"strings"
)

func deleteItem(currentItemsList map[string]itemInfo) map[string]itemInfo {

	var (
		name       string
		notInItems bool = true
	)

	fmt.Println("Delete Item.")

	fmt.Println("Enter item name to delete:")
	fmt.Scanf("%s\n", &name)

	for itemName := range currentItemsList {
		if strings.EqualFold(name, itemName) {
			delete(currentItemsList, itemName)
			fmt.Printf("Deleted %s\n", itemName)
			notInItems = false
		}
	}
	if notInItems {
		fmt.Println("Item not found. Nothing to delete!")
	}
	return currentItemsList
}
