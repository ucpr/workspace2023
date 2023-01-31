package main

import (
	"fmt"

	"github.com/ucpr/workspace2023/wire_variadic_param/provider"
)

func NewFuga(hoges ...provider.Hoge) string {
	var res string
	for i := range hoges {
		res += "/" + fmt.Sprint(hoges[i])
	}
	return res
}
