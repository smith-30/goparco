package gormq

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uint      `json:"-" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt" gorm:"index"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"index"`
	Email     string    `json:"email" gorm:"unique_index"`
	Name      string    `json:"name" gorm:"name"`
}

type Repo struct {
	db *gorm.DB
}

type Query struct {
	SelectFields []string
	Conditions   []func(db *gorm.DB) *gorm.DB
	Preloads     []func(db *gorm.DB) *gorm.DB
	Order        func(db *gorm.DB) *gorm.DB
	Limit        func(db *gorm.DB) *gorm.DB
	ForUpdate    func(db *gorm.DB) *gorm.DB
}

func NewQuery(f []string) *Query {
	if len(f) == 0 {
		f = []string{"*"}
	}
	return &Query{
		SelectFields: f,
	}
}

func (a *Query) AddWhere(cond string, v interface{}) {
	a.Conditions = append(a.Conditions, func(db *gorm.DB) *gorm.DB {
		return db.Where(cond, v)
	})
}

func (a *Query) AddPreload(target string) {
	a.Preloads = append(a.Preloads, func(db *gorm.DB) *gorm.DB {
		return db.Preload(target)
	})
}

func (a *Query) AddOr(cond string, v interface{}) {
	a.Conditions = append(a.Conditions, func(db *gorm.DB) *gorm.DB {
		return db.Or(cond, v)
	})
}

func (a *Query) SetOrder(cond string) {
	a.Order = func(db *gorm.DB) *gorm.DB {
		return db.Order(cond)
	}
}

func (a *Query) SetLimit(v int) {
	a.Limit = func(db *gorm.DB) *gorm.DB {
		return db.Limit(v)
	}
}

func (a *Query) EnableForUpdate() {
	a.ForUpdate = func(db *gorm.DB) *gorm.DB {
		return db.Set("gorm:query_option", "FOR UPDATE")
	}
}

func (a *Query) build(db *gorm.DB) *gorm.DB {
	db = db.Select(a.SelectFields)
	for _, item := range a.Conditions {
		db = item(db)
	}
	for _, item := range a.Preloads {
		db = item(db)
	}
	if a.Order != nil {
		db = a.Order(db)
	}
	if a.Limit != nil {
		db = a.Limit(db)
	}
	if a.ForUpdate != nil {
		db = a.ForUpdate(db)
	}

	return db
}

func (a *Repo) GetUser(q *Query) (User, error) {
	var u User
	db := q.build(a.db)
	err := db.Take(&u).Error
	return u, err
}

func (a *Repo) GetUsers(q *Query) ([]User, error) {
	var us []User
	db := q.build(a.db)
	err := db.Find(&us).Error
	return us, err
}
