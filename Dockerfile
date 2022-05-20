FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN cd cmd/ && go build -o ../online-judge-backend
RUN apk --no-cache add gcc g++ 
EXPOSE 8000

CMD ["./online-judge-backend", "--config=./config/res/kubernetes_config.json"]
