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
	Args:  cobra.MinimumNArgs(0),
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.PersistentFlags().StringP("input", "i", "", "input project name")
	rootCmd.MarkPersistentFlagRequired("input")
	rootCmd.PersistentFlags().StringP("output", "o", "STDOUT", "output file")
	rootCmd.PersistentFlags().IntP("depth", "d", 128, "max plot depth of the dependency tree")
	rootCmd.PersistentFlags().StringP("focus", "f", "", "focus on the specific module")
	rootCmd.PersistentFlags().StringP("search", "s", "", "top directory of searching")
	rootCmd.PersistentFlags().BoolP("leaf", "l", false, "whether leaf nodes are plotted")
	rootCmd.PersistentFlags().BoolP("metrics", "m", false, "display module metrics")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	if err := process(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}

func process() error {
	fmt.Println()
	inputDir := rootCmd.Flag("input").Value.String()
	output := rootCmd.Flag("output").Value.String()
	reserved := rootCmd.Flag("focus").Value.String()
	seek := rootCmd.Flag("search").Value.String()
	leaf, _ := strconv.ParseBool(rootCmd.Flag("leaf").Value.String())
	depth, _ := strconv.Atoi(rootCmd.Flag("depth").Value.String())
	metricsOpt, _ := strconv.ParseBool(rootCmd.Flag("metrics").Value.String())

	dir, err := goimport.ImportDir(inputDir)
	if err != nil {
		return fmt.Errorf("Couldn't fetch import path: %-v", err)
	}
	factory := goimport.ParseRelation(dir, seek, leaf)
	if factory == nil {
		return fmt.Errorf("directory not found: [ %s ]", inputDir)
	}
	root := factory.GetRoot()
	if !root.HasFiles() {
		return fmt.Errorf("%s has no .go files", root.ImportPath)
	}
	if 0 > depth {
		return fmt.Errorf("-d or --depth should have positive int")
	}
	out := getOutputWriter(output)
	if metricsOpt {
		metricsWriter := metrics.New(out)
		metricsWriter.Plot(pathToNode(factory.GetAll()))
		return nil
	}

	writer := dotwriter.New(out)
	writer.MaxDepth = depth
	if reserved == "" {
		writer.PlotGraph(root)
		return nil
	}
	writer.Reversed = true

	rroot := factory.Get(reserved)
	if rroot == nil {
		return fmt.Errorf("-r %s does not exist.\n ", reserved)
	}
	if !rroot.HasFiles() {
		return fmt.Errorf("-r %s has no go files.\n ", reserved)
	}

	writer.PlotGraph(rroot)
	return nil
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
