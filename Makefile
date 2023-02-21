proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		    --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
    	proto/*.proto



CREATE TABLE "data_test" (
  "id" INT  PRIMARY KEY,
  "name" varchar NOT NULL,
  "full_name" varchar NOT NULL
);