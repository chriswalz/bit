// Package main is complete tool for the go command line
package main

import (
	"github.com/chriswalz/bit/cmd"
	"github.com/posener/complete/v2"
	"github.com/posener/complete/v2/predict"
)

var (
	ellipsis   = predict.Set{"./..."}
	//anyPackage = complete.PredictFunc(predictPackages)
	goFiles    = predict.Files("*.go")
	anyFile    = predict.Files("*")
	//anyGo      = predict.Or(goFiles, anyPackage, ellipsis)
)

func main() {
	//build := &complete.Command{
	//	Flags: map[string]complete.Predictor{
	//		"o": anyFile,
	//		"i": predict.Nothing,
	//
	//		"a":             predict.Nothing,
	//		"n":             predict.Nothing,
	//		"p":             predict.Something,
	//		"race":          predict.Nothing,
	//		"msan":          predict.Nothing,
	//		"v":             predict.Nothing,
	//		"work":          predict.Nothing,
	//		"x":             predict.Nothing,
	//		"asmflags":      predict.Something,
	//		"buildmode":     predict.Something,
	//		"compiler":      predict.Something,
	//		"gccgoflags":    predict.Set{"gccgo", "gc"},
	//		"gcflags":       predict.Something,
	//		"installsuffix": predict.Something,
	//		"ldflags":       predict.Something,
	//		"linkshared":    predict.Nothing,
	//		"pkgdir":        anyPackage,
	//		"tags":          predict.Something,
	//		"toolexec":      predict.Something,
	//	},
	//	Args: anyGo,
	//}
	//
	//run := &complete.Command{
	//	Flags: map[string]complete.Predictor{
	//		"exec": predict.Something,
	//	},
	//	Args: goFiles,
	//}
	//

	branchCompletion := &complete.Command{
		Args: complete.PredictFunc(func(prefix string) []string{
			branches := cmd.BranchListSuggestions()
			completion := make([]string, len(branches))
			for i, v := range branches {
				completion[i] = v.Text
			}
			return completion
		}),
	}

	cmds := cmd.AllBitAndGitSubCommands(cmd.ShellCmd)
	completionSubCmdMap := map[string]*complete.Command{}
	for _, v := range cmds {
		completionSubCmdMap[v.Name()] = &complete.Command{}
		if v.Name() == "checkout" || v.Name() == "co" || v.Name() == "switch" {
			completionSubCmdMap[v.Name()] = branchCompletion
		}
	}

	gogo := &complete.Command{
		Sub: completionSubCmdMap,
	}

	gogo.Complete("bit")
}