FROM golang:1.17.5
COPY main.go .
ENTRYPOINT [ "go", "run", "main.go" ]
