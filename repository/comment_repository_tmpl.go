package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_mysqlrepository/entity"
	"strconv"
)

type commentRepositoryTmpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryTmpl{DB: db}
}

func (repository *commentRepositoryTmpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comment(email, comment) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)

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
func (repository *commentRepositoryTmpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id,email,comment FROM comment WHERE id=? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}
func (repository *commentRepositoryTmpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id,email,comment FROM comment"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
