
FROM --platform=$BUILDPLATFORM golang:1.18 AS buildbase
WORKDIR /src

RUN wget https://github.com/tinygo-org/tinygo/releases/download/v0.26.0/tinygo_0.26.0_amd64.deb \
    &&  dpkg -i tinygo_0.26.0_amd64.deb

# This line installs WasmEdge including the AOT compiler
RUN curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash

FROM buildbase AS build
COPY go.mod go.sum .
COPY src ./src 
# Build the Wasm binary
RUN tinygo build -wasm-abi=generic -target=wasi -o main.wasm src/main.go

# This line builds the AOT Wasm binary
RUN /root/.wasmedge/bin/wasmedgec main.wasm main-wasmedge.wasm

FROM scratch
ENTRYPOINT [ "main-wasmedge.wasm" ]
COPY --link --from=build /src/main-wasmedge.wasm /main-wasmedge.wasm