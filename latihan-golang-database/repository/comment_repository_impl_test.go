package repository

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	latihan_golang_database "latiha-golang-database"
	"latiha-golang-database/entity"
	"testing"
)

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(latihan_golang_database.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(latihan_golang_database.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 111)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(latihan_golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:    "repo@test.com",
		Comments: "comment repo",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println("success inserted data, id", result.Id)

}
