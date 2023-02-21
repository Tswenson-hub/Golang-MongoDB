# Golang-MongoDB
Connection to locally hosted Mongo DB with handling for basic CRUD routes. The example use case is basic storage of movies.

Logic is split into 3 different modules. Controller handles helper functions for routing functions driven by requests. Requests are handled within the router folder. 
Main.go initiates the server. Models.go simply maintains the basic struct that models the data for storage in MongoDB.
