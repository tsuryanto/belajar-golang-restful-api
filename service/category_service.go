package service

import (
	"context"

	"github.com/tsuryanto/belajar-golang-restful-api/model/web"
)

/*
	Biasanya jumlah service mengikuti jumlah API nya
	1 function ke 1 service
*/

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
