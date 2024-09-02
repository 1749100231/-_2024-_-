package model

import "time"

type Article struct {
	ID      uint   `json:"article_id"`
	Content string `json:"content"`
	Author  int64  `json:"user_id"`

	Status         int       `json:"status"` //-1 没有被举报 0 举报未审批 1 举报通过 2 举报不通过
	ReportedReason string    `json:"reason"`
	ReporterID     int64     `json:"reporter_id"`
	CreatedAt      time.Time `json:"create_time"`
	UpdatedAt      time.Time `json:"Update_time"`
}
