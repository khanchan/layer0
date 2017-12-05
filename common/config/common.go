package config

import "time"

const (
	FLAG_PORT                      = "port"
	FLAG_JOB_EXPIRY                = "job-expiry"
	FLAG_LOCK_EXPIRY               = "lock-expiry"
	FLAG_DEBUG                     = "debug"
	FLAG_OUTPUT                    = "output"
	FLAG_TIMEOUT                   = "timeout"
	FLAG_INSTANCE                  = "instance"
	FLAG_ENDPOINT                  = "endpoint"
	FLAG_TOKEN                     = "token"
	FLAG_SKIP_VERIFY_VERSION       = "skip-verify-version"
	FLAG_SKIP_VERIFY_SSL           = "skip-ssl-verify"
	FLAG_NO_WAIT                   = "no-wait"
	FLAG_AWS_ACCOUNT_ID            = "account-id"
	FLAG_AWS_ACCESS_KEY            = "access-key"
	FLAG_AWS_SECRET_KEY            = "secret-key"
	FLAG_AWS_REGION                = "region"
	FLAG_AWS_VPC                   = "vpc"
	FLAG_AWS_SSH_KEY_PAIR          = "ssh-key-pair"
	FLAG_AWS_LINUX_AMI             = "linux-ami"
	FLAG_AWS_WINDOWS_AMI           = "windows-ami"
	FLAG_AWS_S3_BUCKET             = "s3-bucket"
	FLAG_AWS_INSTANCE_PROFILE      = "instance-profile"
	FLAG_AWS_PUBLIC_SUBNETS        = "public-subnets"
	FLAG_AWS_PRIVATE_SUBNETS       = "private-subnets"
	FLAG_AWS_DYNAMO_JOB_TABLE      = "job-table"
	FLAG_AWS_DYNAMO_TAG_TABLE      = "tag-table"
	FLAG_AWS_DYNAMO_LOCK_TABLE     = "lock-table"
	FLAG_AWS_LOG_GROUP_NAME        = "log-group-name"
	FLAG_AWS_TIME_BETWEEN_REQUESTS = "aws-time-between-requests"
	FLAG_AWS_MAX_RETRIES           = "aws-max-retries"
)

const (
	ENVVAR_JOB_EXPIRY                = "LAYER0_JOB_EXPIRY"
	ENVVAR_LOCK_EXPIRY               = "LAYER0_LOCK_EXPIRY"
	ENVVAR_PORT                      = "LAYER0_PORT"
	ENVVAR_DEBUG                     = "LAYER0_DEBUG"
	ENVVAR_OUTPUT                    = "LAYER0_OUTPUT"
	ENVVAR_TIMEOUT                   = "LAYER0_TIMEOUT"
	ENVVAR_INSTANCE                  = "LAYER0_INSTANCE"
	ENVVAR_ENDPOINT                  = "LAYER0_API_ENDPOINT"
	ENVVAR_TOKEN                     = "LAYER0_AUTH_TOKEN"
	ENVVAR_NO_WAIT                   = "LAYER0_NO_WAIT"
	ENVVAR_SKIP_VERIFY_VERSION       = "LAYER0_SKIP_VERSION_VERIFY"
	ENVVAR_SKIP_VERIFY_SSL           = "LAYER0_SKIP_SSL_VERIFY"
	ENVVAR_AWS_ACCOUNT_ID            = "LAYER0_AWS_ACCOUNT_ID"
	ENVVAR_AWS_ACCESS_KEY            = "LAYER0_AWS_ACCESS_KEY_ID"
	ENVVAR_AWS_SECRET_KEY            = "LAYER0_AWS_SECRET_ACCESS_KEY"
	ENVVAR_AWS_REGION                = "LAYER0_AWS_REGION"
	ENVVAR_AWS_VPC                   = "LAYER0_AWS_VPC_ID"
	ENVVAR_AWS_LINUX_AMI             = "LAYER0_AWS_LINUX_SERVICE_AMI"
	ENVVAR_AWS_WINDOWS_AMI           = "LAYER0_AWS_WINDOWS_SERVICE_AMI"
	ENVVAR_AWS_S3_BUCKET             = "LAYER0_AWS_S3_BUCKET"
	ENVVAR_AWS_SSH_KEY_PAIR          = "LAYER0_AWS_SSH_KEY_PAIR"
	ENVVAR_AWS_INSTANCE_PROFILE      = "LAYER0_AWS_ECS_INSTANCE_PROFILE"
	ENVVAR_AWS_PUBLIC_SUBNETS        = "LAYER0_AWS_PUBLIC_SUBNETS"
	ENVVAR_AWS_PRIVATE_SUBNETS       = "LAYER0_AWS_PRIVATE_SUBNETS"
	ENVVAR_AWS_DYNAMO_JOB_TABLE      = "LAYER0_AWS_DYNAMO_JOB_TABLE"
	ENVVAR_AWS_DYNAMO_TAG_TABLE      = "LAYER0_AWS_DYNAMO_TAG_TABLE"
	ENVVAR_AWS_DYNAMO_LOCK_TABLE     = "LAYER0_AWS_DYNAMO_LOCK_TABLE"
	ENVVAR_AWS_LOG_GROUP_NAME        = "LAYER0_AWS_LOG_GROUP_NAME"
	ENVVAR_AWS_TIME_BETWEEN_REQUESTS = "LAYER0_AWS_TIME_BETWEEN_REQUESTS"
	ENVVAR_AWS_MAX_RETRIES           = "LAYER0_AWS_MAX_RETRIES"
)

const (
	DefaultAWSRegion           = "us-west-2"
	DefaultTimeBetweenRequests = time.Millisecond * 10
	DefaultJobExpiry           = time.Hour * 1
	DefaultLockExpiry          = time.Hour * 24
	DefaultPort                = 9090
)
