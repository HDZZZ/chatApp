package db

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	AppCommon "github.com/HDDDZ/test/chatApp/common"
	_ "github.com/go-sql-driver/mysql"
	MapStrcuture "github.com/mitchellh/mapstructure"
)

var db *sql.DB

func init() {
	// dbc, err := sql.Open("mysql",
	// 	"root:hanzhi123@tcp(127.0.0.1:3306)/chat_app")
	dbc, err := sql.Open("mysql",
		"root:mysql@tcp(120.79.7.215:3306)/chat_app")
	if err != nil {
		fmt.Println(err)
	}
	db = dbc

	AppCommon.RegisterSubscriber(AppCommon.AppClose, func(params ...any) {
		appClosed()
	})
}

/*
  - 插入数据到数据库的表中
  - tableName: 表名,例如 users
    insertKeys: 插入的值对应的key, 例如  ["username","password"]
    values: 插入的值, 每个[]string 对应一条row
*/
func insertRows(tableName string, insertKeys []string, values ...[]string) (insertIds int64, err error) {

	var syntax string = fmt.Sprintf("INSERT INTO %s(%s) VALUES", tableName, strings.Join(insertKeys, ","))

	for _, rowValue := range values {
		syntax = syntax + "(" + strings.Join(rowValue, ",") + ")"
	}
	log("insert,syntax=", syntax)

	res, err := db.Exec(syntax)
	if err != nil {
		fmt.Println(err)
		return
	}
	insertIds, err = res.LastInsertId()
	if err != nil {
		elog(err)
		return
	}
	return
}

/*
  - 查询数据
  - querySynax: 查询语句
    call: 单次遍历结束后会被调用
    queryValue: 查询的结果会赋予给这些变量
*/
func queryRows(querySynax string, call func(), queryValue ...any) error {
	rows, err := db.Query(querySynax)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(queryValue...)
		if err != nil {
			fmt.Println(err)
			return err
		}
		call()
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

/*
  - 查询数据
  - querySynax: 查询语句
    user: 泛型,确定返回结果的类型, 注意, 需要时struct类型且成员变量的tag
    mapstructure 需要与数据库表中的colum对应,注意:当前只支持 int,string类型 queryStruct
*/
func queryStruct[T interface{}](querySynax string, user T) (results []T, err error) {
	rows, err := db.Query(querySynax)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	colums, _ := rows.Columns()
	myMap := make(map[string]interface{}, len(colums))
	cols := make([]interface{}, len(colums))
	colPtrs := make([]interface{}, len(colums))
	for i := 0; i < len(colums); i++ {
		colPtrs[i] = &cols[i]
	}
	typeUser := reflect.TypeOf(user)
	typeUserMap := make(map[string]reflect.Type, typeUser.NumField())
	for i := 0; i < typeUser.NumField(); i++ {
		if typeUser.Field(i).Tag.Get("mapstructure") == "" {
			continue
		}
		typeUserMap[typeUser.Field(i).Tag.Get("mapstructure")] = typeUser.Field(i).Type
	}

	for rows.Next() {
		err = rows.Scan(colPtrs...)
		for i, col := range cols {
			if col != nil {
				columType := typeUserMap[colums[i]]
				if columType == nil {
					continue
				}
				switch columType.Name() {
				case "string":
					myMap[colums[i]] = string(col.([]uint8))
				case "int":
					myMap[colums[i]], _ = strconv.Atoi(string(col.([]uint8)))
				}
			}
		}

		if err != nil {
			fmt.Println(err)
			return
		}
		err = MapStrcuture.Decode(myMap, &user)
		if err != nil {
			elog("MapStrcuture.Decode,err=", err)
			return
		}
		results = append(results, user)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	return
}

/*
  - 变更表中的数据
  - tableName: 表名,例如 users
    setSyntax: 变更语句
    conditionSyntax: 条件语句
*/
func updateRows(tableName string, setSyntax string, conditionSyntax string) error {

	var syntax = fmt.Sprintf("UPDATE %s set %s where %s", tableName, setSyntax, conditionSyntax)

	log("updateRows,syntax=", syntax)

	_, err := db.Exec(syntax)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

/*
  - 删除表中的rows
  - tableName: 表名,例如 users
    setSyntax: 变更语句
    conditionSyntax: 条件语句
*/
func delRows(tableName string, conditionSyntax string) error {
	var syntax = fmt.Sprintf("DELETE FROM %s where %s", tableName, conditionSyntax)
	log("updateRows,syntax=", syntax)

	_, err := db.Exec(syntax)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

type QueryRowInt interface {
	getAllField() []any
}

func log(logs ...any) {
	fmt.Println("MySQL:", logs)
}

func elog(logs ...any) {
	var mysql any = "MySQL Error:"
	logs = append([]any{"\033[1;31;40m", mysql}, logs, "\033[0m\n")
	fmt.Println(logs...)
}

func Test() {
	elog("demaxiya", "luokesas")

}
