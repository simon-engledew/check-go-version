package main

import (
	"log"
	"os"

	"golang.org/x/mod/modfile"
)

func main() {
	log.SetFlags(0)
	version, ok := os.LookupEnv("GOLANG_VERSION")
	if !ok {
		log.Fatal("GOLANG_VERSION not set")
	}
	data, err := os.ReadFile("go.mod")
	if err != nil {
		log.Fatal(err)
	}
	parsed, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		log.Fatal(err)
	}
	if parsed.Go.Version != version {
		log.Fatalf(`Go version mismatch (%q != %q), please run:

  go mod edit -go=%q

`, parsed.Go.Version, version, version)
	}
}
