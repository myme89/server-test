
# gen code from .proto
# proto-sever:
# 	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
#         --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
# 		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
#     	proto/*.proto


proto-storage:
	protoc --proto_path=server-storage/proto_storage --go_out=server-storage/pb_storage --go_opt=paths=source_relative \
        --go-grpc_out=server-storage/pb_storage --go-grpc_opt=paths=source_relative \
    	server-storage/proto_storage/*.proto


proto-process:
	protoc --proto_path=server-proccess-data/proto_processing --go_out=server-proccess-data/pb_processing --go_opt=paths=source_relative \
        --go-grpc_out=server-proccess-data/pb_processing --go-grpc_opt=paths=source_relative \
    	server-proccess-data/proto_processing/*.proto


# //Create table postgess
# CREATE TABLE "data_test" (
#   "id" INT  PRIMARY KEY,
#   "name" varchar NOT NULL,
#   "full_name" varchar NOT NULL
# );

# cmd run server docker mysql, postgess, mongodb

# postgress: docker run --name postgres -d -p 5432:5432 -e POSTGRES_PASSWORD=1 postgres
# mysql:  docker run --name mysql -e MYSQL_ROOT_PASSWORD=1 -e MYSQL_USER=nhatnt -e MYSQL_DATABASE=test -e MYSQL_PASSWORD=1  -p 3306:3306 -d mysql:latest
# service-mongo: docker run -d -p 27017:27017 --name data-test-mongo mongo:latest
# service-redis: docker run -d --name some-redis -p 6379:6379 redis

#key_release: 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

test:
	go test -v -cover ./database/...


# curl -s https://api.github.com/repos/myme89/server-test/releases/latest | grep "tarball_url" 

# curl -L -o archive.zip $(curl -s https://api.github.com/repos/myme89/server-test/releases/latest | grep "tarball_url" | cut -d : -f 2,3 | tr -d \" |tr -d \,)