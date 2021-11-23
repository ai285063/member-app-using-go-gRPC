
# protoc --go_out=. --go_opt=paths=source_relative memberApp.proto
proto:
	protoc --go_out=plugins=grpc:. *.proto
	protoc *.proto --go_out=plugins=grpc:. --go_opt=paths=source_relative *.proto