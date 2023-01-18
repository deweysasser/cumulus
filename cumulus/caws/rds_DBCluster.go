package caws

import (
	"fmt"
	aws "github.com/aws/aws-sdk-go/aws"
	cumulus "github.com/deweysasser/cumulus/cumulus"
)

// Code generated. DO NOT EDIT.  Code generated from rds_DBCluster.yaml

func (i dbcluster) GeneratedFields(builder cumulus.IFieldBuilder) {

	if i.obj.ActivityStreamKinesisStreamName != nil {
		builder.What("activity_stream_kinesis_stream_name", aws.StringValue(i.obj.ActivityStreamKinesisStreamName), cumulus.DefaultHidden)
	}

	if i.obj.ActivityStreamKmsKeyId != nil {
		builder.What("activity_stream_kms_key_id", aws.StringValue(i.obj.ActivityStreamKmsKeyId), cumulus.DefaultHidden)
	}

	if i.obj.ActivityStreamMode != nil {
		builder.What("activity_stream_mode", aws.StringValue(i.obj.ActivityStreamMode), cumulus.DefaultHidden)
	}

	if i.obj.ActivityStreamStatus != nil {
		builder.What("activity_stream_status", aws.StringValue(i.obj.ActivityStreamStatus), cumulus.DefaultHidden)
	}

	if i.obj.AllocatedStorage != nil {
		builder.What("allocated_storage", fmt.Sprint(aws.Int64Value(i.obj.AllocatedStorage)), cumulus.DefaultHidden)
	}

	if i.obj.AutoMinorVersionUpgrade != nil {
		builder.What("auto_minor_version_upgrade", fmt.Sprint(boolToString(i.obj.AutoMinorVersionUpgrade)), cumulus.DefaultHidden)
	}

	if i.obj.AutomaticRestartTime != nil {
		builder.When("automatic_restart_time", aws.TimeValue(i.obj.AutomaticRestartTime), cumulus.DefaultHidden)
	}

	if i.obj.AvailabilityZones != nil {
		builder.What("availability_zones", fmt.Sprint((i.obj.AvailabilityZones)), cumulus.DefaultHidden)
	}

	if i.obj.BacktrackConsumedChangeRecords != nil {
		builder.What("backtrack_consumed_change_records", fmt.Sprint(aws.Int64Value(i.obj.BacktrackConsumedChangeRecords)), cumulus.DefaultHidden)
	}

	if i.obj.BacktrackWindow != nil {
		builder.What("backtrack_window", fmt.Sprint(aws.Int64Value(i.obj.BacktrackWindow)), cumulus.DefaultHidden)
	}

	if i.obj.BackupRetentionPeriod != nil {
		builder.What("backup_retention_period", fmt.Sprint(aws.Int64Value(i.obj.BackupRetentionPeriod)), cumulus.DefaultHidden)
	}

	if i.obj.Capacity != nil {
		builder.What("capacity", fmt.Sprint(aws.Int64Value(i.obj.Capacity)), cumulus.DefaultHidden)
	}

	if i.obj.CharacterSetName != nil {
		builder.What("character_set_name", aws.StringValue(i.obj.CharacterSetName), cumulus.DefaultHidden)
	}

	if i.obj.CloneGroupId != nil {
		builder.What("clone_group_id", aws.StringValue(i.obj.CloneGroupId), cumulus.DefaultHidden)
	}

	if i.obj.ClusterCreateTime != nil {
		builder.When("cluster_create_time", aws.TimeValue(i.obj.ClusterCreateTime), cumulus.DefaultHidden)
	}

	if i.obj.CopyTagsToSnapshot != nil {
		builder.What("copy_tags_to_snapshot", fmt.Sprint(boolToString(i.obj.CopyTagsToSnapshot)), cumulus.DefaultHidden)
	}

	if i.obj.CrossAccountClone != nil {
		builder.What("cross_account_clone", fmt.Sprint(boolToString(i.obj.CrossAccountClone)), cumulus.DefaultHidden)
	}

	if i.obj.CustomEndpoints != nil {
		builder.What("custom_endpoints", fmt.Sprint((i.obj.CustomEndpoints)), cumulus.DefaultHidden)
	}

	if i.obj.DBClusterArn != nil {
		builder.What("db_cluster_arn", aws.StringValue(i.obj.DBClusterArn), cumulus.DefaultHidden)
	}

	if i.obj.DBClusterIdentifier != nil {
		builder.What("db_cluster_identifier", aws.StringValue(i.obj.DBClusterIdentifier), cumulus.DefaultHidden)
	}

	if i.obj.DBClusterInstanceClass != nil {
		builder.What("db_cluster_instance_class", aws.StringValue(i.obj.DBClusterInstanceClass), cumulus.DefaultHidden)
	}

	if i.obj.DBClusterParameterGroup != nil {
		builder.What("db_cluster_parameter_group", aws.StringValue(i.obj.DBClusterParameterGroup), cumulus.DefaultHidden)
	}

	if i.obj.DBSubnetGroup != nil {
		builder.What("db_subnet_group", aws.StringValue(i.obj.DBSubnetGroup), cumulus.DefaultHidden)
	}

	if i.obj.DatabaseName != nil {
		builder.What("database_name", aws.StringValue(i.obj.DatabaseName), cumulus.DefaultHidden)
	}

	if i.obj.DbClusterResourceId != nil {
		builder.What("db_cluster_resource_id", aws.StringValue(i.obj.DbClusterResourceId), cumulus.DefaultHidden)
	}

	if i.obj.DeletionProtection != nil {
		builder.What("deletion_protection", fmt.Sprint(boolToString(i.obj.DeletionProtection)), cumulus.DefaultHidden)
	}

	if i.obj.EarliestBacktrackTime != nil {
		builder.When("earliest_backtrack_time", aws.TimeValue(i.obj.EarliestBacktrackTime), cumulus.DefaultHidden)
	}

	if i.obj.EarliestRestorableTime != nil {
		builder.When("earliest_restorable_time", aws.TimeValue(i.obj.EarliestRestorableTime), cumulus.DefaultHidden)
	}

	if i.obj.EnabledCloudwatchLogsExports != nil {
		builder.What("enabled_cloudwatch_logs_exports", fmt.Sprint((i.obj.EnabledCloudwatchLogsExports)), cumulus.DefaultHidden)
	}

	if i.obj.Endpoint != nil {
		builder.What("endpoint", aws.StringValue(i.obj.Endpoint), cumulus.DefaultHidden)
	}

	if i.obj.Engine != nil {
		builder.What("engine", aws.StringValue(i.obj.Engine), cumulus.DefaultHidden)
	}

	if i.obj.EngineMode != nil {
		builder.What("engine_mode", aws.StringValue(i.obj.EngineMode), cumulus.DefaultHidden)
	}

	if i.obj.EngineVersion != nil {
		builder.What("engine_version", aws.StringValue(i.obj.EngineVersion), cumulus.DefaultHidden)
	}

	if i.obj.GlobalWriteForwardingRequested != nil {
		builder.What("global_write_forwarding_requested", fmt.Sprint(boolToString(i.obj.GlobalWriteForwardingRequested)), cumulus.DefaultHidden)
	}

	if i.obj.GlobalWriteForwardingStatus != nil {
		builder.What("global_write_forwarding_status", aws.StringValue(i.obj.GlobalWriteForwardingStatus), cumulus.DefaultHidden)
	}

	if i.obj.HostedZoneId != nil {
		builder.What("hosted_zone_id", aws.StringValue(i.obj.HostedZoneId), cumulus.DefaultHidden)
	}

	if i.obj.HttpEndpointEnabled != nil {
		builder.What("http_endpoint_enabled", fmt.Sprint(boolToString(i.obj.HttpEndpointEnabled)), cumulus.DefaultHidden)
	}

	if i.obj.IAMDatabaseAuthenticationEnabled != nil {
		builder.What("iam_database_authentication_enabled", fmt.Sprint(boolToString(i.obj.IAMDatabaseAuthenticationEnabled)), cumulus.DefaultHidden)
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

	if i.obj.MasterUsername != nil {
		builder.What("master_username", aws.StringValue(i.obj.MasterUsername), cumulus.DefaultHidden)
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

	if i.obj.NetworkType != nil {
		builder.What("network_type", aws.StringValue(i.obj.NetworkType), cumulus.DefaultHidden)
	}

	if i.obj.PercentProgress != nil {
		builder.What("percent_progress", aws.StringValue(i.obj.PercentProgress), cumulus.DefaultHidden)
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

	if i.obj.Port != nil {
		builder.What("port", fmt.Sprint(aws.Int64Value(i.obj.Port)), cumulus.DefaultHidden)
	}

	if i.obj.PreferredBackupWindow != nil {
		builder.What("preferred_backup_window", aws.StringValue(i.obj.PreferredBackupWindow), cumulus.DefaultHidden)
	}

	if i.obj.PreferredMaintenanceWindow != nil {
		builder.What("preferred_maintenance_window", aws.StringValue(i.obj.PreferredMaintenanceWindow), cumulus.DefaultHidden)
	}

	if i.obj.PubliclyAccessible != nil {
		builder.What("publicly_accessible", fmt.Sprint(boolToString(i.obj.PubliclyAccessible)), cumulus.DefaultHidden)
	}

	if i.obj.ReadReplicaIdentifiers != nil {
		builder.What("read_replica_identifiers", fmt.Sprint((i.obj.ReadReplicaIdentifiers)), cumulus.DefaultHidden)
	}

	if i.obj.ReaderEndpoint != nil {
		builder.What("reader_endpoint", aws.StringValue(i.obj.ReaderEndpoint), cumulus.DefaultHidden)
	}

	if i.obj.ReplicationSourceIdentifier != nil {
		builder.What("replication_source_identifier", aws.StringValue(i.obj.ReplicationSourceIdentifier), cumulus.DefaultHidden)
	}

	if i.obj.Status != nil {
		builder.What("status", aws.StringValue(i.obj.Status), cumulus.DefaultHidden)
	}

	if i.obj.StorageEncrypted != nil {
		builder.What("storage_encrypted", fmt.Sprint(boolToString(i.obj.StorageEncrypted)), cumulus.DefaultHidden)
	}

	if i.obj.StorageType != nil {
		builder.What("storage_type", aws.StringValue(i.obj.StorageType), cumulus.DefaultHidden)
	}

}
