#!/usr/bin/env bash
mysql -u golang-test-user -p golang-test-pass golang-test-database < "../sql/mysql-table-create.sql"
