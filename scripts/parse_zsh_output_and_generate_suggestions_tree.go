package scripts

import (
	"fmt"
	"github.com/chriswalz/complete/v3"
	"io/ioutil"
	"os"
	"strings"
)

func ParseZshAutocompleteOutput(path string) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
		return
	}
	s := string(b[:])
	//log.Print(s)
	sg := map[string]*complete.CompTree{}
	s = strings.ReplaceAll(s, "#######", "#####")
	s = strings.ReplaceAll(s, "######", "#####")
	parts := strings.Split(s, "#####")
	for i := 0; i < len(parts); i += 2 {
		argPart := parts[i]
		flagPart := parts[i+1]
		argPart = strings.TrimSpace(argPart)
		//fmt.Println(argPart)
		lines := strings.Split(argPart, "\n")
		command := strings.TrimSpace(strings.ReplaceAll(lines[0], "$ ", ""))
		command = strings.ReplaceAll(command, " --", "")
		command = strings.ReplaceAll(command, " -", "")
		//fmt.Println(command)
		var subsubs = map[string]*complete.CompTree{}
		if len(lines) > 1 {
			for _, subb := range lines[1:] {
				if strings.Contains(subb, "%backup%") || strings.Contains(subb, "Applications/") {
					break
				}
				subb = strings.TrimSpace(subb)
				ss := strings.Split(subb, "--")
				desc := strings.TrimSpace(ss[len(ss)-1])
				sub := strings.Fields(strings.TrimSpace(ss[0]))[0]
				subsubs[sub] = &complete.CompTree{
					Desc: desc,
				}
			}
		}
		sub := strings.ReplaceAll(command, "git ", "")

		flagPart = strings.TrimSpace(flagPart)
		lines = strings.Split(flagPart, "\n")

		var flags = map[string]*complete.CompTree{}
		if len(lines) > 1 {
			for _, subb := range lines[1:] {
				if strings.Contains(subb, "%backup%") || strings.Contains(subb, "Applications/") {
					break
				}
				subb = strings.TrimSpace(subb)
				lastDoubleDash := strings.LastIndex(subb, "--")
				ss := strings.Fields(subb[:lastDoubleDash])
				desc := subb[lastDoubleDash+3:]
				for _, flagg := range ss {
					if strings.HasPrefix(flagg, "-") {
						flagg = flagg[1:]
					}
					if strings.HasPrefix(flagg, "-") {
						flagg = flagg[1:]
					}
					flagName := strings.TrimSpace(flagg)
					if flagName == "" {
						continue
					}
					flags[flagName] = &complete.CompTree{
						Desc: desc,
					}
				}
				//ss := strings.Split(subb[2:], "--")
				//desc := strings.TrimSpace(ss[len(ss)-1])

			}
		}

		sg[sub] = &complete.CompTree{
			Desc:  "",
			Sub:   subsubs,
			Flags: flags,
			Args:  nil,
		}
	}
	//fmt.Println("hello")
	bittree := &complete.CompTree{
		Desc:  "",
		Sub:   sg,
		Flags: nil,
		Args:  nil,
	}
	codestring := printSuggestionTreeCode(bittree)
	codestring = `package cmd

import "github.com/chriswalz/complete/v3"

var b = ` + codestring
	f, err := os.Create("cmd/code_generated_src.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.WriteString(codestring)

}

func printSuggestionTreeCode(suggestionTree *complete.CompTree) string {
	flags := ""
	flagsStruct := ""
	if suggestionTree.Flags != nil {
		// {Desc: "quiet flag desc"}
		for flagName, v := range suggestionTree.Flags {
			if flagName == "\"" {
				continue
			}
			flags += "\"" + flagName + "\": {Desc: `" + v.Desc + "`},\n"
		}
		flagsStruct = `
Flags: map[string]*complete.CompTree{
` + flags + `
},
`
	}
	subs := ""
	if suggestionTree.Sub != nil {
		for key, val := range suggestionTree.Sub {
			subs += "\"" + key + "\": " + printSuggestionTreeCode(val)

		}
	}

	return `
&complete.CompTree{ Desc: "` + strings.ReplaceAll(suggestionTree.Desc, "`", "") + `", Sub: map[string]*complete.CompTree{` + subs + `}, ` + flagsStruct + `},
`

}
