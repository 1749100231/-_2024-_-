package dao

import (
	"JH_2024_MJJ/internal/model"
	"context"
)

func (d *Dao) CreateArticle(ctx context.Context, article *model.Article) error {
	return d.orm.WithContext(ctx).Create(article).Error
}

func (d *Dao) GetArticleList(ctx context.Context) ([]*model.Article, error) {
	var articleList []*model.Article
	err := d.orm.WithContext(ctx).Find(&articleList).Error
	return articleList, err
}
func (d *Dao) GetArticleByID(ctx context.Context, articleID int64) (*model.Article, error) {
	var article model.Article
	err := d.orm.WithContext(ctx).First(&article, "id = ?", articleID).Error
	return &article, err
}
func (d *Dao) DelArticle(ctx context.Context, articleID int64) error {
	return d.orm.WithContext(ctx).Delete(&model.Article{}, "id = ?", articleID).Error
}

func (d *Dao) UpdatePost(ctx context.Context, article *model.Article) error {
	return d.orm.WithContext(ctx).Save(article).Error
}

func (d *Dao) QueryReport(ctx context.Context, userID int64) (error, []*model.Article) {
	var articleList []*model.Article
	err := d.orm.WithContext(ctx).Find(&articleList, "author = ? AND  status != -1", userID).Error
	return err, articleList
}
