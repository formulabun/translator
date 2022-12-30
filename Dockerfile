FROM golang:1.18 AS build
WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o translator server/*.go

FROM scratch AS runtime
COPY --from=build /go/src/translator ./
EXPOSE 5092/tcp
ENTRYPOINT ["./translator"]
