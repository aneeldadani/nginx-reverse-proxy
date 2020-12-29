#!/bin/bash
aws_profile=s3-pcap
backup_path=/tmp
bucket_name=my-pcap-bucket

echo "S3 Sync started, sleeping for 30 seconds"
sleep 30
aws s3 sync $backup_path s3://$bucket_name --exclude *.log --include *.pcap* --profile=$aws_profile
echo "S3 Sync ran, check S3 bucket"
