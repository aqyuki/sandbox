package main

import (
	"fmt"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const DatabaseURL = "postgres://postgres:postgres@localhost:5432/database?sslmode=disable"

type User struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

func (u User) String() string {
	return fmt.Sprintf(`User
  ID: %s
  Name: %s
`, u.ID, u.Name)
}

type Note struct {
	ID      string `db:"id"`
	OwnerID string `db:"owner_id"`
	Title   string `db:"title"`
	Body    string `db:"body"`
}

func (n Note) String() string {
	return fmt.Sprintf(`Note
  ID: %s
  OwnerID: %s
  Title: %s
  Body: %s
`, n.ID, n.OwnerID, n.Title, n.Body)
}

func main() {
	user := &User{
		ID:   uuid.New().String(),
		Name: "John Doe",
	}

	note := &Note{
		ID:      uuid.New().String(),
		OwnerID: user.ID,
		Title:   "Hello, World!",
		Body:    "This is a note.",
	}

	note2 := &Note{
		ID:      uuid.New().String(),
		OwnerID: user.ID,
		Title:   "Database testing",
		Body:    "Database testing",
	}

	db, err := sqlx.Connect("pgx", DatabaseURL)
	if err != nil {
		panic(err)
	}

	// Create new user.
	if _, err := db.NamedExec(`insert into users (id, name) values (:id, :name)`, user); err != nil {
		panic(err)
	}

	// Create new note.
	if _, err := db.NamedExec(`insert into notes (id, owner_id, title, body) values (:id, :owner_id, :title, :body)`, note); err != nil {
		panic(err)
	}
	if _, err := db.NamedExec(`insert into notes (id, owner_id, title, body) values (:id, :owner_id, :title, :body)`, note2); err != nil {
		panic(err)
	}

	// fetch notes
	var notes []struct {
		User *User `db:"user"`
		Note *Note `db:"note"`
	}

	// join notes and users
	if err := db.Select(&notes, `select notes.id "note.id", notes.owner_id "note.owner_id", notes.title "note.title", notes.body "note.body", users.id "user.id", users.name "user.name" from notes join users on notes.owner_id = users.id`); err != nil {
		panic(err)
	}

	for _, n := range notes {
		fmt.Println(n.User)
		fmt.Println(n.Note)
	}

	// remove notes
	if _, err := db.Exec(`delete from notes`); err != nil {
		panic(err)
	}

	// remove users
	if _, err := db.Exec(`delete from users`); err != nil {
		panic(err)
	}
}
