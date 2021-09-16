FROM golang:1.17-alpine AS build

ENV CGO_ENABLED=0

RUN apk add --no-cache \
      git \
      curl \
      make

WORKDIR /src
COPY . .
RUN make install

FROM gcr.io/distroless/base

ARG EXECUTABLE
ENV EXECUTABLE=$EXECUTABLE

COPY --from=build /go/bin/${EXECUTABLE} /metacreds
ENTRYPOINT ["/metacreds"]
