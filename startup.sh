# Setup
source .env

# Start the Docker Container and the DB's script
docker compose up -d > /dev/null

echo "docker started, setting up the DB..."
docker cp ./setup.surql surrealdb:/home > /dev/null
docker exec -it surrealdb /surreal import -u "$GOTCHA_DB_USER" -p "$GOTCHA_DB_PASS" --namespace "$GOTCHA_DB_NAMESPACE" --database "$GOTCHA_DB_DATABASE"  --endpoint http://localhost:8000 /home/setup.surql