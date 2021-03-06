package goimport

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

//ImportDir fetches the import dir, to fetch the go import from
func ImportDir(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("Unable to find absolute path for %s: %-v", path, err)
	}
	srcDir, err := goSrc()
	if err != nil {
		return "", fmt.Errorf("Error getting go import dir: %-v", err)
	}
	importDir, err := filepath.Rel(srcDir, absPath)
	if err != nil {
		return "", fmt.Errorf("Error fetching import dir %s: %-v", importDir, err)
	}
	return importDir, nil
}

func ParseRelation(rootPath string, seekPath string, leafVisibility bool) *ImportPathFactory {
	factory := NewImportPathFactory(
		rootPath,
		seekPath,
		leafVisibility,
	)
	factory.Root = factory.Get(rootPath)
	if factory.Root == nil {
		return nil
	}
	return factory

}

type ImportPathFactory struct {
	Root   *ImportPath
	Filter *ImportFilter
	Pool   map[string]*ImportPath
}

func NewImportPathFactory(
	rootPath string, seekPath string, leafVisibility bool) *ImportPathFactory {

	self := &ImportPathFactory{Pool: make(map[string]*ImportPath)}
	filter := NewImportFilter(
		rootPath,
		seekPath,
		leafVisibility,
	)

	self.Filter = filter
	return self
}
func (self *ImportPathFactory) GetRoot() *ImportPath {
	return self.Root
}

func (self *ImportPathFactory) GetAll() []*ImportPath {
	ret := make([]*ImportPath, 0)
	for _, value := range self.Pool {
		ret = append(ret, value)
	}
	return ret
}

func (self *ImportPathFactory) Get(importPath string) *ImportPath {
	// aquire from pool
	pool := self.Pool
	if _, ok := pool[importPath]; ok {
		return pool[importPath]
	}
	filter := self.Filter
	// if not applicable return nullobject
	if !filter.Applicable(importPath) {
		// if invisible return nil
		if !filter.Visible(importPath) {
			return nil
		}
		pool[importPath] = &ImportPath{
			ImportPath: importPath}
		return pool[importPath]
	}

	srcDir, err := goSrc()
	if err != nil {
		fmt.Printf("Error getting go import dir: %-v", err)
	}

	dirPath := filepath.Join(srcDir, importPath)
	if !fileExists(dirPath) {
		// if invisible return nil
		if !filter.Visible(importPath) {
			return nil
		}
		pool[importPath] = &ImportPath{
			ImportPath: importPath}
		return pool[importPath]
	}
	ret := &ImportPath{
		ImportPath: importPath,
	}
	pool[importPath] = ret
	fileNames := glob(dirPath)
	ret.Init(self, fileNames)
	return ret
}

//ImportFilter
type ImportFilter struct {
	root     string
	seekPath string
	plotLeaf bool
}

func NewImportFilter(root string, seekPath string, plotLeaf bool) *ImportFilter {
	if seekPath == "SELF" {
		seekPath = root
	}
	impf := &ImportFilter{
		root:     root,
		seekPath: seekPath,
		plotLeaf: plotLeaf,
	}
	return impf

}

func (self *ImportFilter) Visible(path string) bool {
	return self.plotLeaf
}

func (self *ImportFilter) Applicable(path string) bool {
	if self.seekPath == "" {
		return true
	}
	if strings.Index(path, self.seekPath) == 0 {
		return true
	}
	return false
}

func isMatched(pattern string, target string) bool {
	r, _ := regexp.Compile(pattern)
	return r.MatchString(target)
}

func glob(dirPath string) []string {
	fileNames, err := filepath.Glob(filepath.Join(dirPath, "/*.go"))
	if err != nil {
		panic("no gofiles")
	}

	files := make([]string, 0, len(fileNames))

	for _, v := range fileNames {
		if isMatched("test", v) {
			continue
		}
		if isMatched("example", v) {
			continue
		}
		files = append(files, v)
	}
	return files
}

func goSrc() (string, error) {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		return "", fmt.Errorf("environment variable GOPATH is unset")
	}
	return filepath.Join(os.Getenv("GOPATH"), "src"), nil
}
