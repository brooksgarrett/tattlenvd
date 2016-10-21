package tattlenvd

import "time"

// Nvd is array of Entry
type Nvd struct {
	Entries []Entry `xml:"entry"`
}

// Entry is Root Element
type Entry struct {
	CveID            string      `xml:"id,attr" json:"id"`
	PublishedDate    time.Time   `xml:"published-datetime"`
	LastModifiedDate time.Time   `xml:"last-modified-datetime"`
	Cvss             Cvss        `xml:"cvss>base_metrics" json:"cvss"`
	Products         []string    `xml:"vulnerable-software-list>product"` //CPE
	Summary          string      `xml:"summary"`
	References       []Reference `xml:"references"`
	Cwe              Cwe         `xml:"cwe"`
}

// Cvss is Cvss Score
type Cvss struct {
	Score                 string    `xml:"score"`
	AccessVector          string    `xml:"access-vector"`
	AccessComplexity      string    `xml:"access-complexity"`
	Authentication        string    `xml:"authentication"`
	ConfidentialityImpact string    `xml:"confidentiality-impact"`
	IntegrityImpact       string    `xml:"integrity-impact"`
	AvailabilityImpact    string    `xml:"availability-impact"`
	Source                string    `xml:"source"`
	GeneratedOnDate       time.Time `xml:"generated-on-datetime"`
}

// Cwe has Cwe ID
type Cwe struct {
	ID string `xml:"id,attr"`
}

// Cpe is the full CPE parsed from string form
type Cpe struct {
	Part      string
	Vendor    string
	Product   string
	Version   string
	Update    string
	Edition   string
	SWEdition string
	TargetSW  string
	TargetHW  string
	Language  string
	Other     string
}

// Reference is additional information about the CVE
type Reference struct {
	Type   string `xml:"reference_type,attr"`
	Source string `xml:"source"`
	Link   Link   `xml:"reference"`
}

// Link is additional information about the CVE
type Link struct {
	Value string `xml:",chardata" json:"value"`
	Href  string `xml:"href,attr" json:"href"`
}
