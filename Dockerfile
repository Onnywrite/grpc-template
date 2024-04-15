FROM golang:1.22.1-alpine3.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o ./bin/app.elf ./cmd/main.go

FROM alpine:3.19 AS runner

WORKDIR /lib/testapp

COPY --from=builder /app/bin ./

RUN adduser -DH usr && chown -R usr: /lib/testapp && chmod -R 700 /lib/testapp

USER usr
 
CMD [ "./app.elf" ]
