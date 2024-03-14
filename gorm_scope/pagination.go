package gorm_scope

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"parking-back/obj"
	"parking-back/utils"
	"time"
)

func Paginate(countFunction func(*gorm.DB) *gorm.DB, pagination obj.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Scopes(countFunction).Count(&totalRows)
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

func FlexWhere(value interface{}, search obj.Search) func(db *gorm.DB) *gorm.DB {
	if search.GetSearchBy() == "" || search.GetSearchText() == "" {
		return func(db *gorm.DB) *gorm.DB { return db }
	}
	sql := fmt.Sprintf(getSqlByType(search.GetType()), search.GetSearchBy())
	st := prepare(search.GetType(), search.GetSearchText())
	fmt.Println(sql, st)
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Model(value).
			Where(sql, st)
	}
}

func getSqlByType(t string) string {
	if t == "" || t == "string" {
		return "%s ILIKE TRIM(?)"
	}
	if t == "time" {
		return "%s = (?)::timestamptz"
	}
	if t == "int" {
		return "%s = (?)::int"
	}
	panic("Unknown type for sql matcher")
}

func prepare(t string, value string) string {
	if t == "" || t == "string" {
		return "%" + value + "%"
	}
	if t == "time" {
		parsedTime, err := time.Parse(time.RFC3339Nano, value)
		if err != nil {
			return value
		}
		return parsedTime.String()
	}
	return value
}

func DefaultCFunction(value interface{}) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Model(value)
	}
}
