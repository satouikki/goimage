#!/usr/bin/env bash
mysql -u golang-test-user -pgolang-test-pass golang-test-database < "/docker-entrypoint-initdb.d/mysql-table-create.sql"
