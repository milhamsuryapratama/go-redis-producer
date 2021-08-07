package models

type Category struct {
	CategoryName string `json:"category_name"`
}

func (c *Category) SetCategoryName(categoryName string) *Category {
	c.CategoryName = categoryName
	return c
}

func (c *Category) GetCategoryName() string {
	return c.CategoryName
}
