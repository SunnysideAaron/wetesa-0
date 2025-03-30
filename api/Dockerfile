FROM golang:1.23.6-bookworm
# FROM bitnami/golang:1.23-debian-12

WORKDIR /

# RUN go clean -modcache

COPY . .

WORKDIR /api

# TODO there are other live reload packages look into it.
RUN go install github.com/air-verse/air@v1.61.7
RUN go install honnef.co/go/tools/cmd/staticcheck@v0.6.1
RUN go install github.com/mgechev/revive@v1.7.0
RUN go install golang.org/x/vuln/cmd/govulncheck@v1.1.4

#RUN go mod tidy
