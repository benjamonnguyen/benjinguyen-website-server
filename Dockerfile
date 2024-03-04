FROM golang:1.22.0 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /benjinguyen-website-server

FROM gcr.io/distroless/base-debian11 AS build-release
WORKDIR /
COPY --from=build /benjinguyen-website-server /benjinguyen-website-server
EXPOSE 3000

ENTRYPOINT ["/benjinguyen-website-server"]