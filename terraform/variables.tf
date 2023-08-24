variable "server_port" {
  description = "The port the server will use for HTTP requests"
  type        = number
  default     = 80
}

variable "linux_ami" {
  description = "Ubuntu 20.04 ami"
  type        = string
  default     = "ami-030633f630317131c"
}

variable "region" {
  description = "aws region"
  type        = string
  default     = "eu-west-3"
}

variable "backend_bucket_name" {
  type = string
  default = "terraform-url-shortener-state"
}

variable "backend_locks" {
  type = string
  default = "url-shortener-locks"
}