FROM golang:1.21 AS build

ARG GH_PAT

WORKDIR /app

RUN git config --global url."https://${GH_PAT}:@github.com/".insteadOf "https://github.com/"

ENV GOPRIVATE=github.com/rasteiro11/*

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o mcabank-customer cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/mcabank-customer .

EXPOSE 8080

CMD ["./mcabank-customer"]

