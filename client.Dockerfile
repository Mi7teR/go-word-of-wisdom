FROM golang:latest as builder

WORKDIR /go/src/app
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go mod download

COPY . /go/src/app
RUN go build -ldflags="-s -w" -o client ./cmd/client/client.go


FROM alpine
ENV ADDR=localhost:1337
ENV ITERATIONS=1000000
COPY --from=builder /go/src/app/client client
USER 2000
ENTRYPOINT echo "Running client" && ./client -addr=$ADDR -i=$ITERATIONS