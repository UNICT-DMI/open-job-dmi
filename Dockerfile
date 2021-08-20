FROM golang:alpine
COPY . /app
WORKDIR /app
RUN go build -o open-jon-dmi

ENTRYPOINT [ "open-job-dmi" ]
