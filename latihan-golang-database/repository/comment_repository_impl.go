package repository

import (
	"context"
	"database/sql"
	"errors"
	"latiha-golang-database/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}
func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	//TODO implement me
	var commentId int32 = 0
	script := `INSERT INTO comments(email, comments) VALUES ($1, $2) RETURNING id`
	err := repository.DB.QueryRowContext(ctx, script, comment.Email, comment.Comments).Scan(&commentId)
	if err != nil {
		return comment, err
	}
	comment.Id = commentId
	return comment, err
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	//TODO implement me
	script := "SELECT id, email, comments.comments FROM comments WHERE id =  $1 LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comments)
		defer rows.Close()
		return comment, nil
	} else {
		return comment, errors.New("Id : " + strconv.Itoa(int(id)) + " is not found")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	//TODO implement me
	script := "SELECT id, email, comments.comments FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comments)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
		//defer rows.Close()
	}
	return comments, nil
}
