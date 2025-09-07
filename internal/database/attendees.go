package database

import "database/sql"

type AttendeeModel struct {
	DB *sql.DB
}

type Attendee struct {
	Id       int `json:"id"`
	User_Id  int `json:"user_id"`
	Event_Id int `json:"event_id"`
}
