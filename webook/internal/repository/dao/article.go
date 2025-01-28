package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ArticleDAO interface {
	Insert(ctx context.Context, art Article) (int64, error)
	UpdateById(ctx context.Context, entity Article) error
	Sync(ctx context.Context, entity Article) (int64, error)
}

type ArticleGormDAO struct {
	db *gorm.DB
}

func NewArticleGormDAO(db *gorm.DB) ArticleDAO {
	return &ArticleGormDAO{
		db: db,
	}
}

type Article struct {
	Id      int64  `gorm:"primaryKey,autoIncrement"`
	Title   string `gorm:"type=varchar(4096)"`
	Content string `gorm:"type=BLOB"`
	// 我要根据创作者ID来查询
	AuthorId int64 `gorm:"index"`
	Ctime    int64
	// 更新时间
	Utime int64
}

type PublishedArticle Article

type PublishedArticleV1 struct {
	Article
}

func (a *ArticleGormDAO) Insert(ctx context.Context, art Article) (int64, error) {
	now := time.Now().UnixMilli()
	art.Ctime = now
	art.Utime = now
	err := a.db.WithContext(ctx).Create(&art).Error
	return art.Id, err
}

func (a *ArticleGormDAO) UpdateById(ctx context.Context, art Article) error {
	now := time.Now().UnixMilli()
	res := a.db.WithContext(ctx).Model(&art).
		Where("id = ? AND author_id = ?", art.Id, art.AuthorId).Updates(map[string]any{
		"title":   art.Title,
		"content": art.Content,
		"utime":   now,
	})
	if res.Error != nil {
		return res.Error
	}
	// 怎么知道有没有更新数据
	if res.RowsAffected == 0 {
		// 创作者不对，说明有人在瞎搞
		return errors.New("ID 不对或者创作者不对")
	}
	return nil
}

func (a *ArticleGormDAO) Sync(ctx context.Context, art Article) (int64, error) {
	var id = art.Id
	err := a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var (
			err error
		)
		dao := NewArticleGormDAO(tx)
		if id > 0 {
			err = dao.UpdateById(ctx, art)
		} else {
			id, err = dao.Insert(ctx, art)
		}
		if err != nil {
			return err
		}
		art.Id = id
		now := time.Now().UnixMilli()
		pubArt := PublishedArticle(art)
		pubArt.Ctime = now
		pubArt.Utime = now
		err = tx.Clauses(clause.OnConflict{
			// 对MySQL不起效，但是可以兼容别的方言
			// INSERT xxx ON DUPLICATE KEY SET `title`=?
			// 别的方言：
			// sqlite INSERT XXX ON CONFLICT DO UPDATES WHERE
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"title":   pubArt.Title,
				"content": pubArt.Content,
				"utime":   now,
			}),
		}).Create(&pubArt).Error
		return err
	})
	return id, err
}

func (a *ArticleGormDAO) SyncV1(ctx context.Context, art Article) (int64, error) {
	tx := a.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}
	// 防止后面业务panic
	defer tx.Rollback()

	var (
		id  = art.Id
		err error
	)
	dao := NewArticleGormDAO(tx)
	if id > 0 {
		err = dao.UpdateById(ctx, art)
	} else {
		id, err = dao.Insert(ctx, art)
	}
	if err != nil {
		return 0, err
	}
	art.Id = id
	now := time.Now().UnixMilli()
	pubArt := PublishedArticle(art)
	pubArt.Ctime = now
	pubArt.Utime = now
	err = tx.Clauses(clause.OnConflict{
		// 对MySQL不起效，但是可以兼容别的方言
		// INSERT xxx ON DUPLICATE KEY SET `title`=?
		// 别的方言：
		// sqlite INSERT XXX ON CONFLICT DO UPDATES WHERE
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"title":   pubArt.Title,
			"content": pubArt.Content,
			"utime":   now,
		}),
	}).Create(&pubArt).Error
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return id, nil
}
