# Unite Society Backend Repo

## Generating Go Code

> swagger generate server -t gen -f swagger.yaml --exclude-main -A unite_society

## Testing It
> curl -X GET "http://localhost/v1/users" -H "accept: application/json"

## Running It
CD into Server folder
> dev_appserver.py app.yaml

## Deploying It
gcloud app deploy --version dev
