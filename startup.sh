# Setup
source .env

# Start the Docker Container and the DB's script
docker compose up -d > /dev/null
sleep 5

echo "docker started, setting up the DB..."
docker cp ./setup.surql surrealdb:/home > /dev/null
docker exec -it surrealdb /surreal import -u "$DB_USER" -p "$DB_PASS" --namespace "$DB_NAMESPACE" --database "$DB_DATABASE"  --endpoint http://localhost:8000 /home/setup.surql