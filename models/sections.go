package models

import "fmt"

type Section struct {
	Id string `json:"id"`
	SectionId string `json:"sectionId"`
	Name string `json:"name"`
	ProductContext []string `json:"productContext"`
	CartContext []string `json:"cartContext"`
	CategoryContext []string `json:"categoryContext"`
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

//func AddContext(sectionsId, contextType string, contextItem string) {
//	query := fmt.Sprintf(
//		`INSERT INTO sectionsContexts (sectionsId, type, item)
//			VALUES ('%v', '%v', '%v')`, sectionsId, contextType, contextItem)
//	insert, err := db.Query(query)
//	if err != nil {
//		fmt.Printf("insert err %v\n", err.Error())
//		panic(err.Error())
//	}
//	defer insert.Close()
//}
//
//func DelContext(sectionsId, contextType string, contextItem string) {
//	query := fmt.Sprintf(
//		`INSERT INTO sectionsContexts (sectionsId, type, item)
//			VALUES ('%v', '%v', '%v')`, sectionsId, contextType, contextItem)
//	insert, err := db.Query(query)
//	if err != nil {
//		fmt.Printf("insert err %v\n", err.Error())
//		panic(err.Error())
//	}
//	defer insert.Close()
//}