package template


import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

var templates []string
var patternGlob string

func visit(path string, f os.FileInfo, err error) error {
	status , err := os.Stat(path)
	if err != nil{
		return err
	}
	if !status.IsDir()&&strings.Contains(path, patternGlob){
		templates = append(templates,path)
	}
	return nil
}

// TemplateWalk is used to find all files. Parse files become template.
// root: template location
// pattern: Filename Extension
func TemplateWalk(root, pattern string) *template.Template {
	err := filepath.Walk(root, visit)
	patternGlob =pattern
	if err != nil {
		panic(err)
	}
	return template.Must(template.ParseFiles(templates...))
}
