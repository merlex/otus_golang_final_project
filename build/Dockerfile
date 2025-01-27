# Собираем в гошке
FROM golang:1.22 as build

ENV BIN_FILE /opt/previewer/previewer
ENV APP_DIR /opt/previewer/
ENV CODE_DIR /go/src/

RUN mkdir -p ${APP_DIR}
WORKDIR ${CODE_DIR}

# Кэшируем слои с модулями
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY configs/ ${APP_DIR}
COPY .. ${CODE_DIR}
# Собираем статический бинарник Go (без зависимостей на Си API),
# иначе он не будет работать в alpine образе.
ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/*

# На выходе тонкий образ
FROM alpine:3.9

LABEL ORGANIZATION="OTUS Online Education"
LABEL SERVICE="previewer"
LABEL MAINTAINERS="a@merlex.ru"

ENV BIN_FILE "/opt/previewer/previewer"
ENV APP_DIR /opt/previewer/

RUN mkdir -p ${APP_DIR}

COPY --from=build ${APP_DIR} ${APP_DIR}
WORKDIR ${APP_DIR}

CMD ${BIN_FILE}
