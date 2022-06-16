/*
 * MIT License
 *
 * Copyright (c) 2021 zeromicro
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 */

package parser

import (
	"strings"

	"github.com/miaogaolin/ddlparser/gen"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Token is an abstraction from each lexical element, literal, etc.
type Token interface {
	GetLine() int
	GetColumn() int
	GetText() string
	SetText(s string)
}

type parseOption func(text string) string

func parseToken(t antlr.Token, option ...parseOption) string {
	t.GetText()
	text := t.GetText()
	for _, o := range option {
		text = o(text)
	}
	return text
}

func parseTerminalNode(t antlr.TerminalNode, option ...parseOption) string {
	text := t.GetText()
	for _, o := range option {
		text = o(text)
	}
	return text
}

func withTrim(c string) parseOption {
	return func(text string) string {
		return strings.Trim(text, c)
	}
}

func withUpperCase() parseOption {
	return func(text string) string {
		return strings.ToUpper(text)
	}
}

func withReplacer(oldnew ...string) parseOption {
	return func(text string) string {
		return strings.NewReplacer(oldnew...).Replace(text)
	}
}

func parseDataTypeSource(ctx gen.IDataTypeContext) string {
	if ctx == nil {
		return ""
	}
	source := ""
	switch t := ctx.(type) {
	case *gen.StringDataTypeContext:
		source = t.GetTypeName().GetText()
		if t.LengthOneDimension() != nil {
			source += t.LengthOneDimension().GetText()
		}
		if t.GetBinaryType() != nil {
			source += t.GetBinaryType().GetText()
		}

	case *gen.NationalStringDataTypeContext:
		source = t.GetText()
	case *gen.NationalVaryingStringDataTypeContext:
		source = t.GetText()
	case *gen.DimensionDataTypeContext:
		source = t.GetText()
	case *gen.SimpleDataTypeContext:
		source = t.GetText()
	case *gen.CollectionDataTypeContext:
		source = t.GetTypeName().GetText()
		if t.CollectionOptions() != nil {
			source += t.CollectionOptions().GetText()
		}
	case *gen.SpatialDataTypeContext:
		source = t.GetText()
	case *gen.LongVarcharDataTypeContext:
		source = t.GetTypeName().GetText()
	case *gen.LongVarbinaryDataTypeContext:
		source = t.GetText()
	}
	return source
}
