package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var db *sql.DB

func Set(newDb *sql.DB) {
	db = newDb
}

type Condition struct {
	Col      string
	Operator string
	Val      interface{}
}

func existRowQuery(table string, conditions []Condition) *sql.Row {
	args := []interface{}{}
	where := " true "
	argCount := 1
	for _, condition := range conditions {
		if condition.Operator == "IS NOT NULL" || condition.Operator == "IS NULL" {
			where = where + fmt.Sprintf(" AND %s %s", condition.Col, condition.Operator)
			continue
		}
		args = append(args, condition.Val)
		where = where + fmt.Sprintf(" AND %s %s $%d", condition.Col, condition.Operator, argCount)
		argCount++
	}
	query := `SELECT 1 FROM ` + table + ` WHERE ` + where + `;`
	return db.QueryRow(query, args...)
}

func Exist(t *testing.T, table string, conditions []Condition) {
	row := existRowQuery(table, conditions)
	isExist := false
	assert.Nil(t, row.Scan(&isExist))
}

func CountByPostgresArray(t *testing.T, expected int, columnName string, table string, conditions []Condition) {
	args := []interface{}{}
	where := " true "
	argCount := 1
	for _, condition := range conditions {
		if condition.Operator == "IS NOT NULL" || condition.Operator == "IS NULL" {
			where = where + fmt.Sprintf(" AND %s %s", condition.Col, condition.Operator)
			continue
		}
		args = append(args, condition.Val)
		where = where + fmt.Sprintf(" AND %s %s $%d", condition.Col, condition.Operator, argCount)
		argCount++
	}

	query := fmt.Sprintf("SELECT array_length(%s, 1) FROM ", columnName) + ` "` + table + `" WHERE ` + where + `;`
	row := db.QueryRow(query, args...)
	var count int
	row.Scan(&count)
	fmt.Println(query, count)
	assert.Equal(t, expected, count)
}

//func Count(t *testing.T, expected int, table string, conditions []Condition) {
//	args := []interface{}{}
//	where := " true "
//	argCount := 1
//	for _, condition := range conditions {
//		if condition.Operator == "IS NOT NULL" || condition.Operator == "IS NULL" {
//			where = where + fmt.Sprintf(" AND %s %s", condition.Col, condition.Operator)
//			continue
//		}
//		args = append(args, condition.Val)
//		where = where + fmt.Sprintf(" AND %s %s $%d", condition.Col, condition.Operator, argCount)
//		argCount++
//	}
//	query := `SELECT count(1) FROM "` + table + `" WHERE ` + where + `;`
//	row := db.QueryRow(query, args...)
//	var count int
//	row.Scan(&count)
//	assert.Equal(t, expected, count)
//}
