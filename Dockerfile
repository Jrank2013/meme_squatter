FROM golang:1.23-alpine as base
WORKDIR /github.com/jrank2013/meme_squatter

FROM base as builder
RUN apk add --no-cache make git
# RUN go get -u github.com/jessevdk/go-assets-builder
RUN go install github.com/jessevdk/go-assets-builder@latest

COPY Makefile go.mod go.sum ./
COPY cmd/ ./cmd
COPY pkg/ ./pkg

RUN ls pkg && make build-squat

FROM base

COPY --from=builder /github.com/jrank2013/meme_squatter/out/squat .

EXPOSE 8080
ENV GIN_MODE=release

CMD [ "./squat" ]

