```
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```

```
oapi-codegen -generate "types,server" -package api -o api.gen.go openapi.yml
```
