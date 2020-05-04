# The load a truck service

Create a service that support our parcel company. For this challenge we want to be able to create trucks and load and unload it with parcels as requested. We want to be able to get the number of parcels we have in each truck and the weight of the truck, each parcel can have a different weight it should be defined when it is created.

Some additional non-functional requirements
Your services need to be accessible via HTTP.
It should be built in PHP, Golang or Nodejs(javascript and typescript)
Tests are welcome

## Create Truck
```
curl -X POST http://localhost:8080/trucks -d '
{
 "model": "m1",
 "capacityKg": 100.5
}'
```

## Get Trucks
```
curl -X GET http://localhost:8080/trucks
```

## Get Truck By ID
```
curl -X GET http://localhost:8080/trucks/1
```

## Load Parcel to Truck
```
curl -X POST http://localhost:8080/trucks/1/parcels -d '
{
 "weightKg": 10.5
}'
```

## Unload Parcel from Truck
```
curl -X POST http://localhost:8080/trucks/1/parcels/1
```
