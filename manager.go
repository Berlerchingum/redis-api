package main

import (
	"strconv"
	"encoding/json"
	"time"
)

func getMaxId() int {
	client := GetClient()
	defer client.Close()

	keys, _, err := client.Scan(0, "notes:*", 50).Result()
	handleError(err)
	maxId := len(keys)

	return maxId
}

// insert
// insert a new note in the databse
func insert(Note Note) Note {
	client := GetClient()
	defer client.Close()

	currentNoteId := getMaxId() + 1
	Note.Id = currentNoteId
	Note.CreatedAt = time.Now().Format("02-01-2006")
	dataJson, err := json.Marshal(Note)
	handleError(err)
	result := client.Set("notes:"+strconv.Itoa(Note.Id), dataJson, 0)
	handleError(result.Err())

	return Note
}

// Get
// Retrieve a specific note by id
func get(key string) (Note, error) {
	var note Note

	client := GetClient()
	defer client.Close()

	reply, err := client.Get("notes:" + key).Bytes()
	if len(reply) > 0 {
		err = json.Unmarshal(reply, &note)

		return note, nil
	}

	return note, err
}

// getAll
// Retrieve all notes
func getAll() Notes {
	var notes Notes

	client := GetClient()
	defer client.Close()

	notesIterator := client.Scan(0, "notes:*", 50).Iterator()

	for notesIterator.Next() == true {
		var note Note
		key := notesIterator.Val()
		// Get the data at each key
		reply, err := client.Get(key).Bytes()
		handleError(err)
		// Marshal JSON data to Post
		err = json.Unmarshal(reply, &note)
		handleError(err)
		// Append each to posts slice
		notes = append(notes, note)
	}

	return notes
}

// deleteById
func deleteById(idNote string) error {
	client := GetClient()
	defer client.Close()

	intCmd := client.Del("notes:" + idNote)

	return intCmd.Err()
}
