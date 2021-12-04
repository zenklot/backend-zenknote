# Backend For Zenknote

Go backend for [Zenknote](https://zenk-note.netlify.app/): zenknote is free notes app for everyone. A simple notes application which is built on the basis of the `Create`, `Read`, `Update`, `Delete` process and authentication process with `JWT`.


## Environment
In order to run `backend zenknote`  you need to set the correct environment variables. Since [dotenv](https://www.npmjs.com/package/dotenv) is already included, it is just matter of renaming `.env.example` file to `.env`: this file include all of the mandatory defaults.

## Database
This project uses a PostgreSQL database and uses the [gorm.io](https://gorm.io/) library so you only need to create a database and connect to the database because `schemas and tables are automatically migrated`.

## Tests
Every code additions or fixs on the existing code, has to be tested. To run all tests use `go run test test/unitest_test.go`.

# Documentation
## Restful API
Read more at <a href="https://app.swaggerhub.com/apis-docs/zenklot/zenk-note_restful_api/1.0"><code><b>Zenknote Restful API</b></code></a>.

# License
Released under the [MIT license](LICENSE).