SRC_DIR = "./src/"

clean :
	@rm -rf main
	@echo "message : main.exe clean"

# main.go build 
build :
	make clean
	@go build $(SRC_DIR)main.go
	@echo "message : main.go build success"

# execute main.go (dev) 
dev-run :
	nodemon --exec "go run" $(SRC_DIR)main.go	

# build && execute main.exe (prod)
prod-run :
	make build 

test :
	@go test -v src/files/files_test.go
