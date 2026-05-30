// wave-certify — WAVE Certified hardware validation CLI.
//
// Runs a protocol-correctness battery against a target (NDI/Dante/SRT/MoQ),
// emits a signed certification artifact. Scaffold stage; batteries are
// Wave-1 work tracked in the repo roadmap.
package main

import (
	"errors"
	"fmt"
	"os"
)

const usage = `wave-certify — WAVE Certified hardware validation

Usage:
  wave-certify check --target <addr> --protocol <ndi|dante|srt|moq>
  wave-certify submit <artifact.json>
  wave-certify version

`

var errNotImplemented = errors.New("wave-certify: battery not yet implemented")

func main() {
	if len(os.Args) < 2 {
		fmt.Print(usage)
		os.Exit(2)
	}
	switch os.Args[1] {
	case "check":
		fmt.Fprintln(os.Stderr, errNotImplemented)
		os.Exit(1)
	case "submit":
		fmt.Fprintln(os.Stderr, errNotImplemented)
		os.Exit(1)
	case "version":
		fmt.Println("wave-certify 0.0.0-scaffold")
	default:
		fmt.Print(usage)
		os.Exit(2)
	}
}
