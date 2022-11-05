# Model generation

https://go-zero.dev/docs/goctl/model

## From DB to Go source
The `$db` source string should be adapted based your configuration
```shell
$db = "postgres://postgres:tobechanged@localhost:5432/postgres?sslmode=disable"
goctl model pg datasource -url=$db -table="*"  -dir="model"
```