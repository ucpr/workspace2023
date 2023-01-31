//go:build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/ucpr/workspace2023/wire_variadic_param/provider"
)

func initialize() string {
	wire.Build(
		NewFuga,
		provider.Set,
	)
	return ""
}
