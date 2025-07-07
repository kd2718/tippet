package pipeline

type Source interface {
	Read() // add return types
}

type Sink interface {
	Write() // add return types
}
