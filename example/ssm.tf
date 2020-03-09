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

