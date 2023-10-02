# syntax=docker/dockerfile:1
FROM docker.io/library/golang:1.20 AS build-stage
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=1 GOOS=linux go build -o point_of_sale
FROM docker.io/library/ubuntu:jammy-20230624 AS build-release-stage
WORKDIR /
COPY --from=build-stage /app/point_of_sale /point_of_sale
EXPOSE 8080
ENTRYPOINT ["/point_of_sale"]