FROM golang:1.14.4-alpine3.12 as builder
WORKDIR /go/src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o snsall

FROM gcr.io/cloud-builders/gcloud:latest
WORKDIR /go/src/github_keigooba_snsall
RUN pwd

COPY --from=builder /go/src/snsall /snsall
ENTRYPOINT ["/snsall"]
