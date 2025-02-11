# API DOCUMENTATION OF BROTOTYPE API

This is a simple application built using Golang-Fiber its like a simple wrapper around existing Brototype API Endpoints so you would be able to access all the data in this api Iam listing all the api end points too. If you are planning to access it from the frontend or creating any data around it. I was planning to create a small wrapper around it inorder to fetch all data and list then in a beautiful manner.

## AUTHENTICATION
### ENDPOINT 
POST http://localhost:8080/api/auth

### REQUEST
Content-Type : application/json
```json
{
  "mobile": "+910000000000",
  "password": "sample123"
}
```
### RESPONSE
```json
{
   "token":"asdaujdaskdas.sdasdsadasda.sdasdasdas"
}
```
### CURL REQUEST
```bash
curl -X POST http://localhost:8080/api/auth \
-H "Content-Type: application/json" \
-d '{
  "mobile": "+910000000000",
  "password": "sample123"
}'

```

### NOTE:
Remember to paste that token to fetch data in subsequent requests.<br/>
 <b>Note that the password and username to be used is the one that you have used in brototype's student portal.</b>



## DETAILS

### ENDPOINT
GET http://localhost:8080/api/details

### REQUEST
Authorization: Bearer \<Token from Previous Step>

### RESPONSE
The response would have almost all the data's that you have provided to Brototype

```bash
curl -X GET http://localhost:8080/api/details \
-H "Authorization: Bearer <TOKEN>" \
```


## REVIEWS
### ENDPOINT
GET http://localhost:8080/api/reviews

### REQUEST


```
NO BODY
```
### RESPONSE
```json
{
    "reviews":
    [
        {
            "id": "2e2207e0-1cca-4226-b3ac-636770171d4e",
            "week": "28",
            "preferredTime": "2025-02-10T00:00:00Z",
            "reviewStageCode": 0,
            "ratings": [],
            "scheduledOn": null,
            "status": null,
            "completedOn": null,
            "advisor": "Cilla Joy",
            "meetLink": null,
            "reviewType": "Normal",
            "reviewBadge": "Second Project Hosting",
            "specialType": "normal"
        },
    ]
}
```
### CURL REQUEST
```bash
curl --location --request GET 'http://localhost:8080/api/reviews' \
--header 'Authorization: Bearer <Token>' \
```

## FOUNDATIONS

### ENDPOINT
GET http://localhost:8080/api/foundations?reviewStageCode=value
### PARAMETERS
<table>
<tr>
<th>Parameter</th>
<th>value's</th>
<th>significance</th>
</tr>
<tr>
<td>reviewStageCodeAt</td>
<td>1</td>
<td>Upcoming</td>
</tr>
<tr>
<td>reviewStageCodeAt</td>
<td>2</td>
<td>Conducted</td>
</tr>
<tr>
<td>reviewStageCodeAt</td>
<td>3</td>
<td>Completed</td>
</tr>


</table>

### REQUEST
```json
NO BODY
```

### RESPONSE
```json
{
    "reviews": [
        {
            "id": "67d65a58-6f34-4052-908c-092fa3a1a864",
            "week": "1",
            "preferredTime": "2025-02-11T00:00:00Z",
            "reviewStageCode": 3,
            "scheduledOn": "2025-02-11T05:30:00Z",
            "status": "passed",
            "completedOn": "2025-02-11T06:37:29.334731Z",
            "advisor": "",
            "meetLink": "",
            "reviewType": "Foundation",
            "reviewBadge": null,
            "specialType": "technical",
            "conductedOn": "2025-02-11T06:03:53.563Z",
            "batchName": "FP001",
            "studentName": "",
            "feedback": ""
        },
    ]
}
```

### CURL REQUEST
```bash
curl --location --request GET 'http://localhost:8080/api/foundations?reviewStageCode=3' \

--header 'Authorization: Bearer TOKEN \

```

<b>NOTE:</b>

The end point for the foundations like this is the new program we can access the data this way.


### CONTRIBUTORS GUIDE