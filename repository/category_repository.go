package repository

import (
	"belajar_golang_restful_api/model/domain"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindAll(ctx context.Context, tx *sql.Tx, limit int, offset int) []domain.Category
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error)
	Count(ctx context.Context, tx *sql.Tx) int
}
