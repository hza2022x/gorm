package mysql

var errCodes = map[string]uint16{
	"uniqueConstraint": 1062,
}

func (dialector Dialector) Translate(err error) error {
	//if mysqlErr, ok := err.(*mysql.MySQLError); ok {
	//	if mysqlErr.Number == errCodes["uniqueConstraint"] {
	//		return gorm.ErrDuplicatedKey
	//	}
	//}

	return err
}
