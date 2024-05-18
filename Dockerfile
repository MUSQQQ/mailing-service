ARG image
FROM ${image} as builder
ENV GO_FLAGS="-mod=vendor"
WORKDIR /code
COPY . .
RUN make install

FROM registry.hub.docker.com/library/alpine:latest

WORKDIR /root/

COPY --from=builder /go/bin/mailing-service .

RUN apk add --no-cache bash
RUN apk add gcompat

RUN adduser -D user

USER root 
RUN chmod +x .

USER user

CMD ["/root/mailing-service"]
