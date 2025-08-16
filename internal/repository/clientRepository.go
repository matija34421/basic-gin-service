package repository

import (
	"basic-gin/internal/model"
	"database/sql"
	"errors"
	"fmt"
)

type ClientRepository struct {
	db *sql.DB
}

const query = "id,first_name,last_name,email,residence_address, birth_date"

func NewClientRepository(db *sql.DB) *ClientRepository {
	if db == nil {
		panic("db connection is nil")
	}

	return &ClientRepository{
		db: db,
	}
}

func (r *ClientRepository) GetClients() ([]*model.Client, error) {
	var clients []*model.Client

	statement := fmt.Sprintf("select %s from clients", query)
	rows, err := r.db.Query(statement)

	if err != nil {
		return nil, fmt.Errorf("couldnt retrieve clients from database: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var client model.Client

		if err := rows.Scan(&client.ID, &client.FirstName, &client.LastName, &client.Email, &client.ResidenceAddress, &client.BirthDate); err != nil {
			return nil, fmt.Errorf("error scanning the rows: %w", err)
		}

		clients = append(clients, &client)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row error: %w", err)
	}

	return clients, nil
}

func (r *ClientRepository) GetClientById(id int) (*model.Client, error) {
	var client model.Client

	statement := fmt.Sprintf("select %s from clients where id=$1", query)
	row := r.db.QueryRow(statement, id)

	if err := row.Scan(&client.ID, &client.FirstName, &client.LastName, &client.Email, &client.ResidenceAddress, &client.BirthDate); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("couldnt find client with an id(%d)", id)
		}
		return nil, fmt.Errorf("error scanning the row: %w", err)
	}

	return &client, nil
}

func (r *ClientRepository) CreateClient(client *model.Client) (*model.Client, error) {

	err := r.db.QueryRow(`
		INSERT INTO clients (first_name, last_name, email, residence_address, birth_date)
		VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		client.FirstName, client.LastName, client.Email, client.ResidenceAddress, client.BirthDate,
	).Scan(&client.ID)

	if err != nil {
		return nil, fmt.Errorf("couldnt save client: %w", err)
	}

	return client, nil
}

func (r *ClientRepository) UpdateClient(client *model.Client) (*model.Client, error) {
	res, err := r.db.Exec(`
		UPDATE clients 
		SET first_name = $1, last_name = $2, email = $3, residence_address = $4, birth_date = $5 
		WHERE id = $6`,
		client.FirstName, client.LastName, client.Email, client.ResidenceAddress, client.BirthDate, client.ID,
	)

	if err != nil {
		return nil, fmt.Errorf("couldnt update client with id(%d): %w", client.ID, err)
	}

	if _, err := res.RowsAffected(); err != nil {
		return nil, fmt.Errorf("could not fetch affected rows: %w", err)
	}

	return client, nil
}

func (r *ClientRepository) DeleteClient(id int) error {
	res, err := r.db.Exec("delete from clients where id = $1", id)

	if err != nil {
		return fmt.Errorf("error deleting client with id(%d): %w", id, err)
	}

	if _, err := res.RowsAffected(); err != nil {
		return fmt.Errorf("could not fetch affected rows: %w", err)
	}

	return nil
}
