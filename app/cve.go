package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)


// generated structs
type CVEDataMeta struct {
	ID       string `json:"ID"`
	ASSIGNER string `json:"ASSIGNER"`
}
type VersionData struct {
	VersionValue    string `json:"version_value"`
	VersionAffected string `json:"version_affected"`
}
type Version struct {
	VersionData []VersionData `json:"version_data"`
}
type ProductData struct {
	ProductName string  `json:"product_name"`
	Version     Version `json:"version"`
}
type Product struct {
	ProductData []ProductData `json:"product_data"`
}
type VendorData struct {
	VendorName string  `json:"vendor_name"`
	Product    Product `json:"product"`
}
type Vendor struct {
	VendorData []VendorData `json:"vendor_data"`
}
type Affects struct {
	Vendor Vendor `json:"vendor"`
}
type PTDDescription struct {
	Lang  string `json:"lang"`
	Value string `json:"value"`
}
type ProblemTypeData struct {
	Description []PTDDescription `json:"description"`
}
type ProblemType struct {
	ProblemtypeData []ProblemTypeData `json:"problemtype_data"`
}
type ReferenceData struct {
	URL       string   `json:"url"`
	Name      string   `json:"name"`
	Refsource string   `json:"refsource"`
	Tags      []string `json:"tags"`
}
type References struct {
	ReferenceData []ReferenceData `json:"reference_data"`
}
type DescriptionData struct {
	Lang  string `json:"lang"`
	Value string `json:"value"`
}
type Description struct {
	DescriptionData []DescriptionData `json:"description_data"`
}
type CpeMatch struct {
	Vulnerable bool   `json:"vulnerable"`
	Cpe23URI   string `json:"cpe23Uri"`
}
type Nodes struct {
	Operator string     `json:"operator"`
	CpeMatch []CpeMatch `json:"cpe_match"`
}
type Configurations struct {
	CVEDataVersion string  `json:"CVE_data_version"`
	Nodes          []Nodes `json:"nodes"`
}
type CvssV3 struct {
	Version               string  `json:"version"`
	VectorString          string  `json:"vectorString"`
	AttackVector          string  `json:"attackVector"`
	AttackComplexity      string  `json:"attackComplexity"`
	PrivilegesRequired    string  `json:"privilegesRequired"`
	UserInteraction       string  `json:"userInteraction"`
	Scope                 string  `json:"scope"`
	ConfidentialityImpact string  `json:"confidentialityImpact"`
	IntegrityImpact       string  `json:"integrityImpact"`
	AvailabilityImpact    string  `json:"availabilityImpact"`
	BaseScore             float64 `json:"baseScore"`
	BaseSeverity          string  `json:"baseSeverity"`
}
type BaseMetricV3 struct {
	CvssV3              CvssV3  `json:"cvssV3"`
	ExploitabilityScore float64 `json:"exploitabilityScore"`
	ImpactScore         float64 `json:"impactScore"`
}
type CvssV2 struct {
	Version               string  `json:"version"`
	VectorString          string  `json:"vectorString"`
	AccessVector          string  `json:"accessVector"`
	AccessComplexity      string  `json:"accessComplexity"`
	Authentication        string  `json:"authentication"`
	ConfidentialityImpact string  `json:"confidentialityImpact"`
	IntegrityImpact       string  `json:"integrityImpact"`
	AvailabilityImpact    string  `json:"availabilityImpact"`
	BaseScore             float64 `json:"baseScore"`
}
type BaseMetricV2 struct {
	CvssV2                  CvssV2  `json:"cvssV2"`
	Severity                string  `json:"severity"`
	ExploitabilityScore     float64 `json:"exploitabilityScore"`
	ImpactScore             float64 `json:"impactScore"`
	AcInsufInfo             bool    `json:"acInsufInfo"`
	ObtainAllPrivilege      bool    `json:"obtainAllPrivilege"`
	ObtainUserPrivilege     bool    `json:"obtainUserPrivilege"`
	ObtainOtherPrivilege    bool    `json:"obtainOtherPrivilege"`
	UserInteractionRequired bool    `json:"userInteractionRequired"`
}
type Impact struct {
	BaseMetricV3 BaseMetricV3 `json:"baseMetricV3"`
	BaseMetricV2 BaseMetricV2 `json:"baseMetricV2"`
}


// the important ones
type initCveContext struct {
	Filename string
	CVEFile *CVEFile
	Error error
}

type CVEFile struct {
	CVEDataType         string     `json:"CVE_data_type"`
	CVEDataFormat       string     `json:"CVE_data_format"`
	CVEDataVersion      string     `json:"CVE_data_version"`
	CVEDataNumberOfCVEs string     `json:"CVE_data_numberOfCVEs"`
	CVEDataTimestamp    string     `json:"CVE_data_timestamp"`
	CVEItems            []CVEItems `json:"CVE_Items"`
}


type CVEItem struct {
	DataType    string      `json:"data_type"`
	DataFormat  string      `json:"data_format"`
	DataVersion string      `json:"data_version"`
	CVEDataMeta CVEDataMeta `json:"CVE_data_meta"`
	Affects     Affects     `json:"affects"`
	Problemtype ProblemType `json:"problemtype"`
	References  References  `json:"references"`
	Description Description `json:"description"`
}

type CVEItems struct {
	CVEItem          CVEItem            `json:"cve"`
	Configurations   Configurations `json:"configurations"`
	Impact           Impact         `json:"impact"`
	PublishedDate    string         `json:"publishedDate"`
	LastModifiedDate string         `json:"lastModifiedDate"`
}

// this one is used by other components in the app
type CVEStore struct {
	CVEItems []CVEItem
}


func (c *Config) InitCveStore() (cveStore CVEStore, err error) {

	files, err := ioutil.ReadDir(c.CVEFeedsDir)
	if err != nil {
		return
	}

	filenames := make([]string, len(files))

	for i, f := range files {
		filenames[i] = f.Name()
	}

	jobs    := make(chan *initCveContext, len(filenames))
	results := make(chan *initCveContext, len(filenames))

	for w := 0; w < c.CVEWorkerPoolSize; w++ {
		go c.worker(jobs, results)
	}

	for _, f := range filenames {
		jobs <- &initCveContext{Filename: f}
	}
	close(jobs)


	cveFiles := make([]*CVEFile, len(filenames))

	for r := 0; r < len(filenames); r++ {
		initCveContext := <-results

		if initCveContext.Error != nil {
			return CVEStore{}, initCveContext.Error
		}

		cveFiles[r] = initCveContext.CVEFile
	}

	//TODO: loop through the files and put all CVEItems into the CVEStore

	for _, cf := range cveFiles {
		log.Printf("InitCveStore() len(cf.CVEItems)=%d", len(cf.CVEItems))
	}

	return
}

func (c *Config) worker(jobs <-chan *initCveContext, results chan<- *initCveContext) {

	for icc := range jobs {

		log.Printf("worker() filename=%s", icc.Filename)

		filepath := fmt.Sprintf("%s/%s", c.CVEFeedsDir, icc.Filename)

		contents, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Printf("worker() ReadFile err=%s", err)
			icc.Error = err
		}

		cveFile := &CVEFile{}

		err = json.Unmarshal(contents, cveFile)
		if err != nil {
			log.Printf("worker() Unmarshal err=%s", err)
			icc.Error = err
		}

		icc.CVEFile = cveFile

		results <- icc
	}
}
