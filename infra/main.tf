resource "aws_security_group" "rds_sg" {
  name        = "rds-security-group"
  description = "Allow MySQL inbound traffic"

  ingress {
    description = "Allow MySQL access"
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    description = "Allow all outbound traffic"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "rds-security-group"
  }
}

resource "aws_db_instance" "evermos_db" {
  identifier             = "evermos-db"
  allocated_storage      = 10
  db_name                = var.rds_db_name
  engine                 = "mysql"
  engine_version         = "8.0"
  instance_class         = "db.t3.micro"
  username               = var.rds_username
  password               = var.rds_password
  parameter_group_name   = "default.mysql8.0"
  skip_final_snapshot    = true
  vpc_security_group_ids = [aws_security_group.rds_sg.id]
  publicly_accessible    = true
}

resource "aws_security_group" "github_runner_sg" {
  name        = "github-runner-sg"
  description = "Security group for GitHub Actions Runner"

  ingress {
    description = "Allow SSH"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "Allow port 8000 for application access"
    from_port   = 8000
    to_port     = 8000
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    description = "Allow all outbound traffic"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "github-runner-sg"
  }
}

resource "tls_private_key" "ssh_key" {
  algorithm = "RSA"
  rsa_bits  = 2048
}

resource "aws_key_pair" "generated" {
  key_name   = "my-ec2-key"
  public_key = tls_private_key.ssh_key.public_key_openssh
}

resource "aws_instance" "github_runner" {
  ami             = "ami-0d22ac6a0e117cefe"
  instance_type   = "t3.micro"
  key_name        = aws_key_pair.generated.key_name
  security_groups = [aws_security_group.github_runner_sg.name]

  user_data = <<-EOF
#!/bin/bash
set -e

apt-get update -y
apt-get install -y docker.io curl git tar

systemctl enable --now docker
usermod -aG docker ubuntu

mkdir -p /home/ubuntu/actions-runner && cd /home/ubuntu/actions-runner
curl -o actions-runner-linux-x64-2.322.0.tar.gz -L https://github.com/actions/runner/releases/download/v2.322.0/actions-runner-linux-x64-2.322.0.tar.gz
tar xzf ./actions-runner-linux-x64-2.322.0.tar.gz

chown -R ubuntu:ubuntu /home/ubuntu/actions-runner
chmod +x /home/ubuntu/actions-runner/config.sh
chmod +x /home/ubuntu/actions-runner/run.sh

su - ubuntu -c "/home/ubuntu/actions-runner/config.sh --url https://github.com/nurmuh-alhakim18/evermos-project --token ${var.github_runner_token}"
su - ubuntu -c "/home/ubuntu/actions-runner/run.sh"

EOF
}

resource "aws_s3_bucket" "evermos_bucket" {
  bucket        = "my-evermos-bucket"
  force_destroy = true
}

resource "aws_s3_bucket_ownership_controls" "bucket_ownership_controls" {
  bucket = aws_s3_bucket.evermos_bucket.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_public_access_block" "bucket_public_access_block" {
  bucket = aws_s3_bucket.evermos_bucket.id

  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
}

resource "aws_s3_bucket_acl" "bucket_acl" {
  depends_on = [
    aws_s3_bucket_ownership_controls.bucket_ownership_controls,
    aws_s3_bucket_public_access_block.bucket_public_access_block,
  ]

  bucket = aws_s3_bucket.evermos_bucket.id
  acl    = "public-read"
}
