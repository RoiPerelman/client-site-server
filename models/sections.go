package models

import (
	"fmt"
)

type Section struct {
	Id int `json:"id"`
	SectionId string `json:"sectionId"`
	Name string `json:"name"`
	Contexts Contexts `json:"contexts"`
}

func AddSection(id int, section int) {
	// create user in database
	query := fmt.Sprintf(
		`INSERT INTO sections (userId, sectionId)
			VALUES ('%v', '%v')`, id, section)
	insert, err := db.Query(query)
	if err != nil {
		fmt.Printf("insert err %v\n", err.Error())
		panic(err.Error())
	}
	defer insert.Close()
}

func DelSection(id int, section int) {
	// create user in database
	query := fmt.Sprintf(
		`DELETE FROM sections WHERE userId=%v AND sectionId=%v`, id, section)
	insert, err := db.Query(query)
	if err != nil {
		fmt.Printf("delete err %v\n", err.Error())
		panic(err.Error())
	}
	defer insert.Close()
}

func UpdateIsMultipleSectionFeature(id int, isMulti bool) {
	// create user in database
	query := fmt.Sprintf(
		`UPDATE users
			SET isMultipleSection=%v
			WHERE id='%v'
		`, isMulti, id)
	insert, err := db.Query(query)
	if err != nil {
		fmt.Printf("update err %v\n", err.Error())
		panic(err.Error())
	}
	defer insert.Close()
}

//func getAllUserIdSections(userId int) []Section {
//	sections := make([]Section, 0)
//	query := fmt.Sprintf("Select id, sectionId FROM sections WHERE sections.userId=%v", userId)
//	sectionResults, err := db.Query(query)
//	if err != nil {
//		log.Panic(err)
//	}
//	defer sectionResults.Close()
//
//	for sectionResults.Next() {
//		section := new(Section)
//		section.Contexts.ProductContext = make([]string, 0)
//		section.Contexts.CartContext = make([]string, 0)
//		section.Contexts.CategoryContext = make([]string, 0)
//		err := sectionResults.Scan(&section.Id, &section.SectionId)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		query := fmt.Sprintf("Select type, item FROM contexts WHERE contexts.sectionsId=%v", section.Id)
//		contextResults, err := db.Query(query)
//		if err != nil {
//			log.Panic(err)
//		}
//		defer contextResults.Close()
//
//		for contextResults.Next() {
//			contextItem := new(ContextItem)
//			err := contextResults.Scan(&contextItem.ContextType, &contextItem.Item)
//			if err != nil {
//				log.Fatal(err)
//			}
//
//			switch contextType := contextItem.ContextType; contextType {
//			case "PRODUCT":
//				section.Contexts.ProductContext = append(section.Contexts.ProductContext, contextItem.Item)
//			case "CART":
//				section.Contexts.CartContext = append(section.Contexts.CartContext, contextItem.Item)
//			case "CATEGORY":
//				section.Contexts.CategoryContext = append(section.Contexts.CategoryContext, contextItem.Item)
//			}
//		}
//
//		sections = append(sections, *section)
//	}
//	return sections
//}