package repository

import (
	"context"
	"github.com/Wenkun2001/We-Red-Book/webook/internal/domain"
)

type ArticleAuthorRepository interface {
	Create(ctx context.Context, art domain.Article) (int64, error)
	Update(ctx context.Context, art domain.Article) error
}
