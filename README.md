# apple-maintained

## Usage

### Heroku

1. execute follwing commands

    ```bash
    $ heroku create apple-maintained-bot
    # create psql for application
    $ heroku addons:create  heroku-postgresql:hobby-dev -a apple-maintained-bot
    ...
    Created postgresql-***** as DATABASE_URL

    $ heroku pg:credentials:url postgresql-***** -a apple-maintained-bot
    # ~ credential information ~
    ```

1. copy credential information to .env

    ```bash
    $ cp .env.sample .env
    $ vim .env  # add credential information to .env
    ```

    .env

    ```vi
    HOST       : Host name of psql cluster
    DATABASE   : Database name of psql
    USER_NAME  : User name of psql
    PORT       : Port number of psql
    PASSWORD   : Login password for psql
    ```