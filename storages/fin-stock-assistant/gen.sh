#!/usr/bin/env bash

mysql-orm-gen -sql_file=./fin-stock-assistant.sql -orm_file=./fin-stock-assistant-gen.go -package_name="fin_stock_assistant"