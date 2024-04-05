Provides an Auto Scaling Group resource.

> **Note:** You must specify either `launch_configuration`, `launch_template`, or `mixed_instances_policy`.

> **NOTE on Auto Scaling Groups, Attachments and Traffic Source Attachments:** Pulumi provides standalone Attachment (for attaching Classic Load Balancers and Application Load Balancer, Gateway Load Balancer, or Network Load Balancer target groups) and Traffic Source Attachment (for attaching Load Balancers and VPC Lattice target groups) resources and an Auto Scaling Group resource with `load_balancers`, `target_group_arns` and `traffic_source` attributes. Do not use the same traffic source in more than one of these resources. Doing so will cause a conflict of attachments. A `lifecycle` configuration block can be used to suppress differences if necessary.



{{% examples %}}
## Example Usage
{{% example %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.ec2.PlacementGroup;
import com.pulumi.aws.ec2.PlacementGroupArgs;
import com.pulumi.aws.autoscaling.Group;
import com.pulumi.aws.autoscaling.GroupArgs;
import com.pulumi.aws.autoscaling.inputs.GroupInstanceMaintenancePolicyArgs;
import com.pulumi.aws.autoscaling.inputs.GroupInitialLifecycleHookArgs;
import com.pulumi.aws.autoscaling.inputs.GroupTagArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var test = new PlacementGroup("test", PlacementGroupArgs.builder()        
            .strategy("cluster")
            .build());

        var bar = new Group("bar", GroupArgs.builder()        
            .maxSize(5)
            .minSize(2)
            .healthCheckGracePeriod(300)
            .healthCheckType("ELB")
            .desiredCapacity(4)
            .forceDelete(true)
            .placementGroup(test.id())
            .launchConfiguration(aws_launch_configuration.foobar().name())
            .vpcZoneIdentifiers(            
                aws_subnet.example1().id(),
                aws_subnet.example2().id())
            .instanceMaintenancePolicy(GroupInstanceMaintenancePolicyArgs.builder()
                .minHealthyPercentage(90)
                .maxHealthyPercentage(120)
                .build())
            .initialLifecycleHooks(GroupInitialLifecycleHookArgs.builder()
                .name("foobar")
                .defaultResult("CONTINUE")
                .heartbeatTimeout(2000)
                .lifecycleTransition("autoscaling:EC2_INSTANCE_LAUNCHING")
                .notificationMetadata(serializeJson(
                    jsonObject(
                        jsonProperty("foo", "bar")
                    )))
                .notificationTargetArn("arn:aws:sqs:us-east-1:444455556666:queue1*")
                .roleArn("arn:aws:iam::123456789012:role/S3Access")
                .build())
            .tags(            
                GroupTagArgs.builder()
                    .key("foo")
                    .value("bar")
                    .propagateAtLaunch(true)
                    .build(),
                GroupTagArgs.builder()
                    .key("lorem")
                    .value("ipsum")
                    .propagateAtLaunch(false)
                    .build())
            .timeouts(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:ec2:PlacementGroup
    properties:
      strategy: cluster
  bar:
    type: aws:autoscaling:Group
    properties:
      maxSize: 5
      minSize: 2
      healthCheckGracePeriod: 300
      healthCheckType: ELB
      desiredCapacity: 4
      forceDelete: true
      placementGroup: ${test.id}
      launchConfiguration: ${aws_launch_configuration.foobar.name}
      vpcZoneIdentifiers:
        - ${aws_subnet.example1.id}
        - ${aws_subnet.example2.id}
      instanceMaintenancePolicy:
        minHealthyPercentage: 90
        maxHealthyPercentage: 120
      initialLifecycleHooks:
        - name: foobar
          defaultResult: CONTINUE
          heartbeatTimeout: 2000
          lifecycleTransition: autoscaling:EC2_INSTANCE_LAUNCHING
          notificationMetadata:
            fn::toJSON:
              foo: bar
          notificationTargetArn: arn:aws:sqs:us-east-1:444455556666:queue1*
          roleArn: arn:aws:iam::123456789012:role/S3Access
      tags:
        - key: foo
          value: bar
          propagateAtLaunch: true
        - key: lorem
          value: ipsum
          propagateAtLaunch: false
      timeouts:
        - delete: 15m
```
{{% /example %}}
{{% example %}}
### With Latest Version Of Launch Template

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foobar = new aws.ec2.LaunchTemplate("foobar", {
    namePrefix: "foobar",
    imageId: "ami-1a2b3c",
    instanceType: "t2.micro",
});
const bar = new aws.autoscaling.Group("bar", {
    availabilityZones: ["us-east-1a"],
    desiredCapacity: 1,
    maxSize: 1,
    minSize: 1,
    launchTemplate: {
        id: foobar.id,
        version: "$Latest",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

foobar = aws.ec2.LaunchTemplate("foobar",
    name_prefix="foobar",
    image_id="ami-1a2b3c",
    instance_type="t2.micro")
bar = aws.autoscaling.Group("bar",
    availability_zones=["us-east-1a"],
    desired_capacity=1,
    max_size=1,
    min_size=1,
    launch_template=aws.autoscaling.GroupLaunchTemplateArgs(
        id=foobar.id,
        version="$Latest",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foobar = new Aws.Ec2.LaunchTemplate("foobar", new()
    {
        NamePrefix = "foobar",
        ImageId = "ami-1a2b3c",
        InstanceType = "t2.micro",
    });

    var bar = new Aws.AutoScaling.Group("bar", new()
    {
        AvailabilityZones = new[]
        {
            "us-east-1a",
        },
        DesiredCapacity = 1,
        MaxSize = 1,
        MinSize = 1,
        LaunchTemplate = new Aws.AutoScaling.Inputs.GroupLaunchTemplateArgs
        {
            Id = foobar.Id,
            Version = "$Latest",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		foobar, err := ec2.NewLaunchTemplate(ctx, "foobar", &ec2.LaunchTemplateArgs{
			NamePrefix:   pulumi.String("foobar"),
			ImageId:      pulumi.String("ami-1a2b3c"),
			InstanceType: pulumi.String("t2.micro"),
		})
		if err != nil {
			return err
		}
		_, err = autoscaling.NewGroup(ctx, "bar", &autoscaling.GroupArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-east-1a"),
			},
			DesiredCapacity: pulumi.Int(1),
			MaxSize:         pulumi.Int(1),
			MinSize:         pulumi.Int(1),
			LaunchTemplate: &autoscaling.GroupLaunchTemplateArgs{
				Id:      foobar.ID(),
				Version: pulumi.String("$Latest"),
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
import com.pulumi.aws.ec2.LaunchTemplate;
import com.pulumi.aws.ec2.LaunchTemplateArgs;
import com.pulumi.aws.autoscaling.Group;
import com.pulumi.aws.autoscaling.GroupArgs;
import com.pulumi.aws.autoscaling.inputs.GroupLaunchTemplateArgs;
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
        var foobar = new LaunchTemplate("foobar", LaunchTemplateArgs.builder()        
            .namePrefix("foobar")
            .imageId("ami-1a2b3c")
            .instanceType("t2.micro")
            .build());

        var bar = new Group("bar", GroupArgs.builder()        
            .availabilityZones("us-east-1a")
            .desiredCapacity(1)
            .maxSize(1)
            .minSize(1)
            .launchTemplate(GroupLaunchTemplateArgs.builder()
                .id(foobar.id())
                .version("$Latest")
                .build())
            .build());

    }
}
```
```yaml
resources:
  foobar:
    type: aws:ec2:LaunchTemplate
    properties:
      namePrefix: foobar
      imageId: ami-1a2b3c
      instanceType: t2.micro
  bar:
    type: aws:autoscaling:Group
    properties:
      availabilityZones:
        - us-east-1a
      desiredCapacity: 1
      maxSize: 1
      minSize: 1
      launchTemplate:
        id: ${foobar.id}
        version: $Latest
```
{{% /example %}}
{{% example %}}
### Mixed Instances Policy

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLaunchTemplate = new aws.ec2.LaunchTemplate("exampleLaunchTemplate", {
    namePrefix: "example",
    imageId: data.aws_ami.example.id,
    instanceType: "c5.large",
});
const exampleGroup = new aws.autoscaling.Group("exampleGroup", {
    availabilityZones: ["us-east-1a"],
    desiredCapacity: 1,
    maxSize: 1,
    minSize: 1,
    mixedInstancesPolicy: {
        launchTemplate: {
            launchTemplateSpecification: {
                launchTemplateId: exampleLaunchTemplate.id,
            },
            overrides: [
                {
                    instanceType: "c4.large",
                    weightedCapacity: "3",
                },
                {
                    instanceType: "c3.large",
                    weightedCapacity: "2",
                },
            ],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_launch_template = aws.ec2.LaunchTemplate("exampleLaunchTemplate",
    name_prefix="example",
    image_id=data["aws_ami"]["example"]["id"],
    instance_type="c5.large")
example_group = aws.autoscaling.Group("exampleGroup",
    availability_zones=["us-east-1a"],
    desired_capacity=1,
    max_size=1,
    min_size=1,
    mixed_instances_policy=aws.autoscaling.GroupMixedInstancesPolicyArgs(
        launch_template=aws.autoscaling.GroupMixedInstancesPolicyLaunchTemplateArgs(
            launch_template_specification=aws.autoscaling.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs(
                launch_template_id=example_launch_template.id,
            ),
            overrides=[
                aws.autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs(
                    instance_type="c4.large",
                    weighted_capacity="3",
                ),
                aws.autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs(
                    instance_type="c3.large",
                    weighted_capacity="2",
                ),
            ],
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleLaunchTemplate = new Aws.Ec2.LaunchTemplate("exampleLaunchTemplate", new()
    {
        NamePrefix = "example",
        ImageId = data.Aws_ami.Example.Id,
        InstanceType = "c5.large",
    });

    var exampleGroup = new Aws.AutoScaling.Group("exampleGroup", new()
    {
        AvailabilityZones = new[]
        {
            "us-east-1a",
        },
        DesiredCapacity = 1,
        MaxSize = 1,
        MinSize = 1,
        MixedInstancesPolicy = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyArgs
        {
            LaunchTemplate = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateArgs
            {
                LaunchTemplateSpecification = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs
                {
                    LaunchTemplateId = exampleLaunchTemplate.Id,
                },
                Overrides = new[]
                {
                    new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs
                    {
                        InstanceType = "c4.large",
                        WeightedCapacity = "3",
                    },
                    new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs
                    {
                        InstanceType = "c3.large",
                        WeightedCapacity = "2",
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLaunchTemplate, err := ec2.NewLaunchTemplate(ctx, "exampleLaunchTemplate", &ec2.LaunchTemplateArgs{
			NamePrefix:   pulumi.String("example"),
			ImageId:      pulumi.Any(data.Aws_ami.Example.Id),
			InstanceType: pulumi.String("c5.large"),
		})
		if err != nil {
			return err
		}
		_, err = autoscaling.NewGroup(ctx, "exampleGroup", &autoscaling.GroupArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-east-1a"),
			},
			DesiredCapacity: pulumi.Int(1),
			MaxSize:         pulumi.Int(1),
			MinSize:         pulumi.Int(1),
			MixedInstancesPolicy: &autoscaling.GroupMixedInstancesPolicyArgs{
				LaunchTemplate: &autoscaling.GroupMixedInstancesPolicyLaunchTemplateArgs{
					LaunchTemplateSpecification: &autoscaling.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs{
						LaunchTemplateId: exampleLaunchTemplate.ID(),
					},
					Overrides: autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideArray{
						&autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs{
							InstanceType:     pulumi.String("c4.large"),
							WeightedCapacity: pulumi.String("3"),
						},
						&autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs{
							InstanceType:     pulumi.String("c3.large"),
							WeightedCapacity: pulumi.String("2"),
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
import com.pulumi.aws.ec2.LaunchTemplate;
import com.pulumi.aws.ec2.LaunchTemplateArgs;
import com.pulumi.aws.autoscaling.Group;
import com.pulumi.aws.autoscaling.GroupArgs;
import com.pulumi.aws.autoscaling.inputs.GroupMixedInstancesPolicyArgs;
import com.pulumi.aws.autoscaling.inputs.GroupMixedInstancesPolicyLaunchTemplateArgs;
import com.pulumi.aws.autoscaling.inputs.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs;
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
        var exampleLaunchTemplate = new LaunchTemplate("exampleLaunchTemplate", LaunchTemplateArgs.builder()        
            .namePrefix("example")
            .imageId(data.aws_ami().example().id())
            .instanceType("c5.large")
            .build());

        var exampleGroup = new Group("exampleGroup", GroupArgs.builder()        
            .availabilityZones("us-east-1a")
            .desiredCapacity(1)
            .maxSize(1)
            .minSize(1)
            .mixedInstancesPolicy(GroupMixedInstancesPolicyArgs.builder()
                .launchTemplate(GroupMixedInstancesPolicyLaunchTemplateArgs.builder()
                    .launchTemplateSpecification(GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs.builder()
                        .launchTemplateId(exampleLaunchTemplate.id())
                        .build())
                    .overrides(                    
                        GroupMixedInstancesPolicyLaunchTemplateOverrideArgs.builder()
                            .instanceType("c4.large")
                            .weightedCapacity("3")
                            .build(),
                        GroupMixedInstancesPolicyLaunchTemplateOverrideArgs.builder()
                            .instanceType("c3.large")
                            .weightedCapacity("2")
                            .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleLaunchTemplate:
    type: aws:ec2:LaunchTemplate
    properties:
      namePrefix: example
      imageId: ${data.aws_ami.example.id}
      instanceType: c5.large
  exampleGroup:
    type: aws:autoscaling:Group
    properties:
      availabilityZones:
        - us-east-1a
      desiredCapacity: 1
      maxSize: 1
      minSize: 1
      mixedInstancesPolicy:
        launchTemplate:
          launchTemplateSpecification:
            launchTemplateId: ${exampleLaunchTemplate.id}
          overrides:
            - instanceType: c4.large
              weightedCapacity: '3'
            - instanceType: c3.large
              weightedCapacity: '2'
```
{{% /example %}}
{{% example %}}
### Mixed Instances Policy with Spot Instances and Capacity Rebalance

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLaunchTemplate = new aws.ec2.LaunchTemplate("exampleLaunchTemplate", {
    namePrefix: "example",
    imageId: data.aws_ami.example.id,
    instanceType: "c5.large",
});
const exampleGroup = new aws.autoscaling.Group("exampleGroup", {
    capacityRebalance: true,
    desiredCapacity: 12,
    maxSize: 15,
    minSize: 12,
    vpcZoneIdentifiers: [
        aws_subnet.example1.id,
        aws_subnet.example2.id,
    ],
    mixedInstancesPolicy: {
        instancesDistribution: {
            onDemandBaseCapacity: 0,
            onDemandPercentageAboveBaseCapacity: 25,
            spotAllocationStrategy: "capacity-optimized",
        },
        launchTemplate: {
            launchTemplateSpecification: {
                launchTemplateId: exampleLaunchTemplate.id,
            },
            overrides: [
                {
                    instanceType: "c4.large",
                    weightedCapacity: "3",
                },
                {
                    instanceType: "c3.large",
                    weightedCapacity: "2",
                },
            ],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_launch_template = aws.ec2.LaunchTemplate("exampleLaunchTemplate",
    name_prefix="example",
    image_id=data["aws_ami"]["example"]["id"],
    instance_type="c5.large")
example_group = aws.autoscaling.Group("exampleGroup",
    capacity_rebalance=True,
    desired_capacity=12,
    max_size=15,
    min_size=12,
    vpc_zone_identifiers=[
        aws_subnet["example1"]["id"],
        aws_subnet["example2"]["id"],
    ],
    mixed_instances_policy=aws.autoscaling.GroupMixedInstancesPolicyArgs(
        instances_distribution=aws.autoscaling.GroupMixedInstancesPolicyInstancesDistributionArgs(
            on_demand_base_capacity=0,
            on_demand_percentage_above_base_capacity=25,
            spot_allocation_strategy="capacity-optimized",
        ),
        launch_template=aws.autoscaling.GroupMixedInstancesPolicyLaunchTemplateArgs(
            launch_template_specification=aws.autoscaling.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs(
                launch_template_id=example_launch_template.id,
            ),
            overrides=[
                aws.autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs(
                    instance_type="c4.large",
                    weighted_capacity="3",
                ),
                aws.autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs(
                    instance_type="c3.large",
                    weighted_capacity="2",
                ),
            ],
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleLaunchTemplate = new Aws.Ec2.LaunchTemplate("exampleLaunchTemplate", new()
    {
        NamePrefix = "example",
        ImageId = data.Aws_ami.Example.Id,
        InstanceType = "c5.large",
    });

    var exampleGroup = new Aws.AutoScaling.Group("exampleGroup", new()
    {
        CapacityRebalance = true,
        DesiredCapacity = 12,
        MaxSize = 15,
        MinSize = 12,
        VpcZoneIdentifiers = new[]
        {
            aws_subnet.Example1.Id,
            aws_subnet.Example2.Id,
        },
        MixedInstancesPolicy = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyArgs
        {
            InstancesDistribution = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyInstancesDistributionArgs
            {
                OnDemandBaseCapacity = 0,
                OnDemandPercentageAboveBaseCapacity = 25,
                SpotAllocationStrategy = "capacity-optimized",
            },
            LaunchTemplate = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateArgs
            {
                LaunchTemplateSpecification = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs
                {
                    LaunchTemplateId = exampleLaunchTemplate.Id,
                },
                Overrides = new[]
                {
                    new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs
                    {
                        InstanceType = "c4.large",
                        WeightedCapacity = "3",
                    },
                    new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs
                    {
                        InstanceType = "c3.large",
                        WeightedCapacity = "2",
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLaunchTemplate, err := ec2.NewLaunchTemplate(ctx, "exampleLaunchTemplate", &ec2.LaunchTemplateArgs{
			NamePrefix:   pulumi.String("example"),
			ImageId:      pulumi.Any(data.Aws_ami.Example.Id),
			InstanceType: pulumi.String("c5.large"),
		})
		if err != nil {
			return err
		}
		_, err = autoscaling.NewGroup(ctx, "exampleGroup", &autoscaling.GroupArgs{
			CapacityRebalance: pulumi.Bool(true),
			DesiredCapacity:   pulumi.Int(12),
			MaxSize:           pulumi.Int(15),
			MinSize:           pulumi.Int(12),
			VpcZoneIdentifiers: pulumi.StringArray{
				aws_subnet.Example1.Id,
				aws_subnet.Example2.Id,
			},
			MixedInstancesPolicy: &autoscaling.GroupMixedInstancesPolicyArgs{
				InstancesDistribution: &autoscaling.GroupMixedInstancesPolicyInstancesDistributionArgs{
					OnDemandBaseCapacity:                pulumi.Int(0),
					OnDemandPercentageAboveBaseCapacity: pulumi.Int(25),
					SpotAllocationStrategy:              pulumi.String("capacity-optimized"),
				},
				LaunchTemplate: &autoscaling.GroupMixedInstancesPolicyLaunchTemplateArgs{
					LaunchTemplateSpecification: &autoscaling.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs{
						LaunchTemplateId: exampleLaunchTemplate.ID(),
					},
					Overrides: autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideArray{
						&autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs{
							InstanceType:     pulumi.String("c4.large"),
							WeightedCapacity: pulumi.String("3"),
						},
						&autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs{
							InstanceType:     pulumi.String("c3.large"),
							WeightedCapacity: pulumi.String("2"),
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
import com.pulumi.aws.ec2.LaunchTemplate;
import com.pulumi.aws.ec2.LaunchTemplateArgs;
import com.pulumi.aws.autoscaling.Group;
import com.pulumi.aws.autoscaling.GroupArgs;
import com.pulumi.aws.autoscaling.inputs.GroupMixedInstancesPolicyArgs;
import com.pulumi.aws.autoscaling.inputs.GroupMixedInstancesPolicyInstancesDistributionArgs;
import com.pulumi.aws.autoscaling.inputs.GroupMixedInstancesPolicyLaunchTemplateArgs;
import com.pulumi.aws.autoscaling.inputs.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs;
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
        var exampleLaunchTemplate = new LaunchTemplate("exampleLaunchTemplate", LaunchTemplateArgs.builder()        
            .namePrefix("example")
            .imageId(data.aws_ami().example().id())
            .instanceType("c5.large")
            .build());

        var exampleGroup = new Group("exampleGroup", GroupArgs.builder()        
            .capacityRebalance(true)
            .desiredCapacity(12)
            .maxSize(15)
            .minSize(12)
            .vpcZoneIdentifiers(            
                aws_subnet.example1().id(),
                aws_subnet.example2().id())
            .mixedInstancesPolicy(GroupMixedInstancesPolicyArgs.builder()
                .instancesDistribution(GroupMixedInstancesPolicyInstancesDistributionArgs.builder()
                    .onDemandBaseCapacity(0)
                    .onDemandPercentageAboveBaseCapacity(25)
                    .spotAllocationStrategy("capacity-optimized")
                    .build())
                .launchTemplate(GroupMixedInstancesPolicyLaunchTemplateArgs.builder()
                    .launchTemplateSpecification(GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs.builder()
                        .launchTemplateId(exampleLaunchTemplate.id())
                        .build())
                    .overrides(                    
                        GroupMixedInstancesPolicyLaunchTemplateOverrideArgs.builder()
                            .instanceType("c4.large")
                            .weightedCapacity("3")
                            .build(),
                        GroupMixedInstancesPolicyLaunchTemplateOverrideArgs.builder()
                            .instanceType("c3.large")
                            .weightedCapacity("2")
                            .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleLaunchTemplate:
    type: aws:ec2:LaunchTemplate
    properties:
      namePrefix: example
      imageId: ${data.aws_ami.example.id}
      instanceType: c5.large
  exampleGroup:
    type: aws:autoscaling:Group
    properties:
      capacityRebalance: true
      desiredCapacity: 12
      maxSize: 15
      minSize: 12
      vpcZoneIdentifiers:
        - ${aws_subnet.example1.id}
        - ${aws_subnet.example2.id}
      mixedInstancesPolicy:
        instancesDistribution:
          onDemandBaseCapacity: 0
          onDemandPercentageAboveBaseCapacity: 25
          spotAllocationStrategy: capacity-optimized
        launchTemplate:
          launchTemplateSpecification:
            launchTemplateId: ${exampleLaunchTemplate.id}
          overrides:
            - instanceType: c4.large
              weightedCapacity: '3'
            - instanceType: c3.large
              weightedCapacity: '2'
```
{{% /example %}}
{{% example %}}
### Mixed Instances Policy with Instance level LaunchTemplateSpecification Overrides

When using a diverse instance set, some instance types might require a launch template with configuration values unique to that instance type such as a different AMI (Graviton2), architecture specific user data script, different EBS configuration, or different networking configuration.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLaunchTemplate = new aws.ec2.LaunchTemplate("exampleLaunchTemplate", {
    namePrefix: "example",
    imageId: data.aws_ami.example.id,
    instanceType: "c5.large",
});
const example2 = new aws.ec2.LaunchTemplate("example2", {
    namePrefix: "example2",
    imageId: data.aws_ami.example2.id,
});
const exampleGroup = new aws.autoscaling.Group("exampleGroup", {
    availabilityZones: ["us-east-1a"],
    desiredCapacity: 1,
    maxSize: 1,
    minSize: 1,
    mixedInstancesPolicy: {
        launchTemplate: {
            launchTemplateSpecification: {
                launchTemplateId: exampleLaunchTemplate.id,
            },
            overrides: [
                {
                    instanceType: "c4.large",
                    weightedCapacity: "3",
                },
                {
                    instanceType: "c6g.large",
                    launchTemplateSpecification: {
                        launchTemplateId: example2.id,
                    },
                    weightedCapacity: "2",
                },
            ],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_launch_template = aws.ec2.LaunchTemplate("exampleLaunchTemplate",
    name_prefix="example",
    image_id=data["aws_ami"]["example"]["id"],
    instance_type="c5.large")
example2 = aws.ec2.LaunchTemplate("example2",
    name_prefix="example2",
    image_id=data["aws_ami"]["example2"]["id"])
example_group = aws.autoscaling.Group("exampleGroup",
    availability_zones=["us-east-1a"],
    desired_capacity=1,
    max_size=1,
    min_size=1,
    mixed_instances_policy=aws.autoscaling.GroupMixedInstancesPolicyArgs(
        launch_template=aws.autoscaling.GroupMixedInstancesPolicyLaunchTemplateArgs(
            launch_template_specification=aws.autoscaling.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs(
                launch_template_id=example_launch_template.id,
            ),
            overrides=[
                aws.autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs(
                    instance_type="c4.large",
                    weighted_capacity="3",
                ),
                aws.autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs(
                    instance_type="c6g.large",
                    launch_template_specification=aws.autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationArgs(
                        launch_template_id=example2.id,
                    ),
                    weighted_capacity="2",
                ),
            ],
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleLaunchTemplate = new Aws.Ec2.LaunchTemplate("exampleLaunchTemplate", new()
    {
        NamePrefix = "example",
        ImageId = data.Aws_ami.Example.Id,
        InstanceType = "c5.large",
    });

    var example2 = new Aws.Ec2.LaunchTemplate("example2", new()
    {
        NamePrefix = "example2",
        ImageId = data.Aws_ami.Example2.Id,
    });

    var exampleGroup = new Aws.AutoScaling.Group("exampleGroup", new()
    {
        AvailabilityZones = new[]
        {
            "us-east-1a",
        },
        DesiredCapacity = 1,
        MaxSize = 1,
        MinSize = 1,
        MixedInstancesPolicy = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyArgs
        {
            LaunchTemplate = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateArgs
            {
                LaunchTemplateSpecification = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs
                {
                    LaunchTemplateId = exampleLaunchTemplate.Id,
                },
                Overrides = new[]
                {
                    new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs
                    {
                        InstanceType = "c4.large",
                        WeightedCapacity = "3",
                    },
                    new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs
                    {
                        InstanceType = "c6g.large",
                        LaunchTemplateSpecification = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationArgs
                        {
                            LaunchTemplateId = example2.Id,
                        },
                        WeightedCapacity = "2",
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLaunchTemplate, err := ec2.NewLaunchTemplate(ctx, "exampleLaunchTemplate", &ec2.LaunchTemplateArgs{
			NamePrefix:   pulumi.String("example"),
			ImageId:      pulumi.Any(data.Aws_ami.Example.Id),
			InstanceType: pulumi.String("c5.large"),
		})
		if err != nil {
			return err
		}
		example2, err := ec2.NewLaunchTemplate(ctx, "example2", &ec2.LaunchTemplateArgs{
			NamePrefix: pulumi.String("example2"),
			ImageId:    pulumi.Any(data.Aws_ami.Example2.Id),
		})
		if err != nil {
			return err
		}
		_, err = autoscaling.NewGroup(ctx, "exampleGroup", &autoscaling.GroupArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-east-1a"),
			},
			DesiredCapacity: pulumi.Int(1),
			MaxSize:         pulumi.Int(1),
			MinSize:         pulumi.Int(1),
			MixedInstancesPolicy: &autoscaling.GroupMixedInstancesPolicyArgs{
				LaunchTemplate: &autoscaling.GroupMixedInstancesPolicyLaunchTemplateArgs{
					LaunchTemplateSpecification: &autoscaling.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs{
						LaunchTemplateId: exampleLaunchTemplate.ID(),
					},
					Overrides: autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideArray{
						&autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs{
							InstanceType:     pulumi.String("c4.large"),
							WeightedCapacity: pulumi.String("3"),
						},
						&autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs{
							InstanceType: pulumi.String("c6g.large"),
							LaunchTemplateSpecification: &autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationArgs{
								LaunchTemplateId: example2.ID(),
							},
							WeightedCapacity: pulumi.String("2"),
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
import com.pulumi.aws.ec2.LaunchTemplate;
import com.pulumi.aws.ec2.LaunchTemplateArgs;
import com.pulumi.aws.autoscaling.Group;
import com.pulumi.aws.autoscaling.GroupArgs;
import com.pulumi.aws.autoscaling.inputs.GroupMixedInstancesPolicyArgs;
import com.pulumi.aws.autoscaling.inputs.GroupMixedInstancesPolicyLaunchTemplateArgs;
import com.pulumi.aws.autoscaling.inputs.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs;
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
        var exampleLaunchTemplate = new LaunchTemplate("exampleLaunchTemplate", LaunchTemplateArgs.builder()        
            .namePrefix("example")
            .imageId(data.aws_ami().example().id())
            .instanceType("c5.large")
            .build());

        var example2 = new LaunchTemplate("example2", LaunchTemplateArgs.builder()        
            .namePrefix("example2")
            .imageId(data.aws_ami().example2().id())
            .build());

        var exampleGroup = new Group("exampleGroup", GroupArgs.builder()        
            .availabilityZones("us-east-1a")
            .desiredCapacity(1)
            .maxSize(1)
            .minSize(1)
            .mixedInstancesPolicy(GroupMixedInstancesPolicyArgs.builder()
                .launchTemplate(GroupMixedInstancesPolicyLaunchTemplateArgs.builder()
                    .launchTemplateSpecification(GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs.builder()
                        .launchTemplateId(exampleLaunchTemplate.id())
                        .build())
                    .overrides(                    
                        GroupMixedInstancesPolicyLaunchTemplateOverrideArgs.builder()
                            .instanceType("c4.large")
                            .weightedCapacity("3")
                            .build(),
                        GroupMixedInstancesPolicyLaunchTemplateOverrideArgs.builder()
                            .instanceType("c6g.large")
                            .launchTemplateSpecification(GroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationArgs.builder()
                                .launchTemplateId(example2.id())
                                .build())
                            .weightedCapacity("2")
                            .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleLaunchTemplate:
    type: aws:ec2:LaunchTemplate
    properties:
      namePrefix: example
      imageId: ${data.aws_ami.example.id}
      instanceType: c5.large
  example2:
    type: aws:ec2:LaunchTemplate
    properties:
      namePrefix: example2
      imageId: ${data.aws_ami.example2.id}
  exampleGroup:
    type: aws:autoscaling:Group
    properties:
      availabilityZones:
        - us-east-1a
      desiredCapacity: 1
      maxSize: 1
      minSize: 1
      mixedInstancesPolicy:
        launchTemplate:
          launchTemplateSpecification:
            launchTemplateId: ${exampleLaunchTemplate.id}
          overrides:
            - instanceType: c4.large
              weightedCapacity: '3'
            - instanceType: c6g.large
              launchTemplateSpecification:
                launchTemplateId: ${example2.id}
              weightedCapacity: '2'
```
{{% /example %}}
{{% example %}}
### Mixed Instances Policy with Attribute-based Instance Type Selection

As an alternative to manually choosing instance types when creating a mixed instances group, you can specify a set of instance attributes that describe your compute requirements.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLaunchTemplate = new aws.ec2.LaunchTemplate("exampleLaunchTemplate", {
    namePrefix: "example",
    imageId: data.aws_ami.example.id,
    instanceType: "c5.large",
});
const exampleGroup = new aws.autoscaling.Group("exampleGroup", {
    availabilityZones: ["us-east-1a"],
    desiredCapacity: 1,
    maxSize: 1,
    minSize: 1,
    mixedInstancesPolicy: {
        launchTemplate: {
            launchTemplateSpecification: {
                launchTemplateId: exampleLaunchTemplate.id,
            },
            overrides: [{
                instanceRequirements: {
                    memoryMib: {
                        min: 1000,
                    },
                    vcpuCount: {
                        min: 4,
                    },
                },
            }],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_launch_template = aws.ec2.LaunchTemplate("exampleLaunchTemplate",
    name_prefix="example",
    image_id=data["aws_ami"]["example"]["id"],
    instance_type="c5.large")
example_group = aws.autoscaling.Group("exampleGroup",
    availability_zones=["us-east-1a"],
    desired_capacity=1,
    max_size=1,
    min_size=1,
    mixed_instances_policy=aws.autoscaling.GroupMixedInstancesPolicyArgs(
        launch_template=aws.autoscaling.GroupMixedInstancesPolicyLaunchTemplateArgs(
            launch_template_specification=aws.autoscaling.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs(
                launch_template_id=example_launch_template.id,
            ),
            overrides=[aws.autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs(
                instance_requirements=aws.autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsArgs(
                    memory_mib=aws.autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibArgs(
                        min=1000,
                    ),
                    vcpu_count=aws.autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsVcpuCountArgs(
                        min=4,
                    ),
                ),
            )],
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleLaunchTemplate = new Aws.Ec2.LaunchTemplate("exampleLaunchTemplate", new()
    {
        NamePrefix = "example",
        ImageId = data.Aws_ami.Example.Id,
        InstanceType = "c5.large",
    });

    var exampleGroup = new Aws.AutoScaling.Group("exampleGroup", new()
    {
        AvailabilityZones = new[]
        {
            "us-east-1a",
        },
        DesiredCapacity = 1,
        MaxSize = 1,
        MinSize = 1,
        MixedInstancesPolicy = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyArgs
        {
            LaunchTemplate = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateArgs
            {
                LaunchTemplateSpecification = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs
                {
                    LaunchTemplateId = exampleLaunchTemplate.Id,
                },
                Overrides = new[]
                {
                    new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs
                    {
                        InstanceRequirements = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsArgs
                        {
                            MemoryMib = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibArgs
                            {
                                Min = 1000,
                            },
                            VcpuCount = new Aws.AutoScaling.Inputs.GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsVcpuCountArgs
                            {
                                Min = 4,
                            },
                        },
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLaunchTemplate, err := ec2.NewLaunchTemplate(ctx, "exampleLaunchTemplate", &ec2.LaunchTemplateArgs{
			NamePrefix:   pulumi.String("example"),
			ImageId:      pulumi.Any(data.Aws_ami.Example.Id),
			InstanceType: pulumi.String("c5.large"),
		})
		if err != nil {
			return err
		}
		_, err = autoscaling.NewGroup(ctx, "exampleGroup", &autoscaling.GroupArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-east-1a"),
			},
			DesiredCapacity: pulumi.Int(1),
			MaxSize:         pulumi.Int(1),
			MinSize:         pulumi.Int(1),
			MixedInstancesPolicy: &autoscaling.GroupMixedInstancesPolicyArgs{
				LaunchTemplate: &autoscaling.GroupMixedInstancesPolicyLaunchTemplateArgs{
					LaunchTemplateSpecification: &autoscaling.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs{
						LaunchTemplateId: exampleLaunchTemplate.ID(),
					},
					Overrides: autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideArray{
						&autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs{
							InstanceRequirements: &autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsArgs{
								MemoryMib: &autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibArgs{
									Min: pulumi.Int(1000),
								},
								VcpuCount: &autoscaling.GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsVcpuCountArgs{
									Min: pulumi.Int(4),
								},
							},
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
import com.pulumi.aws.ec2.LaunchTemplate;
import com.pulumi.aws.ec2.LaunchTemplateArgs;
import com.pulumi.aws.autoscaling.Group;
import com.pulumi.aws.autoscaling.GroupArgs;
import com.pulumi.aws.autoscaling.inputs.GroupMixedInstancesPolicyArgs;
import com.pulumi.aws.autoscaling.inputs.GroupMixedInstancesPolicyLaunchTemplateArgs;
import com.pulumi.aws.autoscaling.inputs.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs;
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
        var exampleLaunchTemplate = new LaunchTemplate("exampleLaunchTemplate", LaunchTemplateArgs.builder()        
            .namePrefix("example")
            .imageId(data.aws_ami().example().id())
            .instanceType("c5.large")
            .build());

        var exampleGroup = new Group("exampleGroup", GroupArgs.builder()        
            .availabilityZones("us-east-1a")
            .desiredCapacity(1)
            .maxSize(1)
            .minSize(1)
            .mixedInstancesPolicy(GroupMixedInstancesPolicyArgs.builder()
                .launchTemplate(GroupMixedInstancesPolicyLaunchTemplateArgs.builder()
                    .launchTemplateSpecification(GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs.builder()
                        .launchTemplateId(exampleLaunchTemplate.id())
                        .build())
                    .overrides(GroupMixedInstancesPolicyLaunchTemplateOverrideArgs.builder()
                        .instanceRequirements(GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsArgs.builder()
                            .memoryMib(GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibArgs.builder()
                                .min(1000)
                                .build())
                            .vcpuCount(GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsVcpuCountArgs.builder()
                                .min(4)
                                .build())
                            .build())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleLaunchTemplate:
    type: aws:ec2:LaunchTemplate
    properties:
      namePrefix: example
      imageId: ${data.aws_ami.example.id}
      instanceType: c5.large
  exampleGroup:
    type: aws:autoscaling:Group
    properties:
      availabilityZones:
        - us-east-1a
      desiredCapacity: 1
      maxSize: 1
      minSize: 1
      mixedInstancesPolicy:
        launchTemplate:
          launchTemplateSpecification:
            launchTemplateId: ${exampleLaunchTemplate.id}
          overrides:
            - instanceRequirements:
                memoryMib:
                  min: 1000
                vcpuCount:
                  min: 4
```
{{% /example %}}
{{% example %}}
### Dynamic tagging

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.autoscaling.Group;
import com.pulumi.aws.autoscaling.GroupArgs;
import com.pulumi.aws.autoscaling.inputs.GroupTagArgs;
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
        final var config = ctx.config();
        final var extraTags = config.get("extraTags").orElse(        
            %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference),
            %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference));
        var test = new Group("test", GroupArgs.builder()        
            .maxSize(5)
            .minSize(2)
            .launchConfiguration(aws_launch_configuration.foobar().name())
            .vpcZoneIdentifiers(            
                aws_subnet.example1().id(),
                aws_subnet.example2().id())
            .tags(            
                GroupTagArgs.builder()
                    .key("explicit1")
                    .value("value1")
                    .propagateAtLaunch(true)
                    .build(),
                GroupTagArgs.builder()
                    .key("explicit2")
                    .value("value2")
                    .propagateAtLaunch(true)
                    .build())
            .dynamic(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
            .build());

    }
}
```
```yaml
configuration:
  extraTags:
    type: dynamic
    default:
      - key: Foo
        value: Bar
        propagate_at_launch: true
      - key: Baz
        value: Bam
        propagate_at_launch: true
resources:
  test:
    type: aws:autoscaling:Group
    properties:
      maxSize: 5
      minSize: 2
      launchConfiguration: ${aws_launch_configuration.foobar.name}
      vpcZoneIdentifiers:
        - ${aws_subnet.example1.id}
        - ${aws_subnet.example2.id}
      tags:
        - key: explicit1
          value: value1
          propagateAtLaunch: true
        - key: explicit2
          value: value2
          propagateAtLaunch: true
      dynamic:
        - forEach: ${extraTags}
          content:
            - key: ${tag.value.key}
              propagateAtLaunch: ${tag.value.propagate_at_launch}
              value: ${tag.value.value}
```
{{% /example %}}
{{% example %}}
### Automatically refresh all instances after the group is updated

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAmi = aws.ec2.getAmi({
    mostRecent: true,
    owners: ["amazon"],
    filters: [{
        name: "name",
        values: ["amzn-ami-hvm-*-x86_64-gp2"],
    }],
});
const exampleLaunchTemplate = new aws.ec2.LaunchTemplate("exampleLaunchTemplate", {
    imageId: exampleAmi.then(exampleAmi => exampleAmi.id),
    instanceType: "t3.nano",
});
const exampleGroup = new aws.autoscaling.Group("exampleGroup", {
    availabilityZones: ["us-east-1a"],
    desiredCapacity: 1,
    maxSize: 2,
    minSize: 1,
    launchTemplate: {
        id: exampleLaunchTemplate.id,
        version: exampleLaunchTemplate.latestVersion,
    },
    tags: [{
        key: "Key",
        value: "Value",
        propagateAtLaunch: true,
    }],
    instanceRefresh: {
        strategy: "Rolling",
        preferences: {
            minHealthyPercentage: 50,
        },
        triggers: ["tag"],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_ami = aws.ec2.get_ami(most_recent=True,
    owners=["amazon"],
    filters=[aws.ec2.GetAmiFilterArgs(
        name="name",
        values=["amzn-ami-hvm-*-x86_64-gp2"],
    )])
example_launch_template = aws.ec2.LaunchTemplate("exampleLaunchTemplate",
    image_id=example_ami.id,
    instance_type="t3.nano")
example_group = aws.autoscaling.Group("exampleGroup",
    availability_zones=["us-east-1a"],
    desired_capacity=1,
    max_size=2,
    min_size=1,
    launch_template=aws.autoscaling.GroupLaunchTemplateArgs(
        id=example_launch_template.id,
        version=example_launch_template.latest_version,
    ),
    tags=[aws.autoscaling.GroupTagArgs(
        key="Key",
        value="Value",
        propagate_at_launch=True,
    )],
    instance_refresh=aws.autoscaling.GroupInstanceRefreshArgs(
        strategy="Rolling",
        preferences=aws.autoscaling.GroupInstanceRefreshPreferencesArgs(
            min_healthy_percentage=50,
        ),
        triggers=["tag"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleAmi = Aws.Ec2.GetAmi.Invoke(new()
    {
        MostRecent = true,
        Owners = new[]
        {
            "amazon",
        },
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetAmiFilterInputArgs
            {
                Name = "name",
                Values = new[]
                {
                    "amzn-ami-hvm-*-x86_64-gp2",
                },
            },
        },
    });

    var exampleLaunchTemplate = new Aws.Ec2.LaunchTemplate("exampleLaunchTemplate", new()
    {
        ImageId = exampleAmi.Apply(getAmiResult => getAmiResult.Id),
        InstanceType = "t3.nano",
    });

    var exampleGroup = new Aws.AutoScaling.Group("exampleGroup", new()
    {
        AvailabilityZones = new[]
        {
            "us-east-1a",
        },
        DesiredCapacity = 1,
        MaxSize = 2,
        MinSize = 1,
        LaunchTemplate = new Aws.AutoScaling.Inputs.GroupLaunchTemplateArgs
        {
            Id = exampleLaunchTemplate.Id,
            Version = exampleLaunchTemplate.LatestVersion,
        },
        Tags = new[]
        {
            new Aws.AutoScaling.Inputs.GroupTagArgs
            {
                Key = "Key",
                Value = "Value",
                PropagateAtLaunch = true,
            },
        },
        InstanceRefresh = new Aws.AutoScaling.Inputs.GroupInstanceRefreshArgs
        {
            Strategy = "Rolling",
            Preferences = new Aws.AutoScaling.Inputs.GroupInstanceRefreshPreferencesArgs
            {
                MinHealthyPercentage = 50,
            },
            Triggers = new[]
            {
                "tag",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleAmi, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
			MostRecent: pulumi.BoolRef(true),
			Owners: []string{
				"amazon",
			},
			Filters: []ec2.GetAmiFilter{
				{
					Name: "name",
					Values: []string{
						"amzn-ami-hvm-*-x86_64-gp2",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		exampleLaunchTemplate, err := ec2.NewLaunchTemplate(ctx, "exampleLaunchTemplate", &ec2.LaunchTemplateArgs{
			ImageId:      *pulumi.String(exampleAmi.Id),
			InstanceType: pulumi.String("t3.nano"),
		})
		if err != nil {
			return err
		}
		_, err = autoscaling.NewGroup(ctx, "exampleGroup", &autoscaling.GroupArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-east-1a"),
			},
			DesiredCapacity: pulumi.Int(1),
			MaxSize:         pulumi.Int(2),
			MinSize:         pulumi.Int(1),
			LaunchTemplate: &autoscaling.GroupLaunchTemplateArgs{
				Id:      exampleLaunchTemplate.ID(),
				Version: exampleLaunchTemplate.LatestVersion,
			},
			Tags: autoscaling.GroupTagArray{
				&autoscaling.GroupTagArgs{
					Key:               pulumi.String("Key"),
					Value:             pulumi.String("Value"),
					PropagateAtLaunch: pulumi.Bool(true),
				},
			},
			InstanceRefresh: &autoscaling.GroupInstanceRefreshArgs{
				Strategy: pulumi.String("Rolling"),
				Preferences: &autoscaling.GroupInstanceRefreshPreferencesArgs{
					MinHealthyPercentage: pulumi.Int(50),
				},
				Triggers: pulumi.StringArray{
					pulumi.String("tag"),
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetAmiArgs;
import com.pulumi.aws.ec2.LaunchTemplate;
import com.pulumi.aws.ec2.LaunchTemplateArgs;
import com.pulumi.aws.autoscaling.Group;
import com.pulumi.aws.autoscaling.GroupArgs;
import com.pulumi.aws.autoscaling.inputs.GroupLaunchTemplateArgs;
import com.pulumi.aws.autoscaling.inputs.GroupTagArgs;
import com.pulumi.aws.autoscaling.inputs.GroupInstanceRefreshArgs;
import com.pulumi.aws.autoscaling.inputs.GroupInstanceRefreshPreferencesArgs;
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
        final var exampleAmi = Ec2Functions.getAmi(GetAmiArgs.builder()
            .mostRecent(true)
            .owners("amazon")
            .filters(GetAmiFilterArgs.builder()
                .name("name")
                .values("amzn-ami-hvm-*-x86_64-gp2")
                .build())
            .build());

        var exampleLaunchTemplate = new LaunchTemplate("exampleLaunchTemplate", LaunchTemplateArgs.builder()        
            .imageId(exampleAmi.applyValue(getAmiResult -> getAmiResult.id()))
            .instanceType("t3.nano")
            .build());

        var exampleGroup = new Group("exampleGroup", GroupArgs.builder()        
            .availabilityZones("us-east-1a")
            .desiredCapacity(1)
            .maxSize(2)
            .minSize(1)
            .launchTemplate(GroupLaunchTemplateArgs.builder()
                .id(exampleLaunchTemplate.id())
                .version(exampleLaunchTemplate.latestVersion())
                .build())
            .tags(GroupTagArgs.builder()
                .key("Key")
                .value("Value")
                .propagateAtLaunch(true)
                .build())
            .instanceRefresh(GroupInstanceRefreshArgs.builder()
                .strategy("Rolling")
                .preferences(GroupInstanceRefreshPreferencesArgs.builder()
                    .minHealthyPercentage(50)
                    .build())
                .triggers("tag")
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleGroup:
    type: aws:autoscaling:Group
    properties:
      availabilityZones:
        - us-east-1a
      desiredCapacity: 1
      maxSize: 2
      minSize: 1
      launchTemplate:
        id: ${exampleLaunchTemplate.id}
        version: ${exampleLaunchTemplate.latestVersion}
      tags:
        - key: Key
          value: Value
          propagateAtLaunch: true
      instanceRefresh:
        strategy: Rolling
        preferences:
          minHealthyPercentage: 50
        triggers:
          - tag
  exampleLaunchTemplate:
    type: aws:ec2:LaunchTemplate
    properties:
      imageId: ${exampleAmi.id}
      instanceType: t3.nano
variables:
  exampleAmi:
    fn::invoke:
      Function: aws:ec2:getAmi
      Arguments:
        mostRecent: true
        owners:
          - amazon
        filters:
          - name: name
            values:
              - amzn-ami-hvm-*-x86_64-gp2
```
{{% /example %}}
{{% example %}}
### Auto Scaling group with Warm Pool

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLaunchTemplate = new aws.ec2.LaunchTemplate("exampleLaunchTemplate", {
    namePrefix: "example",
    imageId: data.aws_ami.example.id,
    instanceType: "c5.large",
});
const exampleGroup = new aws.autoscaling.Group("exampleGroup", {
    availabilityZones: ["us-east-1a"],
    desiredCapacity: 1,
    maxSize: 5,
    minSize: 1,
    warmPool: {
        poolState: "Hibernated",
        minSize: 1,
        maxGroupPreparedCapacity: 10,
        instanceReusePolicy: {
            reuseOnScaleIn: true,
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_launch_template = aws.ec2.LaunchTemplate("exampleLaunchTemplate",
    name_prefix="example",
    image_id=data["aws_ami"]["example"]["id"],
    instance_type="c5.large")
example_group = aws.autoscaling.Group("exampleGroup",
    availability_zones=["us-east-1a"],
    desired_capacity=1,
    max_size=5,
    min_size=1,
    warm_pool=aws.autoscaling.GroupWarmPoolArgs(
        pool_state="Hibernated",
        min_size=1,
        max_group_prepared_capacity=10,
        instance_reuse_policy=aws.autoscaling.GroupWarmPoolInstanceReusePolicyArgs(
            reuse_on_scale_in=True,
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleLaunchTemplate = new Aws.Ec2.LaunchTemplate("exampleLaunchTemplate", new()
    {
        NamePrefix = "example",
        ImageId = data.Aws_ami.Example.Id,
        InstanceType = "c5.large",
    });

    var exampleGroup = new Aws.AutoScaling.Group("exampleGroup", new()
    {
        AvailabilityZones = new[]
        {
            "us-east-1a",
        },
        DesiredCapacity = 1,
        MaxSize = 5,
        MinSize = 1,
        WarmPool = new Aws.AutoScaling.Inputs.GroupWarmPoolArgs
        {
            PoolState = "Hibernated",
            MinSize = 1,
            MaxGroupPreparedCapacity = 10,
            InstanceReusePolicy = new Aws.AutoScaling.Inputs.GroupWarmPoolInstanceReusePolicyArgs
            {
                ReuseOnScaleIn = true,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewLaunchTemplate(ctx, "exampleLaunchTemplate", &ec2.LaunchTemplateArgs{
			NamePrefix:   pulumi.String("example"),
			ImageId:      pulumi.Any(data.Aws_ami.Example.Id),
			InstanceType: pulumi.String("c5.large"),
		})
		if err != nil {
			return err
		}
		_, err = autoscaling.NewGroup(ctx, "exampleGroup", &autoscaling.GroupArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-east-1a"),
			},
			DesiredCapacity: pulumi.Int(1),
			MaxSize:         pulumi.Int(5),
			MinSize:         pulumi.Int(1),
			WarmPool: &autoscaling.GroupWarmPoolArgs{
				PoolState:                pulumi.String("Hibernated"),
				MinSize:                  pulumi.Int(1),
				MaxGroupPreparedCapacity: pulumi.Int(10),
				InstanceReusePolicy: &autoscaling.GroupWarmPoolInstanceReusePolicyArgs{
					ReuseOnScaleIn: pulumi.Bool(true),
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
import com.pulumi.aws.ec2.LaunchTemplate;
import com.pulumi.aws.ec2.LaunchTemplateArgs;
import com.pulumi.aws.autoscaling.Group;
import com.pulumi.aws.autoscaling.GroupArgs;
import com.pulumi.aws.autoscaling.inputs.GroupWarmPoolArgs;
import com.pulumi.aws.autoscaling.inputs.GroupWarmPoolInstanceReusePolicyArgs;
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
        var exampleLaunchTemplate = new LaunchTemplate("exampleLaunchTemplate", LaunchTemplateArgs.builder()        
            .namePrefix("example")
            .imageId(data.aws_ami().example().id())
            .instanceType("c5.large")
            .build());

        var exampleGroup = new Group("exampleGroup", GroupArgs.builder()        
            .availabilityZones("us-east-1a")
            .desiredCapacity(1)
            .maxSize(5)
            .minSize(1)
            .warmPool(GroupWarmPoolArgs.builder()
                .poolState("Hibernated")
                .minSize(1)
                .maxGroupPreparedCapacity(10)
                .instanceReusePolicy(GroupWarmPoolInstanceReusePolicyArgs.builder()
                    .reuseOnScaleIn(true)
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleLaunchTemplate:
    type: aws:ec2:LaunchTemplate
    properties:
      namePrefix: example
      imageId: ${data.aws_ami.example.id}
      instanceType: c5.large
  exampleGroup:
    type: aws:autoscaling:Group
    properties:
      availabilityZones:
        - us-east-1a
      desiredCapacity: 1
      maxSize: 5
      minSize: 1
      warmPool:
        poolState: Hibernated
        minSize: 1
        maxGroupPreparedCapacity: 10
        instanceReusePolicy:
          reuseOnScaleIn: true
```
{{% /example %}}
{{% example %}}
### Auto Scaling group with Traffic Sources

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.autoscaling.Group;
import com.pulumi.aws.autoscaling.GroupArgs;
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
        var test = new Group("test", GroupArgs.builder()        
            .vpcZoneIdentifiers(aws_subnet.test().id())
            .maxSize(1)
            .minSize(1)
            .forceDelete(true)
            .dynamic(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
            .build());

    }
}
```
{{% /example %}}
{{% /examples %}}
## Waiting for Capacity

A newly-created ASG is initially empty and begins to scale to `min_size` (or
`desired_capacity`, if specified) by launching instances using the provided
Launch Configuration. These instances take time to launch and boot.

On ASG Update, changes to these values also take time to result in the target
number of instances providing service.

This provider provides two mechanisms to help consistently manage ASG scale up
time across dependent resources.

#### Waiting for ASG Capacity

The first is default behavior. This provider waits after ASG creation for
`min_size` (or `desired_capacity`, if specified) healthy instances to show up
in the ASG before continuing.

If `min_size` or `desired_capacity` are changed in a subsequent update,
this provider will also wait for the correct number of healthy instances before
continuing.

This provider considers an instance "healthy" when the ASG reports `HealthStatus:
"Healthy"` and `LifecycleState: "InService"`. See the [AWS AutoScaling
Docs](https://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/AutoScalingGroupLifecycle.html)
for more information on an ASG's lifecycle.

This provider will wait for healthy instances for up to
`wait_for_capacity_timeout`. If ASG creation is taking more than a few minutes,
it's worth investigating for scaling activity errors, which can be caused by
problems with the selected Launch Configuration.

Setting `wait_for_capacity_timeout` to `"0"` disables ASG Capacity waiting.

#### Waiting for ELB Capacity

The second mechanism is optional, and affects ASGs with attached ELBs specified
via the `load_balancers` attribute or with ALBs specified with `target_group_arns`.

The `min_elb_capacity` parameter causes the provider to wait for at least the
requested number of instances to show up `"InService"` in all attached ELBs
during ASG creation. It has no effect on ASG updates.

If `wait_for_elb_capacity` is set, the provider will wait for exactly that number
of Instances to be `"InService"` in all attached ELBs on both creation and
updates.

These parameters can be used to ensure that service is being provided before
the provider moves on. If new instances don't pass the ELB's health checks for any
reason, the apply will time out, and the ASG will be marked as
tainted (i.e., marked to be destroyed in a follow up run).

As with ASG Capacity, the provider will wait for up to `wait_for_capacity_timeout`
for the proper number of instances to be healthy.

#### Troubleshooting Capacity Waiting Timeouts

If ASG creation takes more than a few minutes, this could indicate one of a
number of configuration problems. See the [AWS Docs on Load Balancer
Troubleshooting](https://docs.aws.amazon.com/ElasticLoadBalancing/latest/DeveloperGuide/elb-troubleshooting.html)
for more information.


## Import

Using `pulumi import`, import Auto Scaling Groups using the `name`. For example:

```sh
 $ pulumi import aws:autoscaling/group:Group web web-asg
```
 