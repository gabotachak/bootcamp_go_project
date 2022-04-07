## Steps to run the application
- Download sql script
- Import the .sql file into your database engine (mysql workbench is recommended)
- Build your .env file with the access data to your mysql database motor.(/bootcamp/.env)
- Example:
~~~
HOST = localhost:8080
DB_USER = my_username
DB_PASS = my_password
DB_NAME = my_db
DB_NAME_TEST = my_db_test
DB_HOST = localhost:3306
~~~
- Located in bootcamp/ run:
> go run cmd/server/main.go
- READY ;)


## To consider
- In some requirements, it is requested to check that the id of some record of an entity does not exist in the database. We define not to do that :), since the database itself (with the autoincremental id option) guarantees that there are no repeated id's.


## Example diagram to visualize api architecture (as an example we use a post)

![Captura de Pantalla 2021-08-18 a la(s) 12 19 31](https://user-images.githubusercontent.com/86376650/129925703-f47f1376-b393-4e85-a356-21d8cac2e070.png)

## Entity relationship diagram - database

![Captura de Pantalla 2021-08-18 a la(s) 11 51 49](https://user-images.githubusercontent.com/86376650/129920634-31bcd89d-be4b-4262-a0ed-62f8bdc0064f.png)
