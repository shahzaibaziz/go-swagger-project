install_swagger:
	GO111MODULE=off | go get -u github.com/go-swagger/go-swagger/cmd/swagger
install_lint:
	GO111MODULE=off | go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
codegen:
	bash ./scripts/swagger.sh
format:
	bash ./scripts/format.sh
check: format install_lint
	bash ./scripts/check.sh
build:
	bash ./scripts/build.sh
test: check db_start
	bash ./scripts/test.sh
help:
	@echo
	@echo 'Usage: make COMMAND'
	@echo
	@echo 'Commands:'
	@echo '  build           Compile project.'
	@echo '  install_swagger          install swagger .'
	@echo '  install_golangci-lint          install go linter .'
	@echo '  check           Run linter.'
	@echo '  format          Format source code.'
	@echo '  codegen         Generate code.'
	@echo '  format          format code.'
	@echo '  test            Run test case'
	@echo