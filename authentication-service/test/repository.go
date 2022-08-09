package test

import (
	"authentication/data"
	"database/sql"
	"time"
)

type PostgresTestRepository struct {
	Conn *sql.DB
}

func NewPostgresTestRepository(db *sql.DB) *PostgresTestRepository {
	return &PostgresTestRepository{
		Conn: db,
	}
}

// GetAll returns a slice of all users, sorted by last name
func (r *PostgresTestRepository) GetAll() ([]*data.User, error) {
	users := []*data.User{}

	return users, nil
}

// GetByEmail returns one user by email
func (r *PostgresTestRepository) GetByEmail(email string) (*data.User, error) {
	user := data.User{
		ID:        1,
		FirstName: "First",
		LastName:  "Last",
		Email:     "me@here.com",
		Password:  "",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}

// GetOne returns one user by id
func (r *PostgresTestRepository) GetOne(id int) (*data.User, error) {
	user := data.User{
		ID:        1,
		FirstName: "First",
		LastName:  "Last",
		Email:     "me@here.com",
		Password:  "",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (r *PostgresTestRepository) Update(user data.User) error {
	return nil
}

// DeleteByID deletes one user from the database, by ID
func (r *PostgresTestRepository) DeleteByID(id int) error {
	return nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (r *PostgresTestRepository) Insert(user data.User) (int, error) {
	return 2, nil
}

// ResetPassword is the method we will use to change a user's password.
func (r *PostgresTestRepository) ResetPassword(password string, user data.User) error {
	return nil
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
// with the hash we have stored for a given user in the database. If the password
// and hash match, we return true; otherwise, we return false.
func (r *PostgresTestRepository) PasswordMatches(plainText string, user data.User) (bool, error) {
	return true, nil
}
