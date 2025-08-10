FROM golang:1.22-alpine AS builder
WORKDIR /build

# Copie apenas os arquivos de dependência primeiro
COPY go.mod ./
RUN go mod download

# Agora copie o restante do código
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o app

FROM alpine:3.19 AS final
RUN adduser --uid 1000 --disabled-password appuser
USER appuser
WORKDIR /app
COPY --from=builder --chown=appuser:appuser /build/app .
EXPOSE 5000
CMD [ "./app" ]