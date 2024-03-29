## Builder Image
FROM golang:1.13-stretch as builder
ENV GO111MODULE=on
COPY . /c4-active
WORKDIR /c4-active
RUN go get github.com/google/wire/cmd/wire
RUN wire
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/application

## Run Image
FROM scratch
COPY --from=builder /c4-active/bin/application application
EXPOSE 8000
ENTRYPOINT ["./application"]