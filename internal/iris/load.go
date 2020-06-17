package iris

import (
	"github.com/andygeiss/pipeline-example/internal/api"
	"google.golang.org/protobuf/proto"
)

// Load ...
func Load(raw []byte) (out proto.Message, err error) {
	var data api.Processed
	if err := proto.Unmarshal(raw, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
