package models

import (
	"fmt"
	"log"
)

type Contexts struct {
	ProductContext []string `json:"product"`
	CartContext []string `json:"cart"`
	CategoryContext []string `json:"category"`
}

type ContextItem struct {
	Id int `json:"id"`
	SectionsId int `json:"sectionsId"`
	SectionId string `json:"sectionId"`
	ContextType string `json:"contextType"`
	Item string `json:"item"`
}

func AddContextTypeItem(contextItem *ContextItem) {
	query := fmt.Sprintf(
		`INSERT INTO contexts (sectionsId, type, item) VALUES ('%v', '%v', '%v')`,
		contextItem.SectionsId, contextItem.ContextType, contextItem.Item)
	_, err := db.Exec(query)
	if err != nil {
		fmt.Printf("insert err %v\n", err.Error())
		panic(err.Error())
	}
}

func DelContextTypeItem(contextItem *ContextItem) {
	query := fmt.Sprintf(
		`DELETE FROM contexts WHERE sectionsId=%v AND type='%v' And item='%v'`,
		contextItem.SectionsId, contextItem.ContextType, contextItem.Item)
	_, err := db.Exec(query)
	if err != nil {
		fmt.Printf("insert err %v\n", err.Error())
		panic(err.Error())
	}
}

func GetContextsBySectionsId(sectionsIdentifier int) Contexts {
	contexts := new(Contexts)
	contexts.ProductContext = make([]string, 0)
	contexts.CartContext = make([]string, 0)
	contexts.CategoryContext = make([]string, 0)

	query := fmt.Sprintf("Select type, item FROM contexts WHERE contexts.sectionsId=%v", sectionsIdentifier)
	contextResults, err := db.Query(query)
	if err != nil {
		log.Panic(err)
	}
	defer contextResults.Close()

	for contextResults.Next() {
		contextItem := new(ContextItem)
		err := contextResults.Scan(&contextItem.ContextType, &contextItem.Item)
		if err != nil {
			log.Fatal(err)
		}

		switch contextType := contextItem.ContextType; contextType {
		case "PRODUCT":
			contexts.ProductContext = append(contexts.ProductContext, contextItem.Item)
		case "CART":
			contexts.CartContext = append(contexts.CartContext, contextItem.Item)
		case "CATEGORY":
			contexts.CategoryContext = append(contexts.CategoryContext, contextItem.Item)
		}
	}
	return *contexts
}