package database

import (
	"context"
	"errors"
	"log"

	"jwt-authentication/models"

	"github.com/jackc/pgconn"
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
	CREATE TABLE IF NOT EXISTS accounts(
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
	_, err := r.db.Exec(ctx, "DROP TABLE accounts")

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "42P01" {
				return ErrTableNotExist
			}
		}
	}
	log.Println(err)
	return err
}

func (r *PostgresRepo) Add(ctx context.Context, user models.User) (*models.User, error) {
	var res models.User
	err := r.db.QueryRow(ctx, "INSERT INTO accounts(name, username, email, password) VALUES($1, $2, $3, $4) RETURNING name, username, email", user.Name, user.Username, user.Email, user.Password).Scan(&res.Name, &res.Username, &res.Email)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}
	return &res, nil
}

func (r *PostgresRepo) Delete(ctx context.Context, username string) error {
	_, err := r.db.Exec(ctx, "DELETE FROM accounts WHERE username=$1", username)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "42P01" {
				return ErrRowNotExists
			}
		}
		log.Println("Delete failed")
		return err
	}
	return nil
}

func (r *PostgresRepo) GetAll(ctx context.Context) ([]models.User, error) {
	res, err := r.db.Query(ctx, "SELECT * FROM accounts")
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
