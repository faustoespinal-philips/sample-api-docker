# Sample API Docker
Repository containing example of contract-first approach developed web service.  The service is packaged as a docker container and has associated simple CI/CD pipeline using Github actions.

The code is auto-generated from the OpenAPI spec YAML contract:
```
openapi-generator-cli generate -i /local/api.yaml -g go-gin-server -o /local/src
```


