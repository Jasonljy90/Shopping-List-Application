package main

import "fmt"

func viewShopList(currentItemList map[string]itemInfo, currentCategories []string) {
	fmt.Println("Shopping List Contents:")
	for itemName, itemInfo := range currentItemList {
		fmt.Printf("Category: %s - Item: %s Quantity: %d Unit Cost: %.2f \n", currentCategories[itemInfo.itemCategory], itemName, itemInfo.itemQuantity, itemInfo.unitCost)
	}
}
