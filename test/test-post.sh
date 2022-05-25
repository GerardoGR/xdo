#!/usr/bin/env bash

set -e

BASE_URL=$1

curl -H "Content-Type: application/x-www-form-urlencoded" "$BASE_URL/todo" -d "todo=test"
