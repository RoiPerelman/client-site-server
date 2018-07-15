package models

import (
	"fmt"
	"log"
)

type Section struct {
	Id        int      `json:"id"`
	SectionId string   `json:"sectionId"`
	Name      string   `json:"name"`
	Contexts  Contexts `json:"contexts"`
}

func (db *DB) GetAllUserIdSections(userId int) map[string]Section {
	sections := make(map[string]Section, 0)
	sectionResults, err := db.Query("Select id, sectionId FROM sections WHERE sections.userId=?", userId)
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
		section.Contexts = db.GetContextsBySectionsId(section.Id)

		sections[section.SectionId] = *section
	}
	return sections
}

func (db *DB) GetUserSectionBySectionsId(sectionsId int) Section {
	section := new(Section)
	sectionResult, err := db.Query("Select id, sectionId FROM sections WHERE sections.id=?", sectionsId)
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
		section.Contexts = db.GetContextsBySectionsId(section.Id)
	}
	return *section
}

func (db *DB) AddSection(userId int, section Section) int {
	// create user in database
	insert, err := db.Exec(
		`INSERT INTO sections (userId, sectionId)
			VALUES (?, ?)`, userId, section.SectionId)
	if err != nil {
		fmt.Printf("insert err %v\n", err.Error())
		panic(err.Error())
	}
	id, err := insert.LastInsertId()
	if err != nil {
		fmt.Printf("insert err %v\n", err.Error())
		panic(err.Error())
	}
	return int(id)
}

func (db *DB) DelSection(id int, section Section) {
	// create user in database
	_, err := db.Exec(
		`DELETE FROM sections WHERE userId=? AND sectionId=?`, id, section.SectionId)
	if err != nil {
		fmt.Printf("delete err %v\n", err.Error())
		panic(err.Error())
	}
}

func (db *DB) UpdateIsMultipleSectionFeature(id int, isMulti bool) {
	// create user in database
	insert, err := db.Query(
		`UPDATE users
			SET isMultipleSection=?
			WHERE id=?
		`, isMulti, id)
	if err != nil {
		fmt.Printf("update err %v\n", err.Error())
		panic(err.Error())
	}
	defer insert.Close()
}
