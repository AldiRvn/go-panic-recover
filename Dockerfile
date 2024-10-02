FROM golang:alpine
RUN apk --no-cache add tzdata

WORKDIR code
COPY . .

ENV TZ=Asia/Jakarta
ENTRYPOINT ["go","run","."]