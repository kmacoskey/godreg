package main

import (
	"fmt"
	_ "github.com/lib/pq"
)

func getNote(noteID int) (Note, error) {
	res := Note{}

	var id int
	var content string

	err := db.QueryRow(`SELECT id, content FROM notes where id = $1`, noteID).Scan(&id, &content)
	if err == nil {
		res = Note{ID: id, Content: content}
	}

	return res, err
}

func allNotes() ([]Note, error) {
	notes := []Note{}

	rows, err := db.Query(`SELECT id, content FROM notes order by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var content string

		err = rows.Scan(&id, &content)
		if err != nil {
			return notes, err
		}

		currentNote := Note{ID: id, Content: content}

		notes = append(notes, currentNote)
	}

	return notes, err
}

func insertNote(content string) (int, error) {
	var noteID int
	err := db.QueryRow(`INSERT INTO notes(content) VALUES($1) RETURNING id`, content).Scan(&noteID)

	if err != nil {
		return 0, err
	}

	fmt.Printf("Last inserted ID: %v\n", noteID)
	return noteID, err
}

func updateNote(id int, content string) (int, error) {
	//Create
	res, err := db.Exec(`UPDATE notes set content=$1 where id=$2 RETURNING id`, content, id)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsUpdated), err
}

func removeNote(noteID int) (int, error) {
	res, err := db.Exec(`delete from notes where id = $1`, noteID)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil
}
