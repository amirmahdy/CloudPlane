# CloudPlane APIs

CloudPlane APIs helps developers to build, edit, monitor and maintain cloud infrastructure at ease.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You need to have Docker installed on your machine. Then you would need to install Go-Migrate [here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) to migrate DB schema to psql.

### Installing

1. Build images
```bash
sudo make build
```
2. Run app and DB
```bash
sudo make start
```
You can stop the application by:
```bash
sudo make stop
```

### Test

To run tests you can manually run following command.
```bash
sudo make test
```