# URL Shortener in Go

## To run the project do the following:

- clone the repo
- `cd` into the project directory
- make sure to have redis spinning by running the following command `redis-server`
- run `go run main.go` to start our server
  - our Gin app will be listening on port 9808
- run the following curl command:

```
curl --request POST \
--data '{
    "long_url": "https://job-boards.greenhouse.io/figmrefa/jobs/5602159004?utm_source=github-vansh-ouckahl",
    "user_id" : "e0dba740-fc4b-4977-872c-d360239e6b10"
}' \
  http://localhost:9808/create-short-url
```

- should get a response that looks like the following:

```
{"message":"short url created successfully","short_url":"http://localhost:9808/P6g2cecw"}
```

- if we go to the referenced shortUrl link we should get redirected to our original link
