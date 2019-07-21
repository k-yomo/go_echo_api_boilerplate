# build
FROM golang:1.12 as build
ENV GO111MODULE=on

WORKDIR /go/src/github.com/k-yomo/go_echo_boilerplate

COPY go.mod .
COPY go.sum .
RUN go mod download && go get github.com/oxequa/realize

COPY . .

RUN make build

# exec
FROM scratch
COPY --from=build /go/src/github.com/k-yomo/go_echo_boilerplate/bin/server /go/bin/go_echo_boilerplate
ENTRYPOINT ["/go/bin/go_echo_boilerplate"]