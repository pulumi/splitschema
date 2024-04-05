Provides an Elastic MapReduce Cluster, a web service that makes it easy to process large amounts of data efficiently. See [Amazon Elastic MapReduce Documentation](https://aws.amazon.com/documentation/elastic-mapreduce/) for more information.

To configure [Instance Groups](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-group-configuration.html#emr-plan-instance-groups) for [task nodes](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-master-core-task-nodes.html#emr-plan-task), see the `aws.emr.InstanceGroup` resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const cluster = new aws.emr.Cluster("cluster", {
    releaseLabel: "emr-4.6.0",
    applications: ["Spark"],
    additionalInfo: `{
  "instanceAwsClientConfiguration": {
    "proxyPort": 8099,
    "proxyHost": "myproxy.example.com"
  }
}
`,
    terminationProtection: false,
    keepJobFlowAliveWhenNoSteps: true,
    ec2Attributes: {
        subnetId: aws_subnet.main.id,
        emrManagedMasterSecurityGroup: aws_security_group.sg.id,
        emrManagedSlaveSecurityGroup: aws_security_group.sg.id,
        instanceProfile: aws_iam_instance_profile.emr_profile.arn,
    },
    masterInstanceGroup: {
        instanceType: "m4.large",
    },
    coreInstanceGroup: {
        instanceType: "c4.large",
        instanceCount: 1,
        ebsConfigs: [{
            size: 40,
            type: "gp2",
            volumesPerInstance: 1,
        }],
        bidPrice: "0.30",
        autoscalingPolicy: `{
"Constraints": {
  "MinCapacity": 1,
  "MaxCapacity": 2
},
"Rules": [
  {
    "Name": "ScaleOutMemoryPercentage",
    "Description": "Scale out if YARNMemoryAvailablePercentage is less than 15",
    "Action": {
      "SimpleScalingPolicyConfiguration": {
        "AdjustmentType": "CHANGE_IN_CAPACITY",
        "ScalingAdjustment": 1,
        "CoolDown": 300
      }
    },
    "Trigger": {
      "CloudWatchAlarmDefinition": {
        "ComparisonOperator": "LESS_THAN",
        "EvaluationPeriods": 1,
        "MetricName": "YARNMemoryAvailablePercentage",
        "Namespace": "AWS/ElasticMapReduce",
        "Period": 300,
        "Statistic": "AVERAGE",
        "Threshold": 15.0,
        "Unit": "PERCENT"
      }
    }
  }
]
}
`,
    },
    ebsRootVolumeSize: 100,
    tags: {
        role: "rolename",
        env: "env",
    },
    bootstrapActions: [{
        path: "s3://elasticmapreduce/bootstrap-actions/run-if",
        name: "runif",
        args: [
            "instance.isMaster=true",
            "echo running on master node",
        ],
    }],
    configurationsJson: `  [
    {
      "Classification": "hadoop-env",
      "Configurations": [
        {
          "Classification": "export",
          "Properties": {
            "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
          }
        }
      ],
      "Properties": {}
    },
    {
      "Classification": "spark-env",
      "Configurations": [
        {
          "Classification": "export",
          "Properties": {
            "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
          }
        }
      ],
      "Properties": {}
    }
  ]
`,
    serviceRole: aws_iam_role.iam_emr_service_role.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

cluster = aws.emr.Cluster("cluster",
    release_label="emr-4.6.0",
    applications=["Spark"],
    additional_info="""{
  "instanceAwsClientConfiguration": {
    "proxyPort": 8099,
    "proxyHost": "myproxy.example.com"
  }
}
""",
    termination_protection=False,
    keep_job_flow_alive_when_no_steps=True,
    ec2_attributes=aws.emr.ClusterEc2AttributesArgs(
        subnet_id=aws_subnet["main"]["id"],
        emr_managed_master_security_group=aws_security_group["sg"]["id"],
        emr_managed_slave_security_group=aws_security_group["sg"]["id"],
        instance_profile=aws_iam_instance_profile["emr_profile"]["arn"],
    ),
    master_instance_group=aws.emr.ClusterMasterInstanceGroupArgs(
        instance_type="m4.large",
    ),
    core_instance_group=aws.emr.ClusterCoreInstanceGroupArgs(
        instance_type="c4.large",
        instance_count=1,
        ebs_configs=[aws.emr.ClusterCoreInstanceGroupEbsConfigArgs(
            size=40,
            type="gp2",
            volumes_per_instance=1,
        )],
        bid_price="0.30",
        autoscaling_policy="""{
"Constraints": {
  "MinCapacity": 1,
  "MaxCapacity": 2
},
"Rules": [
  {
    "Name": "ScaleOutMemoryPercentage",
    "Description": "Scale out if YARNMemoryAvailablePercentage is less than 15",
    "Action": {
      "SimpleScalingPolicyConfiguration": {
        "AdjustmentType": "CHANGE_IN_CAPACITY",
        "ScalingAdjustment": 1,
        "CoolDown": 300
      }
    },
    "Trigger": {
      "CloudWatchAlarmDefinition": {
        "ComparisonOperator": "LESS_THAN",
        "EvaluationPeriods": 1,
        "MetricName": "YARNMemoryAvailablePercentage",
        "Namespace": "AWS/ElasticMapReduce",
        "Period": 300,
        "Statistic": "AVERAGE",
        "Threshold": 15.0,
        "Unit": "PERCENT"
      }
    }
  }
]
}
""",
    ),
    ebs_root_volume_size=100,
    tags={
        "role": "rolename",
        "env": "env",
    },
    bootstrap_actions=[aws.emr.ClusterBootstrapActionArgs(
        path="s3://elasticmapreduce/bootstrap-actions/run-if",
        name="runif",
        args=[
            "instance.isMaster=true",
            "echo running on master node",
        ],
    )],
    configurations_json="""  [
    {
      "Classification": "hadoop-env",
      "Configurations": [
        {
          "Classification": "export",
          "Properties": {
            "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
          }
        }
      ],
      "Properties": {}
    },
    {
      "Classification": "spark-env",
      "Configurations": [
        {
          "Classification": "export",
          "Properties": {
            "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
          }
        }
      ],
      "Properties": {}
    }
  ]
""",
    service_role=aws_iam_role["iam_emr_service_role"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var cluster = new Aws.Emr.Cluster("cluster", new()
    {
        ReleaseLabel = "emr-4.6.0",
        Applications = new[]
        {
            "Spark",
        },
        AdditionalInfo = @"{
  ""instanceAwsClientConfiguration"": {
    ""proxyPort"": 8099,
    ""proxyHost"": ""myproxy.example.com""
  }
}
",
        TerminationProtection = false,
        KeepJobFlowAliveWhenNoSteps = true,
        Ec2Attributes = new Aws.Emr.Inputs.ClusterEc2AttributesArgs
        {
            SubnetId = aws_subnet.Main.Id,
            EmrManagedMasterSecurityGroup = aws_security_group.Sg.Id,
            EmrManagedSlaveSecurityGroup = aws_security_group.Sg.Id,
            InstanceProfile = aws_iam_instance_profile.Emr_profile.Arn,
        },
        MasterInstanceGroup = new Aws.Emr.Inputs.ClusterMasterInstanceGroupArgs
        {
            InstanceType = "m4.large",
        },
        CoreInstanceGroup = new Aws.Emr.Inputs.ClusterCoreInstanceGroupArgs
        {
            InstanceType = "c4.large",
            InstanceCount = 1,
            EbsConfigs = new[]
            {
                new Aws.Emr.Inputs.ClusterCoreInstanceGroupEbsConfigArgs
                {
                    Size = 40,
                    Type = "gp2",
                    VolumesPerInstance = 1,
                },
            },
            BidPrice = "0.30",
            AutoscalingPolicy = @"{
""Constraints"": {
  ""MinCapacity"": 1,
  ""MaxCapacity"": 2
},
""Rules"": [
  {
    ""Name"": ""ScaleOutMemoryPercentage"",
    ""Description"": ""Scale out if YARNMemoryAvailablePercentage is less than 15"",
    ""Action"": {
      ""SimpleScalingPolicyConfiguration"": {
        ""AdjustmentType"": ""CHANGE_IN_CAPACITY"",
        ""ScalingAdjustment"": 1,
        ""CoolDown"": 300
      }
    },
    ""Trigger"": {
      ""CloudWatchAlarmDefinition"": {
        ""ComparisonOperator"": ""LESS_THAN"",
        ""EvaluationPeriods"": 1,
        ""MetricName"": ""YARNMemoryAvailablePercentage"",
        ""Namespace"": ""AWS/ElasticMapReduce"",
        ""Period"": 300,
        ""Statistic"": ""AVERAGE"",
        ""Threshold"": 15.0,
        ""Unit"": ""PERCENT""
      }
    }
  }
]
}
",
        },
        EbsRootVolumeSize = 100,
        Tags = 
        {
            { "role", "rolename" },
            { "env", "env" },
        },
        BootstrapActions = new[]
        {
            new Aws.Emr.Inputs.ClusterBootstrapActionArgs
            {
                Path = "s3://elasticmapreduce/bootstrap-actions/run-if",
                Name = "runif",
                Args = new[]
                {
                    "instance.isMaster=true",
                    "echo running on master node",
                },
            },
        },
        ConfigurationsJson = @"  [
    {
      ""Classification"": ""hadoop-env"",
      ""Configurations"": [
        {
          ""Classification"": ""export"",
          ""Properties"": {
            ""JAVA_HOME"": ""/usr/lib/jvm/java-1.8.0""
          }
        }
      ],
      ""Properties"": {}
    },
    {
      ""Classification"": ""spark-env"",
      ""Configurations"": [
        {
          ""Classification"": ""export"",
          ""Properties"": {
            ""JAVA_HOME"": ""/usr/lib/jvm/java-1.8.0""
          }
        }
      ],
      ""Properties"": {}
    }
  ]
",
        ServiceRole = aws_iam_role.Iam_emr_service_role.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := emr.NewCluster(ctx, "cluster", &emr.ClusterArgs{
			ReleaseLabel: pulumi.String("emr-4.6.0"),
			Applications: pulumi.StringArray{
				pulumi.String("Spark"),
			},
			AdditionalInfo: pulumi.String(`{
  "instanceAwsClientConfiguration": {
    "proxyPort": 8099,
    "proxyHost": "myproxy.example.com"
  }
}
`),
			TerminationProtection:       pulumi.Bool(false),
			KeepJobFlowAliveWhenNoSteps: pulumi.Bool(true),
			Ec2Attributes: &emr.ClusterEc2AttributesArgs{
				SubnetId:                      pulumi.Any(aws_subnet.Main.Id),
				EmrManagedMasterSecurityGroup: pulumi.Any(aws_security_group.Sg.Id),
				EmrManagedSlaveSecurityGroup:  pulumi.Any(aws_security_group.Sg.Id),
				InstanceProfile:               pulumi.Any(aws_iam_instance_profile.Emr_profile.Arn),
			},
			MasterInstanceGroup: &emr.ClusterMasterInstanceGroupArgs{
				InstanceType: pulumi.String("m4.large"),
			},
			CoreInstanceGroup: &emr.ClusterCoreInstanceGroupArgs{
				InstanceType:  pulumi.String("c4.large"),
				InstanceCount: pulumi.Int(1),
				EbsConfigs: emr.ClusterCoreInstanceGroupEbsConfigArray{
					&emr.ClusterCoreInstanceGroupEbsConfigArgs{
						Size:               pulumi.Int(40),
						Type:               pulumi.String("gp2"),
						VolumesPerInstance: pulumi.Int(1),
					},
				},
				BidPrice: pulumi.String("0.30"),
				AutoscalingPolicy: pulumi.String(`{
"Constraints": {
  "MinCapacity": 1,
  "MaxCapacity": 2
},
"Rules": [
  {
    "Name": "ScaleOutMemoryPercentage",
    "Description": "Scale out if YARNMemoryAvailablePercentage is less than 15",
    "Action": {
      "SimpleScalingPolicyConfiguration": {
        "AdjustmentType": "CHANGE_IN_CAPACITY",
        "ScalingAdjustment": 1,
        "CoolDown": 300
      }
    },
    "Trigger": {
      "CloudWatchAlarmDefinition": {
        "ComparisonOperator": "LESS_THAN",
        "EvaluationPeriods": 1,
        "MetricName": "YARNMemoryAvailablePercentage",
        "Namespace": "AWS/ElasticMapReduce",
        "Period": 300,
        "Statistic": "AVERAGE",
        "Threshold": 15.0,
        "Unit": "PERCENT"
      }
    }
  }
]
}
`),
			},
			EbsRootVolumeSize: pulumi.Int(100),
			Tags: pulumi.StringMap{
				"role": pulumi.String("rolename"),
				"env":  pulumi.String("env"),
			},
			BootstrapActions: emr.ClusterBootstrapActionArray{
				&emr.ClusterBootstrapActionArgs{
					Path: pulumi.String("s3://elasticmapreduce/bootstrap-actions/run-if"),
					Name: pulumi.String("runif"),
					Args: pulumi.StringArray{
						pulumi.String("instance.isMaster=true"),
						pulumi.String("echo running on master node"),
					},
				},
			},
			ConfigurationsJson: pulumi.String(`  [
    {
      "Classification": "hadoop-env",
      "Configurations": [
        {
          "Classification": "export",
          "Properties": {
            "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
          }
        }
      ],
      "Properties": {}
    },
    {
      "Classification": "spark-env",
      "Configurations": [
        {
          "Classification": "export",
          "Properties": {
            "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
          }
        }
      ],
      "Properties": {}
    }
  ]
`),
			ServiceRole: pulumi.Any(aws_iam_role.Iam_emr_service_role.Arn),
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
import com.pulumi.aws.emr.Cluster;
import com.pulumi.aws.emr.ClusterArgs;
import com.pulumi.aws.emr.inputs.ClusterEc2AttributesArgs;
import com.pulumi.aws.emr.inputs.ClusterMasterInstanceGroupArgs;
import com.pulumi.aws.emr.inputs.ClusterCoreInstanceGroupArgs;
import com.pulumi.aws.emr.inputs.ClusterBootstrapActionArgs;
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
        var cluster = new Cluster("cluster", ClusterArgs.builder()        
            .releaseLabel("emr-4.6.0")
            .applications("Spark")
            .additionalInfo("""
{
  "instanceAwsClientConfiguration": {
    "proxyPort": 8099,
    "proxyHost": "myproxy.example.com"
  }
}
            """)
            .terminationProtection(false)
            .keepJobFlowAliveWhenNoSteps(true)
            .ec2Attributes(ClusterEc2AttributesArgs.builder()
                .subnetId(aws_subnet.main().id())
                .emrManagedMasterSecurityGroup(aws_security_group.sg().id())
                .emrManagedSlaveSecurityGroup(aws_security_group.sg().id())
                .instanceProfile(aws_iam_instance_profile.emr_profile().arn())
                .build())
            .masterInstanceGroup(ClusterMasterInstanceGroupArgs.builder()
                .instanceType("m4.large")
                .build())
            .coreInstanceGroup(ClusterCoreInstanceGroupArgs.builder()
                .instanceType("c4.large")
                .instanceCount(1)
                .ebsConfigs(ClusterCoreInstanceGroupEbsConfigArgs.builder()
                    .size("40")
                    .type("gp2")
                    .volumesPerInstance(1)
                    .build())
                .bidPrice("0.30")
                .autoscalingPolicy("""
{
"Constraints": {
  "MinCapacity": 1,
  "MaxCapacity": 2
},
"Rules": [
  {
    "Name": "ScaleOutMemoryPercentage",
    "Description": "Scale out if YARNMemoryAvailablePercentage is less than 15",
    "Action": {
      "SimpleScalingPolicyConfiguration": {
        "AdjustmentType": "CHANGE_IN_CAPACITY",
        "ScalingAdjustment": 1,
        "CoolDown": 300
      }
    },
    "Trigger": {
      "CloudWatchAlarmDefinition": {
        "ComparisonOperator": "LESS_THAN",
        "EvaluationPeriods": 1,
        "MetricName": "YARNMemoryAvailablePercentage",
        "Namespace": "AWS/ElasticMapReduce",
        "Period": 300,
        "Statistic": "AVERAGE",
        "Threshold": 15.0,
        "Unit": "PERCENT"
      }
    }
  }
]
}
                """)
                .build())
            .ebsRootVolumeSize(100)
            .tags(Map.ofEntries(
                Map.entry("role", "rolename"),
                Map.entry("env", "env")
            ))
            .bootstrapActions(ClusterBootstrapActionArgs.builder()
                .path("s3://elasticmapreduce/bootstrap-actions/run-if")
                .name("runif")
                .args(                
                    "instance.isMaster=true",
                    "echo running on master node")
                .build())
            .configurationsJson("""
  [
    {
      "Classification": "hadoop-env",
      "Configurations": [
        {
          "Classification": "export",
          "Properties": {
            "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
          }
        }
      ],
      "Properties": {}
    },
    {
      "Classification": "spark-env",
      "Configurations": [
        {
          "Classification": "export",
          "Properties": {
            "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
          }
        }
      ],
      "Properties": {}
    }
  ]
            """)
            .serviceRole(aws_iam_role.iam_emr_service_role().arn())
            .build());

    }
}
```
```yaml
resources:
  cluster:
    type: aws:emr:Cluster
    properties:
      releaseLabel: emr-4.6.0
      applications:
        - Spark
      additionalInfo: |
        {
          "instanceAwsClientConfiguration": {
            "proxyPort": 8099,
            "proxyHost": "myproxy.example.com"
          }
        }
      terminationProtection: false
      keepJobFlowAliveWhenNoSteps: true
      ec2Attributes:
        subnetId: ${aws_subnet.main.id}
        emrManagedMasterSecurityGroup: ${aws_security_group.sg.id}
        emrManagedSlaveSecurityGroup: ${aws_security_group.sg.id}
        instanceProfile: ${aws_iam_instance_profile.emr_profile.arn}
      masterInstanceGroup:
        instanceType: m4.large
      coreInstanceGroup:
        instanceType: c4.large
        instanceCount: 1
        ebsConfigs:
          - size: '40'
            type: gp2
            volumesPerInstance: 1
        bidPrice: '0.30'
        autoscalingPolicy: |
          {
          "Constraints": {
            "MinCapacity": 1,
            "MaxCapacity": 2
          },
          "Rules": [
            {
              "Name": "ScaleOutMemoryPercentage",
              "Description": "Scale out if YARNMemoryAvailablePercentage is less than 15",
              "Action": {
                "SimpleScalingPolicyConfiguration": {
                  "AdjustmentType": "CHANGE_IN_CAPACITY",
                  "ScalingAdjustment": 1,
                  "CoolDown": 300
                }
              },
              "Trigger": {
                "CloudWatchAlarmDefinition": {
                  "ComparisonOperator": "LESS_THAN",
                  "EvaluationPeriods": 1,
                  "MetricName": "YARNMemoryAvailablePercentage",
                  "Namespace": "AWS/ElasticMapReduce",
                  "Period": 300,
                  "Statistic": "AVERAGE",
                  "Threshold": 15.0,
                  "Unit": "PERCENT"
                }
              }
            }
          ]
          }
      ebsRootVolumeSize: 100
      tags:
        role: rolename
        env: env
      bootstrapActions:
        - path: s3://elasticmapreduce/bootstrap-actions/run-if
          name: runif
          args:
            - instance.isMaster=true
            - echo running on master node
      configurationsJson: |2
          [
            {
              "Classification": "hadoop-env",
              "Configurations": [
                {
                  "Classification": "export",
                  "Properties": {
                    "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
                  }
                }
              ],
              "Properties": {}
            },
            {
              "Classification": "spark-env",
              "Configurations": [
                {
                  "Classification": "export",
                  "Properties": {
                    "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
                  }
                }
              ],
              "Properties": {}
            }
          ]
      serviceRole: ${aws_iam_role.iam_emr_service_role.arn}
```

The `aws.emr.Cluster` resource typically requires two IAM roles, one for the EMR Cluster to use as a service role, and another is assigned to every EC2 instance in a cluster and each application process that runs on a cluster assumes this role for permissions to interact with other AWS services. An additional role, the Auto Scaling role, is required if your cluster uses automatic scaling in Amazon EMR.

The default AWS managed EMR service role is called `EMR_DefaultRole` with Amazon managed policy `AmazonEMRServicePolicy_v2` attached. The name of default instance profile role is `EMR_EC2_DefaultRole` with default managed policy `AmazonElasticMapReduceforEC2Role` attached, but it is on the path to deprecation and will not be replaced with another default managed policy. You'll need to create and specify an instance profile to replace the deprecated role and default policy. See the [Configure IAM service roles for Amazon EMR](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-iam-roles.html) guide for more information on these IAM roles. There is also a fully-bootable example Pulumi configuration at the bottom of this page.
{{% /example %}}
{{% example %}}
### Instance Fleet

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.emr.Cluster("example", {
    masterInstanceFleet: {
        instanceTypeConfigs: [{
            instanceType: "m4.xlarge",
        }],
        targetOnDemandCapacity: 1,
    },
    coreInstanceFleet: {
        instanceTypeConfigs: [
            {
                bidPriceAsPercentageOfOnDemandPrice: 80,
                ebsConfigs: [{
                    size: 100,
                    type: "gp2",
                    volumesPerInstance: 1,
                }],
                instanceType: "m3.xlarge",
                weightedCapacity: 1,
            },
            {
                bidPriceAsPercentageOfOnDemandPrice: 100,
                ebsConfigs: [{
                    size: 100,
                    type: "gp2",
                    volumesPerInstance: 1,
                }],
                instanceType: "m4.xlarge",
                weightedCapacity: 1,
            },
            {
                bidPriceAsPercentageOfOnDemandPrice: 100,
                ebsConfigs: [{
                    size: 100,
                    type: "gp2",
                    volumesPerInstance: 1,
                }],
                instanceType: "m4.2xlarge",
                weightedCapacity: 2,
            },
        ],
        launchSpecifications: {
            spotSpecifications: [{
                allocationStrategy: "capacity-optimized",
                blockDurationMinutes: 0,
                timeoutAction: "SWITCH_TO_ON_DEMAND",
                timeoutDurationMinutes: 10,
            }],
        },
        name: "core fleet",
        targetOnDemandCapacity: 2,
        targetSpotCapacity: 2,
    },
});
const task = new aws.emr.InstanceFleet("task", {
    clusterId: example.id,
    instanceTypeConfigs: [
        {
            bidPriceAsPercentageOfOnDemandPrice: 100,
            ebsConfigs: [{
                size: 100,
                type: "gp2",
                volumesPerInstance: 1,
            }],
            instanceType: "m4.xlarge",
            weightedCapacity: 1,
        },
        {
            bidPriceAsPercentageOfOnDemandPrice: 100,
            ebsConfigs: [{
                size: 100,
                type: "gp2",
                volumesPerInstance: 1,
            }],
            instanceType: "m4.2xlarge",
            weightedCapacity: 2,
        },
    ],
    launchSpecifications: {
        spotSpecifications: [{
            allocationStrategy: "capacity-optimized",
            blockDurationMinutes: 0,
            timeoutAction: "TERMINATE_CLUSTER",
            timeoutDurationMinutes: 10,
        }],
    },
    targetOnDemandCapacity: 1,
    targetSpotCapacity: 1,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.emr.Cluster("example",
    master_instance_fleet=aws.emr.ClusterMasterInstanceFleetArgs(
        instance_type_configs=[aws.emr.ClusterMasterInstanceFleetInstanceTypeConfigArgs(
            instance_type="m4.xlarge",
        )],
        target_on_demand_capacity=1,
    ),
    core_instance_fleet=aws.emr.ClusterCoreInstanceFleetArgs(
        instance_type_configs=[
            aws.emr.ClusterCoreInstanceFleetInstanceTypeConfigArgs(
                bid_price_as_percentage_of_on_demand_price=80,
                ebs_configs=[aws.emr.ClusterCoreInstanceFleetInstanceTypeConfigEbsConfigArgs(
                    size=100,
                    type="gp2",
                    volumes_per_instance=1,
                )],
                instance_type="m3.xlarge",
                weighted_capacity=1,
            ),
            aws.emr.ClusterCoreInstanceFleetInstanceTypeConfigArgs(
                bid_price_as_percentage_of_on_demand_price=100,
                ebs_configs=[aws.emr.ClusterCoreInstanceFleetInstanceTypeConfigEbsConfigArgs(
                    size=100,
                    type="gp2",
                    volumes_per_instance=1,
                )],
                instance_type="m4.xlarge",
                weighted_capacity=1,
            ),
            aws.emr.ClusterCoreInstanceFleetInstanceTypeConfigArgs(
                bid_price_as_percentage_of_on_demand_price=100,
                ebs_configs=[aws.emr.ClusterCoreInstanceFleetInstanceTypeConfigEbsConfigArgs(
                    size=100,
                    type="gp2",
                    volumes_per_instance=1,
                )],
                instance_type="m4.2xlarge",
                weighted_capacity=2,
            ),
        ],
        launch_specifications=aws.emr.ClusterCoreInstanceFleetLaunchSpecificationsArgs(
            spot_specifications=[aws.emr.ClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationArgs(
                allocation_strategy="capacity-optimized",
                block_duration_minutes=0,
                timeout_action="SWITCH_TO_ON_DEMAND",
                timeout_duration_minutes=10,
            )],
        ),
        name="core fleet",
        target_on_demand_capacity=2,
        target_spot_capacity=2,
    ))
task = aws.emr.InstanceFleet("task",
    cluster_id=example.id,
    instance_type_configs=[
        aws.emr.InstanceFleetInstanceTypeConfigArgs(
            bid_price_as_percentage_of_on_demand_price=100,
            ebs_configs=[aws.emr.InstanceFleetInstanceTypeConfigEbsConfigArgs(
                size=100,
                type="gp2",
                volumes_per_instance=1,
            )],
            instance_type="m4.xlarge",
            weighted_capacity=1,
        ),
        aws.emr.InstanceFleetInstanceTypeConfigArgs(
            bid_price_as_percentage_of_on_demand_price=100,
            ebs_configs=[aws.emr.InstanceFleetInstanceTypeConfigEbsConfigArgs(
                size=100,
                type="gp2",
                volumes_per_instance=1,
            )],
            instance_type="m4.2xlarge",
            weighted_capacity=2,
        ),
    ],
    launch_specifications=aws.emr.InstanceFleetLaunchSpecificationsArgs(
        spot_specifications=[aws.emr.InstanceFleetLaunchSpecificationsSpotSpecificationArgs(
            allocation_strategy="capacity-optimized",
            block_duration_minutes=0,
            timeout_action="TERMINATE_CLUSTER",
            timeout_duration_minutes=10,
        )],
    ),
    target_on_demand_capacity=1,
    target_spot_capacity=1)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Emr.Cluster("example", new()
    {
        MasterInstanceFleet = new Aws.Emr.Inputs.ClusterMasterInstanceFleetArgs
        {
            InstanceTypeConfigs = new[]
            {
                new Aws.Emr.Inputs.ClusterMasterInstanceFleetInstanceTypeConfigArgs
                {
                    InstanceType = "m4.xlarge",
                },
            },
            TargetOnDemandCapacity = 1,
        },
        CoreInstanceFleet = new Aws.Emr.Inputs.ClusterCoreInstanceFleetArgs
        {
            InstanceTypeConfigs = new[]
            {
                new Aws.Emr.Inputs.ClusterCoreInstanceFleetInstanceTypeConfigArgs
                {
                    BidPriceAsPercentageOfOnDemandPrice = 80,
                    EbsConfigs = new[]
                    {
                        new Aws.Emr.Inputs.ClusterCoreInstanceFleetInstanceTypeConfigEbsConfigArgs
                        {
                            Size = 100,
                            Type = "gp2",
                            VolumesPerInstance = 1,
                        },
                    },
                    InstanceType = "m3.xlarge",
                    WeightedCapacity = 1,
                },
                new Aws.Emr.Inputs.ClusterCoreInstanceFleetInstanceTypeConfigArgs
                {
                    BidPriceAsPercentageOfOnDemandPrice = 100,
                    EbsConfigs = new[]
                    {
                        new Aws.Emr.Inputs.ClusterCoreInstanceFleetInstanceTypeConfigEbsConfigArgs
                        {
                            Size = 100,
                            Type = "gp2",
                            VolumesPerInstance = 1,
                        },
                    },
                    InstanceType = "m4.xlarge",
                    WeightedCapacity = 1,
                },
                new Aws.Emr.Inputs.ClusterCoreInstanceFleetInstanceTypeConfigArgs
                {
                    BidPriceAsPercentageOfOnDemandPrice = 100,
                    EbsConfigs = new[]
                    {
                        new Aws.Emr.Inputs.ClusterCoreInstanceFleetInstanceTypeConfigEbsConfigArgs
                        {
                            Size = 100,
                            Type = "gp2",
                            VolumesPerInstance = 1,
                        },
                    },
                    InstanceType = "m4.2xlarge",
                    WeightedCapacity = 2,
                },
            },
            LaunchSpecifications = new Aws.Emr.Inputs.ClusterCoreInstanceFleetLaunchSpecificationsArgs
            {
                SpotSpecifications = new[]
                {
                    new Aws.Emr.Inputs.ClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationArgs
                    {
                        AllocationStrategy = "capacity-optimized",
                        BlockDurationMinutes = 0,
                        TimeoutAction = "SWITCH_TO_ON_DEMAND",
                        TimeoutDurationMinutes = 10,
                    },
                },
            },
            Name = "core fleet",
            TargetOnDemandCapacity = 2,
            TargetSpotCapacity = 2,
        },
    });

    var task = new Aws.Emr.InstanceFleet("task", new()
    {
        ClusterId = example.Id,
        InstanceTypeConfigs = new[]
        {
            new Aws.Emr.Inputs.InstanceFleetInstanceTypeConfigArgs
            {
                BidPriceAsPercentageOfOnDemandPrice = 100,
                EbsConfigs = new[]
                {
                    new Aws.Emr.Inputs.InstanceFleetInstanceTypeConfigEbsConfigArgs
                    {
                        Size = 100,
                        Type = "gp2",
                        VolumesPerInstance = 1,
                    },
                },
                InstanceType = "m4.xlarge",
                WeightedCapacity = 1,
            },
            new Aws.Emr.Inputs.InstanceFleetInstanceTypeConfigArgs
            {
                BidPriceAsPercentageOfOnDemandPrice = 100,
                EbsConfigs = new[]
                {
                    new Aws.Emr.Inputs.InstanceFleetInstanceTypeConfigEbsConfigArgs
                    {
                        Size = 100,
                        Type = "gp2",
                        VolumesPerInstance = 1,
                    },
                },
                InstanceType = "m4.2xlarge",
                WeightedCapacity = 2,
            },
        },
        LaunchSpecifications = new Aws.Emr.Inputs.InstanceFleetLaunchSpecificationsArgs
        {
            SpotSpecifications = new[]
            {
                new Aws.Emr.Inputs.InstanceFleetLaunchSpecificationsSpotSpecificationArgs
                {
                    AllocationStrategy = "capacity-optimized",
                    BlockDurationMinutes = 0,
                    TimeoutAction = "TERMINATE_CLUSTER",
                    TimeoutDurationMinutes = 10,
                },
            },
        },
        TargetOnDemandCapacity = 1,
        TargetSpotCapacity = 1,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := emr.NewCluster(ctx, "example", &emr.ClusterArgs{
			MasterInstanceFleet: &emr.ClusterMasterInstanceFleetArgs{
				InstanceTypeConfigs: emr.ClusterMasterInstanceFleetInstanceTypeConfigArray{
					&emr.ClusterMasterInstanceFleetInstanceTypeConfigArgs{
						InstanceType: pulumi.String("m4.xlarge"),
					},
				},
				TargetOnDemandCapacity: pulumi.Int(1),
			},
			CoreInstanceFleet: &emr.ClusterCoreInstanceFleetArgs{
				InstanceTypeConfigs: emr.ClusterCoreInstanceFleetInstanceTypeConfigArray{
					&emr.ClusterCoreInstanceFleetInstanceTypeConfigArgs{
						BidPriceAsPercentageOfOnDemandPrice: pulumi.Float64(80),
						EbsConfigs: emr.ClusterCoreInstanceFleetInstanceTypeConfigEbsConfigArray{
							&emr.ClusterCoreInstanceFleetInstanceTypeConfigEbsConfigArgs{
								Size:               pulumi.Int(100),
								Type:               pulumi.String("gp2"),
								VolumesPerInstance: pulumi.Int(1),
							},
						},
						InstanceType:     pulumi.String("m3.xlarge"),
						WeightedCapacity: pulumi.Int(1),
					},
					&emr.ClusterCoreInstanceFleetInstanceTypeConfigArgs{
						BidPriceAsPercentageOfOnDemandPrice: pulumi.Float64(100),
						EbsConfigs: emr.ClusterCoreInstanceFleetInstanceTypeConfigEbsConfigArray{
							&emr.ClusterCoreInstanceFleetInstanceTypeConfigEbsConfigArgs{
								Size:               pulumi.Int(100),
								Type:               pulumi.String("gp2"),
								VolumesPerInstance: pulumi.Int(1),
							},
						},
						InstanceType:     pulumi.String("m4.xlarge"),
						WeightedCapacity: pulumi.Int(1),
					},
					&emr.ClusterCoreInstanceFleetInstanceTypeConfigArgs{
						BidPriceAsPercentageOfOnDemandPrice: pulumi.Float64(100),
						EbsConfigs: emr.ClusterCoreInstanceFleetInstanceTypeConfigEbsConfigArray{
							&emr.ClusterCoreInstanceFleetInstanceTypeConfigEbsConfigArgs{
								Size:               pulumi.Int(100),
								Type:               pulumi.String("gp2"),
								VolumesPerInstance: pulumi.Int(1),
							},
						},
						InstanceType:     pulumi.String("m4.2xlarge"),
						WeightedCapacity: pulumi.Int(2),
					},
				},
				LaunchSpecifications: &emr.ClusterCoreInstanceFleetLaunchSpecificationsArgs{
					SpotSpecifications: emr.ClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationArray{
						&emr.ClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationArgs{
							AllocationStrategy:     pulumi.String("capacity-optimized"),
							BlockDurationMinutes:   pulumi.Int(0),
							TimeoutAction:          pulumi.String("SWITCH_TO_ON_DEMAND"),
							TimeoutDurationMinutes: pulumi.Int(10),
						},
					},
				},
				Name:                   pulumi.String("core fleet"),
				TargetOnDemandCapacity: pulumi.Int(2),
				TargetSpotCapacity:     pulumi.Int(2),
			},
		})
		if err != nil {
			return err
		}
		_, err = emr.NewInstanceFleet(ctx, "task", &emr.InstanceFleetArgs{
			ClusterId: example.ID(),
			InstanceTypeConfigs: emr.InstanceFleetInstanceTypeConfigArray{
				&emr.InstanceFleetInstanceTypeConfigArgs{
					BidPriceAsPercentageOfOnDemandPrice: pulumi.Float64(100),
					EbsConfigs: emr.InstanceFleetInstanceTypeConfigEbsConfigArray{
						&emr.InstanceFleetInstanceTypeConfigEbsConfigArgs{
							Size:               pulumi.Int(100),
							Type:               pulumi.String("gp2"),
							VolumesPerInstance: pulumi.Int(1),
						},
					},
					InstanceType:     pulumi.String("m4.xlarge"),
					WeightedCapacity: pulumi.Int(1),
				},
				&emr.InstanceFleetInstanceTypeConfigArgs{
					BidPriceAsPercentageOfOnDemandPrice: pulumi.Float64(100),
					EbsConfigs: emr.InstanceFleetInstanceTypeConfigEbsConfigArray{
						&emr.InstanceFleetInstanceTypeConfigEbsConfigArgs{
							Size:               pulumi.Int(100),
							Type:               pulumi.String("gp2"),
							VolumesPerInstance: pulumi.Int(1),
						},
					},
					InstanceType:     pulumi.String("m4.2xlarge"),
					WeightedCapacity: pulumi.Int(2),
				},
			},
			LaunchSpecifications: &emr.InstanceFleetLaunchSpecificationsArgs{
				SpotSpecifications: emr.InstanceFleetLaunchSpecificationsSpotSpecificationArray{
					&emr.InstanceFleetLaunchSpecificationsSpotSpecificationArgs{
						AllocationStrategy:     pulumi.String("capacity-optimized"),
						BlockDurationMinutes:   pulumi.Int(0),
						TimeoutAction:          pulumi.String("TERMINATE_CLUSTER"),
						TimeoutDurationMinutes: pulumi.Int(10),
					},
				},
			},
			TargetOnDemandCapacity: pulumi.Int(1),
			TargetSpotCapacity:     pulumi.Int(1),
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
import com.pulumi.aws.emr.Cluster;
import com.pulumi.aws.emr.ClusterArgs;
import com.pulumi.aws.emr.inputs.ClusterMasterInstanceFleetArgs;
import com.pulumi.aws.emr.inputs.ClusterCoreInstanceFleetArgs;
import com.pulumi.aws.emr.inputs.ClusterCoreInstanceFleetLaunchSpecificationsArgs;
import com.pulumi.aws.emr.InstanceFleet;
import com.pulumi.aws.emr.InstanceFleetArgs;
import com.pulumi.aws.emr.inputs.InstanceFleetInstanceTypeConfigArgs;
import com.pulumi.aws.emr.inputs.InstanceFleetLaunchSpecificationsArgs;
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
            .masterInstanceFleet(ClusterMasterInstanceFleetArgs.builder()
                .instanceTypeConfigs(ClusterMasterInstanceFleetInstanceTypeConfigArgs.builder()
                    .instanceType("m4.xlarge")
                    .build())
                .targetOnDemandCapacity(1)
                .build())
            .coreInstanceFleet(ClusterCoreInstanceFleetArgs.builder()
                .instanceTypeConfigs(                
                    ClusterCoreInstanceFleetInstanceTypeConfigArgs.builder()
                        .bidPriceAsPercentageOfOnDemandPrice(80)
                        .ebsConfigs(ClusterCoreInstanceFleetInstanceTypeConfigEbsConfigArgs.builder()
                            .size(100)
                            .type("gp2")
                            .volumesPerInstance(1)
                            .build())
                        .instanceType("m3.xlarge")
                        .weightedCapacity(1)
                        .build(),
                    ClusterCoreInstanceFleetInstanceTypeConfigArgs.builder()
                        .bidPriceAsPercentageOfOnDemandPrice(100)
                        .ebsConfigs(ClusterCoreInstanceFleetInstanceTypeConfigEbsConfigArgs.builder()
                            .size(100)
                            .type("gp2")
                            .volumesPerInstance(1)
                            .build())
                        .instanceType("m4.xlarge")
                        .weightedCapacity(1)
                        .build(),
                    ClusterCoreInstanceFleetInstanceTypeConfigArgs.builder()
                        .bidPriceAsPercentageOfOnDemandPrice(100)
                        .ebsConfigs(ClusterCoreInstanceFleetInstanceTypeConfigEbsConfigArgs.builder()
                            .size(100)
                            .type("gp2")
                            .volumesPerInstance(1)
                            .build())
                        .instanceType("m4.2xlarge")
                        .weightedCapacity(2)
                        .build())
                .launchSpecifications(ClusterCoreInstanceFleetLaunchSpecificationsArgs.builder()
                    .spotSpecifications(ClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationArgs.builder()
                        .allocationStrategy("capacity-optimized")
                        .blockDurationMinutes(0)
                        .timeoutAction("SWITCH_TO_ON_DEMAND")
                        .timeoutDurationMinutes(10)
                        .build())
                    .build())
                .name("core fleet")
                .targetOnDemandCapacity(2)
                .targetSpotCapacity(2)
                .build())
            .build());

        var task = new InstanceFleet("task", InstanceFleetArgs.builder()        
            .clusterId(example.id())
            .instanceTypeConfigs(            
                InstanceFleetInstanceTypeConfigArgs.builder()
                    .bidPriceAsPercentageOfOnDemandPrice(100)
                    .ebsConfigs(InstanceFleetInstanceTypeConfigEbsConfigArgs.builder()
                        .size(100)
                        .type("gp2")
                        .volumesPerInstance(1)
                        .build())
                    .instanceType("m4.xlarge")
                    .weightedCapacity(1)
                    .build(),
                InstanceFleetInstanceTypeConfigArgs.builder()
                    .bidPriceAsPercentageOfOnDemandPrice(100)
                    .ebsConfigs(InstanceFleetInstanceTypeConfigEbsConfigArgs.builder()
                        .size(100)
                        .type("gp2")
                        .volumesPerInstance(1)
                        .build())
                    .instanceType("m4.2xlarge")
                    .weightedCapacity(2)
                    .build())
            .launchSpecifications(InstanceFleetLaunchSpecificationsArgs.builder()
                .spotSpecifications(InstanceFleetLaunchSpecificationsSpotSpecificationArgs.builder()
                    .allocationStrategy("capacity-optimized")
                    .blockDurationMinutes(0)
                    .timeoutAction("TERMINATE_CLUSTER")
                    .timeoutDurationMinutes(10)
                    .build())
                .build())
            .targetOnDemandCapacity(1)
            .targetSpotCapacity(1)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:emr:Cluster
    properties:
      masterInstanceFleet:
        instanceTypeConfigs:
          - instanceType: m4.xlarge
        targetOnDemandCapacity: 1
      coreInstanceFleet:
        instanceTypeConfigs:
          - bidPriceAsPercentageOfOnDemandPrice: 80
            ebsConfigs:
              - size: 100
                type: gp2
                volumesPerInstance: 1
            instanceType: m3.xlarge
            weightedCapacity: 1
          - bidPriceAsPercentageOfOnDemandPrice: 100
            ebsConfigs:
              - size: 100
                type: gp2
                volumesPerInstance: 1
            instanceType: m4.xlarge
            weightedCapacity: 1
          - bidPriceAsPercentageOfOnDemandPrice: 100
            ebsConfigs:
              - size: 100
                type: gp2
                volumesPerInstance: 1
            instanceType: m4.2xlarge
            weightedCapacity: 2
        launchSpecifications:
          spotSpecifications:
            - allocationStrategy: capacity-optimized
              blockDurationMinutes: 0
              timeoutAction: SWITCH_TO_ON_DEMAND
              timeoutDurationMinutes: 10
        name: core fleet
        targetOnDemandCapacity: 2
        targetSpotCapacity: 2
  task:
    type: aws:emr:InstanceFleet
    properties:
      clusterId: ${example.id}
      instanceTypeConfigs:
        - bidPriceAsPercentageOfOnDemandPrice: 100
          ebsConfigs:
            - size: 100
              type: gp2
              volumesPerInstance: 1
          instanceType: m4.xlarge
          weightedCapacity: 1
        - bidPriceAsPercentageOfOnDemandPrice: 100
          ebsConfigs:
            - size: 100
              type: gp2
              volumesPerInstance: 1
          instanceType: m4.2xlarge
          weightedCapacity: 2
      launchSpecifications:
        spotSpecifications:
          - allocationStrategy: capacity-optimized
            blockDurationMinutes: 0
            timeoutAction: TERMINATE_CLUSTER
            timeoutDurationMinutes: 10
      targetOnDemandCapacity: 1
      targetSpotCapacity: 1
```
{{% /example %}}
{{% example %}}
### Enable Debug Logging

[Debug logging in EMR](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-debugging.html) is implemented as a step. It is highly recommended that you utilize the resource options configuration with `ignoreChanges` if other steps are being managed outside of this provider.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// ... other configuration ...
const example = new aws.emr.Cluster("example", {steps: [{
    actionOnFailure: "TERMINATE_CLUSTER",
    name: "Setup Hadoop Debugging",
    hadoopJarStep: {
        jar: "command-runner.jar",
        args: ["state-pusher-script"],
    },
}]});
```
```python
import pulumi
import pulumi_aws as aws

# ... other configuration ...
example = aws.emr.Cluster("example", steps=[aws.emr.ClusterStepArgs(
    action_on_failure="TERMINATE_CLUSTER",
    name="Setup Hadoop Debugging",
    hadoop_jar_step=aws.emr.ClusterStepHadoopJarStepArgs(
        jar="command-runner.jar",
        args=["state-pusher-script"],
    ),
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // ... other configuration ...
    var example = new Aws.Emr.Cluster("example", new()
    {
        Steps = new[]
        {
            new Aws.Emr.Inputs.ClusterStepArgs
            {
                ActionOnFailure = "TERMINATE_CLUSTER",
                Name = "Setup Hadoop Debugging",
                HadoopJarStep = new Aws.Emr.Inputs.ClusterStepHadoopJarStepArgs
                {
                    Jar = "command-runner.jar",
                    Args = new[]
                    {
                        "state-pusher-script",
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := emr.NewCluster(ctx, "example", &emr.ClusterArgs{
			Steps: emr.ClusterStepArray{
				&emr.ClusterStepArgs{
					ActionOnFailure: pulumi.String("TERMINATE_CLUSTER"),
					Name:            pulumi.String("Setup Hadoop Debugging"),
					HadoopJarStep: &emr.ClusterStepHadoopJarStepArgs{
						Jar: pulumi.String("command-runner.jar"),
						Args: pulumi.StringArray{
							pulumi.String("state-pusher-script"),
						},
					},
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
import com.pulumi.aws.emr.Cluster;
import com.pulumi.aws.emr.ClusterArgs;
import com.pulumi.aws.emr.inputs.ClusterStepArgs;
import com.pulumi.aws.emr.inputs.ClusterStepHadoopJarStepArgs;
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
            .steps(ClusterStepArgs.builder()
                .actionOnFailure("TERMINATE_CLUSTER")
                .name("Setup Hadoop Debugging")
                .hadoopJarStep(ClusterStepHadoopJarStepArgs.builder()
                    .jar("command-runner.jar")
                    .args("state-pusher-script")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:emr:Cluster
    properties:
      steps:
        - actionOnFailure: TERMINATE_CLUSTER
          name: Setup Hadoop Debugging
          hadoopJarStep:
            jar: command-runner.jar
            args:
              - state-pusher-script
```
{{% /example %}}
{{% example %}}
### Multiple Node Master Instance Group

Available in EMR version 5.23.0 and later, an EMR Cluster can be launched with three master nodes for high availability. Additional information about this functionality and its requirements can be found in the [EMR Management Guide](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-ha.html).

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// This configuration is for illustrative purposes and highlights
// only relevant configurations for working with this functionality.
// Map public IP on launch must be enabled for public (Internet accessible) subnets
// ... other configuration ...
const exampleSubnet = new aws.ec2.Subnet("exampleSubnet", {mapPublicIpOnLaunch: true});
// ... other configuration ...
const exampleCluster = new aws.emr.Cluster("exampleCluster", {
    releaseLabel: "emr-5.24.1",
    terminationProtection: true,
    ec2Attributes: {
        subnetId: exampleSubnet.id,
    },
    masterInstanceGroup: {
        instanceCount: 3,
    },
    coreInstanceGroup: {},
});
```
```python
import pulumi
import pulumi_aws as aws

# This configuration is for illustrative purposes and highlights
# only relevant configurations for working with this functionality.
# Map public IP on launch must be enabled for public (Internet accessible) subnets
# ... other configuration ...
example_subnet = aws.ec2.Subnet("exampleSubnet", map_public_ip_on_launch=True)
# ... other configuration ...
example_cluster = aws.emr.Cluster("exampleCluster",
    release_label="emr-5.24.1",
    termination_protection=True,
    ec2_attributes=aws.emr.ClusterEc2AttributesArgs(
        subnet_id=example_subnet.id,
    ),
    master_instance_group=aws.emr.ClusterMasterInstanceGroupArgs(
        instance_count=3,
    ),
    core_instance_group=aws.emr.ClusterCoreInstanceGroupArgs())
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // This configuration is for illustrative purposes and highlights
    // only relevant configurations for working with this functionality.
    // Map public IP on launch must be enabled for public (Internet accessible) subnets
    // ... other configuration ...
    var exampleSubnet = new Aws.Ec2.Subnet("exampleSubnet", new()
    {
        MapPublicIpOnLaunch = true,
    });

    // ... other configuration ...
    var exampleCluster = new Aws.Emr.Cluster("exampleCluster", new()
    {
        ReleaseLabel = "emr-5.24.1",
        TerminationProtection = true,
        Ec2Attributes = new Aws.Emr.Inputs.ClusterEc2AttributesArgs
        {
            SubnetId = exampleSubnet.Id,
        },
        MasterInstanceGroup = new Aws.Emr.Inputs.ClusterMasterInstanceGroupArgs
        {
            InstanceCount = 3,
        },
        CoreInstanceGroup = null,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleSubnet, err := ec2.NewSubnet(ctx, "exampleSubnet", &ec2.SubnetArgs{
			MapPublicIpOnLaunch: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = emr.NewCluster(ctx, "exampleCluster", &emr.ClusterArgs{
			ReleaseLabel:          pulumi.String("emr-5.24.1"),
			TerminationProtection: pulumi.Bool(true),
			Ec2Attributes: &emr.ClusterEc2AttributesArgs{
				SubnetId: exampleSubnet.ID(),
			},
			MasterInstanceGroup: &emr.ClusterMasterInstanceGroupArgs{
				InstanceCount: pulumi.Int(3),
			},
			CoreInstanceGroup: nil,
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
import com.pulumi.aws.ec2.Subnet;
import com.pulumi.aws.ec2.SubnetArgs;
import com.pulumi.aws.emr.Cluster;
import com.pulumi.aws.emr.ClusterArgs;
import com.pulumi.aws.emr.inputs.ClusterEc2AttributesArgs;
import com.pulumi.aws.emr.inputs.ClusterMasterInstanceGroupArgs;
import com.pulumi.aws.emr.inputs.ClusterCoreInstanceGroupArgs;
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
        var exampleSubnet = new Subnet("exampleSubnet", SubnetArgs.builder()        
            .mapPublicIpOnLaunch(true)
            .build());

        var exampleCluster = new Cluster("exampleCluster", ClusterArgs.builder()        
            .releaseLabel("emr-5.24.1")
            .terminationProtection(true)
            .ec2Attributes(ClusterEc2AttributesArgs.builder()
                .subnetId(exampleSubnet.id())
                .build())
            .masterInstanceGroup(ClusterMasterInstanceGroupArgs.builder()
                .instanceCount(3)
                .build())
            .coreInstanceGroup()
            .build());

    }
}
```
```yaml
resources:
  # This configuration is for illustrative purposes and highlights
  # only relevant configurations for working with this functionality.

  # Map public IP on launch must be enabled for public (Internet accessible) subnets
  exampleSubnet:
    type: aws:ec2:Subnet
    properties:
      mapPublicIpOnLaunch: true
  exampleCluster:
    type: aws:emr:Cluster
    properties:
      # EMR version must be 5.23.0 or later
      releaseLabel: emr-5.24.1
      # Termination protection is automatically enabled for multiple masters
      #   # To destroy the cluster, this must be configured to false and applied first
      terminationProtection: true
      ec2Attributes:
        subnetId: ${exampleSubnet.id}
      masterInstanceGroup:
        instanceCount: 3
      coreInstanceGroup: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EMR clusters using the `id`. For example:

```sh
 $ pulumi import aws:emr/cluster:Cluster cluster j-123456ABCDEF
```
 Since the API does not return the actual values for Kerberos configurations, environments with those options set will need to use the `lifecycle` configuration block `ignore_changes` argument available to all Pulumi resources to prevent perpetual differences. For example:

