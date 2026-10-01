// SPDX-License-Identifier: Apache-2.0

//go:build ignore

// go-symbols produces a symbol index for this repository's own Go source, in the format
// section 5 of the process definition fixes.
//
// It is this project's own tooling and not part of Xeno, which ships no indexer and names
// none: #149 put the choice of tool on the project, and this repository is a project whose
// language is Go. It sits beside github-settings.sh and module-set.sh for the same reason.
// The build tag keeps it out of the module, so the release and the five binaries never see
// it.
//
//	go run scripts/go-symbols.go > .xeno/local/index/symbols.yaml
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const toolVersion = "0.1.0"

type symbol struct {
	name      string
	kind      string
	file      string
	line      int
	container string
}

func main() {
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}
	syms, err := walk(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	out := os.Stdout
	fmt.Fprintf(out, "tool: go-symbols\n")
	fmt.Fprintf(out, "tool_version: %s\n", toolVersion)
	fmt.Fprintf(out, "produced_at: %q\n", time.Now().UTC().Format(time.RFC3339))
	fmt.Fprintf(out, "symbols:\n")
	for _, s := range syms {
		fmt.Fprintf(out, "  - name: %q\n", s.name)
		fmt.Fprintf(out, "    kind: %s\n", s.kind)
		fmt.Fprintf(out, "    file: %s\n", s.file)
		fmt.Fprintf(out, "    line: %d\n", s.line)
		if s.container != "" {
			fmt.Fprintf(out, "    container: %q\n", s.container)
		}
	}
}

// walk parses every Go file outside vendor and testdata. Function bodies are skipped: the
// format holds where a name is defined and nothing more, so parsing them would be work
// thrown away.
func walk(root string) ([]symbol, error) {
	fset := token.NewFileSet()
	var out []symbol
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, rerr := filepath.Rel(root, path)
		if rerr != nil {
			return rerr
		}
		rel = filepath.ToSlash(rel)
		if d.IsDir() {
			switch d.Name() {
			case "vendor", "testdata", ".git", ".xeno":
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(rel, ".go") {
			return nil
		}
		f, perr := parser.ParseFile(fset, path, nil, parser.SkipObjectResolution)
		if perr != nil {
			// A file that does not parse contributes nothing and stops nothing. An index is
			// allowed to be incomplete, which is cheaper than refusing to index a repository
			// over one generated file.
			fmt.Fprintf(os.Stderr, "skipped %s: %v\n", rel, perr)
			return nil
		}
		out = append(out, fileSymbols(fset, f, rel)...)
		return nil
	})
	return out, err
}

func fileSymbols(fset *token.FileSet, f *ast.File, rel string) []symbol {
	var out []symbol
	at := func(p token.Pos) int { return fset.Position(p).Line }

	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			kind, container := "func", ""
			if d.Recv != nil && len(d.Recv.List) == 1 {
				kind, container = "method", receiver(d.Recv.List[0].Type)
			}
			out = append(out, symbol{d.Name.Name, kind, rel, at(d.Name.Pos()), container})
		case *ast.GenDecl:
			out = append(out, genSymbols(d, rel, at)...)
		}
	}
	return out
}

func genSymbols(d *ast.GenDecl, rel string, at func(token.Pos) int) []symbol {
	var out []symbol
	for _, spec := range d.Specs {
		switch s := spec.(type) {
		case *ast.TypeSpec:
			out = append(out, symbol{s.Name.Name, typeKind(s.Type), rel, at(s.Name.Pos()), ""})
			out = append(out, fields(s, rel, at)...)
		case *ast.ValueSpec:
			kind := "var"
			if d.Tok == token.CONST {
				kind = "const"
			}
			for _, n := range s.Names {
				out = append(out, symbol{n.Name, kind, rel, at(n.Pos()), ""})
			}
		}
	}
	return out
}

// fields indexes a struct's fields and an interface's methods, with the type as their
// container. That is what "enclosing container" is for, and without it a field name is
// unfindable in a repository where twenty types have a Name.
func fields(s *ast.TypeSpec, rel string, at func(token.Pos) int) []symbol {
	var out []symbol
	switch t := s.Type.(type) {
	case *ast.StructType:
		for _, f := range t.Fields.List {
			for _, n := range f.Names {
				out = append(out, symbol{n.Name, "field", rel, at(n.Pos()), s.Name.Name})
			}
		}
	case *ast.InterfaceType:
		for _, m := range t.Methods.List {
			for _, n := range m.Names {
				out = append(out, symbol{n.Name, "method", rel, at(n.Pos()), s.Name.Name})
			}
		}
	}
	return out
}

func typeKind(t ast.Expr) string {
	switch t.(type) {
	case *ast.StructType:
		return "struct"
	case *ast.InterfaceType:
		return "interface"
	default:
		return "type"
	}
}

func receiver(t ast.Expr) string {
	if star, ok := t.(*ast.StarExpr); ok {
		t = star.X
	}
	if id, ok := t.(*ast.Ident); ok {
		return id.Name
	}
	return ""
}
