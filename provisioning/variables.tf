variable "default_tags" {
  type = "map"
  default =  {
    org = "trafficlens"
    product = "trafficlens-go"
  }
}
variable "queue_name" {
  type = "string"
  default = "terraform-example-queue"
}
variable "message_retention_seconds" {
  type = "string"
  default = "1209600"
}
variable "maxReceiveCountBeforeDLQ" {
  type = "string"
  default = "4"
}