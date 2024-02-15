FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app

FROM alpine
WORKDIR /app/
COPY --from=builder /app/app /app/
#EXPOSE
ENTRYPOINT ["./app"]

#docker build -t social-todo-service:1.0 .

#MYSQL_GORM_DB_SOURCE=root:my-secret-pw@tcp(127.0.0.1:3309)/social-todo-list?charset=utf8mb4&parseTime=True&loc=Local;
#MYSQL_GORM_DB_SOURCE=root:my-secret-pw@tcp(my_mysql:3306)/social-todo-list?charset=utf8mb4&parseTime=True&loc=Local;
#MYSQL_GORM_DB_TYPE=mysql;SECRET=iTaskSecret2024;SIMPLE_VALUEF=iTask



docker run -d --name social-todo -p 3000:3000  -e MYSQL_GORM_DB_SOURCE="root:Trankhanh47@tcp(trankhanh-rds.c3kmq2oy4fy8.us-east-1.rds.amazonaws.com:3306)/social-todo-list?charset=utf8mb4&parseTime=True&loc=Local" -e MYSQL_GORM_DB_TYPE=mysql -e SECRET=iTaskSecret2024 social-todo-service