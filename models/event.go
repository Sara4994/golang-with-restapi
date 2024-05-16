package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Date        time.Time `binding:"required"`
	UserID      int64
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, date, user_id)
	VALUES (?,?,?,?,?)
	`
	stmnt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmnt.Close()
	result, err := stmnt.Exec(e.Name, e.Description, e.Location, e.Date, e.UserID)

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) UpdateEvent() error {

	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, date = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.Date, event.ID)

	return err
}

func (event Event) DeleteEvent() error {
	query := `DELETE FROM events WHERE id = ?`

	stmnt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmnt.Close()

	_, err = stmnt.Exec(event.ID)

	return err
}

func (e Event) Register(userId int64) error {
	query := `INSERT INTO registrations(event_id, user_id) VALUES (?, ?)`

	stmnt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmnt.Close()

	_, err = stmnt.Exec(e.ID, userId)

	return err
}

func (e Event) CancelRegister(userId int64) error {
	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`

	stmnt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmnt.Close()

	_, err = stmnt.Exec(e.ID, userId)

	return err
}
