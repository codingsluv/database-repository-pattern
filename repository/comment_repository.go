package repository

import (
	"context"

	"github.com/codingsluv/go-database/entity"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
	Delete(ctx context.Context, id int32) error
}
