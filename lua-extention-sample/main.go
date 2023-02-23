package main

import (
	"fmt"

	"github.com/Shopify/go-lua"
)

func main() {
	l := lua.NewState()
	lua.OpenLibraries(l)

	// TODO: lua.LoadFile に直す
	if err := lua.DoFile(l, "hoge.lua"); err != nil {
		panic(err)
	}

	// フィボナッチ数の計算
	/*
		lua ではこうなる
		result = fibonacci(10)
	*/
	arg := 1 // 引数の数
	res := 1 // 戻り値の数
	l.Global("fibonacci")
	l.PushNumber(10)
	l.Call(arg, res)
	l.SetGlobal("result")

	top := l.Top()
	ans, ok := l.ToInteger(top)
	if !ok {
		panic("failed to cast to integer")
	}
	fmt.Println(ans)

	l.Pop(top) // スタックのお掃除
}
