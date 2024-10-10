package clause

type generator func(values ...interface{}) (string, []interface{})

var generators map[Type]generator
