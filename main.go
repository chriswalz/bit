/*
Copyright © 2020 Chris Walz <walz@reconbuddy.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	bitcmd "github.com/chriswalz/bit/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

// this should be overwritten at compile time
var version string = "v0.8.0"

func main() {
	bitcmd.Bitcomplete()

	// defer needed to handle funkyness with CTRL + C & go-prompt
	defer bitcmd.HandleExit()
	bitcmd.BitCmd.Version = version

	// set debug level
	log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	argsWithoutProg := os.Args[1:]

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	debugIndex := bitcmd.Find(argsWithoutProg, "--debug")
	if debugIndex != -1 {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		argsWithoutProg = append(argsWithoutProg[:debugIndex], argsWithoutProg[debugIndex+1:]...)
	}

	// verify is git repo
	if len(os.Args) >= 2 {
		if os.Args[1] == "--version" || os.Args[1] == "version" {
			fmt.Println("bit version " + version)
			bitcmd.PrintGitVersion()
			return
		}
	}
	if !bitcmd.IsGitRepo() {
		if len(os.Args) >= 2 && (os.Args[1] == "update" || os.Args[1] == "clone" || os.Args[1] == "complete" || os.Args[1] == "init" || os.Args[1] == "send-email" || os.Args[1] == "svn" || os.Args[1] == "config" || os.Args[1] == "bugreport") {
			// do nothing here, proceed to update path
		} else {
			fmt.Println("fatal: not a git repository (or any of the parent directories): .git")
			return
		}
	}

	bitcliCmds := []string{"save", "sync", "help", "info", "release", "update", "pr", "complete", "gitmoji"}
	if len(argsWithoutProg) == 0 || bitcmd.Find(bitcliCmds, argsWithoutProg[0]) != -1 {
		bitcli()
	} else {
		completerSuggestionMap, _ := bitcmd.CreateSuggestionMap(bitcmd.BitCmd)
		yes := bitcmd.HijackGitCommandOccurred(argsWithoutProg, completerSuggestionMap, version)
		if yes {
			return
		}
		bitcmd.RunGitCommandWithArgs(argsWithoutProg)
	}
}

func bitcli() {
	bitcmd.Execute()
}
