#!/usr/bin/env bash

# Credits: based on https://medium.com/@beld_pro/quick-tip-creating-a-postgresql-container-with-default-user-and-password-8bb2adb82342

# This script is used to initialize postgres, after it started running,
# to provide the database(s) and table(s) expected by a connecting
# application.

# In this case, postgres is used by a author-service microservice,
# which expects:

#   - database called publication_manager
#   - within it a table called author

#   * db user with approriate privileges to the database

set -o errexit

PUBLICATION_MANAGER_DB=${PUBLICATION_MANAGER_DB:-publication_manager}
AUTHOR_DB_TABLE=${AUTHOR_DB_TABLE:-authors}
AUTHOR_DB_USER=${AUTHOR_DB_USER:-author_user}
AUTHOR_DB_PASSWORD=${AUTHOR_DB_PASSWORD:-authorpass}
POSTGRES_USER=${POSTGRES_USER:-postgres}

# By default POSTGRES_PASSWORD is an empty string. For security reasons it is advisable
# to set set it up when we start running the container:
#
#   docker run --rm -e POSTGRES_PASSWORD=mypass -p 5432:5432 -d --name author_postgres author_postgres
#   psql -h localhost -p 5432 -U postgres

#       Note that unlike in MySQL, psql does not provide a flag for providing password.
#       The password is provided interactively.
#       The PostgreSQL image sets up trust authentication locally, so password is not required
#       when connecting from localhost (inside the same container). Ie. psql in this script, 
#       that runs after Postgres starts, does not need the authentication. 

POSTGRES_PASSWORD=${POSTGRES_PASSWORD:-}

# Debug ----------------------------------------------------
echo "==> POSTGRES_USER ... $POSTGRES_USER"
echo "==> POSTGRES_DB ... $POSTGRES_DB"
echo "==> PUBLICATION_MANAGER_DB ... $PUBLICATION_MANAGER_DB"
echo "==> AUTHOR_DB_USER ... $AUTHOR_DB_USER"
echo "==> AUTHOR_DB_PASSWORD ... [$AUTHOR_DB_PASSWORD]"
echo "==> AUTHOR_DB_TABLE ... $AUTHOR_DB_TABLE"
echo "==> POSTGRES_PASSWORD = [$POSTGRES_PASSWORD]"
# ----------------------------------------------------------

# What environment variables need to be set up.
#   Environment variable defaults are set up in this case, 
#   however we want to ensure the defaults are not accidentally
#   removed from this file causing a problem.
readonly REQUIRED_ENV_VARS=(
  "PUBLICATION_MANAGER_DB"
  "AUTHOR_DB_USER"
  "AUTHOR_DB_PASSWORD"
  "AUTHOR_DB_TABLE")

# Main execution:
# - verifies all environment variables are set
# - runs SQL code to create user and database
# - runs SQL code to create table
main() {
  check_env_vars_set
  init_user_and_db

  # Comment out if wanting to use the author-service uses gorm AutoMigrate feature
  #   the gorm AutoMigrate feature creates extra columns (xxx_unrecognized, xxx_sizecache)
  #   based on the proto message, which is required for proto message to work
  #   with the table
  # init_db_tables
}

# ----------------------------------------------------------
# HELPER FUNCTIONS

# Check if all of the required environment
# variables are set
check_env_vars_set() {
  for required_env_var in ${REQUIRED_ENV_VARS[@]}; do
    if [[ -z "${!required_env_var}" ]]; then
      echo "Error:
    Environment variable '$required_env_var' not set.
    Make sure you have the following environment variables set:
      ${REQUIRED_ENV_VARS[@]}
Aborting."
      exit 1
    fi
  done
}

# Perform initialization in the already-started PostgreSQL
#   - set up user for the author-service database:
#         this user needs to be able to create a table,
#         to insert/update and delete records
#   - create the database
init_user_and_db() {
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
     CREATE USER $AUTHOR_DB_USER WITH PASSWORD '$AUTHOR_DB_PASSWORD';
     CREATE DATABASE $PUBLICATION_MANAGER_DB;
     GRANT ALL PRIVILEGES ON DATABASE $PUBLICATION_MANAGER_DB TO $AUTHOR_DB_USER;
EOSQL
}

#   - create database tables
init_db_tables() {
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" "$PUBLICATION_MANAGER_DB" <<-EOSQL
    CREATE TABLE $AUTHOR_DB_TABLE(
    ID             CHAR VARYING(60) PRIMARY KEY NOT NULL,
    FIRST_NAME     CHAR VARYING(40) NOT NULL,
    LAST_NAME      CHAR VARYING(60) NOT NULL,
    ADDRESS        CHAR(100),
    COUNTRY        CHAR(70),
    EMAIL          CHAR(70),
    PASSWORD       CHAR VARYING(50),
    TOKEN          TEXT
);
EOSQL
}

# Executes the main routine with environment variables
# passed through the command line. Added for completeness 
# as not used here.
main "$@"
