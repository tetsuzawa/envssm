package internal

import (
	"fmt"
	"io"
	"strings"
)

func (p *Parameter) Build(hasPlaceHolder bool) {
	if hasPlaceHolder {
		for k, v := range p.envMap {
			p.SSMBuf.WriteString(fmt.Sprintf(
				`resource "aws_ssm_parameter" "%s" {
  name           = "%s"
  description    = ""
  type           = "SecureString"
  value          = var.%s
}

`, strings.ToLower(k), k, strings.ToLower(k)))

			p.ValBuf.WriteString(fmt.Sprintf(
				`variable "%s" {
  type           = string
  description    = ""
  default        = ""
}

`, strings.ToLower(k)))

			p.TFVarsBuf.WriteString(fmt.Sprintf(
				`%s = "%s"
`,
				strings.ToLower(k), v))
		}
	} else {
		for k, v := range p.envMap {
			p.SSMBuf.WriteString(fmt.Sprintf(
				`resource "aws_ssm_parameter" "%s" {
  name           = "%s"
  type           = "SecureString"
  value          = "var.%s"
}

`, strings.ToLower(k), k, strings.ToLower(k)))

			p.ValBuf.WriteString(fmt.Sprintf(
				`variable "%s" {
  type           = string
}

`, strings.ToLower(k)))

			p.TFVarsBuf.WriteString(fmt.Sprintf(
				`%s = "%s"
`,
				strings.ToLower(k), v))
		}
	}
}

func (p *Parameter) Write(SSMPath, ValPath, TFVarsPath string) error {
	f, err := createFile(SSMPath)
	if err != nil {
		return fmt.Errorf("createFile: %w", err)
	}
	_, err = io.Copy(f, p.SSMBuf)
	if err != nil {
		return fmt.Errorf("Copy: %w", err)
	}

	f, err = createFile(ValPath)
	if err != nil {
		return fmt.Errorf("createFile: %w", err)
	}
	_, err = io.Copy(f, p.ValBuf)
	if err != nil {
		return fmt.Errorf("Copy: %w", err)
	}

	f, err = createFile(TFVarsPath)
	if err != nil {
		return fmt.Errorf("createFile: %w", err)
	}
	_, err = io.Copy(f, p.TFVarsBuf)
	if err != nil {
		return fmt.Errorf("Copy: %w", err)
	}
	return nil
}
