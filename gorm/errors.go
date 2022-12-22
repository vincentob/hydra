package gorm

import "github.com/go-sql-driver/mysql"

func IsDuplicateErr(err error) bool {
	e, ok := err.(*mysql.MySQLError)

	if ok {
		if e.Number == 1062 {
			return true
		}
	}

	return false
}

func IsNotFoundErr(err error) bool {
	e, ok := err.(*mysql.MySQLError)

	if ok {
		if e.Number == 1032 {
			return true
		}
	}

	return false
}
