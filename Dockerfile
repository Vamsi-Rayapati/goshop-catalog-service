FROM golang:1.24 AS base

# Set the working directory inside the container
WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

EXPOSE 4003

FROM base AS dev
RUN go install github.com/air-verse/air@latest
CMD ["air", "-c", ".air.toml"]

