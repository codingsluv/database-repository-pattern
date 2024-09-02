package repository

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	go_database "github.com/codingsluv/go-database"
	"github.com/codingsluv/go-database/entity"
)

func TestCommandInsert(t *testing.T) {
	// Test case
	commentRepository := NewCommentRepositoryImpl(go_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "test@example.com",
		Comment: "Test, Selamat malam",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommandFindById(t *testing.T) {
	// Test case
	commentRepository := NewCommentRepositoryImpl(go_database.GetConnection())

	ctx := context.Background()
	id := int32(43)
	comment, err := commentRepository.FindById(ctx, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestCommandFindAll(t *testing.T) {
	// Test case
	commentRepository := NewCommentRepositoryImpl(go_database.GetConnection())

	ctx := context.Background()
	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}

func TestCommandDelete(t *testing.T) {
	// Test case
	commentRepository := NewCommentRepositoryImpl(go_database.GetConnection())

	ctx := context.Background()
	id := int32(23)
	err := commentRepository.Delete(ctx, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Berhasil hapus data")
}
