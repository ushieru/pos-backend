# syntax=docker/dockerfile:1
FROM docker.io/library/golang:1.20 AS build-stage
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o point_of_sale
FROM docker.io/library/alpine:latest AS build-release-stage
WORKDIR /app
COPY --from=build-stage /app/point_of_sale point_of_sale
COPY --from=build-stage /app/public public
EXPOSE 8080
ENTRYPOINT ["/app/point_of_sale"]