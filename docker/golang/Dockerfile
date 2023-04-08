################  build stage
ARG GOLANG_VERSION
FROM   golang:${GOLANG_VERSION}-alpine3.16  AS Builder

LABEL maintainer="daniel <adwxsghu@gmail.com>"

WORKDIR /app

COPY . .
RUN go build -o main cmd/main.go

################ run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=Builder /app/main .
COPY .env .
COPY db ./db

ENTRYPOINT [ "/app/main" ]
