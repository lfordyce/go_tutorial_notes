.DEFAULT_GLOBAL := build

##@ build
.PHONY: build
build: ## Build the services
	for subdir in cmd/*; do \
		entrypoint=$$(basename $$subdir); \
		echo "Building $${entrypoint} binary..."; \
		go build -o bin/$${entrypoint} $$(find cmd/$${entrypoint}/*.go); \
	done

##@ Service
.PHONY: run
run: ## Run the service
	@if [ "${nobuild}" = "true" ]; then \
		echo "skipping build"; \
	else \
	  make build; \
 	fi
	@./bin/core

.PHONY: clean
clean: ## Clean up the build artifacts
	@echo Cleaning bin/ directory... && \
		cd bin/ && git clean -f -d -x