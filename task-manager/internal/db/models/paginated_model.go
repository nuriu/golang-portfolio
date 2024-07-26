package models

import (
	"math"

	"gorm.io/gorm"
)

type PaginatedModel struct {
	PageSize       int
	CurrentPage    int
	TotalRowCount  int64
	TotalPageCount int
	Order          string
	Data           interface{}
}

func Paginate(model interface{}, pagination *PaginatedModel, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRowCount int64
	db.Model(model).Count(&totalRowCount)

	pagination.TotalRowCount = totalRowCount
	pagination.TotalPageCount = int(math.Ceil(float64(totalRowCount) / float64(pagination.PageSize)))

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetPageSize()).Order(pagination.GetOrder())
	}
}

func (p *PaginatedModel) GetOffset() int {
	return (p.GetPage() - 1) * p.GetPageSize()
}

func (p *PaginatedModel) GetPageSize() int {
	if p.PageSize == 0 {
		p.PageSize = 10
	}

	return p.PageSize
}

func (p *PaginatedModel) GetPage() int {
	if p.CurrentPage == 0 {
		p.CurrentPage = 1
	}

	return p.CurrentPage
}

func (p *PaginatedModel) GetOrder() string {
	if p.Order == "" {
		p.Order = "Id desc"
	}

	return p.Order
}
