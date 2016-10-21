package tattlenvd

import (
	"fmt"
	"io/ioutil"
)

// RefreshInterestingCpes will pull in all the interesting products we want to monitor
func RefreshInterestingCpes() []Cpe {
	var interesting = []Cpe{
		Cpe{
			Part:      "/a",
			Vendor:    "microsoft",
			Product:   "server",
			Version:   "2008",
			Update:    "",
			Edition:   "",
			SWEdition: "",
			TargetSW:  "",
			TargetHW:  "",
			Language:  "",
			Other:     ""},
		Cpe{
			Part:      "/a",
			Vendor:    "microsoft",
			Product:   "server",
			Version:   "2012",
			Update:    "",
			Edition:   "",
			SWEdition: "",
			TargetSW:  "",
			TargetHW:  "",
			Language:  "",
			Other:     ""},
		Cpe{
			Part:      "/a",
			Vendor:    "microsoft",
			Product:   "sql_server",
			Version:   "2008",
			Update:    "r2_sp2",
			Edition:   "",
			SWEdition: "",
			TargetSW:  "",
			TargetHW:  "",
			Language:  "",
			Other:     ""},
		Cpe{
			Part:      "/a",
			Vendor:    "microsoft",
			Product:   "iis",
			Version:   "8.0",
			Update:    "",
			Edition:   "",
			SWEdition: "",
			TargetSW:  "",
			TargetHW:  "",
			Language:  "",
			Other:     ""},
		Cpe{
			Part:      "/a",
			Vendor:    "hp",
			Product:   "webinspect",
			Version:   "",
			Update:    "",
			Edition:   "",
			SWEdition: "",
			TargetSW:  "",
			TargetHW:  "",
			Language:  "",
			Other:     ""},
		Cpe{
			Part:      "/a",
			Vendor:    "microsoft",
			Product:   "iis",
			Version:   "7.5",
			Update:    "",
			Edition:   "",
			SWEdition: "",
			TargetSW:  "",
			TargetHW:  "",
			Language:  "",
			Other:     ""},
		Cpe{
			Part:      "/o",
			Vendor:    "checkpoint",
			Product:   "gaia_os",
			Version:   "r77.0",
			Update:    "",
			Edition:   "",
			SWEdition: "",
			TargetSW:  "",
			TargetHW:  "",
			Language:  "",
			Other:     ""},
	}

	return interesting
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readConfig(configPath string) {
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

}
