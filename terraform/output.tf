output "public_ip" {
  value       = aws_instance.url_shortener_server.public_dns
  description = "The public IP address of the web server"
}