package postgres

import (
	"blog/blog/storage"
	"context"
	"fmt"
	"log"
)

const insertCategory = `
	INSERT INTO categories(
		title,
	) VALUES(
		:title,
	)RETURNING id;
`

func (s Storage) CreateCat(ctx context.Context, t storage.Post) (int64, error) {
	stmt, err := s.db.PrepareNamed(insertCategory)
	if err != nil {
		return 0, err
	}
	var id int64
	if err := stmt.Get(&id, t); err != nil {
		return 0, err
	}
	log.Println("Post ID: ", id)
	return id, nil
}

func (s Storage) ListCat(ctx context.Context) ([]storage.Post, error) {
	var list []storage.Post

	err := s.db.Select(&list, "SELECT *from posts")
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s Storage) GetCat(ctx context.Context, id int64)(storage.Post, error) {

	var t storage.Post
	err := s.db.Get(&t,"SELECT * from posts WHERE id=$1",id)
	if err != nil {
		return t, err
	}
	fmt.Println(t)
	return t, nil
}


const updateCat = `

UPDATE posts
	SET
		title = :title,
		description= :description
	WHERE 
	id = :id
	RETURNING *;
`

func (s *Storage) UpdateCat(ctx context.Context, t storage.Post) error{

	stmt, err := s.db.PrepareNamed(updateCat)
	log.Println(stmt)

	if err != nil {
		return  err
	}
	var ut storage.Post
	if err := stmt.Get(&ut,t); err != nil {
		return err
	}
	fmt.Println(ut)
	return err
}

func (s Storage) DeleteCat(ctx context.Context, id int64) error {
	// fmt.Println("done")
		var data storage.Post
		return s.db.Get(&data, "DELETE FROM posts WHERE id=$1 RETURNING *", id)
	
	}
	