# go-sfomuseum-bcbp

Experimental Go package to run the `sfomuseum/rs-sfomuseum-bcbp` WebAssembly (WASI) binary.

## Important

This does not work yet.

```
$> go test -v ./...
?   	github.com/sfomuseum/go-bcbp/wasm	[no test files]
=== RUN   TestParse
    bcbp_test.go:17: Failed to run WASM binary, Failed to instantiate module, module[__wbindgen_placeholder__] not instantiated
--- FAIL: TestParse (0.17s)
FAIL
FAIL	github.com/sfomuseum/go-bcbp	0.457s
FAIL
```

## Compiling the WASI binary

```
$> cd /usr/local/sfomuseum/rs-sfomuseum-bcbp
$> cargo build --target wasm32-wasi --release
   Compiling cfg-if v1.0.0
   Compiling lexical-core v0.7.6
   Compiling memchr v2.5.0
   Compiling bitflags v1.3.2
   Compiling static_assertions v1.1.0
   Compiling arrayvec v0.5.2
   Compiling ryu v1.0.13
   Compiling wasm-bindgen v0.2.84
   Compiling arrayvec v0.4.12
   Compiling nom v5.1.2
   Compiling nodrop v0.1.14
   Compiling js-sys v0.3.61
   Compiling iata_bcbp v1.0.0
   Compiling rs-sfomuseum-bcbp v0.1.0 (/usr/local/sfomuseum/rs-sfomuseum-bcbp)
    Finished release [optimized] target(s) in 10.51s
```

Then copy the `target/wasm32-wasi/release/rs_sfomuseum_bcbp.wasm` file to `wasm/rs_sfomuseum_bcbp.wasm`.

## See also

* https://github.com/sfomuseum/rs-sfomuseum-bcbp
* https://github.com/aaronland/go-wasi