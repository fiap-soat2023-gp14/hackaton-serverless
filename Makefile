## testing
.PHONY: test
test: 
	@mkdir -p coverage
	@cd src && go test -coverprofile=../coverage/profile.cov -coverpkg=./... ./... && go tool cover -func=../coverage/profile.cov && go tool cover -html ../coverage/profile.cov -o=../coverage/index.html
