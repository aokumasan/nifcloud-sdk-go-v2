// +build codegen

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
)

func main() {
	var svcPath string

	flag.StringVar(&svcPath, "path", "service",
		"The `path` to service clients.",
	)
	flag.Parse()

	svcPath = filepath.FromSlash(svcPath)

	filepath.Walk(svcPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		filename := filepath.Base(path)
		if isAPIFile(filename) {
			if err := rewriteAPIFile(path); err != nil {
				panic(err)
			}
		} else if isServiceFile(filename) {
			if err := rewriteServiceFile(path); err != nil {
				panic(err)
			}
		} else if isInterfaceFile(filename) {
			if err := rewriteInterfaceFile(path); err != nil {
				panic(err)
			}
		}

		return nil
	})
}

func isAPIFile(filename string) bool {
	return strings.HasPrefix(filename, "api_op") || filename == "api_enums.go" || filename == "api_types.go"
}

func isServiceFile(filename string) bool {
	return filename == "api_client.go"
}

func isInterfaceFile(filename string) bool {
	return filename == "interface.go"
}

func rewrite(path string, imports []map[string]string, replaces []map[string]string) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	for _, imp := range imports {
		for org, replace := range imp {
			if org == "" {
				_ = astutil.AddImport(fset, f, replace)
			} else {
				_ = astutil.RewriteImport(fset, f, org, replace)
			}
		}
	}

	buf := &bytes.Buffer{}
	pp := &printer.Config{Tabwidth: 8, Mode: printer.UseSpaces | printer.TabIndent}
	pp.Fprint(buf, fset, f)

	code := buf.String()
	for _, rep := range replaces {
		for org, replace := range rep {
			code = strings.Replace(code, org, replace, -1)
		}
	}

	return ioutil.WriteFile(path, []byte(code), 0644)
}

func rewriteAPIFile(path string) error {
	imports := []map[string]string{
		{"github.com/aws/aws-sdk-go-v2/internal/awsutil": "github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"},
	}
	replaces := []map[string]string{
		{"awsutil": "nifcloudutil"},
	}

	return rewrite(path, imports, replaces)
}

func rewriteServiceFile(path string) error {
	imports := []map[string]string{
		{"": "github.com/aokumasan/nifcloud-sdk-go-v2/nifcloud"},
		{"github.com/aws/aws-sdk-go-v2/private/protocol/computing": "github.com/aokumasan/nifcloud-sdk-go-v2/private/protocol/computing"},
		{"github.com/aws/aws-sdk-go-v2/private/protocol/script": "github.com/aokumasan/nifcloud-sdk-go-v2/private/protocol/script"},
		{"github.com/aws/aws-sdk-go-v2/aws/signer/v2": "github.com/aokumasan/nifcloud-sdk-go-v2/nifcloud/signer/v2"},
	}
	replaces := []map[string]string{
		{"aws.Config": "nifcloud.Config"},
		{"config,": "config.AWSConfig(),"},
	}

	return rewrite(path, imports, replaces)
}

func rewriteInterfaceFile(path string) error {
	serviceName := filepath.Base(filepath.Dir(filepath.Dir(path)))
	imports := []map[string]string{
		{"github.com/aws/aws-sdk-go-v2/aws": "github.com/aokumasan/nifcloud-sdk-go-v2/nifcloud"},
		{fmt.Sprintf("github.com/aws/aws-sdk-go-v2/service/%s", serviceName): fmt.Sprintf("github.com/aokumasan/nifcloud-sdk-go-v2/service/%s", serviceName)},
	}

	return rewrite(path, imports, nil)
}
