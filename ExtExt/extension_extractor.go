package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Entry struct {
	name    string
	include string
}

type Catalog map[string]([]Entry)

var templateInclude = `	"bitbucket.org/ntti3/esi-wan-gohan/[includePath]"
`
var templateInit = "	[extName].Init(env)\n"

var templateDoc = `
package main

import (
	"bitbucket.org/ntti3/esi-wan-gohan/common/extensions/goodies"
	"bitbucket.org/ntti3/esi-wan-gohan/esi/resources"
	"github.com/cloudwan/gohan/extension/goext"
[includes]
)

func Init(env goext.IEnvironment) error {
[inits]
}
`

func createNameAndInclude(yamlPath, url string) (name, include string, err error) {
	start := strings.LastIndex(url, "/") + 1
	startIncl := strings.Index(url, "./") + 1
	if startIncl == 0 {
		startIncl = strings.Index(url, "//") + 1
	}
	end := strings.LastIndex(url, ".")
	endUrl := strings.LastIndex(yamlPath, "/")
	endIncl := strings.LastIndex(url, "/")

	if start == -1 || startIncl == -1 || end == -1 || endIncl == -1 {
		return "", "", fmt.Errorf("Could not extract name or include path")
	}
	if endIncl < startIncl {
		endIncl = startIncl
	}
	name = url[start:end]
	include = yamlPath[:endUrl] + url[startIncl:endIncl]
	return
}

func extractExtensions(s string, cat Catalog) {
	inputContent, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Println("Could not open file ", s, " ", err)
		return
	}
	data := map[interface{}]interface{}{}
	err = yaml.Unmarshal(inputContent, &data)
	if err != nil {
		fmt.Printf("Could not unmarshall %s %v\n", s, err)
		return
	}
	extensions, ok := data["extensions"].([]interface{})
	if !ok {
		fmt.Println("No extensions section in yaml")
		return
	}

	for i := range extensions {
		fmt.Printf("ext: %v\n", extensions[i])

		path := extensions[i].(map[interface{}]interface{})["path"].(string)
		url := extensions[i].(map[interface{}]interface{})["url"].(string)
		cType := extensions[i].(map[interface{}]interface{})["code_type"].(string)
		if cType != "goext" {
			continue
		}

		name, include, err := createNameAndInclude(s, url)
		if err != nil {
			fmt.Println("Problem during parsing extension entry ", s, " ", url, " ", err)
			continue
		}

		curr, ok := cat[path]
		if ok {
			cat[path] = append(curr, Entry{name, include})
		} else {
			cat[path] = []Entry{Entry{name, include}}
		}
	}
}

func parseSchemaFiles(schemaFiles []interface{}) Catalog {
	fmt.Printf("Schema files:\n%v\n\n", schemaFiles)

	cat := make(Catalog)
	for i := range schemaFiles {
		s := schemaFiles[i]
		fmt.Printf("\nParsing schema file: %d : %s\n", i, s)
		extractExtensions(s.(string), cat)
	}

	return cat
}

func parseInput(input string) (res Catalog) {
	inputContent, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println("Could not open file ", input, " ", err)
		return nil
	}
	data := map[interface{}]interface{}{}
	err = yaml.Unmarshal(inputContent, &data)
	if err != nil {
		fmt.Printf("Could not unmarshall %s %v\n", input, err)
		return nil
	}
	schemaFiles, ok := data["schemas"].([]interface{})
	if !ok {
		fmt.Println("No schemas section in file")
		return nil
	}

	return parseSchemaFiles(schemaFiles)
}

func createOutputFiles(cat Catalog, outDir string) {
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		fmt.Println("Creating output directory...")
		os.Mkdir(outDir, os.ModePerm)
	}
	for key := range cat {
		if key == "" {
			key = "empty"
		}
		file := strings.Replace(key, ".", "_", -1)
		file = strings.Replace(file, "/", "_", -1)
		file = outDir + "/" + file + ".go"
		f, err := os.Create(file)
		if err != nil {
			fmt.Println("Failed creating file ", file, err)
			continue
		}

		if key != "empty" {
			printFile(append(cat[""], cat[key]...), f)

		} else {
			printFile(cat[""], f)
		}
		defer f.Close()
	}
}

func printFile(entries []Entry, file *os.File) {
	doc := templateDoc
	includes := createIncludes(entries)
	inits := createInits(entries)

	doc = strings.Replace(doc, "[includes]", includes, 1)
	doc = strings.Replace(doc, "[inits]", inits, 1)
	file.WriteString(doc)
}

func createIncludes(entries []Entry) (res string) {
	for i := range entries {
		init := strings.Replace(templateInclude, "[includePath]", entries[i].include, 1)
		res = res + init
	}
	return
}

func createInits(entries []Entry) (res string) {
	for i := range entries {
		init := strings.Replace(templateInit, "[extName]", entries[i].name, 1)
		res = res + init
	}
	return
}

func main() {

	params := os.Args
	if len(params) < 3 {
		fmt.Println("Please specify inputFile and outputDir. Like: ./extension_extractor in.yaml out\n")
		return
	}

	res := parseInput(params[1])
	if res == nil {
		fmt.Println("No input parsed, exiting")
		return
	}
	fmt.Printf("\nOut:\n%v\n\n", res)
	createOutputFiles(res, params[2])
}
