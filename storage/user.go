package postgres

import (
	"database/sql"
	"todo/models"
)

func (s *DB) CreateUser(req models.CreateUserReq) (*models.User, error) {
	var u models.User
	err := s.DB.QueryRow(`
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, name, email, password, created_at
	`, req.Name, req.Email, req.Password).Scan(
		&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *DB) GetUserByEmail(email string) (*models.User, error) {
	var u models.User
	err := s.DB.QueryRow(`
		SELECT id, name, email, password, created_at
		FROM users
		WHERE email = $1
	`, email).Scan(
		&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil // foydalanuvchi topilmadi
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *DB) GetUserByID(id int) (*models.User, error) {
	var u models.User
	err := s.DB.QueryRow(`
		SELECT id, name, email, password, created_at
		FROM users
		WHERE id = $1
	`, id).Scan(
		&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}