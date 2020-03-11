---
title: .env から AWS パラメータストアの Terraform ファイルを生成するツールを作った
tags: .env AWS SSM Terraform
author: tetsuzawa
slide: false
---

# envssm

Docker/docker-compose で.env を使っていると本番環境のシークレットの管理がめんどうですよね。

シークレットは AWS のパラメータストアに登録して、ECS など環境変数として利用するのが一般的のようですが、一つ一つコンソールに打ち込むのは骨が折れます。

なので.env ファイルからパラメータストアの Terraform ファイルを生成するツールを作りました。

# 対象読者

- Terraform の基本がわかっている方
- docker-compose などで.env を使っている方

Terraform や Docker の使い方には触れませんのでご了承ください。

# インストール

```terminal
go get github.com/tetsuzawa/envssm
```

## Example

### 1. .env を用意

```terminal
$ tree -a
.
└── .env
```

```.env:.env
DB_USER=user
DB_PASSWORD=password
```

### 2. envssm を実行

```terminal
$ envssm
```

### 3. 出力ファイルを確認

```terminal
$ tree -a
.
├── .env
├── ssm.tf              # generated
├── terraform.tfvars    # generated
└── variable.tf         # generated
```

```hcl-terraform:ssm.tf
resource "aws_ssm_parameter" "db_user" {
  name  = "DB_USER"
  type  = "SecureString"
  value = var.db_user
}

resource "aws_ssm_parameter" "db_password" {
  name  = "DB_PASSWORD"
  type  = "SecureString"
  value = var.db_password
}
```

```hcl-terraform:variables.tf
variable "db_user" {
  type = string
}

variable "db_password" {
  type = string
}
```

```hcl-terraform:terraform.tfvars
db_user     = "user"
db_password = "password"
```

## オプション

- -f: .env ファイルのパス (prod.env や ./build/dev.env など)
- -d: `description = ""`などのプレースホルダを生成
- -so: 出力される ssm.tf のパス
- -vo: 出力される variables.tf のパス
- -to: 出力される terraform.tfvars のパス

# 説明

- 出力ファイルを Terraform の構成に加えていただければパラメータストアに環境変数を登録することができます。
- パラメータストアから環境変数を読み込むために EC2 や ECS 側で参照設定を忘れないようにしてください

# 終わりに

私の場合、パラメータストアを使うことで環境変数の注入を AWS 側に任せることができるようになり、GitHub 上での CD の方法に悩まされることがなくなりました。

設計など適当に短時間で作ってしまったため、修正依頼お待ちしてます。
https://github.com/tetsuzawa/envssm
