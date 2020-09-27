// Package main is complete tool for the go command line
package main

import (
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
	//vet := &complete.Command{
	//	Flags: map[string]complete.Predictor{
	//		"n": predict.Nothing,
	//		"x": predict.Nothing,
	//	},
	//	Args: anyGo,
	//}
	//
	//list := &complete.Command{
	//	Flags: map[string]complete.Predictor{
	//		"e":    predict.Nothing,
	//		"f":    predict.Something,
	//		"json": predict.Nothing,
	//	},
	//	Args: predict.Or(anyPackage, ellipsis),
	//}

	//doc := &complete.Command{
	//	Flags: map[string]complete.Predictor{
	//		"c":   predict.Nothing,
	//		"cmd": predict.Nothing,
	//		"u":   predict.Nothing,
	//	},
	//	Args: anyPackage,
	//}
	//
	//clean := &complete.Command{
	//	Args: predict.Or(anyPackage, ellipsis),
	//}
	//
	//env := &complete.Command{
	//	Args: predict.Something,
	//}
	//
	//fix := &complete.Command{
	//	Args: anyGo,
	//}

	gogo := &complete.Command{
		Sub: map[string]*complete.Command{
			"checkout": {},
			"status": {},
			"pull": {},
			"push": {},
			"stash": {},
			"commit": {},
			"add": {},
		},
	}

	gogo.Complete("bit")
}