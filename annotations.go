// +build ignore

// annotations.go generates JSON annotations on struct and enum strings.
// go run annotations.go
package main

// Automated JSON tagging inspired by https://github.com/betacraft/easytags (MIT)

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

func walk(dir string, ext string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) == ext {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func main() {
	files, err := walk("**", ".go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	for _, f := range files {
		GenerateTags(f)
	}
}

// GenerateTags generates snake case json tags so that you won't need to write them. Can be also extended to xml or sql tags
func GenerateTags(fileName string) {
	fset := token.NewFileSet() // positions are relative to fset
	// Parse the file given in arguments
	f, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		fmt.Printf("Error parsing file %v", err)
		return
	}

	// range over the objects in the scope of this generated AST and check for StructType. Then range over fields
	// contained in that struct.

	ast.Inspect(f, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.StructType:
			processTags(t)
			return false
		}
		return true
	})

	// overwrite the file with modified version of ast.
	write, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error opening file %v", err)
		return
	}
	defer write.Close()
	w := bufio.NewWriter(write)
	err = format.Node(w, fset, f)
	if err != nil {
		fmt.Printf("Error formating file %s", err)
		return
	}
	w.Flush()
}

var existingTagReg = regexp.MustCompile("json:\"[^\"]+\"")

func parseTags(field *ast.Field) string {
	var tagValues []string
	fieldName := field.Names[0].String()
	var value string
	existingTag := existingTagReg.FindString(field.Tag.Value)
	if existingTag == "" {
		name := ToCamel(fieldName)
		value = fmt.Sprintf("json:\"%s\"", name)
		tagValues = append(tagValues, value)
	}
	updatedTags := strings.Fields(strings.Trim(field.Tag.Value, "`"))

	if len(tagValues) > 0 {
		updatedTags = append(updatedTags, tagValues...)
	}
	newValue := "`" + strings.Join(updatedTags, " ") + "`"

	return newValue
}

func processTags(x *ast.StructType) {
	for _, field := range x.Fields.List {
		if len(field.Names) == 0 {
			continue
		}
		if !unicode.IsUpper(rune(field.Names[0].String()[0])) {
			// not exported
			continue
		}

		if field.Tag == nil {
			field.Tag = &ast.BasicLit{}
			field.Tag.ValuePos = field.Type.Pos() + 1
			field.Tag.Kind = token.STRING
		}

		newTags := parseTags(field)
		field.Tag.Value = newTags
	}
}

// ToLowerCamel convert the given string to camelCase
func ToCamel(in string) string {
	runes := []rune(in)
	length := len(runes)

	var i int
	for i = 0; i < length; i++ {
		if unicode.IsLower(runes[i]) {
			break
		}
		runes[i] = unicode.ToLower(runes[i])
	}
	if i != 1 && i != length {
		i--
		runes[i] = unicode.ToUpper(runes[i])
	}
	return string(runes)
}
