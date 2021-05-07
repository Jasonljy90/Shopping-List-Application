package main

import (
	"fmt"
)

func deleteCategory(currentItemsList map[string]itemInfo, currentCategories []string) (map[string]itemInfo, []string) {

	var (
		categoryToDelete int
		newCategories    []string
	)

	newItemList := make(map[string]itemInfo)

	fmt.Println("Delete Category.")
	fmt.Println("Enter category to delete:")
	printCategoriesAsNumberedList(currentCategories)
	fmt.Scanf("%d\n", &categoryToDelete)
	if categoryToDelete >= len(currentCategories) {
		fmt.Println("Invalid Input!")
	} else {
		newCategories = append(currentCategories[:categoryToDelete], currentCategories[categoryToDelete+1:]...)
		for itemName, currentItem := range currentItemsList {
			if currentItem.itemCategory != categoryToDelete {
				if currentItem.itemCategory > categoryToDelete {
					currentItem.itemCategory = currentItem.itemCategory - 1
				}
				newItemList[itemName] = itemInfo{
					itemCategory: currentItem.itemCategory,
					itemQuantity: currentItem.itemQuantity,
					unitCost:     currentItem.unitCost,
				}
			}
		}
	}
	return newItemList, newCategories
}
