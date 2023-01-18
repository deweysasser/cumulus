package caws

import (
	"fmt"
	aws "github.com/aws/aws-sdk-go/aws"
	cumulus "github.com/deweysasser/cumulus/cumulus"
)

// Code generated. DO NOT EDIT.  Code generated from rds_DBInstance.yaml

func (i dbinstance) GeneratedFields(builder cumulus.IFieldBuilder) {

	if i.obj.ActivityStreamEngineNativeAuditFieldsIncluded != nil {
		builder.What("activity_stream_engine_native_audit_fields_included", fmt.Sprint(boolToString(i.obj.ActivityStreamEngineNativeAuditFieldsIncluded)), cumulus.DefaultHidden)
	}

	if i.obj.ActivityStreamKinesisStreamName != nil {
		builder.What("activity_stream_kinesis_stream_name", aws.StringValue(i.obj.ActivityStreamKinesisStreamName), cumulus.DefaultHidden)
	}

	if i.obj.ActivityStreamKmsKeyId != nil {
		builder.What("activity_stream_kms_key_id", aws.StringValue(i.obj.ActivityStreamKmsKeyId), cumulus.DefaultHidden)
	}

	if i.obj.ActivityStreamMode != nil {
		builder.What("activity_stream_mode", aws.StringValue(i.obj.ActivityStreamMode), cumulus.DefaultHidden)
	}

	if i.obj.ActivityStreamPolicyStatus != nil {
		builder.What("activity_stream_policy_status", aws.StringValue(i.obj.ActivityStreamPolicyStatus), cumulus.DefaultHidden)
	}

	if i.obj.ActivityStreamStatus != nil {
		builder.What("activity_stream_status", aws.StringValue(i.obj.ActivityStreamStatus), cumulus.DefaultHidden)
	}

	if i.obj.AutoMinorVersionUpgrade != nil {
		builder.What("auto_minor_version_upgrade", fmt.Sprint(boolToString(i.obj.AutoMinorVersionUpgrade)), cumulus.DefaultHidden)
	}

	if i.obj.AutomaticRestartTime != nil {
		builder.When("automatic_restart_time", aws.TimeValue(i.obj.AutomaticRestartTime), cumulus.DefaultHidden)
	}

	if i.obj.AutomationMode != nil {
		builder.What("automation_mode", aws.StringValue(i.obj.AutomationMode), cumulus.DefaultHidden)
	}

	if i.obj.AvailabilityZone != nil {
		builder.What("availability_zone", aws.StringValue(i.obj.AvailabilityZone))
	}

	if i.obj.AwsBackupRecoveryPointArn != nil {
		builder.What("aws_backup_recovery_point_arn", aws.StringValue(i.obj.AwsBackupRecoveryPointArn), cumulus.DefaultHidden)
	}

	if i.obj.BackupRetentionPeriod != nil {
		builder.What("backup_retention_period", fmt.Sprint(aws.Int64Value(i.obj.BackupRetentionPeriod)), cumulus.DefaultHidden)
	}

	if i.obj.BackupTarget != nil {
		builder.What("backup_target", aws.StringValue(i.obj.BackupTarget), cumulus.DefaultHidden)
	}

	if i.obj.CACertificateIdentifier != nil {
		builder.What("ca_certificate_identifier", aws.StringValue(i.obj.CACertificateIdentifier), cumulus.DefaultHidden)
	}

	if i.obj.CharacterSetName != nil {
		builder.What("character_set_name", aws.StringValue(i.obj.CharacterSetName), cumulus.DefaultHidden)
	}

	if i.obj.CopyTagsToSnapshot != nil {
		builder.What("copy_tags_to_snapshot", fmt.Sprint(boolToString(i.obj.CopyTagsToSnapshot)), cumulus.DefaultHidden)
	}

	if i.obj.CustomIamInstanceProfile != nil {
		builder.What("custom_iam_instance_profile", aws.StringValue(i.obj.CustomIamInstanceProfile), cumulus.DefaultHidden)
	}

	if i.obj.CustomerOwnedIpEnabled != nil {
		builder.What("customer_owned_ip_enabled", fmt.Sprint(boolToString(i.obj.CustomerOwnedIpEnabled)), cumulus.DefaultHidden)
	}

	if i.obj.DBClusterIdentifier != nil {
		builder.What("db_cluster_identifier", aws.StringValue(i.obj.DBClusterIdentifier), cumulus.DefaultHidden)
	}

	if i.obj.DBInstanceArn != nil {
		builder.Who("db_instance_arn", aws.StringValue(i.obj.DBInstanceArn))
	}

	if i.obj.DBInstanceClass != nil {
		builder.What("db_instance_class", aws.StringValue(i.obj.DBInstanceClass))
	}

	if i.obj.DBInstanceIdentifier != nil {
		builder.What("db_instance_identifier", aws.StringValue(i.obj.DBInstanceIdentifier))
	}

	if i.obj.DBInstanceStatus != nil {
		builder.What("db_instance_status", aws.StringValue(i.obj.DBInstanceStatus))
	}

	if i.obj.DBName != nil {
		builder.Who("db_name", aws.StringValue(i.obj.DBName))
	}

	if i.obj.DbInstancePort != nil {
		builder.What("db_instance_port", fmt.Sprint(aws.Int64Value(i.obj.DbInstancePort)), cumulus.DefaultHidden)
	}

	if i.obj.DbiResourceId != nil {
		builder.What("dbi_resource_id", aws.StringValue(i.obj.DbiResourceId), cumulus.DefaultHidden)
	}

	if i.obj.DeletionProtection != nil {
		builder.What("deletion_protection", fmt.Sprint(boolToString(i.obj.DeletionProtection)), cumulus.DefaultHidden)
	}

	if i.obj.EnabledCloudwatchLogsExports != nil {
		builder.What("enabled_cloudwatch_logs_exports", fmt.Sprint((i.obj.EnabledCloudwatchLogsExports)), cumulus.DefaultHidden)
	}

	if i.obj.Engine != nil {
		builder.What("engine", aws.StringValue(i.obj.Engine))
	}

	if i.obj.EngineVersion != nil {
		builder.What("engine_version", aws.StringValue(i.obj.EngineVersion))
	}

	if i.obj.EnhancedMonitoringResourceArn != nil {
		builder.What("enhanced_monitoring_resource_arn", aws.StringValue(i.obj.EnhancedMonitoringResourceArn), cumulus.DefaultHidden)
	}

	if i.obj.IAMDatabaseAuthenticationEnabled != nil {
		builder.What("iam_database_authentication_enabled", fmt.Sprint(boolToString(i.obj.IAMDatabaseAuthenticationEnabled)), cumulus.DefaultHidden)
	}

	if i.obj.InstanceCreateTime != nil {
		builder.When("instance_create_time", aws.TimeValue(i.obj.InstanceCreateTime), cumulus.DefaultHidden)
	}

	if i.obj.Iops != nil {
		builder.What("iops", fmt.Sprint(aws.Int64Value(i.obj.Iops)), cumulus.DefaultHidden)
	}

	if i.obj.KmsKeyId != nil {
		builder.What("kms_key_id", aws.StringValue(i.obj.KmsKeyId), cumulus.DefaultHidden)
	}

	if i.obj.LatestRestorableTime != nil {
		builder.When("latest_restorable_time", aws.TimeValue(i.obj.LatestRestorableTime), cumulus.DefaultHidden)
	}

	if i.obj.LicenseModel != nil {
		builder.What("license_model", aws.StringValue(i.obj.LicenseModel), cumulus.DefaultHidden)
	}

	if i.obj.MasterUsername != nil {
		builder.What("master_username", aws.StringValue(i.obj.MasterUsername), cumulus.DefaultHidden)
	}

	if i.obj.MaxAllocatedStorage != nil {
		builder.What("max_allocated_storage", fmt.Sprint(aws.Int64Value(i.obj.MaxAllocatedStorage)), cumulus.DefaultHidden)
	}

	if i.obj.MonitoringInterval != nil {
		builder.What("monitoring_interval", fmt.Sprint(aws.Int64Value(i.obj.MonitoringInterval)), cumulus.DefaultHidden)
	}

	if i.obj.MonitoringRoleArn != nil {
		builder.What("monitoring_role_arn", aws.StringValue(i.obj.MonitoringRoleArn), cumulus.DefaultHidden)
	}

	if i.obj.MultiAZ != nil {
		builder.What("multi_az", fmt.Sprint(boolToString(i.obj.MultiAZ)), cumulus.DefaultHidden)
	}

	if i.obj.NcharCharacterSetName != nil {
		builder.What("nchar_character_set_name", aws.StringValue(i.obj.NcharCharacterSetName), cumulus.DefaultHidden)
	}

	if i.obj.NetworkType != nil {
		builder.What("network_type", aws.StringValue(i.obj.NetworkType), cumulus.DefaultHidden)
	}

	if i.obj.PerformanceInsightsEnabled != nil {
		builder.What("performance_insights_enabled", fmt.Sprint(boolToString(i.obj.PerformanceInsightsEnabled)), cumulus.DefaultHidden)
	}

	if i.obj.PerformanceInsightsKMSKeyId != nil {
		builder.What("performance_insights_kms_key_id", aws.StringValue(i.obj.PerformanceInsightsKMSKeyId), cumulus.DefaultHidden)
	}

	if i.obj.PerformanceInsightsRetentionPeriod != nil {
		builder.What("performance_insights_retention_period", fmt.Sprint(aws.Int64Value(i.obj.PerformanceInsightsRetentionPeriod)), cumulus.DefaultHidden)
	}

	if i.obj.PreferredBackupWindow != nil {
		builder.What("preferred_backup_window", aws.StringValue(i.obj.PreferredBackupWindow), cumulus.DefaultHidden)
	}

	if i.obj.PreferredMaintenanceWindow != nil {
		builder.What("preferred_maintenance_window", aws.StringValue(i.obj.PreferredMaintenanceWindow), cumulus.DefaultHidden)
	}

	if i.obj.PromotionTier != nil {
		builder.What("promotion_tier", fmt.Sprint(aws.Int64Value(i.obj.PromotionTier)), cumulus.DefaultHidden)
	}

	if i.obj.PubliclyAccessible != nil {
		builder.What("publicly_accessible", fmt.Sprint(boolToString(i.obj.PubliclyAccessible)), cumulus.DefaultHidden)
	}

	if i.obj.ReadReplicaDBClusterIdentifiers != nil {
		builder.What("read_replica_db_cluster_identifiers", fmt.Sprint((i.obj.ReadReplicaDBClusterIdentifiers)), cumulus.DefaultHidden)
	}

	if i.obj.ReadReplicaDBInstanceIdentifiers != nil {
		builder.What("read_replica_db_instance_identifiers", fmt.Sprint((i.obj.ReadReplicaDBInstanceIdentifiers)), cumulus.DefaultHidden)
	}

	if i.obj.ReadReplicaSourceDBInstanceIdentifier != nil {
		builder.What("read_replica_source_db_instance_identifier", aws.StringValue(i.obj.ReadReplicaSourceDBInstanceIdentifier), cumulus.DefaultHidden)
	}

	if i.obj.ReplicaMode != nil {
		builder.What("replica_mode", aws.StringValue(i.obj.ReplicaMode), cumulus.DefaultHidden)
	}

	if i.obj.ResumeFullAutomationModeTime != nil {
		builder.When("resume_full_automation_mode_time", aws.TimeValue(i.obj.ResumeFullAutomationModeTime), cumulus.DefaultHidden)
	}

	if i.obj.SecondaryAvailabilityZone != nil {
		builder.What("secondary_availability_zone", aws.StringValue(i.obj.SecondaryAvailabilityZone), cumulus.DefaultHidden)
	}

	if i.obj.StorageEncrypted != nil {
		builder.What("storage_encrypted", fmt.Sprint(boolToString(i.obj.StorageEncrypted)), cumulus.DefaultHidden)
	}

	if i.obj.StorageType != nil {
		builder.What("storage_type", aws.StringValue(i.obj.StorageType), cumulus.DefaultHidden)
	}

	if i.obj.TdeCredentialArn != nil {
		builder.What("tde_credential_arn", aws.StringValue(i.obj.TdeCredentialArn), cumulus.DefaultHidden)
	}

	if i.obj.Timezone != nil {
		builder.What("timezone", aws.StringValue(i.obj.Timezone), cumulus.DefaultHidden)
	}

}
