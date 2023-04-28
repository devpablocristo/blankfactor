# Event Service

Run run the project: $ sudo make up

Stop the project:  $ sudo make down

## Endpoints

### Get all events

curl --request GET localhost:8080/api/v1/events/get-all

### Create event

curl --header "Content-Type: application/json" \
     --request POST \
     --data '{"start_time": "2023-05-16T23:30:00.000Z", "end_time": "2023-05-17T00:00:00.000Z"}' \
     http://localhost:8080/api/v1/events/create

### Get Overlaping events

curl --request GET localhost:8080/api/v1/events/get-overlaping




orm