FROM golang:1.10-alpine
ADD . /go/src/github.com/atomdata/go-man-page
RUN go install github.com/atomdata/go-man-page

FROM alpine:latest
COPY --from=0 /go/bin/go-man-page .
ENV PORT 8080
CMD ["./go-man-page"]