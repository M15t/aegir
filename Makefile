# Generates a help message. Borrowed from https://github.com/pydanny/cookiecutter-djangopackage.
help: ## Display this help message
	@echo "Please use \`make <target>' where <target> is one of"
	@perl -nle'print $& if m{^[\.a-zA-Z_-]+:.*?## .*$$}' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-25s\033[0m %s\n", $$1, $$2}'

migrate: ## Run database migrations
	go run functions/migration/main.go

scratch: ## Run scratch function
	go run functions/scratch/main.go