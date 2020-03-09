# envssm
envssm is a Terraform file generator for AWS SSM written in Go.

## Installation

```terminal
go get github.com/tetsuzawa/envssm
```

## Example

### 1. Prepare .env file

```terminal
$ tree -a
.
└── .env
```

```.env:.env
DB_USER=user
DB_PASSWORD=password
```

### 2. Run envssm

```terminal
$ envssm
```

### 3. Check the output file

```terminal
$ tree -a
.
├── .env
├── ssm.tf              # generated
├── terraform.tfvars    # generated
└── variable.tf         # generated
```

```hcl-terraform:ssm.tf
# ssm.tf

resource "aws_ssm_parameter" "db_user" {
  name        = "DB_USER"
  description = ""
  type        = "SecureString"
  value       = "var.db_user"
}

resource "aws_ssm_parameter" "db_password" {
  name        = "DB_PASSWORD"
  description = ""
  type        = "SecureString"
  value       = "var.db_password"
}
```

```hcl-terraform:variables.tf
# variables.tf

variable "db_user" {
  type        = string
  description = ""
  default     = ""
}

variable "db_password" {
  type        = string
  description = ""
  default     = ""
}
```

```hcl-terraform:terraform.tfvars
# terraform.tfvars

db_user     = "user"
db_password = "password"
```

## Options

- -f: path of environment variables file
- -so: path of output SSM terraform file
- -vo: path of output variables terraform file
- -to: path of output tfvars terraform file

