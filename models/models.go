package models

import "time"

type Item struct {
	ID        int64     `json:"id" faker:"uuid_hyphenated`
	Name      string    `json:"name" faker:"name"`
	Image     string    `json:"image"`
	Price     int       `json:"price" faker:"oneof: 15, 27, 61"`
	Quantity  int       `json:"quantity" faker:"oneof: 15, 27, 61"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ShopRefer int64     `json:"shop_refer"`
}

type Shop struct {
	ID          int64     `json:"id" faker:"uuid_hyphenated`
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	Items       []Item    `json:"shop_items" gorm:"foreignKey:ShopRefer;references:ID"`
}

type Image struct {
	ID        int64     `json:"id"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}
