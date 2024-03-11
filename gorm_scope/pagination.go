package gorm_scope

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"parking-back/obj"
	"parking-back/utils"
)

func Paginate(value interface{}, pagination obj.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)
	pagination.SetTotalRows(totalRows)
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.SetTotalPage(totalPages)
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Offset(utils.GetOffset(pagination.GetPage(), pagination.GetLimit())).
			Limit(pagination.GetLimit()).
			Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: true})
	}
}
