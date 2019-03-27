package app

import (
	"io/ioutil"
	"log"
)

type initCveContext struct {
	Filename string
	CVEFile CVEFile
	Error error
}

type CVEFile struct {
	CVEItems []CVEItem
}

type CVEStore struct {
	CVEItems []CVEItem
}

type CVEItem struct {
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

	jobs := make(chan initCveContext, len(filenames))
	results := make(chan initCveContext, len(filenames))

	for w := 0; w < c.CVEWorkerPoolSize; w++ {
		go worker(jobs, results)
	}

	for _, f := range filenames {
		jobs <- initCveContext{Filename: f}
	}
	close(jobs)


	cveFiles := make([]CVEFile, len(filenames))

	for r := 0; r < len(filenames); r++ {
		initCveContext := <-results
		cveFiles[r] = initCveContext.CVEFile
	}

	//TODO: loop through the files and put all CVEItems into the CVEStore

	return
}

func worker(jobs <-chan initCveContext, results chan<- initCveContext) {

	for icc := range jobs {

		log.Printf("worker() filename=%s", icc.Filename)

		results <- icc
	}
}
