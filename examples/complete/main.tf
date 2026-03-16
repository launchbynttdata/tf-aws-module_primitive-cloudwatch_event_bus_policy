// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

  region                = join("", split("-", data.aws_region.current.name))
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
