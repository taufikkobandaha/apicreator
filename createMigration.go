package main

import (
	"log"
	"os"
	"text/template"

	helper "github.com/shouva/dailyhelper"
)

func createMigrations(models []string) string {
	strtemplate := `
	package main
	func migrate(g *gorm.DB) {
		{{range $index, $model := .}}g.AutoMigrate(&{{$model}}{})
		{{end}}}
	`
	tmpl := template.New("create api template")
	tmpl, err := tmpl.Parse(strtemplate)
	if err != nil {
		log.Fatal("Parse: ", err)
		return ""
	}

	// openfile
	filename := helper.GetCurrentPath(false) + "/out/migrations.go"
	f, err := os.Create(filename)
	if err != nil {
		log.Println("create file: ", err)
		return ""
	}

	// var strout string
	err = tmpl.Execute(f, models)

	if err != nil {
		log.Fatal("Execute: ", err)
		return ""
	}
	f.Close()
	completer(filename)
	return ""
}
