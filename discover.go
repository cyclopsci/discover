package discover

import (
	"os"
	"fmt"
	"regexp"
	"strings"
	"path/filepath"
)

var (
	tree []string
	languages = []language{
		puppetFiles,
		puppetModule,
		yamlFiles,
	}
)

type language struct {
	Key		string
	Extensions	[]string
	Paths		[]string
	Matchers	[]string
	IgnoredDirs	[]string
}

// Run returns all matches of a language type from the root of the specified tree
func Run(root string) map[string][]string {
	os.Chdir(root)
	walkDirectory(root)
	results := analyzeTree(languages, tree)
	results["root"] = []string{root}
	return results
}

func walkDirectory(root string) {
	filepath.Walk(root, visitFile)
}

func visitFile(path string, file os.FileInfo, err error) error {
	if !file.IsDir() {
		tree = append(tree, path)
	}
	return nil
}

func analyzeTree(languages []language, tree []string) map[string][]string {
	var matches []string
	var candidates = make(map[string][]string)
	for _, f := range tree {
		for _, lang := range languages {
			match := matchLanguage(lang, f)
			if match != "" {
				matches = append(matches, match)
				candidates[lang.Key] = append(candidates[lang.Key], match)
			}
		}
	}
	for _, lang := range languages {
		candidates[lang.Key] = deduplicate(candidates[lang.Key], matches)
	}
	return candidates
}

func deduplicate(languageMatches []string, totalMatches []string) []string {
	var results []string
	var found bool
	for _, l := range languageMatches {
		found = false
		for _, m := range totalMatches {
			if strings.Contains(l, m) && l != m {
				found = true
			}
		}
		if !found && !stringInSlice(l, results) {
			results = append(results, l)
		}
	}
	return results
}

func matchLanguage(lang language, file string) string {
	for _, value := range lang.IgnoredDirs {
		if search(value, file) {
			return ""
		}
	}
	for _, value := range lang.Paths {
		if search(value, file) {
			return strings.Replace(file, value, "", 1)
		}
	}
	for _, value := range lang.Matchers {
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
