FROM golang:1.17-alpine

WORKDIR /app
RUN apk --no-cache add gcc g++ linux-headers

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN cd cmd/ && go build -o ../online-judge-backend

EXPOSE 8000

CMD ["./online-judge-backend", "--config=./config/res/kubernetes_config.json"]
