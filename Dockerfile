FROM golang
WORKDIR /appgo
COPY . /appgo
CMD ["go","run","app.go"]