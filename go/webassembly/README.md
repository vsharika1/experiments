# Learn WebAssembly

Learn WebAssembly by following this [Go Wiki](https://github.com/golang/go/wiki/WebAssembly).

## Getting Started
This page assumes a functional Go 1.11 or newer installation. 

To compile a basic Go package for the web,

Set `GOOS=js` and `GOARCH=wasm` environment variables to compile for
WebAssembly:

```sh
$ GOOS=js GOARCH=wasm go build -o main.wasm
```

That will build the package and produce an executable WebAssembly module file
named `main.wasm`. The `.wasm` file extension will make it easier to serve it
over HTTP with the correct Content-Type header later on.

Note that you can only compile main packages. Otherwise, you will get an object
file that cannot be run in WebAssembly. If you have a package that you want to
be able to use with WebAssembly, convert it to a main package and build a
binary.

To execute `main.wasm` in a browser, we’ll also need a JavaScript support file,
and a HTML page to connect everything together.

Copy the JavaScript support file:

```sh
$ cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

Create an `index.html` file:

```html
<html>
	<head>
		<meta charset="utf-8"/>
		<script src="wasm_exec.js"></script>
		<script>
			const go = new Go();
			WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
				go.run(result.instance);
			});
		</script>
	</head>
	<body></body>
</html>
```

If your browser doesn’t yet support `WebAssembly.instantiateStreaming`, you can
use a
[polyfill](https://github.com/golang/go/blob/b2fcfc1a50fbd46556f7075f7f1fbf600b5c9e5d/misc/wasm/wasm_exec.html#L17-L22).

Then serve the three files (`index.html`, `wasm_exec.js`, and `main.wasm`) from
a web server. For example, with
[`goexec`](https://github.com/shurcooL/goexec#goexec):

```sh
# install goexec: go get -u github.com/shurcooL/goexec
$ goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
```

Or use your own [basic HTTP server command](https://play.golang.org/p/pZ1f5pICVbV).

## Executing WebAssembly with Node.js

It’s possible to execute compiled WebAssembly modules using Node.js rather than
a browser, which can be useful for testing and automation.

With Node installed and in your `PATH`, set the `-exec` flag to the location of
`go_js_wasm_exec` when you execute `go run` or `go test`.

By default, `go_js_wasm_exec` is in the `misc/wasm` directory of your Go
installation.

```sh
$ GOOS=js GOARCH=wasm go run -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec" .
Hello, WebAssembly!
$ GOOS=js GOARCH=wasm go test -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec"
PASS
ok  	example.org/my/pkg	0.800s
```

Adding `go_js_wasm_exec` to your `PATH` will allow `go run` and `go test` to
work for `js/wasm` without having to manually provide the `-exec` flag each
time:

```sh
$ export PATH="$PATH:$(go env GOROOT)/misc/wasm"
$ GOOS=js GOARCH=wasm go run .
Hello, WebAssembly!
$ GOOS=js GOARCH=wasm go test
PASS
ok  	example.org/my/pkg	0.800s
```

## Interacting with the DOM

See https://pkg.go.dev/syscall/js.

## WebAssembly in Chrome

If you run a newer version of Chrome there is a flag
(`chrome://flags/#enable-webassembly-baseline`) to enable Liftoff, their new
compiler, which should significantly improve load times. Further info
[here](https://v8.dev/blog/liftoff).

## Debugging

WebAssembly doesn’t yet have any support for debuggers, so you’ll need to use
the good 'ol `println()` approach for now to display output on the JavaScript
console.

An official WebAssembly Debugging Subgroup has been created to address this,
with some initial investigation and proposals under way:
- WebAssembly Debugging Capabilities Living Standard
- DWARF for WebAssembly Target

### Analysing the structure of a WebAssembly file

[WebAssembly Code Explorer](https://wasdk.github.io/wasmcodeexplorer/) is useful
for visualising the structure of a WebAssembly file.
- Clicking on a hex value to the left will highlight the section it is part of,
  and the corresponding text representation on the right
- Clicking a line on the right will highlight the hex byte representations for
  it on the left

## Reducing the size of Wasm files

At present, Go generates large Wasm files, with the smallest possible size being
around ~2MB. If your Go code imports libraries, this file size can increase
dramatically. 10MB+ is common.

There are two main ways (for now) to reduce this file size:

1. Manually compress the .wasm file.

    1. Using `gz` compression reduces the ~2MB (minimum file size) example WASM
       file down to around 500kB. It may be better to use
       [Zopfli](https://github.com/google/zopfli) to do the gzip compression, as
       it gives better results than `gzip --best`, however it does take much
       longer to run.

    2. Using [Brotli](https://github.com/google/brotli) for compression, the
       file sizes are markedly better than both Zopfli and `gzip --best`, and
       compression time is somewhere in between the two, too. This [(new) Brotli
       compressor](https://github.com/andybalholm/brotli) looks reasonable.

    Use something like https://github.com/lpar/gzipped to automatically serve
    compressed files with correct headers, when available.

2. Use [TinyGo](https://github.com/tinygo-org/tinygo) to generate the Wasm file
   instead.

    TinyGo supports a subset of the Go language targeted for embedded devices,
    and has a WebAssembly output target.

    While it does have limitations (not yet a full Go implementation), it is
    still fairly capable and the generated Wasm files are…​ tiny. ~10kB isn’t
    unusual. The "Hello world" example is 575 bytes. If you `gz -6` that, it
    drops down to 408 bytes. :wink:

    This project is also very actively developed, so its capabilities are
    expanding out quickly. See https://tinygo.org/docs/guides/webassembly/ for
    more information on using WebAssembly with TinyGo.

## Other WebAssembly resources

- [Awesome-Wasm](https://github.com/mbasso/awesome-wasm) - An extensive list of
  further Wasm resources. Not Go specific.