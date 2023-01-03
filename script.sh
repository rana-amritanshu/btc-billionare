#!/bin/bash

function openBrowser() {
    if [[ "$OSTYPE" == "darwin"* ]]; then
        open -a "Google Chrome" http://localhost:1323/swagger-ui/index.html
    fi
}

go test btc/pkg/http && docker compose up -d && openBrowser
