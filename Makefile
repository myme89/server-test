
# gen code from .proto
# proto:
# 	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
#         --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
# 		    --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
#     	proto/*.proto



# //Create table postgess
# CREATE TABLE "data_test" (
#   "id" INT  PRIMARY KEY,
#   "name" varchar NOT NULL,
#   "full_name" varchar NOT NULL
# );

# cmd run server docker mysql, postgess, mongodb

# postgress: docker run --name postgres -d -p 5432:5432 -e POSTGRES_PASSWORD=1 postgres
# mysql:  docker run --name mysql -e MYSQL_ROOT_PASSWORD=1 -e MYSQL_USER=nhatnt -e MYSQL_DATABASE=test -e MYSQL_PASSWORD=1  -p 3306:3306 -d mysql:latest
# mongo: docker run -d -p 27017:27017 --name data-test-mongo mongo:latest
# redis: docker run -d --name some-redis -p 6379:6379 redis



test:
	go test -v -cover ./database/db
	go test -v -cover ./database/levedb
	go test -v -cover ./database/mysql
