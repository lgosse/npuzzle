FROM        golang

WORKDIR     /go/src/github.com/lgosse/npuzzle
COPY        . .

RUN apt-get update && apt-get install -y zsh && make