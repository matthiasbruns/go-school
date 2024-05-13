## Start MySQL in Docker

```bash
docker run --name some-mysql -e MYSQL_DATABASE=go-school -e MYSQL_ROOT_PASSWORD=my-secret-pw -p 3306:3306 -d mysql:latest 
```