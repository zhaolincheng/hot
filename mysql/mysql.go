package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"hot/common/util"
	"log"
	"math"
	"strconv"
	"strings"
)

var globalDb *sql.DB

type MySql struct {
	table string  // 表名
	field string  // 字段
	where string  // where语句
	limit string  // 限制条数
	order string  // 排序条件
	exec  string  // 执行sql语句
	conn  *sql.DB // 数据库连接
}

// 初始化连接池
func init() {
	config := GetConfig()
	db, err := sql.Open(config.Driver, config.Source)
	if err != nil {
		util.Error.Fatalln(err)
	}
	db.SetMaxOpenConns(config.MaxOpenConns)       // 最大链接
	db.SetMaxIdleConns(config.MaxIdleConns)       // 空闲连接，也就是连接池里面的数量
	db.SetConnMaxLifetime(config.ConnMaxLifetime) // 连接最大生命周期
	globalDb = db
}

func (MySql MySql) GetConn() *MySql {
	MySql.conn = globalDb
	return &MySql
}

func (MySql *MySql) Close() {
	err := MySql.conn.Close()
	if err != nil {
		util.Error.Fatalln(err)
	}
}

/**
查询方法
*/
func (MySql *MySql) Select(table string, field []string) *MySql {
	var allField string
	allField = strings.Join(field, ",")
	MySql.field = "select " + allField + " from " + table
	MySql.table = table
	return MySql
}

/**
where子句
*/
func (MySql *MySql) Where(cond map[string]string) *MySql {
	var where = ""
	if len(cond) != 0 {
		where = " where "
		for key, value := range cond {
			if !strings.Contains(key, "=") && !strings.Contains(key, ">") && !strings.Contains(key, "<") {
				key += "="
			}
			where += key + "'" + value + "'" + " and "
		}
	}
	// 删除所有字段最后一个,
	where = strings.TrimSuffix(where, "and ")
	MySql.where = where
	return MySql
}

func (MySql *MySql) Limit(limit int) *MySql {
	MySql.limit = " limit " + strconv.Itoa(limit)
	return MySql
}

func (MySql *MySql) OrderBy(order ...string) *MySql {
	if len(order) > 2 || len(order) <= 0 {
		log.Fatal("传入参数错误")
	} else if len(order) == 1 {
		MySql.order = " ORDER BY " + order[0] + " ASC"
	} else {
		MySql.order = " ORDER BY " + order[0] + " " + order[1]
	}
	return MySql
}

/**
更新方法
*/
func (MySql MySql) Update(table string, str map[string]string) int64 {
	var tempStr = ""
	var allValue []interface{}
	for key, value := range str {
		tempStr += key + "=" + "?" + ","
		allValue = append(allValue, value)
	}
	tempStr = strings.TrimSuffix(tempStr, ",")
	MySql.exec = "update " + table + " set " + tempStr
	var allStr = MySql.exec + MySql.where
	stmt, err := MySql.conn.Prepare(allStr)
	if err != nil {
		util.Error.Fatalln(err)
	}
	res, err := stmt.Exec(allValue...)
	if err != nil {
		util.Error.Fatalln(err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		util.Error.Fatalln(err)
	}
	return rows

}

/**
删除方法
*/
func (MySql MySql) Delete(table string) int64 {
	var tempStr = ""
	tempStr = "delete from " + table + MySql.where
	fmt.Println(tempStr)
	stmt, err := MySql.conn.Prepare(tempStr)
	if err != nil {
		util.Error.Fatalln(err)
	}
	res, err := stmt.Exec()
	if err != nil {
		util.Error.Fatalln(err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		util.Error.Fatalln(err)
	}
	return rows
}

/**
插入方法
*/
func (MySql MySql) Insert(table string, data map[string]string) int64 {
	var allField = ""
	var allValue = ""
	var allTrueValue []interface{}
	if len(data) != 0 {
		for key, value := range data {
			allField += key + ","
			allValue += "?" + ","
			allTrueValue = append(allTrueValue, value)
		}
	}
	allValue = strings.TrimSuffix(allValue, ",")
	allField = strings.TrimSuffix(allField, ",")
	allValue = "(" + allValue + ")"
	allField = "(" + allField + ")"
	var theStr = "insert into " + table + " " + allField + " values " + allValue
	stmt, err := MySql.conn.Prepare(theStr)
	if err != nil {
		util.Error.Fatalln(err)
	}
	res, err := stmt.Exec(allTrueValue...)
	if err != nil {
		util.Error.Fatalln(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		util.Error.Fatalln(err)
	}
	return id
}

/**
分页查询
*/
func (MySql MySql) Pagination(Page int, Limit int) map[string]interface{} {
	res := MySql.GetConn().Select(MySql.table, []string{"count(*) as count"}).QueryRow()
	count, _ := strconv.Atoi(res["count"])
	// 计算总页码数
	totalPage := int(math.Ceil(float64(count) / float64(Limit)))
	if Page > totalPage {
		Page = totalPage
	}
	if Page <= 0 {
		Page = 1
	}
	// 计算偏移量
	setOff := (Page - 1) * Limit
	queryStr := MySql.field + MySql.where + MySql.order + " limit " + strconv.Itoa(setOff) + "," + strconv.Itoa(Limit)
	rows, err := MySql.conn.Query(queryStr)
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		util.Error.Fatalln(err)
	}
	Column, err := rows.Columns()
	if err != nil {
		util.Error.Fatalln(err)
	}
	// 创建一个查询字段类型的slice
	values := make([]sql.RawBytes, len(Column))
	// 创建一个任意字段类型的slice
	scanArgs := make([]interface{}, len(values))
	// 创建一个slice保存所以的字段
	var allRows []interface{}
	for i := range values {
		// 把values每个参数的地址存入scanArgs
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		// 把存放字段的元素批量放进去
		err = rows.Scan(scanArgs...)
		if err != nil {
			util.Error.Fatalln(err)
		}
		tempRow := make(map[string]string, len(Column))
		for i, col := range values {
			var key = Column[i]
			tempRow[key] = string(col)
		}
		allRows = append(allRows, tempRow)
	}
	returnData := make(map[string]interface{})
	returnData["totalPage"] = totalPage
	returnData["currentPage"] = Page
	returnData["rows"] = allRows
	return returnData
}

func (MySql MySql) QueryAll() []map[string]string {
	var queryStr = MySql.field + MySql.where + MySql.order + MySql.limit
	rows, err := MySql.conn.Query(queryStr)
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		util.Error.Fatalln(err)
	}
	Column, err := rows.Columns()
	if err != nil {
		util.Error.Fatalln(err)
	}
	// 创建一个查询字段类型的slice
	values := make([]sql.RawBytes, len(Column))
	// 创建一个任意字段类型的slice
	scanArgs := make([]interface{}, len(values))
	// 创建一个slice保存所以的字段
	var allRows []map[string]string
	for i := range values {
		// 把values每个参数的地址存入scanArgs
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		// 把存放字段的元素批量放进去
		err = rows.Scan(scanArgs...)
		if err != nil {
			util.Error.Fatalln(err)
		}
		tempRow := make(map[string]string, len(Column))
		for i, col := range values {
			var key = Column[i]
			tempRow[key] = string(col)
		}
		allRows = append(allRows, tempRow)
	}
	return allRows
}

func (MySql MySql) ExecSql(queryStr string) []map[string]string {
	rows, err := MySql.conn.Query(queryStr)
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		util.Error.Fatalln(err)
	}
	Column, err := rows.Columns()
	if err != nil {
		util.Error.Fatalln(err)
	}
	// 创建一个查询字段类型的slice
	values := make([]sql.RawBytes, len(Column))
	// 创建一个任意字段类型的slice
	scanArgs := make([]interface{}, len(values))
	// 创建一个slice保存所以的字段
	var allRows []map[string]string
	for i := range values {
		// 把values每个参数的地址存入scanArgs
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		// 把存放字段的元素批量放进去
		err = rows.Scan(scanArgs...)
		if err != nil {
			util.Error.Fatalln(err)
		}
		tempRow := make(map[string]string, len(Column))
		for i, col := range values {
			var key = Column[i]
			tempRow[key] = string(col)
		}
		allRows = append(allRows, tempRow)
	}
	return allRows
}

/**
查询单行
*/
func (MySql MySql) QueryRow() map[string]string {
	var queryStr = MySql.field + MySql.where + MySql.order + MySql.limit
	result, err := MySql.conn.Query(queryStr)
	if result != nil {
		defer result.Close()
	}
	if err != nil {
		util.Error.Fatalln(err)
	}
	Column, err := result.Columns()
	// 创建一个查询字段类型的slice的键值对
	values := make([]sql.RawBytes, len(Column))
	// 创建一个任意字段类型的slice的键值对
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		// 把values每个参数的地址存入scanArgs
		scanArgs[i] = &values[i]
	}

	for result.Next() {
		err = result.Scan(scanArgs...)
		if err != nil {
			util.Error.Fatalln(err)
		}
	}
	tempRow := make(map[string]string, len(Column))
	for i, col := range values {
		var key = Column[i]
		tempRow[key] = string(col)
	}
	return tempRow

}
