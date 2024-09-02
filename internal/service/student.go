package service

import (
	"JH_2024_MJJ/internal/model"
)

func PublishPost(Content string, Author int64) error {
	return d.CreateArticle(ctx, &model.Article{
		Content: Content,
		Author:  Author,
		Status:  -1,
	})
}

func GetArticleList() ([]*model.Article, error) {
	return d.GetArticleList(ctx)
}

func GetArticleByID(articleID int64) (*model.Article, error) {
	return d.GetArticleByID(ctx, articleID)
}

func DelArticle(articleID int64) error {
	return d.DelArticle(ctx, articleID)
}

func ReportArticle(articleID int64, reason string, reporterID int64) error {
	article, err := GetArticleByID(articleID)
	if err != nil {
		return err
	}
	article.ReportedReason = reason
	// 文章状态记录为 被举报，未审批
	article.Status = 0
	article.ReporterID = reporterID
	return d.UpdatePost(ctx, article)
}

func UpdateArticle(PostID int64, Content string) error {
	article, err := GetArticleByID(PostID)
	if err != nil {
		return err
	}
	// 更新帖子内容
	article.Content = Content
	return d.UpdatePost(ctx, article)
}

func QueryReport(UserID int64) (error, []*model.Article) {
	return d.QueryReport(ctx, UserID)
}
