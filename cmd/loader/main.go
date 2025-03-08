package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/ryota1119/gin_webapi/internal/domain"
)

func main() {
	stmts, err := gormschema.New("mysql").Load(&domain.Post{}, &domain.User{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	_, err = io.WriteString(os.Stdout, stmts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to write output: %v\n", err)
		os.Exit(1)
	}
}
