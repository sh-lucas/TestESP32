# Dockerfile

# Estágio 1: Builder - sem alterações
FROM golang:latest AS builder

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0

# Apenas para referência, o nome do binário é 'server'
RUN go build -o server .

# Estágio 2: Imagem Final - com as correções
FROM debian:bookworm-slim

# Crie e defina o diretório de trabalho
WORKDIR /web

# Copie os arquivos para o diretório de trabalho que acabamos de definir
# O binário vai para /web/server
# A pasta frontend vai para /web/frontend
COPY --from=builder /app/server .
COPY --from=builder /app/frontend ./frontend

EXPOSE 8080

# Execute o binário a partir do diretório de trabalho atual (/web)
ENTRYPOINT ["./server"]
