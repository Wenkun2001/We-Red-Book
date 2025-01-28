package dao

import (
	"context"
	"gorm.io/gorm"
)

type ArticleReaderDAO interface {
	// Upsert INSERT or UPDATE
	Upsert(ctx context.Context, art Article) error
	UpsertV2(ctx context.Context, art PublishedArticle) error
}

type ArticleGormReaderDAO struct {
	db *gorm.DB
}

func (a *ArticleGormReaderDAO) Upsert(ctx context.Context, art Article) error {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleGormReaderDAO) UpsertV2(ctx context.Context, art PublishedArticle) error {
	//TODO implement me
	panic("implement me")
}

func NewArticleGORMReaderDAO(db *gorm.DB) ArticleReaderDAO {
	return &ArticleGormReaderDAO{db: db}
}
