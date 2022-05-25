#!/usr/bin/env bash

set -e

jq -ncM 'while(true; .+1) | {method: "POST", url: "http://localhost:8000", body: ("todo=todo-no-" + (. | tostring) ) | @base64, header: {"Content-Type": ["application/x-www-form-urlencoded"]}}' | \
  vegeta attack -rate=1000/s -format=json -duration=60s -lazy | \
  vegeta report
