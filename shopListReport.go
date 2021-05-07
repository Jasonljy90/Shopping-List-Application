package main

import "fmt"

var (
	choice int
	sum    float64
)

func shopListReport(currentItemList map[string]itemInfo, currentCategories []string) {
	fmt.Println("Generate Report")
	fmt.Println("1. Total Cost of each category.")
	fmt.Println("2. List of item by category.")
	fmt.Printf("3. Main Menu.\n\n")
	fmt.Println("Choose your report:")
	fmt.Scanf("%d\n", &choice)

	switch choice {
	case 1:
		for i := 0; i < len(currentCategories); i++ {
			for _, itemInfo := range currentItemList {
				if itemInfo.itemCategory == i {
					sum += (float64(itemInfo.itemQuantity) * itemInfo.unitCost)
				}
			}
			fmt.Printf("%s cost: %.2f\n", currentCategories[i], sum)
			sum = 0
		}

	case 2:
		fmt.Println("List of item by category")
		for j := 0; j < len(currentCategories); j++ {
			for itemName, itemInfo := range currentItemList {
				if itemInfo.itemCategory == j {
					fmt.Printf("Category: %s - Item: %s Quantity: %d Unit Cost: %.2f \n", currentCategories[j], itemName, itemInfo.itemQuantity, itemInfo.unitCost)
				}
			}
		}
	case 3:

	default:
		fmt.Println("Invalid Input!")
	}
}
