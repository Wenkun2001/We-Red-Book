package repository

import (
	"context"
	"github.com/Wenkun2001/We-Red-Book/webook/internal/domain"
)

type ArticleReaderRepository interface {
	// Save 有则更新，无则插入，也就是 insert or update 语义
	Save(ctx context.Context, art domain.Article) error
}
