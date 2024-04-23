//go:build wireinject
// +build wireinject

package todo

import "github.com/google/wire"

func InitApp() {
	panic(wire.Build())
}
