package helpers

import (
	"fmt"
	"reflect"
)

type Builder struct {
	Type reflect.Type
}

func (b *Builder) BuildInsertion(record interface{}) {
	t := record.TypeOf(&record).Elem()
	fmt.Println("Type: ", t)
}
