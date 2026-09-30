package repository

import (
	"context"

	"github.com/DevLucasHenrique/mymock-backend/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	dbPool *pgxpool.Pool
}

func NewUserRepository(dbPool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		dbPool: dbPool,
	}
}

func (ur *UserRepository) GetAllUsers(ctx context.Context) ([]entity.UserResponse, error) {
	query := "SELECT name, username, bio, email FROM users;"
	rows, err := ur.dbPool.Query(ctx, query)
	if err!=nil {
		return []entity.UserResponse{}, err
	}

	defer rows.Close()

	var Products []entity.UserResponse
	for rows.Next() {
		var p entity.UserResponse
		if err:=rows.Scan(&p.Name, &p.Username, &p.Bio, &p.Email); err!=nil {
			return []entity.UserResponse{}, err
		}
		Products = append(Products, p)
	}

	return Products, nil
}