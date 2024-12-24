package codegen

import (
	"bytes"
	"fmt"
	"text/template"

	"go/ast"
	"go/format"
	"go/token"
	"log"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/tools/imports"
)

type Generator struct {
	buf     bytes.Buffer
	models  []TlvModel
	tagRe   *regexp.Regexp
	pkgName string
}

func NewGenerator() *Generator {
	return &Generator{
		models: make([]TlvModel, 0),
		tagRe:  regexp.MustCompile(`tlv:"(?P<typ>[0-9a-fA-FxX]+)"`),
	}
}

func (g *Generator) parseTag(tag *ast.BasicLit) uint64 {
	if tag == nil {
		return 0
	}
	matches := g.tagRe.FindStringSubmatch(tag.Value)
	if len(matches) <= 1 {
		return 0
	}
	typVal, err := strconv.ParseUint(matches[1], 0, 0)
	if err != nil {
		return 0
	}
	return typVal
}

func (g *Generator) parseDoc(doc *ast.CommentGroup, indicator string) string {
	if doc == nil {
		return ""
	}

	prefixes := []string{
		fmt.Sprintf("//+%s:", indicator),
		fmt.Sprintf("// +%s:", indicator),
	}
	for _, c := range doc.List {
		for _, prefix := range prefixes {
			if c != nil && strings.HasPrefix(c.Text, prefix) {
				return c.Text[len(prefix):]
			}
		}
	}

	return ""
}

func ParseField(name string, typeNum uint64, fieldStr string, model *TlvModel) (TlvField, error) {
	fieldType := fieldStr
	annotation := ""
	if i := strings.Index(fieldStr, ":"); i >= 0 {
		fieldType = fieldStr[:i]
		annotation = fieldStr[i+1:]
	}
	return CreateField(fieldType, name, typeNum, annotation, model)
}

// This is the entry point of the codegen module
func (g *Generator) ProcessDecl(node ast.Node) bool {
	if g.pkgName == "" {
		pkg, ok := node.(*ast.Ident)
		if ok {
			g.pkgName = strings.TrimSuffix(pkg.Name, "_test")
			return true
		}
	}

	decl, ok := node.(*ast.GenDecl)
	if !ok || decl.Tok != token.TYPE {
		// We only care about type declarations.
		return true
	}
	typSpec, ok := decl.Specs[0].(*ast.TypeSpec)
	if !ok {
		// Skip other declarations
		return false
	}
	stru, ok := typSpec.Type.(*ast.StructType)
	if !ok || stru.Fields == nil {
		// Skip other declarations
		return false
	}

	// Create the TLV model
	model := TlvModel{
		Name:   typSpec.Name.Name,
		Fields: make([]TlvField, 0),
	}

	// Parse model options from the docstring
	modelOptions := g.parseDoc(decl.Doc, "tlv-model")
	if modelOptions != "" {
		opts := strings.Split(modelOptions, ",")
		for _, opt := range opts {
			model.ProcessOption(opt)
		}
	}

	// Parse all fields
	for _, f := range stru.Fields.List {
		if len(f.Names) <= 0 {
			continue
		}

		// Parse field parameters
		fieldName := f.Names[0].Name
		tlvTypNum := g.parseTag(f.Tag)
		fieldStr := g.parseDoc(f.Doc, "field")
		if fieldStr == "" {
			continue
		}

		// Dispatch to specific field for parsing
		f, err := ParseField(fieldName, tlvTypNum, fieldStr, &model)
		if err != nil {
			log.Printf("Failed to parse field %s: %v\n", fieldName, err)
			continue
		}
		model.Fields = append(model.Fields, f)
	}

	// Skip empty models
	if len(model.Fields) > 0 {
		g.models = append(g.models, model)
	}

	return false
}

func (g *Generator) Generate(packName string) {
	if packName == "" {
		packName = g.pkgName
	}

	// header docs: https://pkg.go.dev/cmd/go/internal/generate
	templ := template.Must(template.New("TemplateHeader").Parse(`
		// Code generated by ndn tlv codegen DO NOT EDIT.
		package {{.}}

		import (
			"bytes"
			"encoding/binary"
			"io"
			"strings"
			"time"

			enc "github.com/named-data/ndnd/std/encoding"
			"github.com/named-data/ndnd/std/ndn"
		)
	`))
	err := templ.Execute(&g.buf, packName)
	if err != nil {
		log.Fatal(err)
	}
	for _, m := range g.models {
		err = m.Generate(&g.buf)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Result returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) Result(filename string) []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}

	src2, err := imports.Process(filename, src, nil)
	if err != nil {
		log.Printf("warning: internal error: goimports failed to format code: %s", err)
		return src
	}

	return src2
}
