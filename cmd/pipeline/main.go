package main

import (
	"github.com/andygeiss/pipeline"
	"github.com/andygeiss/pipeline-example/internal/iris"
	"log"
)

func main() {
	p := new(pipeline.Pipeline)
	p.Gather("https://archive.ics.uci.edu/ml/machine-learning-databases/iris/iris.data", "data/external/iris.csv").
		Organize("data/external/iris.csv", iris.Organize).
		Evaluate("reports/plot_scatter.png", iris.PlotScatter).
		Validate(nil, iris.PrintStats)
	if err := p.Error(); err != nil {
		log.Fatal(err)
	}
}
