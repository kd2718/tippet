package pipeline

import "github.com/apache/arrow-go/v18/arrow"

type Source interface {
	Read() (arrow.Record, error)
}

type Sink interface {
	Write(arrow.Record) (int64, error)
}
