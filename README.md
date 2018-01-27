# Unite Society Backend Repo

## Generating Go Code

> swagger generate server -f swagger.yaml --exclude-main -A unite_society
> go run main.go --port 3000

## Testing The Rest API

> curl -H "Accept:application/json" -H "Content-type: application/json" -H "api_key: custom-library-164415" -X POST http://localhost:3000/v1/users -d "{ \"firstname\": \"John\", \"middlename\": null, \"lastname\": \"Doe\", \"email_ids\": [ \"user@example.com\" ], \"phone_numbers\": [ \"9998887777\" ], \"dob\": \"2018-01-26\", \"role\": \"member\", \"image_url\": \"http://example.com/profile.jpg\", \"blood_group\": \"a_positive\", \"society_uuid\": \"2488999e-d664-4e07-8078-55c5d29a7e97\", \"family_uuid\": \"34beb905-aae9-4e21-951d-6748168b40b6\", \"date_created\": \"2018-01-26T13:26:53.557Z\", \"date_updated\": \"2018-01-26T13:26:53.557Z\", \"status\": \"guest\"}"
