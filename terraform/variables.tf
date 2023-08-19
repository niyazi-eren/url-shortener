variable "server_port" {
  description = "The port the server will use for HTTP requests"
  type        = number
  default     = 8080
}

variable "linux_ami" {
  description = "Ubuntu 20.04 ami"
  type        = string
  default     = "ami-030633f630317131c"
}