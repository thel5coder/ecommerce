FROM golang:1.14.0

RUN apt-get update && apt-get install -y

ENV GO111MODULE=on
ENV PKG_NAME=github.com/thel5coder/ecommerce
ENV PKG_PATH=$GOPATH/src/$PKG_NAME
WORKDIR $PKG_PATH/

COPY . $PKG_PATH/

RUN echo $PWD

WORKDIR $PKG_PATH/product-service/server/http/
RUN echo $PWD

RUN go build main.go
EXPOSE 3000
CMD ["sh", "-c", "./main"]