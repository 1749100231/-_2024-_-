package dao

import (
	"JH_2024_MJJ/internal/model"
	"context"
)

func (d *Dao) QueryUnhandledReport(ctx context.Context, userID int64) (error, []*model.Article) {
	var articleList []*model.Article
	err := d.orm.WithContext(ctx).Find(&articleList, "author = ? AND  status = 0", userID).Error
	return err, articleList
}
func (d *Dao) DeleteReport(ctx context.Context, article *model.Article) error {
	return d.orm.WithContext(ctx).Delete(article).Error
}
