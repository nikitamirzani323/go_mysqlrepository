package main

import (
	"context"
	"fmt"
	"go_mysqlrepository/database"
	"go_mysqlrepository/entity"
	"go_mysqlrepository/repository"
)

func main() {
	commentRepository := repository.NewCommentRepository(database.GetCon())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repo@gmail.com",
		Comment: "comment repo",
	}
	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	commentresult, err := commentRepository.FindById(ctx, 3)

	if err != nil {
		panic(err)
	}

	fmt.Println(commentresult)

	commentall, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, resultco := range commentall {
		fmt.Println(resultco)
	}

}
