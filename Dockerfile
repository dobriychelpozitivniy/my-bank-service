FROM golang

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o bank-app ./cmd/main.go

EXPOSE 8000

CMD ["./bank-app"]

