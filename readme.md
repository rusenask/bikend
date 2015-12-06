# BikeDrop system

# Installation

It's quite simple with some container magic. Default configuration is HAProxy loadbalancing on Web applications (powered by Go)
that provide REST API for frontend (Angularjs/Ionic framework) and also are deploying it. Database persistence is kept via MongoDB, 
mapping services provided by ESRI. Authentication done via Facebook OAuth.

# Use docker compose

    docker-compose up
    
# deploy it on Tutum

[![Deploy to Tutum](https://s.tutum.co/deploy-to-tutum.svg)](https://dashboard.tutum.co/stack/deploy/)


# Install locally

Dependencies managed through glide 

    brew install glide

Install packages

    glide up
    
Use vendor experiment

    GO15VENDOREXPERIMENT=1
    
Build it

    go build
    
Run it

    ./bikend
    

# API reference (incomplete)


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
		"hostingPlaces": [{
			"id": "56635db837dd121ce8f0b257",
			"host": "karolis@rusenas4.com",
			"space": 3,
			"active": true,
			"long": 0,
			"lat": 0,
			"bookings": []
		}],
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


add new hosting place: POST /api/places
```javascript
{
    "host": "karolis@rusenas4.com",
    "space": 3,
    "long": "44.44",
    "lat": "32.23",
    "address": "string here",
     "active": true,
}
```

add new booking: POST /api/bookings

```javascript
{
    "host": "karolis@rusenas4.com",
    "user": "karolis@rusenas2.com",
}
```

get booking: GET /api/bookings?user=user@email.com