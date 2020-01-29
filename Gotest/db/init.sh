docker-compose exec db bash -c "chmod 0775 ./sql/init-db.sh"
docker-compose exec db bash -c "./sql/init-db.sh"

