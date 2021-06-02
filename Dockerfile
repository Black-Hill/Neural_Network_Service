FROM tensorflow/tensorflow:latest
RUN apt-get -y update
RUN apt-get -y install git
RUN apt-get -y install golang

EXPOSE 8075

ENV port = 8075

WORKDIR /go/src/neural_network_service

COPY . .

RUN go get -d -v ./...

#RUN go install -v ./...

RUN go build -o neural-network-service

CMD  ./neural-network-service