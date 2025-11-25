package service

import (
	"belajar_golang_restful_api/model/web"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, id int)
	FindAll(ctx context.Context, page int, size int) ([]web.CategoryResponse, int)
	FindById(ctx context.Context, id int) web.CategoryResponse
}
