FROM golang:1.25.6-trixie AS ws

RUN echo "PS1='\[\e[97m\]workspace\[\e[0m\] \[\e[96;1m\]\w\[\e[0m\] \[\e[92m\]➜\[\e[0m\] '" > /root/.bashrc && \
    echo "alias ls='ls --color=auto'" >> /root/.bashrc && \
    echo "alias art='php artisan'" >> /root/.bashrc

RUN apt update && \
    apt upgrade && \
    go install github.com/spf13/cobra-cli@latest && \
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && \
    go install github.com/pressly/goose/v3/cmd/goose@latest
