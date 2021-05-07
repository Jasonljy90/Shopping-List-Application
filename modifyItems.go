package main

import (
	"fmt"
	"strings"
)

var (
	name        string
	oldName     string
	newName     string  = ""
	category    int     = -1
	quantity    int     = -1
	cost        float64 = -1
	notInItems  bool
	oldQuantity int
	oldCategory int
	oldCost     float64
)

func modifyItems(currentItemsList map[string]itemInfo, currentCategories []string) map[string]itemInfo {
	notInItems = true
	fmt.Println("Modify Items.")
	fmt.Println("Which item do you wish to modify?")
	fmt.Scanf("%s\n", &name)

	for itemName, itemInfo := range currentItemsList {
		oldName = itemName
		oldQuantity = itemInfo.itemQuantity
		oldCost = itemInfo.unitCost
		if strings.EqualFold(name, itemName) {
			fmt.Printf("Current item name is %s - Category is %s - Quantity is - %d Unit Cost - %.2f \n", itemName, currentCategories[itemInfo.itemCategory], itemInfo.itemQuantity, itemInfo.unitCost)
			notInItems = false
			fmt.Println("Enter new name. Enter for no change")
			fmt.Scanf("%s\n", &newName)
			newName = strings.TrimSpace(newName)
			if newName == "" {
				newName = oldName
				fmt.Println("No changes to name made.")
			} else {
				delete(currentItemsList, itemName)
			}

			fmt.Println("Enter new Category. Enter for no change")
			printCategoriesAsNumberedList(currentCategories)
			fmt.Scanf("%d\n", &category)
			if category > (len(currentCategories) - 1) {
				fmt.Println("Invalid Input!")
				break
			} else if category == -1 {
				category = oldCategory
				fmt.Println("No changes to category made.")
			}
			fmt.Println("Enter new Quantity. Enter for no change")
			fmt.Scanf("%d\n", &quantity)
			if quantity == -1 {
				quantity = oldQuantity
				fmt.Println("No changes to quantity made.")
			}
			fmt.Println("Enter new Unit cost. Enter for no change")
			fmt.Scanf("%f\n", &cost)
			if cost == -1 {
				cost = oldCost
				fmt.Println("No changes to Unit cost made.")
			}
			break
		}

	}
	if notInItems {
		fmt.Println("Item does not exists!")
	}

	currentItemsList[newName] = itemInfo{
		itemCategory: category,
		itemQuantity: quantity,
		unitCost:     cost,
	}
	return currentItemsList
}
