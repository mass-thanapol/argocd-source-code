FROM golang:1.20 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/app .

EXPOSE 3001

CMD ["./app"]

# docker build -t app-a:v1 .
# docker run -p 3001:3001 app-a:v1
