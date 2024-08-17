# API Car Go

How to run this project:
1. Clone this repository
2. Open your terminal
3. For run the application using command `docker compose up -d` or for linux `sudo docker compose up -d`
4. if you want to check logs the application just using command `docker logs -f container_id` or `sudo docker logs -f container_id`

## Step to access endpoint

### Auth

**Register**
http : `http://localhost:5050/api/auth/register`

#### Example Request
```
     {
          "name": "admin",
          "phone": "12345",
          "password": "admin"
     }
```

#### Example Response
```
     {
          "message": "Success register user",
          "user": {
               "id": 1,
               "name": "admin",
               "phone": "12345",
               "password": "$2a$10$1/1Foj153lJjp.EAs9DCM./up5FDts0KpPwP9WchuTPp81vvHwSmG"
          }
     }
```

**Login**
http : `http://localhost:5050/api/auth/login`
#### Example Request
```
     {
          "phone": "12345",
          "password": "admin"
     }
```

#### Example Response
```
     {
          "message": "success login user",
          "user": {
               "id": 1,
               "name": "admin",
               "phone": "12345",
               "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjQxMTczMjgsImlkIjoxLCJuYW1lIjoiYWRtaW4iLCJwaG9uZSI6IjEyMzQ1In0.Js8BdFiNhRyIdQ3cMs1gxESvv3zZ8bsHY55gYhwJzwg"
          }
     }
```

### Car
#### make sure for copy token from response login before access API CRUD car

**Get All Car**
http : `http://localhost:5050/api/cars`

#### Example Response
##### if not have a car
```
     {
          "cars": [],
          "message": "success get cars"
     }
```

##### if have a car
```
     {
          "cars": [
          {
            "id": 1,
            "name_car": "Toyota",
            "plate_number": "F 1234 FB",
            "owner_id": 1
          }
          ],
          "message": "success get cars"
     }
```

**Get Car By ID**
http : `http://localhost:5050/api/car/1`

#### Example Response
##### if not have a car
```
     {
          "message": "internal server error"
     }
```

##### if have a car
```
     {
          "car": {
               "id": 1,
               "name_car": "Toyota",
               "plate_number": "F 1234 FB",
               "owner_id": 1
          },
          "message": "success get car"
     }
```

**Create a Car**
http : `http://localhost:5050/api/car`

#### Example Request
```
     {
          "name_car": "Toyota",
          "plate_number": "F 1234 FB"
     }    
```

#### Example Response
```
     {
          "car": {
               "id": 1,
               "name_car": "Toyota",
               "plate_number": "F 1234 FB",
               "owner_id": 1
          },
          "message": "success create car"
     }
```

**Update a Car**
http : `http://localhost:5050/api/car/1`

#### Example Request
```
     {
          "name_car": "Honda",
          "plate_number": "F 1234 FB"
     }
```

#### Example Response
```
     {
          "car": {
               "id": 1,
               "name_car": "Honda",
               "plate_number": "F 1234 FB",
               "owner_id": 1
          },
          "message": "success update car"
     }
```

**Delete a Car**
http : `http://localhost:5050/api/car/1`

#### Example Response
```
     {
          "message": "success delete car"
     }
```