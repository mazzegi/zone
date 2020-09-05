package zone

import (
	"fmt"
	"os"
)

func Generate(fileName string, pkgName string, locsVarName string) error {
	locs, err := Locations()
	if err != nil {
		return err
	}
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, "package %s\n\n", pkgName)
	fmt.Fprintf(f, "var %s = []string{\n", locsVarName)
	for _, loc := range locs {
		fmt.Fprintf(f, "    \"%s\",\n", loc)
	}
	fmt.Fprintf(f, "}")
	return nil
}
