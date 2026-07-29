FROM golang:1.26-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /comicread .

FROM scratch

WORKDIR /app

COPY --from=builder /comicread /usr/local/bin/comicread

ENV XDG_CONFIG_HOME=/app/.config

EXPOSE 55566

ENTRYPOINT ["comicread", "--web"]
