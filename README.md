Run the docker compose file to start DB
```bash
docker-compose up
```
Run the application and then navigate to http://localhost:8080/swagger/index.html

Run Application in development mode
```bash
ENV=development go run main.go
```

Run Application in production mode
```bash
ENV=production go run main.go
ENV=production ./main
```