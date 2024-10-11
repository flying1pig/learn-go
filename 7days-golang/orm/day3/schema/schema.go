package schema

import (
	"go/ast"
	"learn-go/7days-golang/orm/day3/dialect"
	"reflect"
)

type Field struct {
	Name string
	Type string
	Tag  string
}

type Schema struct {
	//被映射的对象
	Model interface{}
	//表名
	Name string
	//字段
	Fields []*Field
	//包含所有的字段名(列名)
	FieldNames []string
	//记录字段名和字段的映射关系
	fieldMap map[string]*Field
}

func (s *Schema) GetField(name string) *Field {
	return s.fieldMap[name]
}

func (s *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var filedValues []interface{}
	for _, field := range s.Fields {
		filedValues = append(filedValues, destValue.FieldByName(field.Name).Interface())
	}
	return filedValues
}

func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("orm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}
