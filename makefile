# we want bash behaviour in all shell invocations
SHELL := bash
EMAIL ?=

help: # Show this help
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  help               Show this help"
	@echo "  generate-report    Generate csv with account data"
	@echo ""
	@echo "     EMAIL=example@example.com "

generate-report:
	@if [ -z "$(EMAIL)" ]; then \
		echo "Error: Invalid email. example: make generate-report EMAIL=example@example.com"; \
	else \
		echo "Generating report for email: $(EMAIL)"; \
		wget --no-check-certificate --quiet \
			--method POST \
			--timeout=0 \
			--header 'Content-Type: application/json' \
			--header 'x-api-key: key' \
			--body-data '{"account": "976133242399", "year": 2024, "receiver_email": "$(EMAIL)"}' \
			'http://localhost:8080/api/v1/transactions/generate_report'; \
		echo "Report sent to $(EMAIL)"; \
	fi