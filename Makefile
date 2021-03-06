.PHONY: clean build packing

clean: 
	rm -vf ./phonebook/*.pb.go
	rm -vf ./video/*.pb.go
	rm -vf ./userpost/*.pb.go
	rm -vf ./auth/*.pb.go

proto: clean
	@echo "--- Preparing proto output directories ---"
	@mkdir -p ./phonebook
	@mkdir -p ./video
	@mkdir -p ./userpost
	@mkdir -p ./auth
	@cd ./phonebook && protoc *.proto --go_out=plugins=grpc:.
	@cd ./video && protoc *.proto --go_out=plugins=grpc:.
	@cd ./userpost && protoc *.proto --go_out=plugins=grpc:.
	@cd ./auth && protoc *.proto --go_out=plugins=grpc:.
	@echo "--- Finished generate proto file ---"