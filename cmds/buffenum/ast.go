package main

import "github.com/alecthomas/participle"

type fileAst struct {
	Package string     `parser:"'package' @Ident"`
	Enums   []*enumAst `parser:"@@*"`
}

type enumAst struct {
	Name  string        `parser:"'enum' @Ident"`
	Items []enumItemAst `parser:"'{' @@* '}'"`
}

type enumItemAst struct {
	Name      string `parser:"@Ident"`
	IntValue  int    `parser:"'(' @Int ')'"`
	StrValue  string `parser:"'=' @String"`
	IsDefault bool   `parser:"@'default'?"`
	SelectLabel string `parser:"('label' '=' @String)? ';'"`
}

func (e fileAst) toModel() *File {
	items := make([]*Enum, 0)
	for _, i := range e.Enums {
		items = append(items, i.toModel())
	}

	return &File{Package: e.Package, Enums: items}
}

func (e enumAst) toModel() *Enum {
	items := make([]EnumItem, 0)
	for _, i := range e.Items {
		items = append(items, i.toModel())
	}

	return &Enum{Name: e.Name, Items: items}
}

func (e enumItemAst) toModel() EnumItem {
	return EnumItem{Name: e.Name, IntValue: e.IntValue, StringValue: e.StrValue, IsDefault: e.IsDefault, SelectLabel: e.SelectLabel}
}

var parser = participle.MustBuild(&fileAst{})
