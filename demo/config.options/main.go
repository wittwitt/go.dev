package main

import "fmt"

type Options struct {
	Name string
	Age  int
}

type Option func(*Options)

func OptName(name string) Option {
	return func(ops *Options) {
		ops.Name = name
	}
}

func OptAge(age int) Option {
	return func(ops *Options) {
		ops.Age = age
	}
}

type Some struct {
	opts *Options
	addr string
}

func NewSome(addr string, options ...Option) *Some {
	opts := &Options{
		Name: "default",
	}

	for _, opt := range options {
		opt(opts)
	}

	return &Some{opts: opts, addr: addr}
}

func main() {
	{
		s := NewSome("addr")
		fmt.Println(s.opts.Name)
	}
	{
		s := NewSome("addr", OptName("lisi"))
		fmt.Println(s.opts.Name)
	}

	{
		s := NewSome("addr", OptName("lisi"), OptAge(10))
		fmt.Println(s.opts.Name)
	}
}
