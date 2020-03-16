package gormq

import (
	"database/sql/driver"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func emptyConn() *gorm.DB {
	db, _, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	gdb, _ := gorm.Open("mysql", db)
	// gdb.LogMode(true)
	return gdb
}

func TestRepo_GetUser(t *testing.T) {
	type fields struct {
		db func() *gorm.DB
	}
	type args struct {
		q func() *Query
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    User
		wantErr bool
	}{
		{
			fields: fields{
				db: func() *gorm.DB {
					db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
					mock.ExpectQuery("SELECT * FROM `users` LIMIT 1").WillReturnRows(sqlmock.NewRows([]string{}))
					gdb, _ := gorm.Open("mysql", db)
					return gdb
				},
			},
			args: args{
				q: func() *Query {
					q := NewQuery([]string{})
					return q
				},
			},
		},
		{
			fields: fields{
				db: func() *gorm.DB {
					db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
					mock.ExpectQuery("SELECT * FROM `users`  WHERE (name = ?) AND (email = ?) LIMIT 1").WithArgs(driver.Value("test"), driver.Value("test@gmail.com")).WillReturnRows(sqlmock.NewRows([]string{}))
					gdb, _ := gorm.Open("mysql", db)
					return gdb
				},
			},
			args: args{
				q: func() *Query {
					q := NewQuery([]string{})
					q.AddWhere("name = ?", "test")
					q.AddWhere("email = ?", "test@gmail.com")
					return q
				},
			},
		},
		{
			fields: fields{
				db: func() *gorm.DB {
					db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
					mock.ExpectQuery("SELECT * FROM `users`  WHERE (name IN (?,?)) LIMIT 1").WithArgs(driver.Value("1"), driver.Value("2")).WillReturnRows(sqlmock.NewRows([]string{}))
					gdb, _ := gorm.Open("mysql", db)
					return gdb
				},
			},
			args: args{
				q: func() *Query {
					q := NewQuery([]string{})
					q.AddWhere("name IN (?)", []string{"1", "2"})
					return q
				},
			},
		},
		{
			fields: fields{
				db: func() *gorm.DB {
					db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
					mock.ExpectQuery("SELECT id, name FROM `users`  WHERE (name IN (?)) OR (email = ?) LIMIT 1").WithArgs(driver.Value("1"), driver.Value("test@com")).WillReturnRows(sqlmock.NewRows([]string{}))
					gdb, _ := gorm.Open("mysql", db)
					return gdb
				},
			},
			args: args{
				q: func() *Query {
					q := NewQuery([]string{"id, name"})
					q.AddWhere("name IN (?)", []string{"1"})
					q.AddOr("email = ?", "test@com")
					return q
				},
			},
		},
		{
			fields: fields{
				db: func() *gorm.DB {
					db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
					mock.ExpectQuery("SELECT id, name FROM `users` LIMIT 1").WillReturnRows(sqlmock.NewRows([]string{}))
					gdb, _ := gorm.Open("mysql", db)
					return gdb
				},
			},
			args: args{
				q: func() *Query {
					q := NewQuery([]string{"id, name"})
					q.AddPreload("Profile")
					return q
				},
			},
		},
		{
			fields: fields{
				db: func() *gorm.DB {
					db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
					mock.ExpectQuery("SELECT id, name FROM `users` ORDER BY id desc LIMIT 1").WillReturnRows(sqlmock.NewRows([]string{}))
					gdb, _ := gorm.Open("mysql", db)
					return gdb
				},
			},
			args: args{
				q: func() *Query {
					q := NewQuery([]string{"id, name"})
					q.SetOrder("id desc")
					return q
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Repo{
				db: tt.fields.db(),
			}
			got, err := a.GetUser(tt.args.q())
			if (err != nil) != tt.wantErr {
				if err != gorm.ErrRecordNotFound {
					t.Errorf("Repo.GetUser() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repo.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_GetUsers(t *testing.T) {
	type fields struct {
		db func() *gorm.DB
	}
	type args struct {
		q func() *Query
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    User
		wantErr bool
	}{
		{
			fields: fields{
				db: func() *gorm.DB {
					db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
					mock.ExpectQuery("SELECT id, name FROM `users` LIMIT 2").WillReturnRows(sqlmock.NewRows([]string{}))
					gdb, _ := gorm.Open("mysql", db)
					return gdb
				},
			},
			args: args{
				q: func() *Query {
					q := NewQuery([]string{"id, name"})
					q.SetLimit(2)
					return q
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Repo{
				db: tt.fields.db(),
			}
			_, err := a.GetUsers(tt.args.q())
			if (err != nil) != tt.wantErr {
				if err != gorm.ErrRecordNotFound {
					t.Errorf("Repo.GetUser() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}
