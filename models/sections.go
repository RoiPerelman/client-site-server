package models

import (
	"fmt"
	"log"
)

type Section struct {
	Id int `json:"id"`
	SectionId string `json:"sectionId"`
	Name string `json:"name"`
	Contexts Contexts `json:"contexts"`
}

func AddSection(id int, section Section) {
	// create user in database
	query := fmt.Sprintf(
		`INSERT INTO sections (userId, sectionId)
			VALUES ('%v', '%v')`, id, section.SectionId)
	insert, err := db.Query(query)
	if err != nil {
		fmt.Printf("insert err %v\n", err.Error())
		panic(err.Error())
	}
	defer insert.Close()
}

func DelSection(id int, section Section) {
	// create user in database
	query := fmt.Sprintf(
		`DELETE FROM sections WHERE userId=%v AND sectionId=%v`, id, section.SectionId)
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

func GetAllUserIdSections(userId int) map[string]Section {
	sections := make(map[string]Section, 0)
	query := fmt.Sprintf("Select id, sectionId FROM sections WHERE sections.userId=%v", userId)
	sectionResults, err := db.Query(query)
	if err != nil {
		log.Panic(err)
	}
	defer sectionResults.Close()

	for sectionResults.Next() {
		section := new(Section)
		err := sectionResults.Scan(&section.Id, &section.SectionId)
		if err != nil {
			log.Fatal(err)
		}
		section.Contexts = GetContextsBySectionsIdentifier(section.Id)

		sections[section.SectionId] = *section
	}
	return sections
}

func GetUserIdSectionBySectionId(userId int, sectionId string) Section {
	section := new(Section)
	query := fmt.Sprintf("Select id, sectionId FROM sections WHERE sections.userId=%v AND sections.sectionId='%v'",
		userId, sectionId)
	sectionResult, err := db.Query(query)
	if err != nil {
		log.Panic(err)
	}
	defer sectionResult.Close()
	found := sectionResult.Next()
	if found {
		err := sectionResult.Scan(&section.Id, &section.SectionId)
		if err != nil {
			log.Fatal(err)
		}
		section.Contexts = GetContextsBySectionsIdentifier(section.Id)
	}
	return *section
}