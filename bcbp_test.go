package bcbp

import (
	"context"
	"testing"
)

func TestParse(t *testing.T) {

	ctx := context.Background()

	data := "M1DESMARAIS/LUC       EABC123 YULFRAAC 0834 326J001A0025 100"
	
	err := Parse(ctx, data)

	if err != nil {
		t.Fatal(err)
	}
	
}
