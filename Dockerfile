# O build roda na arquitetura da máquina (rápido) e compila para a arquitetura alvo.
FROM --platform=$BUILDPLATFORM golang:1.25-alpine AS build
ARG TARGETOS TARGETARCH
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -trimpath -ldflags="-s -w" -o /out/api-golang .

# Imagem final mínima, sem shell, rodando como usuário sem privilégios.
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/api-golang /api-golang
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/api-golang"]
