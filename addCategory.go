package main

import (
	"fmt"
	"strings"
)

var (
	addCate             string
	notExistingCategory bool = true
)

func initCategories() []string {
	return []string{"Household", "Food", "Drinks"}
}

func addCategory(existingCategories []string) []string {
	fmt.Println("Add New Category name")
	fmt.Println("What is new category name to add?")
	fmt.Scanf("%s\n", &addCate)

	for i, v := range existingCategories {
		if strings.EqualFold(addCate, v) {
			notExistingCategory = false
			fmt.Println(addCate, "already exists at index", i, "!")
			break
		} else if addCate == "" {
			notExistingCategory = false
			fmt.Println("No Input Found!")
			break
		}
	}
	if notExistingCategory {
		existingCategories = append(existingCategories, addCate)
		fmt.Println("New Category: ", addCate, "added at index ", (len(existingCategories) - 1))
	}
	return existingCategories
}
