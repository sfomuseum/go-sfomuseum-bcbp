package bcbp

import (
	"context"
	"fmt"
	
	"github.com/sfomuseum/go-bcbp/wasm"
	"github.com/aaronland/go-wasi"
)

func Parse(ctx context.Context, data string) error {

	wasm_r, err := wasm.FS.Open("rs_sfomuseum_bcbp.wasm")

	if err != nil {
		return fmt.Errorf("Failed to open WASM binary, %w", err)
	}

	defer wasm_r.Close()
	
	result, err := wasi.Run(ctx, wasm_r, data)

	if err != nil {
		return fmt.Errorf("Failed to run WASM binary, %w", err)
	}

	fmt.Println(result)
	return nil
}
