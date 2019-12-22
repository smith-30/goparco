package conn_pool

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const query = "select sleep(0.01)"
const slowQuery = "select sleep(5)"

func doQuery(db *gorm.DB) error {
	q := fmt.Sprintf("select sleep(%v)", random(0.001, 0.03))
	return db.Exec(q).Error
}

func fetch(db *gorm.DB, q string) error {
	return db.Exec(q).Error
}

// type handler struct {
// 	db *gorm.DB
// }

// func (a *handler) reveive() error {
// 	doQuery(a.db)
// }

func random(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(max-min) + min
}

func noConnPool(db *gorm.DB) {
	db.Exec(query)
}

func connPool(db *gorm.DB) {
	db.Exec(query)
}

func connPoolTx(db *gorm.DB) {
	tx := db.Begin()
	tx.Exec(query)
	defer tx.Rollback()
}

func connPoolTxCommit(db *gorm.DB) error {
	tx := db.Begin()
	tx.Exec(query)
	return tx.Commit().Error
}

func genSetting(username, password, host, port, dbName string) string {
	address := host + ":" + port
	setting := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, address, dbName)
	return setting
}

func openDB(setting string) *gorm.DB {
	db, err := gorm.Open("mysql", setting)
	if err != nil {
		panic(err)
	}
	db.LogMode(false)
	return db
}
