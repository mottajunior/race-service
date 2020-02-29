FROM golang

ADD . /go/src/raceService

RUN env GIT_TERMINAL_PROMPT=1 go get github.com/examplesite/myprivaterepo
RUN go get -u ./...
RUN go install raceService

ENTRYPOINT /go/bin/raceService

EXPOSE 3000