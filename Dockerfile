FROM golang:alpine
COPY . /app
WORKDIR /app
RUN go build -o open-job-dmi

ENTRYPOINT [ "/app/open-job-dmi" ]
