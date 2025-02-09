run command: `go run main.go` <br />
http server runs on localhost:8080 <br />
one example calling the /receipts/process endpoint: 
```
curl --location --request POST 'localhost:8080/receipts/process' \
--header 'Content-Type: application/json' \
--data-raw '{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}'
```

after receiving an id (just an example do not use this directly)
```
{"id":"857cfc6f-12b9-4db8-a8f1-d3eb069f47e2"}
```

call the /receipts/{id}/points endpoint with the id:
```
curl --location --request GET 'localhost:8080/receipts/857cfc6f-12b9-4db8-a8f1-d3eb069f47e2/points'
```

should get a response below:
```
{"points":109}
```
