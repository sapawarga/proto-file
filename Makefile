.PHONY: clean build packing

clean: 
	rm -vf ./phonebook/*.pb.go

proto: clean
	@echo "--- Preparing proto output directories ---"
	@mkdir -p ./phonebook
	@cd ./phonebook && protoc *.proto --go_out=plugins=grpc:.
	@echo "--- Finished generate proto file ---"