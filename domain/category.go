package domain

import "context"

type Category struct {
	ID   string `json:"id"`
	Name string `json:"category_name"`
}

type CategoryRepository interface {
	Create(c context.Context, category Category) error
	GetByName(c context.Context, catName string) (Category, error)
	GetAll(c context.Context) ([]Category, error)
	Update(c context.Context, updateRequest Category) error
}

