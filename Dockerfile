FROM golang:1.21.4 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /benjinguyen-website-server

FROM gcr.io/distroless/base-debian11 AS build-release
WORKDIR /
COPY --from=build /benjinguyen-website-server /benjinguyen-website-server
COPY --from=build /app/public /public
EXPOSE 3000

ENTRYPOINT ["/benjinguyen-website-server"]