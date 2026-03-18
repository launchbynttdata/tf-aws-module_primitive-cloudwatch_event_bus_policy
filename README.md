# Terraform AWS Module: CloudWatch Event Bus Policy

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: CC BY-NC-ND 4.0](https://img.shields.io/badge/License-CC_BY--NC--ND_4.0-lightgrey.svg)](https://creativecommons.org/licenses/by-nc-nd/4.0/)

## Overview

This Terraform module wraps the [aws_cloudwatch_event_bus_policy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_event_bus_policy) resource to manage EventBridge (CloudWatch Events) event bus resource policies. It enables cross-account event access and organization-wide permissions for event buses.

**Note:** This resource is incompatible with `aws_cloudwatch_event_permission` and will overwrite permissions if both are used.

## Pre-Commit Hooks

The [.pre-commit-config.yaml](.pre-commit-config.yaml) file defines pre-commit hooks for Terraform, Go, and common linting tasks.

The `commitlint` hook enforces conventional commit messages. See [commitlint-config-conventional](https://github.com/conventional-changelog/commitlint/tree/master/@commitlint/config-conventional#type-enum) for the format.

The `detect-secrets-hook` prevents new secrets from being introduced into the baseline. See [pre-commit](https://pre-commit.com/) for installation.

To install the commit-msg hook manually:

```
pre-commit install --hook-type commit-msg
```

## Testing Locally

1. Run `make configure` to install dependencies and tooling.
2. For AWS: ensure `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, and `AWS_DEFAULT_REGION` (or equivalent) are set.
3. Create `provider.tf` and `terraform.tfvars` in the example directory as needed for your environment.
4. Run `make check` to run lint, validate, plan, and tests.

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.5 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.14 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.100.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_cloudwatch_event_bus_policy.policy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_event_bus_policy) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_policy"></a> [policy](#input\_policy) | The text of the policy document. This should be a JSON policy document (typically created using the aws\_iam\_policy\_document data source). | `string` | n/a | yes |
| <a name="input_event_bus_name"></a> [event\_bus\_name](#input\_event\_bus\_name) | The name of the event bus to set permissions on. If omitted, permissions are set on the default event bus. | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | The ID of the resource (same as the event bus name). |
<!-- END_TF_DOCS -->
