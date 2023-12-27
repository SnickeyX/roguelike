import os

#  works with 3.x
# requires Go; follow https://go.dev/doc/install#tarball
if __name__ == "__main__":
    os.system('env GOOS=js GOARCH=wasm go build -o rogue.wasm')
    os.system('cp $(go env GOROOT)/misc/wasm/wasm_exec.js .')
    os.system('python -m http.server 8080')
