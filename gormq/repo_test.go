package gormq

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func emptyConn() *gorm.DB {
	db, _, _ := sqlmock.New()
	gdb, _ := gorm.Open("mysql", db)
	gdb.LogMode(true)
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
					return emptyConn()
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
					return emptyConn()
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
					return emptyConn()
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
					return emptyConn()
				},
			},
			args: args{
				q: func() *Query {
					q := NewQuery([]string{"id, name"})
					q.AddWhere("name IN (?)", []string{"1", "2"})
					q.AddOr("email = ?", "test@com")
					return q
				},
			},
		},
		{
			fields: fields{
				db: func() *gorm.DB {
					return emptyConn()
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
					return emptyConn()
				},
			},
			args: args{
				q: func() *Query {
					q := NewQuery([]string{"id, name"})
					q.AddWhere("id IN (?)", []int{1, 2})
					q.EnableForUpdate()
					return q
				},
			},
		},
		{
			fields: fields{
				db: func() *gorm.DB {
					return emptyConn()
				},
			},
			args: args{
				q: func() *Query {
					q := NewQuery([]string{"id, name"})
					q.AddWhere("id IN (?)", []int{1, 2})
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
			got, _ := a.GetUser(tt.args.q())
			// if (err != nil) != tt.wantErr {
			// 	t.Errorf("Repo.GetUser() error = %v, wantErr %v", err, tt.wantErr)
			// 	return
			// }
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repo.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
