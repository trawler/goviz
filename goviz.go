package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/trawler/goviz/dotwriter"
	"github.com/trawler/goviz/goimport"
	"github.com/trawler/goviz/metrics"
)

var rootCmd = &cobra.Command{
	Use:   "goviz",
	Short: "A visualization tool for golang project dependency",
}

func init() {
	rootCmd.PersistentFlags().StringP("input", "i", "", "input project name")
	rootCmd.MarkFlagRequired("input")
	rootCmd.PersistentFlags().StringP("output", "o", "STDOUT", "output file")
	rootCmd.PersistentFlags().IntP("depth", "d", 128, "max plot depth of the dependency tree")
	rootCmd.PersistentFlags().StringP("focus", "f", "", "focus on the specific module")
	rootCmd.PersistentFlags().StringP("search", "s", "", "top directory of searching")
	rootCmd.PersistentFlags().BoolP("leaf", "l", false, "whether leaf nodes are plotted")
	rootCmd.PersistentFlags().BoolP("metrics", "m", false, "display module metrics")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred: %v\n", err)
		os.Exit(1)
	}

	res := process()
	os.Exit(res)

}

func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
}

func process() int {
	fmt.Println()
	inputDir := rootCmd.Flag("input").Value.String()
	output := rootCmd.Flag("output").Value.String()
	reserved := rootCmd.Flag("focus").Value.String()
	leaf, _ := strconv.ParseBool(rootCmd.Flag("leaf").Value.String())
	depth, _ := strconv.Atoi(rootCmd.Flag("depth").Value.String())
	metricsOpt, _ := strconv.ParseBool(rootCmd.Flag("metrics").Value.String())

	factory := goimport.ParseRelation(inputDir, output, leaf)
	goimport.ImportDir(inputDir)

	if factory == nil {
		//fmt.Printf("inputdir does not exist.\n go get %s", inputDir)
		return 1
	}
	root, err := goimport.ImportDir(inputDir)
	if err != nil {
		fmt.Printf("Couldn't fetch import path: %-v", err)
	}
	//if !root.HasFiles() {
	//	errorf("%s has no .go files\n", root.ImportPath)
	//	return 1
	//}
	if 0 > depth {
		errorf("-d or --depth should have positive int\n")
		return 1
	}
	out := getOutputWriter(output)
	if metricsOpt {
		metricsWriter := metrics.New(out)
		metricsWriter.Plot(pathToNode(factory.GetAll()))
		return 0
	}

	writer := dotwriter.New(out)
	writer.MaxDepth = depth
	if reserved == "" {
		writer.PlotGraph(root)
		return 0
	}
	writer.Reversed = true

	rroot := factory.Get(reserved)
	if rroot == nil {
		errorf("-r %s does not exist.\n ", reserved)
		return 1
	}
	if !rroot.HasFiles() {
		errorf("-r %s has no go files.\n ", reserved)
		return 1
	}

	writer.PlotGraph(rroot)
	return 0
}

func pathToNode(pathes []*goimport.ImportPath) []dotwriter.IDotNode {
	r := make([]dotwriter.IDotNode, len(pathes))

	for i := range pathes {
		r[i] = pathes[i]
	}
	return r
}
func getOutputWriter(name string) *os.File {
	if name == "STDOUT" {
		return os.Stdout
	}
	if name == "STDERR" {
		return os.Stderr
	}
	f, _ := os.Create(name)
	return f
}
