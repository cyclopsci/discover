package discover

import (
	"os"
	"fmt"
	"bufio"
	"regexp"
	"strings"
	"path/filepath"
)

type language struct {
	Key          string
	Extensions   []string
	Paths        []string
	PathMatchers []string
	ContentRegex []string
	IgnoredDirs  []string
}

// Run returns all matches of a language type from the root of the specified tree
func Run(root string, displayRoot string) map[string][]string {
	root = strings.TrimSuffix(root, "/")
	displayRoot = strings.TrimSuffix(displayRoot, "/")

	languages := []language{
		puppetManifest,
		puppetModule,
		ansibleRole,
		ansiblePlaybook,
	}

	tree := walkDirectory(root)

	results := analyzeTree(root, displayRoot, languages, tree)
	results["root"] = []string{root}

	return results
}

func walkDirectory(root string) []string {
	tree := []string{}

	filepath.Walk(root, func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			tree = append(tree, path)
		}
		return nil
	})

	return tree
}

func analyzeTree(root string, displayRoot string, languages []language, tree []string) map[string][]string {
	var matches = []string{}
	var results = make(map[string][]string)

	for _, f := range tree {
		for _, lang := range languages {
			match := findLanguageMatch(lang, f)
			if match != "" {
				if root != displayRoot {
					match = strings.Replace(match, root, displayRoot, 1)
				}
				matches = append(matches, match)
				if !stringInSlice(match, results[lang.Key]) {
					results[lang.Key] = append(results[lang.Key], match)
				}
			}
		}
	}

	return results
}

func findLanguageMatch(lang language, file string) string {
	for _, value := range lang.IgnoredDirs {
		s := strings.Split(file, "/")
		if stringInSlice(value, s) {
			return ""
		}
	}

	for _, value := range lang.Paths {
		if search(value, file) {
			return strings.Replace(file, value, "", 1)
		}
	}

	for _, value := range lang.PathMatchers {
		match, _ := regexp.MatchString(fmt.Sprintf("^.*(/)?%s$", value), file)
		if match {
			re := regexp.MustCompile("(" + value + ")")
			return re.ReplaceAllString(file, "")
		}
	}

	for _, value := range lang.Extensions {
		ext := strings.Replace(filepath.Ext(file), ".", "", 1)
		if value == ext {
			return file
		}
	}

	for _, value := range lang.ContentRegex {
		if searchContent(value, file) {
			return strings.Replace(file, value, "", 1)
		}
	}

	return ""
}

func search(value string, target string) bool {
	exp := fmt.Sprintf("^.*(/)?%s(/)?", value)
	matched, _ := regexp.MatchString(exp, target)
	if matched {
		return true
	}
	return false
}

func stringInSlice(value string, slice []string) bool {
	for _, s := range slice {
		if s == value {
			return true
		}
	}
	return false
}

func searchContent(expression string, file string) bool {
	regex, err := regexp.Compile(expression)
	if err != nil {
		return false
	}
	fh, err := os.Open(file)
	f := bufio.NewReader(fh)
	if err != nil {
		return false
	}
	defer fh.Close()
	buf := make([]byte, 1024)
	for {
		buf, _ , err = f.ReadLine()
		if err != nil {
			return false
		}
		s := string(buf)
		if regex.MatchString(s) {
			return true
		}
	}
	return false
}
