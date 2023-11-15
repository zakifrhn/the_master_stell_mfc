#! buat container
FROM golang:1.20.4-alpine

#! buat folder untuk menyimpan code
WORKDIR /goapp

#! Copy semua file
#? titik yg kedua lokasi penyimpanan
COPY . .

#! instal depedency
#? run kedua berfungsi untuk build app
RUN go mod download
RUN go build -v -o /gotestdev/goback ./cmd/main.go

#!open port
EXPOSE 8590

#! run app
ENTRYPOINT [ "/gotestdev/goback" ]

#! docker run --name <nama container> --net <info di networks> -e DB_HOST=<nama image postgres di docker> -p <port luar>:<port dalam> <nama image>