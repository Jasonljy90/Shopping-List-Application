package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func saveShoppingList(currentItemsList map[string]itemInfo) {
	writeItemSetJSON("shoppingList.json", currentItemsList)
	fmt.Printf("Shopping List Saved\n")
}

func retrieveShoppingList() map[string]itemInfo {
	readFile := readItemSetJSON("shoppingList.json")
	fmt.Printf("Shopping List Retrieved\n")
	return readFile
}

func saveCategoryList(currentCategories []string) {
	writeCategorySetJSON("shoppingListCategories.json", currentCategories)
}

func retrieveCategoryList() []string {
	readCatagories := readCategorySetJSON("shoppingListCategories.json")
	return readCatagories
}

func writeItemSetJSON(fileName string, itemWrite map[string]itemInfo) {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	file, _ := os.OpenFile(fileName, os.O_CREATE, os.ModePerm)
	defer file.Close()

	write(file, "{")
	var comma = ""
	for name, writeItems := range itemWrite {
		write(file, comma)
		write(file, "\"")
		write(file, name)
		write(file, "\"")
		write(file, ":")
		writeItemFile(file, writeItems)
		comma = ","
	}
	write(file, "}")
}

func writeCategorySetJSON(fileName string, categoryWrite []string) {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println(err)
	}

	file, _ := os.OpenFile(fileName, os.O_CREATE, os.ModePerm)
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(categoryWrite)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func write(file *os.File, s string) {
	_, err := file.WriteString(s)
	if err != nil {
		log.Println(err)
	}
}

func writeItemFile(file *os.File, writeItems itemInfo) {
	write(file, "{")
	write(file, "\"Category\":\"")
	var categoryAsString = strconv.Itoa(writeItems.itemCategory)
	write(file, categoryAsString)
	write(file, "\", \"Quantity\":\"")
	var quantityAsString = strconv.Itoa(writeItems.itemQuantity)
	write(file, quantityAsString)
	write(file, "\", \"Unit Cost\":\"")
	var costAsString = strconv.Itoa(int(writeItems.unitCost))
	write(file, costAsString)
	write(file, "\"}")
}

func readItemSetJSON(fileName string) map[string]itemInfo {
	var bytes, err = ioutil.ReadFile(fileName)
	var itemListFromFile = make(map[string]map[string]string)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(bytes, &itemListFromFile)
	if err != nil {
		log.Println(err)
	}
	var itemFromFile = convert(itemListFromFile)
	return itemFromFile
}

func readCategorySetJSON(fileName string) []string {
	var categories []string
	file, _ := os.Open(fileName)
	defer file.Close()

	decoder := json.NewDecoder(file)
	if decoder.More() {
		if err := decoder.Decode(&categories); err != nil {
			fmt.Println(err)
		}
	}
	return categories

}

func convert(input map[string]map[string]string) map[string]itemInfo {
	var itemFromFile = make(map[string]itemInfo)
	for itemName, itemObj := range input {
		category, err := strconv.Atoi(itemObj["Category"])
		if err != nil {
			log.Println(err)
		}
		quantity, err1 := strconv.Atoi(itemObj["Quantity"])
		if err1 != nil {
			log.Println(err1)
		}
		cost, err2 := strconv.Atoi(itemObj["Unit Cost"])
		if err2 != nil {
			log.Println(err2)
		}
		itemFromFile[itemName] = itemInfo{
			itemCategory: category,
			itemQuantity: quantity,
			unitCost:     float64(cost),
		}
	}
	return itemFromFile
}
