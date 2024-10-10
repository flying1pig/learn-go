package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

type Dialect interface {
	// DataTypeOf 将go语言的类型转换为数据库的数据类型
	DataTypeOf(typ reflect.Value) string
	// TableExistSQL 返回某个表是否存在的SQL语句
	TableExistSQL(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
