package provider

import "github.com/google/wire"

type Func func() Hoge

var Set = wire.NewSet(
	wire.Value([]Func{NewHoge1, NewHoge2, NewHoge3}),
)
