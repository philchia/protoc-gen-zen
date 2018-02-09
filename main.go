package main

import (
	"github.com/lyft/protoc-gen-star"
	"github.com/philchia/protoc-gen-zen/zengen"
)

func main() {
	pgs.Init(pgs.IncludeGo()).
		RegisterPlugin(zengen.New("zengen")).
		RegisterPostProcessor(pgs.GoFmt()).
		Render()
}
