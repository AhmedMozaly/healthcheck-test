FROM golang:1.17
COPY main.go .
ENTRYPOINT [ "go", "run", "main.go" ]