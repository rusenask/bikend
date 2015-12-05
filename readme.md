API


add new user: POST /api/users
```javascript
{
    "userID": "karolis@rusenas4.com",
    "profilePic": "http://somehwere",
    "firstName": "karolis",
    "lastName": "rusenas"
}
```

get all users: GET /api/users
```javascript
{
	"data": [{
		"id": "5663430637dd12e2022d258d",
		"bikeLocation": {
			"id": "",
			"host": "",
			"space": 0,
			"active": false,
			"long": 0,
			"lat": 0,
			"bookings": null
		},
		"userID": "karolis@rusenas2.com",
		"profilePic": "http://somehwere",
		"firstName": "karolis",
		"lastName": "rusenas"
	}, {
		"id": "566346a337dd12f268241c0c",
		"bikeLocation": {
			"id": "",
			"host": "",
			"space": 0,
			"active": false,
			"long": 0,
			"lat": 0,
			"bookings": null
		},
		"userID": "karolis@rusenas4.com",
		"profilePic": "http://somehwere",
		"firstName": "karolis",
		"lastName": "rusenas"
	}]
}
```


get specific user: GET /api/users?q=karolis@rusenas4.com
```javascript
{
	"data": {
		"id": "566346a337dd12f268241c0c",
		"bikeLocation": {
			"id": "",
			"host": "",
			"space": 0,
			"active": false,
			"long": 0,
			"lat": 0,
			"bookings": null
		},
		"userID": "karolis@rusenas4.com",
		"profilePic": "http://somehwere",
		"firstName": "karolis",
		"lastName": "rusenas"
	}
}
```