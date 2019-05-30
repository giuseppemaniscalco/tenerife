FROM artifactory/golang:1.12

ADD go.mod go.sum /m/

RUN cd /m && go mod download
