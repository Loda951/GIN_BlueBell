package models

import "time"

// 内存对齐

type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

type APIPostDetial struct {
	AuthorName       string `json:"author_name"`
	VotesNumber      int64  `json:"votes_number"`
	*Post            `json:"post"`
	*CommunityDetail `json:"community"`
}
