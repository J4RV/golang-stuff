package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
	"sync"

	"github.com/j4rv/gostuff/log"

	"github.com/j4rv/gostuff/stopwatch"
)

var fileRegex *regexp.Regexp
var containsFilter string
var ignoreCasing bool

func main() {
	var fileRegexStr string
	var verbose bool
	flag.StringVar(&fileRegexStr, "n", "^.*$", "Name regex: files which filenames don't pass this Regex will be discarded")
	flag.StringVar(&containsFilter, "c", "", "Contains: only files that contains the input text in a single line will be logged")
	flag.BoolVar(&verbose, "v", false, "Verbose: more logging")
	flag.BoolVar(&ignoreCasing, "i", false, "Ignore casing")
	flag.Parse()

	if verbose {
		log.SetLevel(log.ALL)
	} else {
		log.SetLevel(log.INFO)
	}
	if ignoreCasing {
		containsFilter = strings.ToLower(containsFilter)
	}

	fileRegex = regexp.MustCompile(fileRegexStr)
	execute()
}

func execute() {
	stopTimer := stopwatch.Start()
	log.Info("Searching...")
	files := traverseAndReturnFilePaths(".")
	log.Info("Files that passed the filter:", len(files))
	log.Info("Printing filenames:")
	for f := range files {
		fmt.Println("\t./" + f)
	}
	elapsed := stopTimer()
	log.Debug("Done in", elapsed.Seconds(), "seconds")
}

func traverseAndReturnFilePaths(dirPath string) chan string {
	log.Info("Traversing the directory recursively and locating files...")
	files := traverseAndReturnFilenamesChan(dirPath)
	log.Info("Total files amount:", len(*files))

	// go workers!
	var wg sync.WaitGroup
	var result = make(chan string, len(*files))
	for i := 0; i < runtime.NumCPU(); i++ {
		log.Debug("Initiating worker number", i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				file, ok := <-*files
				if !ok {
					return
				}
				checkFile(file, &result)
			}
		}()
	}
	wg.Wait()
	close(result)

	return result
}

func traverseAndReturnFilenamesChan(dirPath string) *chan string {
	var filesSlice []string
	traverseRecursive(dirPath, &filesSlice)

	var filesChan = make(chan string, len(filesSlice))
	for _, f := range filesSlice {
		filesChan <- f
	}
	close(filesChan)

	return &filesChan
}

func traverseRecursive(dirPath string, filePaths *[]string) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Error("Could not read directory", dirPath)
		return
	}
	for _, f := range files {
		childPath := path.Join(dirPath, f.Name())
		if f.IsDir() {
			traverseRecursive(childPath, filePaths)
		} else {
			*filePaths = append(*filePaths, childPath)
		}
	}
}

func checkFile(filePath string, result *chan string) {
	if !fileRegex.Match([]byte(filePath)) {
		return
	}
	if containsFilter == "" {
		*result <- filePath
	}
	file, err := os.Open(filePath)
	if err != nil {
		log.Error("Could not open", filePath)
		return
	}
	checkFileContent(filePath, file, result)
	err = file.Close()
	if err != nil {
		log.Error("Could not close", filePath)
	}
}

func checkFileContent(filePath string, file *os.File, result *chan string) {
	contentBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error("Could not read", filePath)
		return
	}
	content := string(contentBytes)
	if ignoreCasing {
		content = strings.ToLower(content)
	}
	if strings.Contains(content, containsFilter) {
		*result <- filePath
	}
}
