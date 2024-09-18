#!/bin/bash

# Run the MockGen commands
mockgen -source=core/core.go -destination=mocks/http_mock.go -package=mocks HTTPClient