# Complete Example

This example creates an EventBridge event bus, an IAM policy document allowing the same account to put events, and applies the policy via the cloudwatch_event_bus_policy primitive module.

## Usage

```hcl
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

module "resource_names" {
  source   = "terraform.registry.launch.nttdata.com/module_library/resource_name/launch"
  version  = "~> 2.0"
  for_each = var.resource_names_map

  logical_product_family  = var.logical_product_family
  logical_product_service = var.logical_product_service
  class_env               = var.class_env
  instance_env            = var.instance_env
  instance_resource       = var.instance_resource
  cloud_resource_type     = each.value.name
  maximum_length          = each.value.max_length

  region               = join("", split("-", data.aws_region.current.name))
  use_azure_region_abbr = false
}

resource "aws_cloudwatch_event_bus" "bus" {
  name = module.resource_names["event_bus"].standard
}

data "aws_iam_policy_document" "event_bus_policy" {
  statement {
    sid    = "AllowSameAccountPutEvents"
    effect = "Allow"
    actions = [
      "events:PutEvents"
    ]
    resources = [
      aws_cloudwatch_event_bus.bus.arn
    ]
    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"]
    }
  }
}

module "event_bus_policy" {
  source = "../.."

  policy         = data.aws_iam_policy_document.event_bus_policy.json
  event_bus_name = aws_cloudwatch_event_bus.bus.name
}
```

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_logical_product_family"></a> [logical_product_family](#input_logical_product_family) | Logical product family for resource naming. | `string` | n/a | yes |
| <a name="input_logical_product_service"></a> [logical_product_service](#input_logical_product_service) | Logical product service for resource naming. | `string` | n/a | yes |
| <a name="input_class_env"></a> [class_env](#input_class_env) | Class environment for resource naming (e.g., dev, prod). | `string` | n/a | yes |
| <a name="input_instance_env"></a> [instance_env](#input_instance_env) | Instance environment number for resource naming. | `string` | n/a | yes |
| <a name="input_instance_resource"></a> [instance_resource](#input_instance_resource) | Instance resource identifier for resource naming. | `string` | n/a | yes |
| <a name="input_resource_names_map"></a> [resource_names_map](#input_resource_names_map) | Map of key to resource_name configuration for the resource naming module. | `map(object({ name = string, max_length = optional(number, 60) }))` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output_id) | The ID of the event bus policy (same as the event bus name). |
| <a name="output_region"></a> [region](#output_region) | The AWS region where the resources were deployed. |

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.5 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.14 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_event_bus_policy"></a> [event\_bus\_policy](#module\_event\_bus\_policy) | ../.. | n/a |
| <a name="module_resource_names"></a> [resource\_names](#module\_resource\_names) | terraform.registry.launch.nttdata.com/module_library/resource_name/launch | ~> 2.0 |

## Resources

| Name | Type |
|------|------|
| [aws_cloudwatch_event_bus.bus](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_event_bus) | resource |
| [aws_caller_identity.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity) | data source |
| [aws_iam_policy_document.event_bus_policy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document) | data source |
| [aws_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_class_env"></a> [class\_env](#input\_class\_env) | Class environment for resource naming (e.g., dev, prod). | `string` | n/a | yes |
| <a name="input_instance_env"></a> [instance\_env](#input\_instance\_env) | Instance environment number for resource naming. | `string` | n/a | yes |
| <a name="input_instance_resource"></a> [instance\_resource](#input\_instance\_resource) | Instance resource identifier for resource naming. | `string` | n/a | yes |
| <a name="input_logical_product_family"></a> [logical\_product\_family](#input\_logical\_product\_family) | Logical product family for resource naming. | `string` | n/a | yes |
| <a name="input_logical_product_service"></a> [logical\_product\_service](#input\_logical\_product\_service) | Logical product service for resource naming. | `string` | n/a | yes |
| <a name="input_resource_names_map"></a> [resource\_names\_map](#input\_resource\_names\_map) | Map of key to resource\_name configuration for the resource naming module. | <pre>map(object({<br/>    name       = string<br/>    max_length = optional(number, 60)<br/>  }))</pre> | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_event_bus_name"></a> [event\_bus\_name](#output\_event\_bus\_name) | The event bus name for test assertions. |
| <a name="output_id"></a> [id](#output\_id) | The ID of the event bus policy (same as the event bus name). |
| <a name="output_region"></a> [region](#output\_region) | The AWS region where the resources were deployed. |
<!-- END_TF_DOCS -->
