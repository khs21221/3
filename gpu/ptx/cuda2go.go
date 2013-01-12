// +build ignore

/*
 This program generates Go wrappers for cuda sources.
 The cuda file should contain exactly one __global__ void.
*/
package main

import (
	"code.google.com/p/mx3/core"
	"flag"
	"os"
	"text/scanner"
	"text/template"
)

func main() {
	flag.Parse()
	for _, fname := range flag.Args() {
		cuda2go(fname)
	}
}

// generate cuda wrapper for file.
func cuda2go(fname string) {
	// open cuda file
	f, err := os.Open(fname)
	core.Fatal(err)
	defer f.Close()

	// read tokens
	var token []string
	var s scanner.Scanner
	s.Init(f)
	tok := s.Scan()
	for tok != scanner.EOF {
		if !filter(s.TokenText()) {
			token = append(token, s.TokenText())
		}
		tok = s.Scan()
	}

	// find function name and arguments
	funcname := ""
	argstart, argstop := -1, -1
	for i := 3; i < len(token); i++ {
		if token[i] == "__global__" {
			funcname = token[i+2]
			argstart = i + 4
		}
		if argstart > 0 && token[i] == ")" {
			argstop = i + 1
			break
		}
	}
	argl := token[argstart:argstop]

	// isolate individual arguments
	var args [][]string
	start := 0
	for i, a := range argl {
		if a == "," || a == ")" {
			args = append(args, argl[start:i])
			start = i + 1
		}
	}

	// separate arg names/types and make pointers Go-style
	argn := make([]string, len(args))
	argt := make([]string, len(args))
	for i := range args {
		if args[i][1] == "*" {
			args[i] = []string{args[i][0] + "*", args[i][2]}
		}
		argt[i] = typemap(args[i][0])
		argn[i] = args[i][1]
	}
	wrapgen(fname, funcname, argt, argn)
}

var tm = map[string]string{"float*": "cu.DevicePtr", "float": "float32", "int": "int"}

// translate C type to Go type.
func typemap(ctype string) string {
	if gotype, ok := tm[ctype]; ok {
		return gotype
	}
	core.Fatalf("cuda2go: unsupported cuda type: %v", ctype)
	return "" // unreachable
}

// template data
type Kernel struct {
	Name string
	ArgT []string
	ArgN []string
}

//func (k *Kernel) Args() string {
//	var a bytes.Buffer
//	for _, arg := range k.args {
//		fmt.Fprint(&a, arg[1], " ", arg[0], ", ")
//	}
//	return a.String()
//}

//func (k *Kernel) NArg() int { return len(k.args) }

// generate wrapper code from template
func wrapgen(filename, funcname string, argt, argn []string) {
	//fmt.Println("wrapgen", filename, funcname, args)
	kernel := &Kernel{funcname, argt, argn}
	core.Fatal(templ.Execute(os.Stdout, kernel))
}

// wrapper code template text
const templText = `
package gpu

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import(
	"unsafe"
	"github.com/barnex/cuda5/cu"
	"sync"
)

var( 
	{{.Name}}_stream = cu.StreamCreate() // only after init?
	{{range $i, $n := .ArgN}} {{$.Name}}_arg_{{$n}}
	{{end}} 
	argp_{{.Name}} [{{len .ArgN}}]unsafe.Pointer
	lock_{{.Name}} sync.Mutex
)

// CUDA kernel wrapper for {{.Name}}.
// The kernel is launched in a separate stream so that it can be parallel with memcpys etc.
// The stream is synchronized before this call returns.
func K_{{.Name}} ( {{range $i, $t := .ArgT}}{{$t}}, {{end}} gridDim, blockDim cu.Dim3) {
	//lock_{{.Name}}.Lock()

	cu.LaunchKernel(code, gridDim.X, gridDim.Y, gridDim.Z, blockDim.X, blockDim.Y, blockDim.Z, 0, stream_{{.Name}}, args)
	stream_{{.Name}}.Synchronize()
	lock_{{.Name}}.Unlock()
}
`

// wrapper code template
var templ = template.Must(template.New("wrap").Parse(templText))

// should token be filtered out of stream?
func filter(token string) bool {
	switch token {
	case "__restrict__":
		return true
	}
	return false
}
