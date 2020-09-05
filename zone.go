package zone

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var locUTC *time.Location

func init() {
	locUTC, _ = time.LoadLocation("UTC")
}

func OffsetToUTC(loc *time.Location) time.Duration {
	locT := time.Date(2006, 1, 2, 15, 4, 5, 0, loc)
	utcT := time.Date(2006, 1, 2, 15, 4, 5, 0, locUTC)
	return utcT.Sub(locT).Round(time.Minute)
}

func Locations() ([]*time.Location, error) {
	cmd := exec.Command("go", "env", "-json")
	b, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	vs := map[string]interface{}{}
	err = json.Unmarshal(b, &vs)
	if err != nil {
		return nil, err
	}
	grv, _ := vs["GOROOT"]
	gr, _ := grv.(string)

	fmt.Printf("found GOROOT: %q\n", gr)
	zi := filepath.Join(gr, "lib", "time", "zoneinfo.zip")
	ziF, err := zip.OpenReader(zi)
	if err != nil {
		return nil, errors.Wrapf(err, "open-zip-reader %q", zi)
	}
	defer ziF.Close()

	locs := []*time.Location{}
	for _, f := range ziF.File {
		if !strings.HasSuffix(f.Name, "/") {
			loc, err := time.LoadLocation(f.Name)
			if err != nil {
				fmt.Printf("ERROR locading location %q: %v\n", f.Name, err)
				continue
			}
			locs = append(locs, loc)
		}
	}
	return locs, nil
}
