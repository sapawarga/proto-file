.PHONY: clean build packing

clean: 
	rm -vf ./phonebook/*.pb.go
	rm -vf ./video/*.pb.go

proto: clean
	@echo "--- Preparing proto output directories ---"
	@mkdir -p ./phonebook
	@mkdir -p ./video
	@cd ./phonebook && protoc *.proto --go_out=plugins=grpc:.
	@cd ./video && protoc *.proto --go_out=plugins=grpc:.
	@echo "--- Finished generate proto file ---"