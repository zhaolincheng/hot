package mysql

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"hot/utils"
)

var db *sql.DB

// 初始化连接池
func init() {
	var err error
	db, err = sql.Open(viper.GetString("mysql.driver"), viper.GetString("mysql.source"))
	if err != nil {
		utils.Error.Println(err)
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns"))            // 最大链接
	db.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns"))            // 空闲连接，也就是连接池里面的数量
	db.SetConnMaxLifetime(viper.GetDuration("mysql.connMaxLifetime")) // 连接最大生命周期
}

func Ping() {
	_ = db.Ping()
}

func Exec(query string, args ...interface{}) {
	_, _ = db.Exec(query, args)
}

func Query(query string, args ...interface{}) {
	_, _ = db.Query(query, args)
}

func Begin() {
	_, _ = db.Begin()
}

func BeginTx(ctx context.Context, opts *sql.TxOptions) {
	_, _ = db.BeginTx(ctx, opts)
}
