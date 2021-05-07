package main

import (
	"fmt"
	"strings"
)

var (
	categoryName    string
	newCategoryName string
	oldCategoryName string
)

func modifyCategory(existingCategories []string) []string {
	fmt.Println("Modify Category.")
	fmt.Println("Which category do you wish to modify?")
	fmt.Println(existingCategories)
	fmt.Scanf("%s\n", &categoryName)
	categoryName = strings.TrimSpace(categoryName)
	if categoryName == "" {
		fmt.Println("Invalid Input!")
		notExistingCategory = false
	}
	for i, currentCategoryName := range existingCategories {
		oldCategoryName = currentCategoryName
		if strings.EqualFold(categoryName, currentCategoryName) {
			fmt.Printf("Current category name is %s\n", currentCategoryName)
			notExistingCategory = false
			fmt.Println("Enter new category name. Enter for no change")
			fmt.Scanf("%s\n", &newCategoryName)
			newCategoryName = strings.TrimSpace(newCategoryName)
			if newCategoryName == "" {
				fmt.Println("No changes to category name made.")
				notExistingCategory = false
				newCategoryName = oldCategoryName
			}
			existingCategories[i] = newCategoryName
		}
	}
	if notExistingCategory {
		fmt.Println("Nothing to modify!")
	}
	return existingCategories
}
