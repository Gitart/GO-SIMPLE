![image](https://user-images.githubusercontent.com/3950155/200340883-64ad4700-f9d9-483b-9a8d-c526ed987e4c.png)

# LINTER
**Install linter**   
https://golangci-lint.run/usage/install/     

**Documentattion**
https://sparkbox.com/foundry/go_vet_gofmt_golint_to_code_check_in_Go   

## Install
```go
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
```

## Jet brain  
https://www.jetbrains.com/help/go/integration-with-go-tools.html   

## GitHub Actions
We recommend using our GitHub Action for running golangci-lint in CI for GitHub projects. It's fast and uses smart caching inside and it can be much faster than the simple binary installation

## Architecture 
There are the following golangci-lint execution steps:  
https://golangci-lint.run/contributing/architecture/   

## Workflow
https://golangci-lint.run/contributing/workflow/



## Usage 
```go
rem Linter   
rem https://sparkbox.com/foundry/go_vet_gofmt_golint_to_code_check_in_Go    
rem go vet   
rem golangci-lint run --disable-all -E errcheck   
```
