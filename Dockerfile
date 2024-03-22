FROM golang:1.20-alpine3.19 AS build_base
RUN apk add --no-cache git
RUN apk add make

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY Makefile .

RUN go mod tidy && go mod vendor && go mod download

COPY . .

RUN make build

FROM alpine:3.9 
RUN apk add ca-certificates
WORKDIR /app

COPY --from=build_base /app/build/test-be ./
COPY --from=build_base /app/config/config.staging.yml ./config/
COPY --from=build_base /app/config/config.production.yml ./config/
COPY --from=build_base /app/docs/swagger.yml ./docs/
COPY --from=build_base /app/.env.production ./
COPY --from=build_base /app/.env.staging ./


CMD ["./test-be"]