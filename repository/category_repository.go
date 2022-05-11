package repository

import (
	"context"
	"database/sql"

	"github.com/tsuryanto/belajar-golang-restful-api/model/domain"
)

/*  Biasakan Membuat kontrak interface
Isinya function yang bisa mensupport API spec yang telah dibuat
CREATE, GET, PUT, DELETE, DET BY ID, DLL
*/

type CategoryRepository interface {
	// Biasakan semua param selalu pakai context
	// Ke 2 : SQl Transactional pointer
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
