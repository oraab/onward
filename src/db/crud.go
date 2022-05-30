package db

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

type Crud struct {
	dbPath string
}
// select
// method that receives arguments and constructs a query `SELECT <arg1> FROM <arg2> WHERE <arg3>`

func NewCrud(dbPath string) *Crud {
	return &Crud{dbPath: dbPath}
}

func (c *Crud) Select(columns []string, table string, condition []string) [][]interface{} {
	query := constructQuery(columns, table, condition)
	log.Infof("query: %v", query)
	return make([][]interface{},0)
}

// insert

// update

// delete

func constructQuery(columns []string, table string, condition []string) string {
	query := "SELECT "
	constructedColumns := "*"
	if len(columns) > 0 {
		constructedColumns = strings.Join(columns,", ")
	}
	query += constructedColumns
	query += " FROM "+table+" "
	constructedCondition := ""
	if len(condition) > 0 {
		constructedCondition = strings.Join(condition," AND ")
	}
	if constructedCondition != "" {
		query += "WHERE "+constructedCondition
	}
	query += ";"
	return query
}