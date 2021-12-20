package storage

type Post struct {
	ID          int64  `db:"id"`
	Description string `db:"description"`
	Title       string `db:"title"`
	IsCompleted bool   `db:"is_completed"`
}

type Category struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
}
