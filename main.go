package main

import (
	"fmt"
	"strings"
)

func main() {

	var (
		mainInput   int
		itemInfoMap map[string]itemInfo
		categories  []string
	)

	itemInfoMap = make(map[string]itemInfo)
	categories = initCategories()

	for {
		fmt.Println("Shopping List Application")
		fmt.Println(strings.Repeat("=", 25))
		fmt.Println("1. View Entire Shopping List.")
		fmt.Println("2. Generate Shopping List Report")
		fmt.Println("3. Add Items.")
		fmt.Println("4. Modify Items.")
		fmt.Println("5. Delete Item.")
		fmt.Println("6. Print Current Data.")
		fmt.Println("7. Add New Category Name")
		fmt.Println("8. Modify Category")
		fmt.Println("9. Delete Category")
		fmt.Println("10. Save Shopping List")
		fmt.Println("11. Previous Shopping List")
		fmt.Println("Select your choice:")
		fmt.Scanf("%d\n", &mainInput)

		switch mainInput {
		case 1:
			viewShopList(itemInfoMap, categories)
		case 2:
			shopListReport(itemInfoMap, categories)
		case 3:
			itemInfoMap = addItem(itemInfoMap, categories)
		case 4:
			itemInfoMap = modifyItems(itemInfoMap, categories)
		case 5:
			itemInfoMap = deleteItem(itemInfoMap)
		case 6:
			printCurrentData(itemInfoMap)
		case 7:
			categories = addCategory(categories)
		case 8:
			categories = modifyCategory(categories)
		case 9:
			itemInfoMap, categories = deleteCategory(itemInfoMap, categories)
		case 10:
			saveShoppingList(itemInfoMap)
			saveCategoryList(categories)
		case 11:
			itemInfoMap = retrieveShoppingList()
			categories = retrieveCategoryList()
		default:
			fmt.Println("Invalid Input")
		}
	}
}
