package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/codingsluv/go-database/entity"
)

type CommentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepositoryImpl(db *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: db}
}

func (repository *CommentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	scr := "INSERT INTO comments (email, comment) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, scr, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repository *CommentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	scr := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, scr, id)
	coment := entity.Comment{}
	if err != nil {
		return coment, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&coment.Id, &coment.Email, &coment.Comment)
		return coment, nil
	} else {
		// jika tidak ada
		return coment, errors.New("Id" + strconv.Itoa(int(id)) + "tidak di temukan")
	}
}

func (repository *CommentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	scr := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, scr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (repository *CommentRepositoryImpl) Delete(ctx context.Context, id int32) error {
	scr := "DELETE FROM comments WHERE id =?"
	_, err := repository.DB.ExecContext(ctx, scr, id)
	return err
}
