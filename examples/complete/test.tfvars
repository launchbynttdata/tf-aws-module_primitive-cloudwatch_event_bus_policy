logical_product_family  = "launch"
logical_product_service = "eventbridge"
class_env               = "dev"
instance_env            = "001"
instance_resource       = "000"

resource_names_map = {
  event_bus = {
    name       = "eb"
    max_length = 64
  }
}
