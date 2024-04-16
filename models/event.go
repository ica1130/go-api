package models

import (
	"rest-api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

func (e Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)` // ? is a placeholder for the actual values but they are safe against SQL injection attacks
	stmt, err := db.DB.Prepare(query) // if we use Prepare() we are storing the query in the database so it can be executed multiple times. This is useful if we are going to execute the same query multiple times
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Exec() is used when you want to manipulate data, create data, update data or delete data.
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id //set the ID field of the struct to the ID of the newly inserted row
	return err
}

func GetAllEvents() ([]Event, error) {
	// Query() is used when you want to retrive data from the database.
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	// defer is used to ensure that the rows are closed after the function returns
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
