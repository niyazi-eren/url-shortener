
provider "aws" {
  region = "eu-west-3"
}

terraform {
  required_version = ">= 0.13"
}

resource "aws_instance" "url_shortener_server" {
  ami                    = var.linux_ami
  instance_type          = "t2.micro"
  vpc_security_group_ids = [aws_security_group.instance.id]
  key_name               = "url-shortener"

  user_data = data.template_file.user_data.rendered

  user_data_replace_on_change = true

  tags = {
    Name = "url-shortener"
  }
}

resource "aws_security_group" "instance" {
  name = "url-shortener-sg"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = var.server_port
    to_port     = var.server_port
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group_rule" "allow_server_all_outbound" {
  type              = "egress"
  security_group_id = aws_security_group.instance.id

  from_port   = 0
  to_port     = 0
  protocol    = -1
  cidr_blocks = ["0.0.0.0/0"]
}

data "template_file" "user_data" {
  template = file("${path.module}/user-data.sh")
}