FROM golang:1.21.4 AS build
WORKDIR /work
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

FROM alpine/curl
COPY --from=build /work/main /usr/bin/.
ENTRYPOINT ["main"]
