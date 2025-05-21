package postgres

import "todo/models"

func (s *DB) CreateTodo(req models.CreateToDoReq) (*models.Todo, error) {
	var t models.Todo
	err := s.DB.QueryRow(`
		INSERT INTO todos (user_id, title, description)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, title, description, is_completed, created_at, updated_at
	`, req.UserId, req.Title, req.Description).Scan(
		&t.ID, &t.UserId, &t.Title, &t.Description, &t.IsCompleted, &t.CreatedAt, &t.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *DB) GetTodos(userID int) ([]models.Todo, error) {
	rows, err := s.DB.Query(`
		SELECT id, user_id, title, description, is_completed, created_at, updated_at
		FROM todos
		WHERE user_id = $1
		ORDER BY created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var t models.Todo
		err := rows.Scan(&t.ID, &t.UserId, &t.Title, &t.Description, &t.IsCompleted, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func (s *DB) UpdateTodo(req models.UpdateToDoReq) error {
	_, err := s.DB.Exec(`
		UPDATE todos
		SET title = $1,
		    description = $2,
		    is_completed = $3,
		    updated_at = CURRENT_TIMESTAMP
		WHERE id = $4 AND user_id = $5
	`, req.Title, req.Description, req.IsCompleted, req.ID, req.UserId)
	return err
}

func (s *DB) DeleteTodo(id int, userID int) error {
	_, err := s.DB.Exec(`
		DELETE FROM todos
		WHERE id = $1 AND user_id = $2
	`, id, userID)
	return err
}