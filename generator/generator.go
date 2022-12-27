package main

import (
	"bytes"
	"fmt"
	"github.com/jwasinger/evmmax-arith/templates"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"errors"
)

type TemplateParams struct {
	LimbCount int
	LimbBits  int
}

func loadTextFile(file_name string) string {
	content, err := ioutil.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	return string(content)
}

// from Bavard
func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

var funcs = template.FuncMap{
	"intRange": func(start, end int) []int {
		n := end - start
		result := make([]int, n)
		for i := 0; i < n; i++ {
			result[i] = start + i
		}
		return result
	},
	"add": func(val, v2 int) int {
		return val + v2
	},
	"sub": func(val, v2 int) int {
		return val - v2
	},
	"max": func(val, val2 int) int {
		if val > val2 {
			return val
		} else {
			return val2
		}
	},
	"gte": func(v1, v2 int) bool {
		return v1 >= v2
	},
	"mul": func(val, v2 int) int {
		return val * v2
	},
    "mulp1": func(val, v2 int) int {
        return (val + 1 ) * v2
    },
	// returns "[x]uint64 {0, 0, 0, ......, 0}"
	"makeZeroedLimbs": func(numLimbs int) string {
		result := fmt.Sprintf("[%d]uint64 {", numLimbs)
		return result + strings.Repeat(" 0,", numLimbs-1) + " 0}"
	},
	"dict": dict,
}

func aggregate(values []string) string {
	var sb strings.Builder
	for _, v := range values {
		sb.WriteString(v)
	}
	return sb.String()
}

var tmplDeps []string = []string{
	templates.GTE,
}

func prependDeps(tmplContent string) string {
	return aggregate([]string{aggregate(tmplDeps), tmplContent})
}

func buildTemplate(dest_path, template_path string, params *TemplateParams) {
	f, err := os.Create(dest_path)
	if err != nil {
		log.Println("create file: ", err)
		panic("")
	}

	template_content := loadTextFile(template_path)

	tmpl := template.Must(template.New("").Funcs(funcs).Parse(template_content))

	if err := tmpl.Execute(f, *params); err != nil {
		log.Fatal(err)
		panic("")
	}

	f.Close()
}

func genAddMod(addModType string, maxLimbs int) {
	headerTemplateContent := loadTextFile("templates/addmodsubmodheader.go.template")
	headerTemplate := template.Must(template.New("").Funcs(funcs).Parse(headerTemplateContent))

	params := TemplateParams{0, 64}
	buf := new(bytes.Buffer)

	f, err := os.Create(fmt.Sprintf("generated_addmod_%s.go", addModType))
	if err != nil {
		log.Fatal(err)
		panic("")
	}

	if err := headerTemplate.Execute(buf, params); err != nil {
		log.Fatal(err)
		panic("")
	}

	addModTemplateContent := loadTextFile(fmt.Sprintf("templates/addmod_%s.go.template", addModType))
	addModTemplate := template.Must(template.New("").Funcs(funcs).Parse(prependDeps(addModTemplateContent)))

	// TODO account for the implementations at 1 limb
	for i := 1; i <= maxLimbs; i++ {
		params = TemplateParams{i, 64}
		if err := addModTemplate.Execute(buf, params); err != nil {
			log.Fatal(err)
			panic("")
		}
	}

	if n, err := f.Write(buf.Bytes()); err != nil || n != len(buf.Bytes()) {
		panic(err)
	}
}

// TODO merge genSubMod/genAddMod into same function
func genSubMod(subModType string, maxLimbs int) {
	fmt.Println("submod start")
	headerTemplateContent := loadTextFile("templates/addmodsubmodheader.go.template")
	headerTemplate := template.Must(template.New("").Funcs(funcs).Parse(headerTemplateContent))

	params := TemplateParams{0, 64}
	buf := new(bytes.Buffer)

	f, err := os.Create(fmt.Sprintf("generated_submod_%s.go", subModType))
	if err != nil {
		log.Fatal(err)
		panic("")
	}

	if err := headerTemplate.Execute(buf, params); err != nil {
		log.Fatal(err)
		panic("")
	}

	subModTemplateContent := loadTextFile(fmt.Sprintf("templates/submod_%s.go.template", subModType))
	subModTemplate := template.Must(template.New("").Funcs(funcs).Parse(prependDeps(subModTemplateContent)))

	fmt.Println("submod loop")
	for i := 1; i <= maxLimbs; i++ {
		fmt.Println("iteration")
		params = TemplateParams{i, 64}
		if err := subModTemplate.Execute(buf, params); err != nil {
			log.Fatal(err)
			panic("")
		}
	}

	if n, err := f.Write(buf.Bytes()); err != nil || n != len(buf.Bytes()) {
		panic(err)
	}
}

func genMulMont(mulmontType string, maxLimbs int) {
	headerTemplateContent := loadTextFile("templates/mulmontheader.go.template")
	headerTemplate := template.Must(template.New("").Funcs(funcs).Parse(headerTemplateContent))

	params := TemplateParams{maxLimbs, 64}
	buf := new(bytes.Buffer)

	f, err := os.Create(fmt.Sprintf("generated_mulmont_%s.go", mulmontType))
	if err != nil {
		log.Fatal(err)
		panic("")
	}

	if err := headerTemplate.Execute(buf, params); err != nil {
		log.Fatal(err)
		panic("")
	}

	mulMontTemplateContent := loadTextFile(fmt.Sprintf("templates/mulmont_%s.go.template", mulmontType))
	mulMontTemplate := template.Must(template.New("").Funcs(funcs).Parse(prependDeps(mulMontTemplateContent)))

	for i := 1; i <= maxLimbs; i++ {
		params = TemplateParams{i, 64}
		if err := mulMontTemplate.Execute(buf, params); err != nil {
			log.Fatal(err)
			panic("")
		}
	}

	if n, err := f.Write(buf.Bytes()); err != nil || n != len(buf.Bytes()) {
		panic(err)
	}
}

func genInverse() {

}

func genPresets(maxLimbs int) {
	params := TemplateParams{maxLimbs, 64}
	buildTemplate("generated_presets.go", "templates/presets.go.template", &params)
}

func main() {
    // TODO document this (it's the point where we have definitively switched over to the generic algo):
    maxLimbs := 16
	genPresets(maxLimbs)
	genMulMont("nonunrolled", maxLimbs)
	genAddMod("nonunrolled", maxLimbs)
	genSubMod("nonunrolled", maxLimbs)
}
