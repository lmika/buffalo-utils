package main

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"io"
)

type codeGen struct {
	file *File

	unknownStringValue string

	typeName *jen.Statement
}

func newCodeGen(file *File) *codeGen {
	return &codeGen{
		file:               file,
		unknownStringValue: "unknown",
	}
}

func (cg *codeGen) generate(w io.Writer) error {
	f := jen.NewFile(cg.file.Package)

	for _, e := range cg.file.Enums {
		cg.generateEnum(f, e)
	}

	return f.Render(w)
}

func (cg *codeGen) generateEnum(f *jen.File, enum *Enum) {
	cg.typeName = jen.Id(enum.Name)

	f.Type().Add(cg.typeName).Int()
	f.Const().DefsFunc(func(g *jen.Group) {
		for _, e := range enum.Items {
			g.Id(cg.enumItemName(e, enum)).Id(enum.Name).Op("=").Lit(e.IntValue)
		}
	})

	// Array of valid values
	f.Var().Id(enum.Name + "Values").Op("=").Index().Id(enum.Name).ValuesFunc(func(g *jen.Group) {
		for _, e := range enum.Items {
			g.Id(cg.enumItemName(e, enum))
		}
	})

	cg.generateStringMethod(f, enum)
	cg.generateScanMethod(f, enum)
	cg.generateUnmarshalTextMethod(f, enum)
	cg.generateSelectValueMethod(f, enum)
	cg.generateSelectLabelMethod(f, enum)
}

func (cg *codeGen) generateStringMethod(f *jen.File, enum *Enum) {
	receiver := jen.Id("e").Add(cg.typeName)
	f.Func().Params(receiver).Id("String").Params().String().BlockFunc(func(g *jen.Group) {
		g.Switch(jen.Id("e")).BlockFunc(func(g *jen.Group) {
			for _, e := range enum.Items {
				g.Case(jen.Id(cg.enumItemName(e, enum))).Block(
					jen.Return(jen.Lit(e.StringValue)),
				)
			}
		})
		g.Return(jen.Lit(cg.unknownStringValue))
	})
}

func (cg *codeGen) generateScanMethod(f *jen.File, enum *Enum) {
	receiver := jen.Id("e").Op("*").Add(cg.typeName)

	f.Func().Params(receiver).Id("Scan").Params(
		jen.Id("src").Interface(),
	).Error().BlockFunc(func(g *jen.Group) {
		g.Id("i").Op(",").Id("isInt").Op(":=").Id("src").Op(".").Params(jen.Id("int64"))

		g.If(jen.Op("!").Id("isInt")).BlockFunc(func(g *jen.Group) {
			g.Return(jen.Qual("errors", "New").Params(
				jen.Lit(fmt.Sprintf("%v is not an int", enum.Name))),
			)
		})

		g.Op("*").Id("e").Op("=").Id(enum.Name).Params(jen.Id("i"))
		g.Return(jen.Nil())
	})
}

func (cg *codeGen) generateUnmarshalTextMethod(f *jen.File ,enum *Enum) {
	receiver := jen.Id("e").Op("*").Add(cg.typeName)
	textParam := jen.Id("text")

	f.Func().Params(receiver).Id("UnmarshalText").Params(
		textParam.Index().Byte(),
	).Error().BlockFunc(func(g *jen.Group) {
		// TODO: allow case insensitive values
		textAsStr := jen.String().Params(jen.Id("text"))

		g.Switch(textAsStr).BlockFunc(func(g *jen.Group) {
			for _, e := range enum.Items {
				g.Case(jen.Lit(e.StringValue)).BlockFunc(func(g *jen.Group) {
					g.Op("*").Id("e").Op("=").Id(cg.enumItemName(e, enum))
					g.Return(jen.Nil())
				})
			}
		})

		// Unrecognised case
		if defaultItem, hasDefaultItem := enum.DefaultItem(); hasDefaultItem {
			// Set it to the default item
			g.Op("*").Id("e").Op("=").Id(cg.enumItemName(defaultItem, enum))
			g.Return(jen.Nil())
		} else {
			// Return an error
			g.Return(jen.Qual("errors", "New").Params(
				jen.Lit(fmt.Sprintf("invalid value for %v", enum.Name))),
			)
		}
	})
}

func (cg *codeGen) generateSelectValueMethod(f *jen.File, enum *Enum) {
	receiver := jen.Id("e").Add(cg.typeName)
	f.Func().Params(receiver).Id("SelectValue").Params().Interface().BlockFunc(func(g *jen.Group) {
		g.Return(jen.Id("e").Dot("String").Call())
	})
}

func (cg *codeGen) generateSelectLabelMethod(f *jen.File, enum *Enum) {
	receiver := jen.Id("e").Add(cg.typeName)
	f.Func().Params(receiver).Id("SelectLabel").Params().String().BlockFunc(func(g *jen.Group) {
		g.Switch(jen.Id("e")).BlockFunc(func(g *jen.Group) {
			for _, e := range enum.Items {
				g.Case(jen.Id(cg.enumItemName(e, enum))).Block(
					jen.Return(jen.Lit(cg.enumSelectLabel(e, enum))),
				)
			}
		})
		g.Return(jen.Lit(cg.unknownStringValue))
	})
}

func (cg *codeGen) enumItemName(enumItem EnumItem, enum *Enum) string {
	// TODO: configure
	return enumItem.Name + enum.Name
}

func (cg *codeGen) enumSelectLabel(enumItem EnumItem, enum *Enum) string {
	if enumItem.SelectLabel != "" {
		return enumItem.SelectLabel
	}
	return enumItem.Name
}