package tattlenvd

import (
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/parnurzeal/gorequest"
)

func fetchFeedFile(url string) (nvd Nvd, err error) {
	var body string
	var errs []error
	var resp *http.Response

	resp, body, errs = gorequest.New().Get(url).End()
	//  defer resp.Body.Close()
	if len(errs) > 0 || resp == nil || resp.StatusCode != 200 {
		return nvd, fmt.Errorf(
			"HTTP error. errs: %v, url: %s", errs, url)
	}

	b := bytes.NewBufferString(body)
	reader, err := gzip.NewReader(b)
	defer reader.Close()
	if err != nil {
		return nvd, fmt.Errorf(
			"Failed to decompress NVD feedfile. url: %s, err: %s", url, err)
	}

	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nvd, fmt.Errorf(
			"Failed to Read NVD feedfile. url: %s, err: %s", url, err)
	}

	if err = xml.Unmarshal(bytes, &nvd); err != nil {
		return nvd, fmt.Errorf(
			"Failed to unmarshal. url: %s, err: %s", url, err)
	}
	return nvd, nil
}

// Interesting determines if a given Entry is "Interesting" to the examiner
func Interesting(entry *Entry) bool {
	interestingCpes := RefreshInterestingCpes()

	for _, cpeString := range entry.Products {
		cpe := ParseCpe(cpeString)
		Trace.Printf("%s, %s: %s", cpe.Vendor, cpe.Product, cpeString)
		for _, interesting := range interestingCpes {
			if cpe.Match(&interesting) {
				return true
			}
		}
	}
	return false
}

// ParseCpe takes a Cpe string (likely from "products" in NVD CVEs) and returns a CPE struct
func ParseCpe(cpeString string) Cpe {
	components := strings.Split(cpeString, ":")
	Trace.Printf("%d: %s", len(components), cpeString)
	slice := StrSlice(strings.Split(cpeString, ":"))
	return Cpe{slice.Get(1), slice.Get(2), slice.Get(3), slice.Get(4), slice.Get(5), slice.Get(6), slice.Get(7), slice.Get(8), slice.Get(9), slice.Get(10), slice.Get(11)}
}

// Match will check if a CPE is a match
func (cpeMatch Cpe) Match(cpe *Cpe) bool {
	if (cpe.Part != cpeMatch.Part) ||
		cpe.Vendor != cpeMatch.Vendor {
		Trace.Printf("Failed at Vendor |%s|%s|", cpe.Vendor, cpeMatch.Vendor)
		return false
	}
	// Treat empty strings as ANY
	if (cpe.Product != cpeMatch.Product) &&
		(cpe.Product != "" && cpeMatch.Product != "") {
		Trace.Printf("Failed at Product |%s|%s|", cpe.Product, cpeMatch.Product)
		return false
	}
	if (cpe.Version != cpeMatch.Version) &&
		(cpe.Version != "" && cpeMatch.Version != "") {
		return false
	}
	if (cpe.Update != cpeMatch.Update) &&
		(cpe.Update != "" && cpeMatch.Update != "") {
		return false
	}
	if (cpe.Edition != cpeMatch.Edition) &&
		(cpe.Edition != "" && cpeMatch.Edition != "") {
		return false
	}
	if (cpe.SWEdition != cpeMatch.SWEdition) &&
		(cpe.SWEdition != "" && cpeMatch.SWEdition != "") {
		return false
	}
	if (cpe.TargetHW != cpeMatch.TargetHW) &&
		(cpe.TargetHW != "" && cpeMatch.TargetHW != "") {
		return false
	}
	if (cpe.TargetSW != cpeMatch.TargetSW) &&
		(cpe.TargetSW != "" && cpeMatch.TargetSW != "") {
		return false
	}
	if (cpe.Language != cpeMatch.Language) &&
		(cpe.Language != "" && cpeMatch.Language != "") {
		return false
	}
	if (cpe.Other != cpeMatch.Other) &&
		(cpe.Other != "" && cpeMatch.Other != "") {
		return false
	}

	return true

}

// StrSlice is an array of strings
type StrSlice []string

// Get provides the setting from a char split string or otherwise the ANY key
func (s StrSlice) Get(i int) string {
	if i >= 0 && i < len(s) {
		return s[i]
	}
	return ""
}
