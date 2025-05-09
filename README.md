# Dockerised Go HTTP Server with PostgreSQL and Redis Caching
summary : A backend system built in Go language that have two endpoints /reverse and /check , where the /reverse endpoint reverses the input strings and stores them in database using PostgreSQL , and /check endpoint searches the input string in database , it also uses the redis caching to speed up the search. Containerized using Docker Compose for modular deployment practices .
## Features 
- Reverses-string using HTTP endpoint (`/reverse`)
- PostgreSQL integration for storing entries ( stores input output pairs with timestamps )
- check-string using HTTP endpoint (`/check`)
- uses Redis caching for faster access
- Docker Compose integration for containerzation and modular practices
- multiple versions of servers seperated in multiple branches with different ports to listen ( ports : 8080 , 8090 , and 9090 are used where port 8090 is active for latest containerised server access )
## Project Structure 
reverse_string_app
├── Dockerfile # builds Go binary into container.
├── check_redis.go # it has the request handler function for /check endpoint : func check_function(w http.ResponseWriter , r *http.Request) {//}.
├── db.go # it contains all the functions related for the database and psql.
├── docker-compose.yml # defines Go + PostgreSQL + Redis services , hostnames , image names , dependencies and environment for docker compose step and containerization. It also has the mapping 8090:8080 which makes it able to listen to port 8090 from outside the container.
├── go.mod # Go module dependencies.
├── go.sum # Go module dependencies.
├── redis.go # It has the functions that makes the connection to redis using pointer of global connection pool and the function of checking connection by sending ping.
├── redis_utils.go # it has two functions that sets the key-value pair to cache and gets the pair from it.
├── reverse_handler.go # it has the request handler function for /reverse endpoint : func response_sender_function(w http.ResponseWriter , r *http.Request){//}.
├── reverse_string.go # it has the main function that calls request handles functions of both /reverse and /check endpoint , then starts listening to port :8080
├── run_daily.sh # its a script file that will be used to automatically turn on the server once its totally complete. and its set to use to crontab -e command where the logs will be sent to a seperate folder outside of this one.
└──closing_server_daily.sh # its a script file that will be used to automatically turn off the server once its totally complete.
## Getting Started 
### prerequisites 
- Docker and Docker Compose
- Go ( optional , for development outside container )
- PostgreSQL Client ( optional for testing )
### Running the server
```bash
docker-compose up
```
server will be available at port localhost: 8090 , send the request to this server
### Example for curl request :
for /reverse endpoint : ```bash curl -X POST http://localhost:8090/reverse -d "text=sometext" ```
output: "txetemos"
for /check endpoint : ```bash curl -X POST http://localhost:8090/check -d "text:hello" ```


### future improvements :
- will be adding more features like tests integration , better observability of logs, metrics, alerts.
- making it properly deployment ready by applying proper directories for each technologies
- some authentication methods.
- etc.

