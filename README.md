# tftl

[![Go version](https://img.shields.io/badge/go-1.24+-blue.svg)](https://go.dev/)

**tftl** is a simple command-line utility to quickly list Terraform/OpenTofu resource targets from specified `.tf` files. It can be easily integrated with Terraform/OpenTofu commands such as `terraform plan` and `terraform apply`, making lives easier for Terraform/OpenTofu automation.

---

## Why?

If you're doing incremental applies/plans (e.g. `terraform apply -target=...`) frequently, this tool helps you extract resource names from Terraform/OpenTofu files and automatically generates the correct `-target` statements.

No manual copy-pasting resource names anymore!

---

## Installation

Make sure you have Go installed, then run:

```shell
go install github.com/tiulpin/tftl@latest
```

Or build from the source:

```shell
git clone https://github.com/tiulpin/tftl.git
cd tftl
go build -o tftl main.go
```

---

## Usage

**tftl** accepts multiple Terraform/OpenTofu files with the `-f` (or long form `--file`) flag.
Use the `-s` (or `--string`) flag if you want the resource names formatted as Terraform/OpenTofu command-line compatible arguments:

```shell
# run terraform plan quickly targeting resources from main.tf and other.tf
terraform plan $(tftl -f main.tf -f other.tf -s)

# apply specific file changes quickly
terraform apply $(tftl -f deployment.tf -s)
```