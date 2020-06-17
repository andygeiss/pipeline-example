package iris

import (
	"github.com/andygeiss/pipeline-example/internal/api"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"google.golang.org/protobuf/proto"
	"os"
)

// PlotScatter ...
func PlotScatter(in interface{}, data proto.Message) error {
	filename := in.(string)
	d := data.(*api.Interim)
	// Prepare the internal pipeline-example (select, filter, ...) for plotting.
	x := make([]float64, 0)
	y := make([]float64, 0)
	for i := range d.SepalLength {
		// Filter
		// if d.Species[i] == "Iris-setosa" {
		//	continue
		// }
		x = append(x, d.SepalLength[i])
		y = append(y, d.SepalWidth[i])
	}
	xys := make([]plotter.XY, 0)
	for i := 0; i < len(x); i++ {
		xys = append(xys, plotter.XY{X: x[i], Y: y[i]})
	}
	// Create a new plot.
	plt, err := plot.New()
	if err != nil {
		return err
	}
	// Create a new scatter based on the values.
	sca, err := plotter.NewScatter(plotter.XYs(xys))
	if err != nil {
		return err
	}
	// Add the scatter to the plot.
	plt.Add(sca)
	// Create a new writer.
	wri, err := plt.WriterTo(512, 512, "png")
	if err != nil {
		return err
	}
	// Create a new file for the output.
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	// Write the scatter plot to the file.
	if _, err := wri.WriteTo(file); err != nil {
		return err
	}
	return nil
}
