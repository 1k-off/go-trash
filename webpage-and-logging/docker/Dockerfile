FROM golang:1-alpine as build
WORKDIR /app
RUN apk update && apk add --no-cache ca-certificates git
COPY go.mod ./
RUN go mod download
COPY . .
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/ugg-demo

FROM scratch
LABEL maintainer="Bogdan Kosarevskyi <bogdan.kosarevskyi@gmail.com>"
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/ugg-demo /app/ugg-demo
EXPOSE 8080
CMD ["/app/ugg-demo"]