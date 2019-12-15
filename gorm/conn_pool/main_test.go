package conn_pool

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func BenchmarkConn(b *testing.B) {
	setting := genSetting("root", "root", "localhost", "13306", "test_database")
	tests := []struct {
		name         string
		maxOpenConns int
		maxIdleConns int
	}{
		{
			maxIdleConns: 10,
			maxOpenConns: 100,
		},
	}
	for _, tt := range tests {
		connPoolDB := openDB(setting)
		// connPoolTxDB := openDB(setting)

		// b.Run(tt.name, func(b *testing.B) {
		// 	b.ResetTimer()
		// 	for i := 0; i < b.N; i++ {
		// 		db := openDB(setting)
		// 		noConnPool(db)
		// 	}
		// })
		// b.Run(tt.name, func(b *testing.B) {
		// 	b.ResetTimer()
		// 	for i := 0; i < b.N; i++ {
		// 		db := openDB(setting)
		// 		connPoolTx(db)
		// 	}
		// })

		connPoolDB.DB().SetMaxIdleConns(tt.maxIdleConns)
		connPoolDB.DB().SetMaxOpenConns(tt.maxOpenConns)
		connPoolDB.DB().SetConnMaxLifetime(time.Hour)

		go func() {
			ticker := time.NewTicker(200 * time.Millisecond)
			for {
				select {
				case <-ticker.C:
					fmt.Printf("%#v\n", connPoolDB.DB().Stats())
				}
			}
		}()

		b.Run(tt.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				connPool(connPoolDB)
			}
		})

		// connPoolTxDB.DB().SetMaxIdleConns(tt.maxIdleConns)
		// connPoolTxDB.DB().SetMaxOpenConns(tt.maxOpenConns)
		// connPoolTxDB.DB().SetConnMaxLifetime(time.Hour)
		// b.Run(tt.name, func(b *testing.B) {
		// 	b.ResetTimer()
		// 	for i := 0; i < b.N; i++ {
		// 		connPoolTx(connPoolTxDB)
		// 	}
		// })

		// connPoolTxDB.DB().SetMaxIdleConns(tt.maxIdleConns)
		// connPoolTxDB.DB().SetMaxOpenConns(tt.maxOpenConns)
		// connPoolTxDB.DB().SetConnMaxLifetime(time.Hour)
		// b.Run(tt.name, func(b *testing.B) {
		// 	b.ResetTimer()
		// 	for i := 0; i < b.N; i++ {
		// 		connPoolTxCommit(connPoolTxDB)
		// 	}
		// })
	}
}

func Test_maxConn(t *testing.T) {
	setting := genSetting("root", "root", "localhost", "13306", "test_database")
	type args struct {
		connCount int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "接続数が上限のとき",
			args: args{
				connCount: 100,
			},
		},
		{
			name: "接続数が上限を超えているとき",
			args: args{
				connCount: 101,
			},
		},
	}
	for _, tt := range tests {
		connPoolDB := openDB(setting)
		if tt.args.connCount > 100 {
			connPoolDB.DB().SetMaxOpenConns(100)
		}

		t.Run(tt.name, func(t *testing.T) {
			db := connPoolDB
			wg := &sync.WaitGroup{}
			for index := 0; index < tt.args.connCount; index++ {
				go func() {
					wg.Add(1)
					defer wg.Done()
					if err := doQuery(db); err != nil {
						t.Errorf("%v\n", err)
					}
				}()
			}
			wg.Wait()
			connPoolDB.Close()
		})
	}
}

func BenchmarkIdleConn(b *testing.B) {
	setting := genSetting("root", "root", "localhost", "13306", "test_database")
	type args struct {
		jobCount    int
		loopCount   int
		maxConn     int
		maxIdleConn int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "job 5",
			args: args{
				jobCount:    5,
				loopCount:   100,
				maxConn:     100,
				maxIdleConn: 30,
			},
		},
	}
	for _, tt := range tests {
		connPoolDB := openDB(setting)
		connPoolDB.DB().SetMaxOpenConns(tt.args.maxConn)
		connPoolDB.DB().SetMaxIdleConns(tt.args.maxIdleConn)
		connPoolDB.DB().SetConnMaxLifetime(time.Hour)

		b.Run(tt.name, func(b *testing.B) {
			sem := make(chan struct{}, tt.args.jobCount)
			go func(db *gorm.DB) {
				ticker := time.NewTicker(200 * time.Millisecond)
				for {
					select {
					case <-ticker.C:
						fmt.Printf("%#v\n", db.DB().Stats())
					}
				}
			}(connPoolDB)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for index := 0; index < tt.args.loopCount; index++ {
					sem <- struct{}{}
					go func(db *gorm.DB) {
						defer func() {
							<-sem
						}()
						if err := doQuery(db); err != nil {
							b.Errorf("%v\n", err)
						}
					}(connPoolDB)
				}
			}
		})
		connPoolDB.Close()
	}
}

func Test_useIdleConn(t *testing.T) {
	setting := genSetting("root", "root", "localhost", "13306", "test_database")
	connPoolDB := openDB(setting)
	connPoolDB.DB().SetMaxOpenConns(100)
	connPoolDB.DB().SetMaxIdleConns(5)
	connPoolDB.DB().SetConnMaxLifetime(time.Hour)

	sem := make(chan struct{}, 5)
	qs := getQueries(10)
	for _, item := range qs {
		sem <- struct{}{}
		go func(db *gorm.DB, item string) {
			defer func() {
				<-sem
			}()
			if err := fetch(db, item); err != nil {
				panic(err)
			}
		}(connPoolDB, item)
	}
}

func getQueries(num int) []string {
	qs := make([]string, 0, num)
	for index := 0; index < num; index++ {
		qs = append(qs, fmt.Sprintf("select sleep(%v)", random(0.001, 0.03)))
	}
	return qs
}
