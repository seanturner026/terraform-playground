# terraform-playground

Simple helper script that creates the files needed for a terraform stack. Useful whenever I need to do a quick Terraform sandbox environment.

```shell
$ go build .

$ tree
.
├── README.md
├── go.mod
├── go.sum
├── main.go
├── templates
│   ├── main.tf
│   ├── outputs.tf
│   ├── provider.tf
│   ├── terraform.tfvars
│   └── variables.tf
└── terraform-playground

$ ./terraform-playground -s cloudfront-functions
Templated stack cloudfront-functions successfully.

$ tree
.
├── README.md
├── cloudfront-functions
│   ├── main.tf
│   ├── outputs.tf
│   ├── provider.tf
│   ├── terraform.tfvars
│   └── variables.tf
├── go.mod
├── go.sum
├── main.go
├── templates
│   ├── main.tf
│   ├── outputs.tf
│   ├── provider.tf
│   ├── terraform.tfvars
│   └── variables.tf
└── terraform-playground
```
