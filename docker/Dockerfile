# builder stage

FROM golang:latest as builder

ARG DOCKER_RUNNING=yes
ARG DEBIAN_FRONTEND=noninteractive
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1

COPY . /project

RUN /project/build/build-server.sh prod

# clean stage

FROM alpine

COPY --from=builder /project/dist/prod/hermes-next /hermes

ENV STATIC_DIR="/static"

EXPOSE 30023

CMD ["/hermes-server"]
