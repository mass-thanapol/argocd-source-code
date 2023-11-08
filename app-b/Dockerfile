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

# docker build -t app-b:v1 .
# docker run -p 3002:3002 app-a:v1
