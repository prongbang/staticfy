package cmd

import "github.com/prongbang/staticfy/pkg/staticfy"

func Run() {
	handle := staticfy.NewHandler()
	handle.Register()
}
