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

variable "policy" {
  description = "The text of the policy document. This should be a JSON policy document (typically created using the aws_iam_policy_document data source)."
  type        = string
}

variable "event_bus_name" {
  description = "The name of the event bus to set permissions on. If omitted, permissions are set on the default event bus."
  type        = string
  default     = null
}
