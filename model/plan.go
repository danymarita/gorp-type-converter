package model

import "time"

type Plan struct {
	ID        string    `db:"id, primarykey" json:"id"`
	UserID    int64     `db:"user_id" json:"user_id"`
	Campaigns Campaigns `db:"campaigns" json:"campaigns"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Campaigns struct {
	Data []CampaignItem `json:"data"`
}

type CampaignItem struct {
	CategoryID int     `json:"category_id"`
	NetAmount  float64 `json:"net_amount"`
}
