package service

import "JH_2024_MJJ/internal/model"

func QueryUnhandledReport(UserID int64) (error, []*model.Article) {
	return d.QueryUnhandledReport(ctx, UserID)
}

func HandleReport(article *model.Article) error {
	return d.UpdatePost(ctx, article)
}
func DeleteReport(article *model.Article) error {
	return d.DeleteReport(ctx, article)
}
