Manages an Amazon MSK cluster.

> **Note:** This resource manages _provisioned_ clusters. To manage a _serverless_ Amazon MSK cluster, use the `aws.msk.ServerlessCluster` resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const vpc = new aws.ec2.Vpc("vpc", {cidrBlock: "192.168.0.0/22"});
const azs = aws.getAvailabilityZones({
    state: "available",
});
const subnetAz1 = new aws.ec2.Subnet("subnetAz1", {
    availabilityZone: azs.then(azs => azs.names?.[0]),
    cidrBlock: "192.168.0.0/24",
    vpcId: vpc.id,
});
const subnetAz2 = new aws.ec2.Subnet("subnetAz2", {
    availabilityZone: azs.then(azs => azs.names?.[1]),
    cidrBlock: "192.168.1.0/24",
    vpcId: vpc.id,
});
const subnetAz3 = new aws.ec2.Subnet("subnetAz3", {
    availabilityZone: azs.then(azs => azs.names?.[2]),
    cidrBlock: "192.168.2.0/24",
    vpcId: vpc.id,
});
const sg = new aws.ec2.SecurityGroup("sg", {vpcId: vpc.id});
const kms = new aws.kms.Key("kms", {description: "example"});
const test = new aws.cloudwatch.LogGroup("test", {});
const bucket = new aws.s3.BucketV2("bucket", {});
const bucketAcl = new aws.s3.BucketAclV2("bucketAcl", {
    bucket: bucket.id,
    acl: "private",
});
const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["firehose.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const firehoseRole = new aws.iam.Role("firehoseRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const testStream = new aws.kinesis.FirehoseDeliveryStream("testStream", {
    destination: "extended_s3",
    extendedS3Configuration: {
        roleArn: firehoseRole.arn,
        bucketArn: bucket.arn,
    },
    tags: {
        LogDeliveryEnabled: "placeholder",
    },
});
const example = new aws.msk.Cluster("example", {
    kafkaVersion: "3.2.0",
    numberOfBrokerNodes: 3,
    brokerNodeGroupInfo: {
        instanceType: "kafka.m5.large",
        clientSubnets: [
            subnetAz1.id,
            subnetAz2.id,
            subnetAz3.id,
        ],
        storageInfo: {
            ebsStorageInfo: {
                volumeSize: 1000,
            },
        },
        securityGroups: [sg.id],
    },
    encryptionInfo: {
        encryptionAtRestKmsKeyArn: kms.arn,
    },
    openMonitoring: {
        prometheus: {
            jmxExporter: {
                enabledInBroker: true,
            },
            nodeExporter: {
                enabledInBroker: true,
            },
        },
    },
    loggingInfo: {
        brokerLogs: {
            cloudwatchLogs: {
                enabled: true,
                logGroup: test.name,
            },
            firehose: {
                enabled: true,
                deliveryStream: testStream.name,
            },
            s3: {
                enabled: true,
                bucket: bucket.id,
                prefix: "logs/msk-",
            },
        },
    },
    tags: {
        foo: "bar",
    },
});
export const zookeeperConnectString = example.zookeeperConnectString;
export const bootstrapBrokersTls = example.bootstrapBrokersTls;
```
```python
import pulumi
import pulumi_aws as aws

vpc = aws.ec2.Vpc("vpc", cidr_block="192.168.0.0/22")
azs = aws.get_availability_zones(state="available")
subnet_az1 = aws.ec2.Subnet("subnetAz1",
    availability_zone=azs.names[0],
    cidr_block="192.168.0.0/24",
    vpc_id=vpc.id)
subnet_az2 = aws.ec2.Subnet("subnetAz2",
    availability_zone=azs.names[1],
    cidr_block="192.168.1.0/24",
    vpc_id=vpc.id)
subnet_az3 = aws.ec2.Subnet("subnetAz3",
    availability_zone=azs.names[2],
    cidr_block="192.168.2.0/24",
    vpc_id=vpc.id)
sg = aws.ec2.SecurityGroup("sg", vpc_id=vpc.id)
kms = aws.kms.Key("kms", description="example")
test = aws.cloudwatch.LogGroup("test")
bucket = aws.s3.BucketV2("bucket")
bucket_acl = aws.s3.BucketAclV2("bucketAcl",
    bucket=bucket.id,
    acl="private")
assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["firehose.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
firehose_role = aws.iam.Role("firehoseRole", assume_role_policy=assume_role.json)
test_stream = aws.kinesis.FirehoseDeliveryStream("testStream",
    destination="extended_s3",
    extended_s3_configuration=aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationArgs(
        role_arn=firehose_role.arn,
        bucket_arn=bucket.arn,
    ),
    tags={
        "LogDeliveryEnabled": "placeholder",
    })
example = aws.msk.Cluster("example",
    kafka_version="3.2.0",
    number_of_broker_nodes=3,
    broker_node_group_info=aws.msk.ClusterBrokerNodeGroupInfoArgs(
        instance_type="kafka.m5.large",
        client_subnets=[
            subnet_az1.id,
            subnet_az2.id,
            subnet_az3.id,
        ],
        storage_info=aws.msk.ClusterBrokerNodeGroupInfoStorageInfoArgs(
            ebs_storage_info=aws.msk.ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs(
                volume_size=1000,
            ),
        ),
        security_groups=[sg.id],
    ),
    encryption_info=aws.msk.ClusterEncryptionInfoArgs(
        encryption_at_rest_kms_key_arn=kms.arn,
    ),
    open_monitoring=aws.msk.ClusterOpenMonitoringArgs(
        prometheus=aws.msk.ClusterOpenMonitoringPrometheusArgs(
            jmx_exporter=aws.msk.ClusterOpenMonitoringPrometheusJmxExporterArgs(
                enabled_in_broker=True,
            ),
            node_exporter=aws.msk.ClusterOpenMonitoringPrometheusNodeExporterArgs(
                enabled_in_broker=True,
            ),
        ),
    ),
    logging_info=aws.msk.ClusterLoggingInfoArgs(
        broker_logs=aws.msk.ClusterLoggingInfoBrokerLogsArgs(
            cloudwatch_logs=aws.msk.ClusterLoggingInfoBrokerLogsCloudwatchLogsArgs(
                enabled=True,
                log_group=test.name,
            ),
            firehose=aws.msk.ClusterLoggingInfoBrokerLogsFirehoseArgs(
                enabled=True,
                delivery_stream=test_stream.name,
            ),
            s3=aws.msk.ClusterLoggingInfoBrokerLogsS3Args(
                enabled=True,
                bucket=bucket.id,
                prefix="logs/msk-",
            ),
        ),
    ),
    tags={
        "foo": "bar",
    })
pulumi.export("zookeeperConnectString", example.zookeeper_connect_string)
pulumi.export("bootstrapBrokersTls", example.bootstrap_brokers_tls)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var vpc = new Aws.Ec2.Vpc("vpc", new()
    {
        CidrBlock = "192.168.0.0/22",
    });

    var azs = Aws.GetAvailabilityZones.Invoke(new()
    {
        State = "available",
    });

    var subnetAz1 = new Aws.Ec2.Subnet("subnetAz1", new()
    {
        AvailabilityZone = azs.Apply(getAvailabilityZonesResult => getAvailabilityZonesResult.Names[0]),
        CidrBlock = "192.168.0.0/24",
        VpcId = vpc.Id,
    });

    var subnetAz2 = new Aws.Ec2.Subnet("subnetAz2", new()
    {
        AvailabilityZone = azs.Apply(getAvailabilityZonesResult => getAvailabilityZonesResult.Names[1]),
        CidrBlock = "192.168.1.0/24",
        VpcId = vpc.Id,
    });

    var subnetAz3 = new Aws.Ec2.Subnet("subnetAz3", new()
    {
        AvailabilityZone = azs.Apply(getAvailabilityZonesResult => getAvailabilityZonesResult.Names[2]),
        CidrBlock = "192.168.2.0/24",
        VpcId = vpc.Id,
    });

    var sg = new Aws.Ec2.SecurityGroup("sg", new()
    {
        VpcId = vpc.Id,
    });

    var kms = new Aws.Kms.Key("kms", new()
    {
        Description = "example",
    });

    var test = new Aws.CloudWatch.LogGroup("test");

    var bucket = new Aws.S3.BucketV2("bucket");

    var bucketAcl = new Aws.S3.BucketAclV2("bucketAcl", new()
    {
        Bucket = bucket.Id,
        Acl = "private",
    });

    var assumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "firehose.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
            },
        },
    });

    var firehoseRole = new Aws.Iam.Role("firehoseRole", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var testStream = new Aws.Kinesis.FirehoseDeliveryStream("testStream", new()
    {
        Destination = "extended_s3",
        ExtendedS3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationArgs
        {
            RoleArn = firehoseRole.Arn,
            BucketArn = bucket.Arn,
        },
        Tags = 
        {
            { "LogDeliveryEnabled", "placeholder" },
        },
    });

    var example = new Aws.Msk.Cluster("example", new()
    {
        KafkaVersion = "3.2.0",
        NumberOfBrokerNodes = 3,
        BrokerNodeGroupInfo = new Aws.Msk.Inputs.ClusterBrokerNodeGroupInfoArgs
        {
            InstanceType = "kafka.m5.large",
            ClientSubnets = new[]
            {
                subnetAz1.Id,
                subnetAz2.Id,
                subnetAz3.Id,
            },
            StorageInfo = new Aws.Msk.Inputs.ClusterBrokerNodeGroupInfoStorageInfoArgs
            {
                EbsStorageInfo = new Aws.Msk.Inputs.ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs
                {
                    VolumeSize = 1000,
                },
            },
            SecurityGroups = new[]
            {
                sg.Id,
            },
        },
        EncryptionInfo = new Aws.Msk.Inputs.ClusterEncryptionInfoArgs
        {
            EncryptionAtRestKmsKeyArn = kms.Arn,
        },
        OpenMonitoring = new Aws.Msk.Inputs.ClusterOpenMonitoringArgs
        {
            Prometheus = new Aws.Msk.Inputs.ClusterOpenMonitoringPrometheusArgs
            {
                JmxExporter = new Aws.Msk.Inputs.ClusterOpenMonitoringPrometheusJmxExporterArgs
                {
                    EnabledInBroker = true,
                },
                NodeExporter = new Aws.Msk.Inputs.ClusterOpenMonitoringPrometheusNodeExporterArgs
                {
                    EnabledInBroker = true,
                },
            },
        },
        LoggingInfo = new Aws.Msk.Inputs.ClusterLoggingInfoArgs
        {
            BrokerLogs = new Aws.Msk.Inputs.ClusterLoggingInfoBrokerLogsArgs
            {
                CloudwatchLogs = new Aws.Msk.Inputs.ClusterLoggingInfoBrokerLogsCloudwatchLogsArgs
                {
                    Enabled = true,
                    LogGroup = test.Name,
                },
                Firehose = new Aws.Msk.Inputs.ClusterLoggingInfoBrokerLogsFirehoseArgs
                {
                    Enabled = true,
                    DeliveryStream = testStream.Name,
                },
                S3 = new Aws.Msk.Inputs.ClusterLoggingInfoBrokerLogsS3Args
                {
                    Enabled = true,
                    Bucket = bucket.Id,
                    Prefix = "logs/msk-",
                },
            },
        },
        Tags = 
        {
            { "foo", "bar" },
        },
    });

    return new Dictionary<string, object?>
    {
        ["zookeeperConnectString"] = example.ZookeeperConnectString,
        ["bootstrapBrokersTls"] = example.BootstrapBrokersTls,
    };
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/msk"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		vpc, err := ec2.NewVpc(ctx, "vpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("192.168.0.0/22"),
		})
		if err != nil {
			return err
		}
		azs, err := aws.GetAvailabilityZones(ctx, &aws.GetAvailabilityZonesArgs{
			State: pulumi.StringRef("available"),
		}, nil)
		if err != nil {
			return err
		}
		subnetAz1, err := ec2.NewSubnet(ctx, "subnetAz1", &ec2.SubnetArgs{
			AvailabilityZone: *pulumi.String(azs.Names[0]),
			CidrBlock:        pulumi.String("192.168.0.0/24"),
			VpcId:            vpc.ID(),
		})
		if err != nil {
			return err
		}
		subnetAz2, err := ec2.NewSubnet(ctx, "subnetAz2", &ec2.SubnetArgs{
			AvailabilityZone: *pulumi.String(azs.Names[1]),
			CidrBlock:        pulumi.String("192.168.1.0/24"),
			VpcId:            vpc.ID(),
		})
		if err != nil {
			return err
		}
		subnetAz3, err := ec2.NewSubnet(ctx, "subnetAz3", &ec2.SubnetArgs{
			AvailabilityZone: *pulumi.String(azs.Names[2]),
			CidrBlock:        pulumi.String("192.168.2.0/24"),
			VpcId:            vpc.ID(),
		})
		if err != nil {
			return err
		}
		sg, err := ec2.NewSecurityGroup(ctx, "sg", &ec2.SecurityGroupArgs{
			VpcId: vpc.ID(),
		})
		if err != nil {
			return err
		}
		kms, err := kms.NewKey(ctx, "kms", &kms.KeyArgs{
			Description: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		test, err := cloudwatch.NewLogGroup(ctx, "test", nil)
		if err != nil {
			return err
		}
		bucket, err := s3.NewBucketV2(ctx, "bucket", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAclV2(ctx, "bucketAcl", &s3.BucketAclV2Args{
			Bucket: bucket.ID(),
			Acl:    pulumi.String("private"),
		})
		if err != nil {
			return err
		}
		assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"firehose.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"sts:AssumeRole",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		firehoseRole, err := iam.NewRole(ctx, "firehoseRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		testStream, err := kinesis.NewFirehoseDeliveryStream(ctx, "testStream", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("extended_s3"),
			ExtendedS3Configuration: &kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationArgs{
				RoleArn:   firehoseRole.Arn,
				BucketArn: bucket.Arn,
			},
			Tags: pulumi.StringMap{
				"LogDeliveryEnabled": pulumi.String("placeholder"),
			},
		})
		if err != nil {
			return err
		}
		example, err := msk.NewCluster(ctx, "example", &msk.ClusterArgs{
			KafkaVersion:        pulumi.String("3.2.0"),
			NumberOfBrokerNodes: pulumi.Int(3),
			BrokerNodeGroupInfo: &msk.ClusterBrokerNodeGroupInfoArgs{
				InstanceType: pulumi.String("kafka.m5.large"),
				ClientSubnets: pulumi.StringArray{
					subnetAz1.ID(),
					subnetAz2.ID(),
					subnetAz3.ID(),
				},
				StorageInfo: &msk.ClusterBrokerNodeGroupInfoStorageInfoArgs{
					EbsStorageInfo: &msk.ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs{
						VolumeSize: pulumi.Int(1000),
					},
				},
				SecurityGroups: pulumi.StringArray{
					sg.ID(),
				},
			},
			EncryptionInfo: &msk.ClusterEncryptionInfoArgs{
				EncryptionAtRestKmsKeyArn: kms.Arn,
			},
			OpenMonitoring: &msk.ClusterOpenMonitoringArgs{
				Prometheus: &msk.ClusterOpenMonitoringPrometheusArgs{
					JmxExporter: &msk.ClusterOpenMonitoringPrometheusJmxExporterArgs{
						EnabledInBroker: pulumi.Bool(true),
					},
					NodeExporter: &msk.ClusterOpenMonitoringPrometheusNodeExporterArgs{
						EnabledInBroker: pulumi.Bool(true),
					},
				},
			},
			LoggingInfo: &msk.ClusterLoggingInfoArgs{
				BrokerLogs: &msk.ClusterLoggingInfoBrokerLogsArgs{
					CloudwatchLogs: &msk.ClusterLoggingInfoBrokerLogsCloudwatchLogsArgs{
						Enabled:  pulumi.Bool(true),
						LogGroup: test.Name,
					},
					Firehose: &msk.ClusterLoggingInfoBrokerLogsFirehoseArgs{
						Enabled:        pulumi.Bool(true),
						DeliveryStream: testStream.Name,
					},
					S3: &msk.ClusterLoggingInfoBrokerLogsS3Args{
						Enabled: pulumi.Bool(true),
						Bucket:  bucket.ID(),
						Prefix:  pulumi.String("logs/msk-"),
					},
				},
			},
			Tags: pulumi.StringMap{
				"foo": pulumi.String("bar"),
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("zookeeperConnectString", example.ZookeeperConnectString)
		ctx.Export("bootstrapBrokersTls", example.BootstrapBrokersTls)
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetAvailabilityZonesArgs;
import com.pulumi.aws.ec2.Subnet;
import com.pulumi.aws.ec2.SubnetArgs;
import com.pulumi.aws.ec2.SecurityGroup;
import com.pulumi.aws.ec2.SecurityGroupArgs;
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.cloudwatch.LogGroup;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.kinesis.FirehoseDeliveryStream;
import com.pulumi.aws.kinesis.FirehoseDeliveryStreamArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamExtendedS3ConfigurationArgs;
import com.pulumi.aws.msk.Cluster;
import com.pulumi.aws.msk.ClusterArgs;
import com.pulumi.aws.msk.inputs.ClusterBrokerNodeGroupInfoArgs;
import com.pulumi.aws.msk.inputs.ClusterBrokerNodeGroupInfoStorageInfoArgs;
import com.pulumi.aws.msk.inputs.ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs;
import com.pulumi.aws.msk.inputs.ClusterEncryptionInfoArgs;
import com.pulumi.aws.msk.inputs.ClusterOpenMonitoringArgs;
import com.pulumi.aws.msk.inputs.ClusterOpenMonitoringPrometheusArgs;
import com.pulumi.aws.msk.inputs.ClusterOpenMonitoringPrometheusJmxExporterArgs;
import com.pulumi.aws.msk.inputs.ClusterOpenMonitoringPrometheusNodeExporterArgs;
import com.pulumi.aws.msk.inputs.ClusterLoggingInfoArgs;
import com.pulumi.aws.msk.inputs.ClusterLoggingInfoBrokerLogsArgs;
import com.pulumi.aws.msk.inputs.ClusterLoggingInfoBrokerLogsCloudwatchLogsArgs;
import com.pulumi.aws.msk.inputs.ClusterLoggingInfoBrokerLogsFirehoseArgs;
import com.pulumi.aws.msk.inputs.ClusterLoggingInfoBrokerLogsS3Args;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var vpc = new Vpc("vpc", VpcArgs.builder()        
            .cidrBlock("192.168.0.0/22")
            .build());

        final var azs = AwsFunctions.getAvailabilityZones(GetAvailabilityZonesArgs.builder()
            .state("available")
            .build());

        var subnetAz1 = new Subnet("subnetAz1", SubnetArgs.builder()        
            .availabilityZone(azs.applyValue(getAvailabilityZonesResult -> getAvailabilityZonesResult.names()[0]))
            .cidrBlock("192.168.0.0/24")
            .vpcId(vpc.id())
            .build());

        var subnetAz2 = new Subnet("subnetAz2", SubnetArgs.builder()        
            .availabilityZone(azs.applyValue(getAvailabilityZonesResult -> getAvailabilityZonesResult.names()[1]))
            .cidrBlock("192.168.1.0/24")
            .vpcId(vpc.id())
            .build());

        var subnetAz3 = new Subnet("subnetAz3", SubnetArgs.builder()        
            .availabilityZone(azs.applyValue(getAvailabilityZonesResult -> getAvailabilityZonesResult.names()[2]))
            .cidrBlock("192.168.2.0/24")
            .vpcId(vpc.id())
            .build());

        var sg = new SecurityGroup("sg", SecurityGroupArgs.builder()        
            .vpcId(vpc.id())
            .build());

        var kms = new Key("kms", KeyArgs.builder()        
            .description("example")
            .build());

        var test = new LogGroup("test");

        var bucket = new BucketV2("bucket");

        var bucketAcl = new BucketAclV2("bucketAcl", BucketAclV2Args.builder()        
            .bucket(bucket.id())
            .acl("private")
            .build());

        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("firehose.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var firehoseRole = new Role("firehoseRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var testStream = new FirehoseDeliveryStream("testStream", FirehoseDeliveryStreamArgs.builder()        
            .destination("extended_s3")
            .extendedS3Configuration(FirehoseDeliveryStreamExtendedS3ConfigurationArgs.builder()
                .roleArn(firehoseRole.arn())
                .bucketArn(bucket.arn())
                .build())
            .tags(Map.of("LogDeliveryEnabled", "placeholder"))
            .build());

        var example = new Cluster("example", ClusterArgs.builder()        
            .kafkaVersion("3.2.0")
            .numberOfBrokerNodes(3)
            .brokerNodeGroupInfo(ClusterBrokerNodeGroupInfoArgs.builder()
                .instanceType("kafka.m5.large")
                .clientSubnets(                
                    subnetAz1.id(),
                    subnetAz2.id(),
                    subnetAz3.id())
                .storageInfo(ClusterBrokerNodeGroupInfoStorageInfoArgs.builder()
                    .ebsStorageInfo(ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs.builder()
                        .volumeSize(1000)
                        .build())
                    .build())
                .securityGroups(sg.id())
                .build())
            .encryptionInfo(ClusterEncryptionInfoArgs.builder()
                .encryptionAtRestKmsKeyArn(kms.arn())
                .build())
            .openMonitoring(ClusterOpenMonitoringArgs.builder()
                .prometheus(ClusterOpenMonitoringPrometheusArgs.builder()
                    .jmxExporter(ClusterOpenMonitoringPrometheusJmxExporterArgs.builder()
                        .enabledInBroker(true)
                        .build())
                    .nodeExporter(ClusterOpenMonitoringPrometheusNodeExporterArgs.builder()
                        .enabledInBroker(true)
                        .build())
                    .build())
                .build())
            .loggingInfo(ClusterLoggingInfoArgs.builder()
                .brokerLogs(ClusterLoggingInfoBrokerLogsArgs.builder()
                    .cloudwatchLogs(ClusterLoggingInfoBrokerLogsCloudwatchLogsArgs.builder()
                        .enabled(true)
                        .logGroup(test.name())
                        .build())
                    .firehose(ClusterLoggingInfoBrokerLogsFirehoseArgs.builder()
                        .enabled(true)
                        .deliveryStream(testStream.name())
                        .build())
                    .s3(ClusterLoggingInfoBrokerLogsS3Args.builder()
                        .enabled(true)
                        .bucket(bucket.id())
                        .prefix("logs/msk-")
                        .build())
                    .build())
                .build())
            .tags(Map.of("foo", "bar"))
            .build());

        ctx.export("zookeeperConnectString", example.zookeeperConnectString());
        ctx.export("bootstrapBrokersTls", example.bootstrapBrokersTls());
    }
}
```
```yaml
resources:
  vpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 192.168.0.0/22
  subnetAz1:
    type: aws:ec2:Subnet
    properties:
      availabilityZone: ${azs.names[0]}
      cidrBlock: 192.168.0.0/24
      vpcId: ${vpc.id}
  subnetAz2:
    type: aws:ec2:Subnet
    properties:
      availabilityZone: ${azs.names[1]}
      cidrBlock: 192.168.1.0/24
      vpcId: ${vpc.id}
  subnetAz3:
    type: aws:ec2:Subnet
    properties:
      availabilityZone: ${azs.names[2]}
      cidrBlock: 192.168.2.0/24
      vpcId: ${vpc.id}
  sg:
    type: aws:ec2:SecurityGroup
    properties:
      vpcId: ${vpc.id}
  kms:
    type: aws:kms:Key
    properties:
      description: example
  test:
    type: aws:cloudwatch:LogGroup
  bucket:
    type: aws:s3:BucketV2
  bucketAcl:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${bucket.id}
      acl: private
  firehoseRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  testStream:
    type: aws:kinesis:FirehoseDeliveryStream
    properties:
      destination: extended_s3
      extendedS3Configuration:
        roleArn: ${firehoseRole.arn}
        bucketArn: ${bucket.arn}
      tags:
        LogDeliveryEnabled: placeholder
  example:
    type: aws:msk:Cluster
    properties:
      kafkaVersion: 3.2.0
      numberOfBrokerNodes: 3
      brokerNodeGroupInfo:
        instanceType: kafka.m5.large
        clientSubnets:
          - ${subnetAz1.id}
          - ${subnetAz2.id}
          - ${subnetAz3.id}
        storageInfo:
          ebsStorageInfo:
            volumeSize: 1000
        securityGroups:
          - ${sg.id}
      encryptionInfo:
        encryptionAtRestKmsKeyArn: ${kms.arn}
      openMonitoring:
        prometheus:
          jmxExporter:
            enabledInBroker: true
          nodeExporter:
            enabledInBroker: true
      loggingInfo:
        brokerLogs:
          cloudwatchLogs:
            enabled: true
            logGroup: ${test.name}
          firehose:
            enabled: true
            deliveryStream: ${testStream.name}
          s3:
            enabled: true
            bucket: ${bucket.id}
            prefix: logs/msk-
      tags:
        foo: bar
variables:
  azs:
    fn::invoke:
      Function: aws:getAvailabilityZones
      Arguments:
        state: available
  assumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - firehose.amazonaws.com
            actions:
              - sts:AssumeRole
outputs:
  zookeeperConnectString: ${example.zookeeperConnectString}
  bootstrapBrokersTls: ${example.bootstrapBrokersTls}
```
{{% /example %}}
{{% example %}}
### With volume_throughput argument

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.msk.Cluster("example", {
    kafkaVersion: "2.7.1",
    numberOfBrokerNodes: 3,
    brokerNodeGroupInfo: {
        instanceType: "kafka.m5.4xlarge",
        clientSubnets: [
            aws_subnet.subnet_az1.id,
            aws_subnet.subnet_az2.id,
            aws_subnet.subnet_az3.id,
        ],
        storageInfo: {
            ebsStorageInfo: {
                provisionedThroughput: {
                    enabled: true,
                    volumeThroughput: 250,
                },
                volumeSize: 1000,
            },
        },
        securityGroups: [aws_security_group.sg.id],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.msk.Cluster("example",
    kafka_version="2.7.1",
    number_of_broker_nodes=3,
    broker_node_group_info=aws.msk.ClusterBrokerNodeGroupInfoArgs(
        instance_type="kafka.m5.4xlarge",
        client_subnets=[
            aws_subnet["subnet_az1"]["id"],
            aws_subnet["subnet_az2"]["id"],
            aws_subnet["subnet_az3"]["id"],
        ],
        storage_info=aws.msk.ClusterBrokerNodeGroupInfoStorageInfoArgs(
            ebs_storage_info=aws.msk.ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs(
                provisioned_throughput=aws.msk.ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughputArgs(
                    enabled=True,
                    volume_throughput=250,
                ),
                volume_size=1000,
            ),
        ),
        security_groups=[aws_security_group["sg"]["id"]],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Msk.Cluster("example", new()
    {
        KafkaVersion = "2.7.1",
        NumberOfBrokerNodes = 3,
        BrokerNodeGroupInfo = new Aws.Msk.Inputs.ClusterBrokerNodeGroupInfoArgs
        {
            InstanceType = "kafka.m5.4xlarge",
            ClientSubnets = new[]
            {
                aws_subnet.Subnet_az1.Id,
                aws_subnet.Subnet_az2.Id,
                aws_subnet.Subnet_az3.Id,
            },
            StorageInfo = new Aws.Msk.Inputs.ClusterBrokerNodeGroupInfoStorageInfoArgs
            {
                EbsStorageInfo = new Aws.Msk.Inputs.ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs
                {
                    ProvisionedThroughput = new Aws.Msk.Inputs.ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughputArgs
                    {
                        Enabled = true,
                        VolumeThroughput = 250,
                    },
                    VolumeSize = 1000,
                },
            },
            SecurityGroups = new[]
            {
                aws_security_group.Sg.Id,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/msk"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := msk.NewCluster(ctx, "example", &msk.ClusterArgs{
			KafkaVersion:        pulumi.String("2.7.1"),
			NumberOfBrokerNodes: pulumi.Int(3),
			BrokerNodeGroupInfo: &msk.ClusterBrokerNodeGroupInfoArgs{
				InstanceType: pulumi.String("kafka.m5.4xlarge"),
				ClientSubnets: pulumi.StringArray{
					aws_subnet.Subnet_az1.Id,
					aws_subnet.Subnet_az2.Id,
					aws_subnet.Subnet_az3.Id,
				},
				StorageInfo: &msk.ClusterBrokerNodeGroupInfoStorageInfoArgs{
					EbsStorageInfo: &msk.ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs{
						ProvisionedThroughput: &msk.ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughputArgs{
							Enabled:          pulumi.Bool(true),
							VolumeThroughput: pulumi.Int(250),
						},
						VolumeSize: pulumi.Int(1000),
					},
				},
				SecurityGroups: pulumi.StringArray{
					aws_security_group.Sg.Id,
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.msk.Cluster;
import com.pulumi.aws.msk.ClusterArgs;
import com.pulumi.aws.msk.inputs.ClusterBrokerNodeGroupInfoArgs;
import com.pulumi.aws.msk.inputs.ClusterBrokerNodeGroupInfoStorageInfoArgs;
import com.pulumi.aws.msk.inputs.ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs;
import com.pulumi.aws.msk.inputs.ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughputArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var example = new Cluster("example", ClusterArgs.builder()        
            .kafkaVersion("2.7.1")
            .numberOfBrokerNodes(3)
            .brokerNodeGroupInfo(ClusterBrokerNodeGroupInfoArgs.builder()
                .instanceType("kafka.m5.4xlarge")
                .clientSubnets(                
                    aws_subnet.subnet_az1().id(),
                    aws_subnet.subnet_az2().id(),
                    aws_subnet.subnet_az3().id())
                .storageInfo(ClusterBrokerNodeGroupInfoStorageInfoArgs.builder()
                    .ebsStorageInfo(ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs.builder()
                        .provisionedThroughput(ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughputArgs.builder()
                            .enabled(true)
                            .volumeThroughput(250)
                            .build())
                        .volumeSize(1000)
                        .build())
                    .build())
                .securityGroups(aws_security_group.sg().id())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:msk:Cluster
    properties:
      kafkaVersion: 2.7.1
      numberOfBrokerNodes: 3
      brokerNodeGroupInfo:
        instanceType: kafka.m5.4xlarge
        clientSubnets:
          - ${aws_subnet.subnet_az1.id}
          - ${aws_subnet.subnet_az2.id}
          - ${aws_subnet.subnet_az3.id}
        storageInfo:
          ebsStorageInfo:
            provisionedThroughput:
              enabled: true
              volumeThroughput: 250
            volumeSize: 1000
        securityGroups:
          - ${aws_security_group.sg.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import MSK clusters using the cluster `arn`. For example:

```sh
 $ pulumi import aws:msk/cluster:Cluster example arn:aws:kafka:us-west-2:123456789012:cluster/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3
```
 