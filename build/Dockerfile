FROM golang:1.24 AS builder
RUN apk add --no-cache git
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GO111MODULE=on go build -v -o /main -installsuffix cgo -ldflags="-w -s" main.go

# second stage
FROM python:3.10-alpine

# install git
RUN apk add --no-cache git

# app directory
WORKDIR /app

RUN git clone https://github.com/dragonsecurity/ptaas-scanner.git

# update work directory
WORKDIR /app/scanner

# install deps
RUN pip install -r requirements.txt

# copy executable file
COPY --from=builder main .

# run app executeable file
CMD ["./main", "core"]
