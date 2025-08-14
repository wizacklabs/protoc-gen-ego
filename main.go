package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

const genGoDocURL = "https://protobuf.dev/reference/go/go-generated"
const grpcDocURL = "https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code"

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		showVersion()
		os.Exit(0)
	}

	if len(os.Args) == 2 && os.Args[1] == "--help" {
		fmt.Fprintf(os.Stdout, "See "+genGoDocURL+" for usage information.\n")
		os.Exit(0)
	}

	var (
		flags flag.FlagSet
		_     = flags.String("plugins", "", "deprecated option")
		_     = flags.String("enum", "", "enum constants generate option")
		//importPrefix = flags.String("import_prefix", "", "prefix to prepend to import paths")
		//experimentalStripNonFunctionalCodegen = flags.Bool("experimental_strip_nonfunctional_codegen", false, "experimental_strip_nonfunctional_codegen true means that the plugin will not emit certain parts of the generated code in order to make it possible to compare a proto2/proto3 file with its equivalent (according to proto spec) editions file. Primarily, this is the encoded descriptor.")
	)

	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-plugins" || os.Args[i] == "--plugins" {
			if i < len(os.Args)-1 && strings.Contains(strings.ToLower(os.Args[i+1]), "grpc") {
				fmt.Fprintf(os.Stderr, "protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC\n\nSee %s for more information.\n", grpcDocURL)
				os.Exit(1)
			}
		}
	}

	opts := &protogen.Options{
		ParamFunc: flags.Set,
	}

	err := run(os.Stdin, opts)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", filepath.Base(os.Args[0]), err)
		os.Exit(1)
	}
}
