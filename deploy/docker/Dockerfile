FROM golang:1.18-alpine as build

WORKDIR /app
RUN mkdir build
COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go test -v ./...

RUN go build -ldflags="-w -s" shrtener.go

RUN ls -la

FROM scratch as shrtener

COPY --from=build /app/shrtener .

EXPOSE 8000

CMD [ "./shrtener"]


FROM golang:1.18-alpine as schema

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY /schema .

ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-w -s" schema.go

CMD ["./schema"]

