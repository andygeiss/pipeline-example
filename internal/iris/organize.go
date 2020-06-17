package iris

import (
	"github.com/andygeiss/pipeline-example/internal/api"
	"google.golang.org/protobuf/proto"
	"strconv"
)

// Organize ...
func Organize(records [][]string) (out proto.Message, err error) {
	interim := new(api.Interim)
	for _, record := range records {
		petalLength, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			return nil, err
		}
		petalWidth, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, err
		}
		sepalLength, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, err
		}
		sepalWidth, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			return nil, err
		}
		interim.PetalLength = append(interim.PetalLength, petalLength)
		interim.PetalWidth = append(interim.PetalWidth, petalWidth)
		interim.SepalLength = append(interim.SepalLength, sepalLength)
		interim.SepalWidth = append(interim.SepalWidth, sepalWidth)
		interim.Species = append(interim.Species, record[4])
	}
	return interim, nil
}
