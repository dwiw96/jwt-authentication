package database

import (
	"context"
	"log"

	"jwt-authentication/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresRepo struct {
	db *pgxpool.Pool
}

func NewPostgresRepo(db *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{
		db: db,
	}
}

func (r *PostgresRepo) Migrate(ctx context.Context) error {
	query := `
	CREATE TABLE IF NOT EXISTS bookworm(
		name TEXT NOT NULL,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL);
		`
	_, err := r.db.Exec(ctx, query)
	log.Println("Migrate Execute")
	return err
}

func (r *PostgresRepo) DeleteTable(ctx context.Context) error {
	_, err := r.db.Exec(ctx, "DROP TABLE bookworm")
	log.Println(err)
	return err
}

func (r *PostgresRepo) Add(ctx context.Context, user models.User) (*models.User, error) {
	var res models.User
	err := r.db.QueryRow(ctx, "INSERT INTO bookworm(name, username, email, password) VALUES($1, $2, $3, $4) RETURNING name, username, email", user.Name, user.Username, user.Email, user.Password).Scan(&res.Name, &res.Username, &res.Email)
	if err != nil {
		log.Println("Add book failed")
		return nil, err
	}
	return &res, nil
}

func (r *PostgresRepo) GetAll(ctx context.Context) ([]models.User, error) {
	res, err := r.db.Query(ctx, "SELECT * FROM bookworm")
	if err != nil {
		log.Println("GetAll failed")
		return nil, err
	}
	defer res.Close()

	var all []models.User
	for res.Next() {
		var user models.User
		if err := res.Scan(&user.Username, &user.Username, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		all = append(all, user)
	}
	return all, nil
}
