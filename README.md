# bTaskeeTest
#### Prerequisites 
- Go version: 1.21+
- Mongodb
- Docker and Docker compose

#### How to run project
1. Clone the repository to locally
2. Move to project folder
```shell 
    cd ./bTaskeeTest
```
3. Init database
```shell 
    docker-compose up -d
```
4. Start Pricing service
  - Move to folder pricing-service
```shell 
    cd ./pricing-service
```
  
  - Install dependencies:
```shell
    go mod tidy
```

  - Run service
```shell
    go run main.go
```

- Server is running at http://localhost:8071

5. Start Booking service
- Step by step as Pricing service

- Server is running at http://localhost:8072

6. Start Sending service
- Step by step as Pricing service

- Server is running at http://localhost:8073

### Database

- Can check database by connect URI `mongodb://localhost:27019` in MongoDB Compass.

### API Document

# API Get Price by date
- URL: http://localhost:8071/api/prices?date=1710981711&job_type=cleaning&date_type=normal&duration=3
- Method: GET
- Note:
    - All query params is required.
    - `date` is Unix timestamps in seconds
    - `job_type` is `cleaning` or `babysitting` .
    - `date_type` is `normal`, `peak` or `holiday`.
    - `duration` is hours work.

Example 
- Request: 
```shell
curl --location 'http://localhost:8071/api/prices?date=1710981711&job_type=cleaning&date_type=normal&duration=3'
``` 
- Response:
```json
{"price":300000}
```

# API Create Job (Booking service)
- URL: http://localhost:8072/api/jobs
- Method: POST

Example 
- Request: 
```shell
curl --location 'http://localhost:8072/api/jobs' \
    --header 'Content-Type: application/json' \
    --data '{
      "date": 1710981711,
	    "type":"cleaning",
	    "description": "Don nha",
	    "duration": 3,
	    "price": 300000, 
	    "address": "phu nhuan TPHCM" 
    }'
``` 
- Response:
```json
{
    "ID": "65fb8edfbedaa5cd786b2153",
    "Date": "2024-03-20T17:00:00Z",
    "Description": "Don nha",
    "Type": "cleaning",
    "Status": "pending",
    "Price": 300000,
    "Address": "phu nhuan TPHCM",
    "Duration": 3,
    "CreateAt": 1710984927607,
    "UpdateAt": 1710984927607,
    "DeleteAt": 0
}
```

# API Get Job by ID (Booking service)
- URL: http://localhost:8072/api/jobs/:id
- Method: GET

Example 
- Request: 
```shell
curl --location 'http://localhost:8072/api/jobs/65fb8edfbedaa5cd786b2153'
``` 
- Response:
```json
{
    "ID": "65fb8edfbedaa5cd786b2153",
    "Date": "2024-03-20T17:00:00Z",
    "Description": "Don nha",
    "Type": "cleaning",
    "Status": "pending",
    "Price": 300000,
    "Address": "phu nhuan TPHCM",
    "Duration": 3,
    "CreateAt": 1710984927607,
    "UpdateAt": 1710984927607,
    "DeleteAt": 0
}
```

# API Assign Job to Helper (Sending service)
- URL: http://localhost:8073/api/assignments
- Method: POST

Example 
- Request: 
```shell
curl --location 'http://localhost:8073/api/assignments' \
  --header 'Content-Type: application/json' \
  --data '{
    "helper_id": "65fb94efdcc5d08c316eed72",
    "job_id": "65fb90ce1c3f3a48af94f5a5"
  }'
``` 
- Response:
```json
{
  "success"
}
```

- Note: 
  - `helper_id` can be retrieved from the Helpers table in the Pricing database or use API API GET List Helper.

# API GET List Helper (Sending service)
- URL: http://localhost:8073/api/helpers
- Method: GET

Example 
- Request: 
```shell
curl --location 'http://localhost:8073/api/helpers'
``` 
- Response:
```json
{
   "helpers": [
        {
            "ID": "65fb94efdcc5d08c316eed71",
            "Name": "alex",
            "Skill": "cleaning",
            "Address": "district 1 TP HCM",
            "Phone": 84142234345,
            "Rating": 5,
            "CreateAt": 1710986479574,
            "UpdateAt": 1710986479574,
            "DeleteAt": 0
        },
        {
            "ID": "65fb94efdcc5d08c316eed72",
            "Name": "john",
            "Skill": "babysitting",
            "Address": "district 1 TP HCM",
            "Phone": 8442234345,
            "Rating": 5,
            "CreateAt": 1710986479574,
            "UpdateAt": 1710986479574,
            "DeleteAt": 0
        }
    ]
}
```