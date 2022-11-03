# About migration

## Requirement
For manual migration the `migrate` cli is required. 
At startup the api will migrate automatically if required.

```shell
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4
```

## To create empty migration

To creat up/down _sql_ with specified _NAME_ :
```shell
migrate create -ext sql [NAME] 
```