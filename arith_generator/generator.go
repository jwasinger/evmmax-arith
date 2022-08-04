package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
    "github.com/jwasinger/mont-arith/templates"

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

func genAddMod(minLimbs, maxLimbs int) {
	headerTemplateContent := loadTextFile("templates/addmodsubmodheader.go.template")
	headerTemplate := template.Must(template.New("").Funcs(funcs).Parse(headerTemplateContent))

	params := TemplateParams{0, 64}
	buf := new(bytes.Buffer)

	f, err := os.Create("generated_addmod.go")
	if err != nil {
		log.Fatal(err)
		panic("")
	}

	if err := headerTemplate.Execute(buf, params); err != nil {
		log.Fatal(err)
		panic("")
	}

	addModNonUnrolledTemplateContent := loadTextFile("templates/addmod.go.template")
	addModNonUnrolledTemplate := template.Must(template.New("").Funcs(funcs).Parse(prependDeps(addModNonUnrolledTemplateContent)))

	for i := minLimbs; i <= maxLimbs; i++ {
		params = TemplateParams{i, 64}
		if err := addModNonUnrolledTemplate.Execute(buf, params); err != nil {
			log.Fatal(err)
			panic("")
		}
	}

	if n, err := f.Write(buf.Bytes()); err != nil || n != len(buf.Bytes()) {
		panic(err)
	}
}

func genSubMod(minLimbs, maxLimbs int) {
	headerTemplateContent := loadTextFile("templates/addmodsubmodheader.go.template")
	headerTemplate := template.Must(template.New("").Funcs(funcs).Parse(headerTemplateContent))

	params := TemplateParams{0, 64}
	buf := new(bytes.Buffer)

	f, err := os.Create("generated_submod.go")
	if err != nil {
		log.Fatal(err)
		panic("")
	}

	if err := headerTemplate.Execute(buf, params); err != nil {
		log.Fatal(err)
		panic("")
	}

	subModNonUnrolledTemplateContent := loadTextFile("templates/submod.go.template")
	subModNonUnrolledTemplate := template.Must(template.New("").Funcs(funcs).Parse(prependDeps(subModNonUnrolledTemplateContent)))

	for i := minLimbs; i <= maxLimbs; i++ {
		params = TemplateParams{i, 64}
		if err := subModNonUnrolledTemplate.Execute(buf, params); err != nil {
			log.Fatal(err)
			panic("")
		}
	}

	if n, err := f.Write(buf.Bytes()); err != nil || n != len(buf.Bytes()) {
		panic(err)
	}
}

func genMulMont(maxLimbs int) {
	headerTemplateContent := loadTextFile("templates/mulmontheader.go.template")
	headerTemplate := template.Must(template.New("").Funcs(funcs).Parse(headerTemplateContent))

	params := TemplateParams{maxLimbs, 64}
	buf := new(bytes.Buffer)

	f, err := os.Create("generated_mulmont.go")
	if err != nil {
		log.Fatal(err)
		panic("")
	}

	if err := headerTemplate.Execute(buf, params); err != nil {
		log.Fatal(err)
		panic("")
	}

	mulMontTemplateContent := loadTextFile("templates/mulmont.go.template")
	mulMontTemplate := template.Must(template.New("").Funcs(funcs).Parse(prependDeps(mulMontTemplateContent)))

	for i := 2; i <= maxLimbs; i++ {
		params = TemplateParams{i, 64}
		if err := mulMontTemplate.Execute(buf, params); err != nil {
			log.Fatal(err)
			panic("")
		}
	}

	if n, err := f.Write(buf.Bytes()); err != nil || n != len(buf.Bytes()) {
		panic(err)
	}

	mulMontGenericContent, err := ioutil.ReadFile("templates/mulmont_generic.go.template")
	if err != nil {
		panic(err)
	}

	if n, err := f.Write(mulMontGenericContent); err != nil || n != len(mulMontGenericContent) {
		panic(err)
	}
}

func genPresets(maxLimbs int) {
    params := TemplateParams{maxLimbs, 64}
    buildTemplate("generated_presets.go", "templates/presets.go.template", &params)
}

func main() {
	maxLimbs := 12
	genMulMont(maxLimbs)
    genPresets(maxLimbs)
    genAddMod(1, maxLimbs)
    genSubMod(1, maxLimbs)
}
