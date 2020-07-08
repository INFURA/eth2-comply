package testcases

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/INFURA/eth2-comply/pkg/eth2spec"
)

// TestsCasesOpts are used to configure statically defined test cases with
// dynamic values at runtime.
type TestsCasesOpts struct {
	// target is a URL hosting an Ethereum 2.0 API where test requests should
	// be sent. For example, http://localhost:5051
	Target string
	// TestsRemote is a URL to a zip file containing a directory tree
	// containing well-specified JSON tests cases.
	TestsRemote string
	// OutDir is the directory where zip files will be downloaded, and
	// unzipped.
	OutDir string
	// testsRoot is a file path to a directory tree containing well-specified
	// JSON tests cases.
	TestsRoot string
	// oapiClient is an instantiated Ethereum 2.0 OAPI client which will be
	// used to conduct OAPI operations.
	OapiClient *eth2spec.APIClient
}

// All returns an array of executable test cases for the directory tree
// specified by testsRoot. Tests will be executable against the given target.
//
// Cases recursively traverses down the testsRoot directory and collects all
// test cases it finds. *Any* file in the dir tree which is not a directory or
// a well-formed JSON test case will cause this function to return an error.
func All(opts *TestsCasesOpts) ([]*Case, error) {
	configs := []CaseConfig{}
	cases := []*Case{}

	if opts.TestsRemote != "" {
		filePath, err := getRemoteTestsFile(opts.TestsRemote, opts.OutDir)
		if err != nil {
			return nil, err
		}
		err = Unzip(filePath, opts.OutDir)
		if err != nil {
			return nil, err
		}
		// Overwrite opts.TestsRoot
		opts.TestsRoot = filepath.Join(opts.OutDir, "tests")
	}

	configs, err := getConfigsFromDirTree(opts.TestsRoot, configs)
	if err != nil {
		return nil, err
	}

	for _, config := range configs {
		if err != nil {
			return nil, err
		}
		c := NewCase(config, opts.OapiClient)

		cases = append(cases, c)

	}

	return cases, nil

}

type TestSpecificationError struct {
	Filepath string
	Err      error
}

func (e TestSpecificationError) Error() string {
	return fmt.Sprintf("Error parsing %s: %s", e.Filepath, e.Err.Error())
}

// getConfigsFromDirTree recursively traverses down the root directory and
// collects test cases it finds. *Any* file in the dir tree which is not a
// directory or a well-formed JSON test case will cause this function to
// return an error.
func getConfigsFromDirTree(root string, configs []CaseConfig) ([]CaseConfig, error) {
	rootDir, err := os.Open(root)
	if err != nil {
		return nil, err
	}

	filesInDir, err := rootDir.Readdir(0)
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range filesInDir {
		if fileInfo.IsDir() {
			// Recursive case, found a directory
			configs, err = getConfigsFromDirTree(root+"/"+fileInfo.Name(), configs)
			if err != nil {
				return nil, err
			}
		} else {
			// Base case, found files
			config := CaseConfig{}
			filePath := root + "/" + fileInfo.Name()
			file, err := os.Open(filePath)
			if err != nil {
				return nil, TestSpecificationError{
					Filepath: filePath,
					Err:      err,
				}
			}

			fileData, err := ioutil.ReadAll(file)
			if err != nil {
				return nil, TestSpecificationError{
					Filepath: filePath,
					Err:      err,
				}
			}

			err = json.Unmarshal(fileData, &config)
			if err != nil {
				return nil, TestSpecificationError{
					Filepath: filePath,
					Err:      err,
				}
			}

			configs = append(configs, config)
		}
	}

	return configs, nil
}

// Unzip unzips a zip file into the provided outDir
func Unzip(zipFilePath, outDir string) error {
	zipFileReader, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}

	// Create root tests directory
	err = os.MkdirAll(outDir, 0777)
	if err != nil {
		zipFileReader.Close()
		return err
	}

	// Write all the tests out into the tests directory
	for _, zippedFile := range zipFileReader.File {
		fileName := filepath.Join(outDir, zippedFile.Name)

		if zippedFile.FileInfo().IsDir() {
			// Create inner directory
			err = os.MkdirAll(fileName, 0777)
			if err != nil {
				return err
			}
		} else {
			// Create file
			err = os.MkdirAll(filepath.Dir(fileName), 0777)
			if err != nil {
				return err
			}
			f, err := os.Create(fileName)
			if err != nil {
				return err
			}

			// Write file contents
			fileReader, err := zippedFile.Open()
			if err != nil {
				return err
			}
			_, err = io.Copy(f, fileReader)
			if err != nil {
				return err
			}

			f.Close()
			fileReader.Close()
		}
	}

	zipFileReader.Close()
	return nil
}

// getRemoteTestsFile fetches a ZIP file from the location specified in
// testsRemote and places it in the outDir.
func getRemoteTestsFile(testsRemote, outDir string) (string, error) {
	resp, err := http.Get(testsRemote)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Place remote zip into /tmp directory
	pathParts := strings.Split(testsRemote, "/")
	fileName := filepath.Join(outDir, pathParts[len(pathParts)-1])
	f, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	_, err = f.Write(body)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
