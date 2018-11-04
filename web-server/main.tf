provider "aws" {
  region = "eu-central-1"
}

resource "aws_instance" "web_server" {
  ami                    = "ami-0dd0be70cc0d493b7"
  instance_type          = "t2.micro"
  key_name               = "kacdab"
  vpc_security_group_ids = ["${aws_security_group.web_server.id}"]

  user_data = "${file("user-data.txt")}"
}

resource "aws_security_group" "web_server" {
  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

output "url" {
  value = "http://${aws_instance.web_server.public_ip}:8080"
}
