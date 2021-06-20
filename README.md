# Qashir Backend 
Click here for a [Demo](https://qashir.ranggarifqi.com)

This repo contains backend logic for a really simple url shortener app.

And i think i'm gonna make this repo as a template for my golang starter code

* It has an easy to use migration tool, 
* No cyclic dependency so far.
* Unit test & mock included
* It also follows the Clean Arch, well..kind of
* Dockerized
* And it has CI/CD pipeline using Travis CI, and deployment to Heroku.

I really want to use AWS Elastic Beanstalk & RDS though, but given the monthly cost and the simplicity of this app, i think it isn't worth the cost. lol

## How to Install
1. Clone this repo somewhere outside of gopath
2. `cp .env.example .env`
3. Fill .env config. If you specify DATABASE_URL, then the all the PG_ variable will be ignored
If you want to use docker, just use this .env configuration
```
GO_ENV=development

PORT=3000

PG_HOST=postgres
PG_PORT=5432
PG_USER=postgres
PG_PASSWORD=postgres_password
PG_DBNAME=postgres
PG_SSLMODE=disable
```

## How to Run
### The Easy Way
1. `make run-docker`
2. done. lol

### The Not As Easy As Docker Way
1. Create a new PostgreSQL db
2. Update the .env file
3. `make run`
4. done

## Note
1. Use an authentication method if you want to build more complex app using this repo
2. Swagger would be good
3. Might need to delete travis.yml & deploy.sh for a different project