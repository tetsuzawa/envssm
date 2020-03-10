resource "aws_ssm_parameter" "db_password" {
  name  = "DB_PASSWORD"
  type  = "SecureString"
  value = "var.db_password"
}

resource "aws_ssm_parameter" "db_user" {
  name  = "DB_USER"
  type  = "SecureString"
  value = "var.db_user"
}

