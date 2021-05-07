package main

import (
	"fmt"
	"strings"
)

type itemInfo struct {
	itemCategory int
	itemQuantity int
	unitCost     float64
}

func addItem(currentItemsList map[string]itemInfo, currentCategories []string) map[string]itemInfo {

	var (
		name     string
		category int
		quantity int
		cost     float64
	)

	fmt.Println("What is the name of your item?")
	fmt.Scanf("%s\n", &name)
	name = strings.TrimSpace(name)
	if name == "" {
		fmt.Println("Invalid Input!")
	} else {
		fmt.Println("What category does it belong to?")
		printCategoriesAsNumberedList(currentCategories)
		fmt.Scanf("%d\n", &category)
		if category >= len(currentCategories) {
			fmt.Println("Invalid Input!")
		} else {
			fmt.Println("How many units are there?")
			fmt.Scanf("%d\n", &quantity)
			fmt.Println("How much does it cost per unit?")
			fmt.Scanf("%f\n", &cost)
			currentItemsList[name] = itemInfo{
				itemCategory: category,
				itemQuantity: quantity,
				unitCost:     cost,
			}
		}
	}
	return currentItemsList
}

func printCategoriesAsNumberedList(currentCategories []string) {
	for i, s := range currentCategories {
		fmt.Printf("%d. %s\n", i, s)
	}
}
