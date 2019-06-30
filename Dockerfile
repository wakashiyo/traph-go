# step 1: build
FROM golang:1.12.4-alpine3.9 as build-step

# for go mod download
RUN apk add --update --no-cache ca-certificates git

RUN mkdir /traph-go
WORKDIR /traph-go
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o /go/bin/traph-go

# -----------------------------------------------------------------------------
# step 2: exec
FROM scratch
COPY --from=build-step /go/bin/traph-go /go/bin/traph-go
ENTRYPOINT ["/go/bin/traph-go"]