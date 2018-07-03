package models

import "fmt"

type Contexts struct {
	ProductContext []string `json:"productContext"`
	CartContext []string `json:"cartContext"`
	CategoryContext []string `json:"categoryContext"`
}

type ContextItem struct {
	Id int `json:"id"`
	SectionsId int `json:"sectionsId"`
	ContextType string `json:"contextType"`
	Item string `json:"item"`
}

func AddContextTypeItem(contextItem *ContextItem) {
	query := fmt.Sprintf(
		`INSERT INTO contexts (sectionsId, type, item) VALUES ('%v', '%v', '%v')`,
		contextItem.SectionsId, contextItem.ContextType, contextItem.Item)
	insert, err := db.Query(query)
	if err != nil {
		fmt.Printf("insert err %v\n", err.Error())
		panic(err.Error())
	}
	defer insert.Close()
}

func DelContextTypeItem(contextItem *ContextItem) {
	query := fmt.Sprintf(
		`DELETE FROM contexts WHERE sectionsId=%v AND type='%v' And item='%v'`,
		contextItem.SectionsId, contextItem.ContextType, contextItem.Item)
	insert, err := db.Query(query)
	if err != nil {
		fmt.Printf("insert err %v\n", err.Error())
		panic(err.Error())
	}
	defer insert.Close()
}
