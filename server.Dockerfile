FROM golang:latest as builder

WORKDIR /go/src/app
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go mod download

COPY . /go/src/app
RUN go build -ldflags="-s -w" -o server ./cmd/server/server.go


FROM alpine
ENV FILE=wisdoms.txt
ENV PORT=1337
ENV BITS=5
ENV ITERATIONS=10000000
COPY --from=builder /go/src/app/server server
COPY --from=builder /go/src/app/wisdoms.txt wisdoms.txt
USER 2000
ENTRYPOINT ./server -p=$PORT -f=$FILE -b=$BITS -i=$ITERATIONS
