output "rds_address" {
  value = aws_db_instance.evermos_db.address
}

output "rds_port" {
  value = aws_db_instance.evermos_db.port
}

output "rds_db_name" {
  value = aws_db_instance.evermos_db.db_name
}

output "s3_bucket_name" {
  value = aws_s3_bucket.evermos_bucket.id
}

output "ec2_address" {
  value = aws_instance.github_runner.public_ip
}