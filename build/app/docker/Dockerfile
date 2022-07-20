FROM golang:1.18-alpine AS dependencies
WORKDIR /tweets-timeline-challenge
RUN apk update && apk add ca-certificates git && rm -rf /var/cache/apk/*
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

FROM golang:1.18-alpine AS build
WORKDIR /tweets-timeline-challenge
COPY --from=dependencies $GOPATH/pkg/mod $GOPATH/pkg/mod
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -gcflags "all=-N -l" -o tweets-timeline ./main.go

FROM golang:1.18-alpine AS tweets-timeline
WORKDIR /tweets-timeline-challenge
COPY --from=build /tweets-timeline-challenge/tweets-timeline tweets-timeline
RUN ls -la
CMD ["./tweets-timeline"]