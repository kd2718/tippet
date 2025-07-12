package main

import (
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

func main() {
	pool := memory.NewGoAllocator()

	NameBuilder := array.NewStringBuilder(pool)
	defer NameBuilder.Release()

	IdBuilder := array.NewInt32Builder(pool)
	defer IdBuilder.Release()

	IdBuilder.Append(1)
	NameBuilder.Append("Kory")
	IdBuilder.Append(2)
	NameBuilder.Append("Cheri")

	IdArray := IdBuilder.NewArray()
	defer IdArray.Release()

	NameArray := NameBuilder.NewArray()
	defer NameArray.Release()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "MyId", Type: arrow.PrimitiveTypes.Int32},
			{Name: "MyName", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	builder := array.NewRecordBuilder(pool, schema)
	defer builder.Release()

	record := array.NewRecord(schema, []arrow.Array{IdArray, NameArray}, -1)
	defer record.Release()

	fmt.Println(record)

}
