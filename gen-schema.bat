@echo off
setlocal EnableDelayedExpansion

set DB_NAME=temp_schema_db
set DB_USER=postgres
set DB_PASSWORD=password
set MIGRATIONS_DIR=migrations
set OUTPUT_FILE=schema.sql
set CONTAINER_NAME=temp_postgres

echo Starting container...
docker run --rm -d --name "%CONTAINER_NAME%" -e POSTGRES_PASSWORD="%DB_PASSWORD%" postgres:18.0-alpine3.22
timeout /t 5 /nobreak >nul

docker exec "%CONTAINER_NAME%" psql -U "%DB_USER%" -c "CREATE DATABASE %DB_NAME%;"

echo Applying migrations...
for /f "delims=" %%f in ('dir "%MIGRATIONS_DIR%\*.sql" /b /on') do (
    echo Processing: %%f
    docker cp "%MIGRATIONS_DIR%\%%f" "%CONTAINER_NAME%":/tmp/migration.sql
    docker exec "%CONTAINER_NAME%" psql -U "%DB_USER%" -d "%DB_NAME%" -q -f /tmp/migration.sql
)

echo Dumping schema...
docker exec "%CONTAINER_NAME%" pg_dump -U "%DB_USER%" -d "%DB_NAME%" ^
  --schema-only ^
  --no-owner ^
  --no-privileges ^
  --no-comments ^
  --no-publications ^
  --no-subscriptions ^
  --no-security-labels ^
  --no-tablespaces ^
  --no-blobs ^
  --no-unlogged-table-data ^
  --no-sync ^
  -n public > "%OUTPUT_FILE%"

docker stop "%CONTAINER_NAME%" >nul

echo Done!
echo File generated: %OUTPUT_FILE%