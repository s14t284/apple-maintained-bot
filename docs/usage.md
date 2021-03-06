# Usage

## Heroku Setup

1. execute following commands

    ```bash
    $ heroku create apple-maintained-bot
    $ heroku git:remote -a apple-maintained-bot
    # create psql for application
    $ heroku addons:create heroku-postgresql:hobby-dev -a apple-maintained-bot
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
    PSQL_HOST              : Host name of psql cluster
    PSQL_DATABASE          : Database name of psql
    PSQL_USER_NAME         : User name of psql
    PSQL_PORT              : Port number of psql
    PSQL_PASSWORD          : Login password for psql
    SLACK_CHANNEL          : Channel of Slack (default. #random)
    SLACK_USER_NAME        : Bot User name (default. AppleMaintainedBot)
    SLACK_ICON             : Bot Icon String (default. :apple:)
    SLACK_WEBHOOK_URL      : Slack Notification URL
    ```

## Deploy Setup

1. execute following commands

    ```bash
    $ heroku plugins:install heroku-config
    $ heroku config:push -a apple-maintained-bot  # reflect environment variables in .env
    $ heroku stack:set container
    $ heroku container:login
    $ heroku container:push web
    $ heroku container:release web
    ```

2. (optional) setup aws lambda function

    ```
    # (This is optional setting. but this application deployed heroku sleep 30 minites without access.)
    1. add `aws_lambda_handler.py` to aws lambda function
    2. set `EventBridge(CloudWatch Events)` triger to lambda function (cron(50 0-13,22-23 ? * * *))
    ```