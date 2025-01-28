package domain

type Author struct {
	Id   int64
	Name string
}

type Article struct {
	Id      int64
	Title   string
	Content string
	Author  Author
}
