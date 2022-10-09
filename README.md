# Checkbox Backend API

## Setting up the database

1. The database being used is Postgres and the configuration for the DB URL is under "/pkg/config/envs/.envs"
2. If the database configurations for yours is different please update the DB URL as required.
3. The database name that is being used currently by default is checkbox_db

## Running the application

1. Once the database connection is setup correctly, just run "make server" and the application will be running on port 8000!
2. Currently there is only one router group available which is "/tasks"
