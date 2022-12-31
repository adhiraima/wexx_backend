FROM golang
RUN mkdir /wexx
ADD . /wexx
WORKDIR /wexx
RUN go build -o main .
CMD ["/wexx/main"]