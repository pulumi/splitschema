Provides an Amazon MQ broker resource. This resources also manages users for the broker.

> For more information on Amazon MQ, see [Amazon MQ documentation](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/welcome.html).

> **NOTE:** Amazon MQ currently places limits on **RabbitMQ** brokers. For example, a RabbitMQ broker cannot have: instances with an associated IP address of an ENI attached to the broker, an associated LDAP server to authenticate and authorize broker connections, storage type `EFS`, or audit logging. Although this resource allows you to create RabbitMQ users, RabbitMQ users cannot have console access or groups. Also, Amazon MQ does not return information about RabbitMQ users so drift detection is not possible.

> **NOTE:** Changes to an MQ Broker can occur when you change a parameter, such as `configuration` or `user`, and are reflected in the next maintenance window. Because of this, the provider may report a difference in its planning phase because a modification has not yet taken place. You can use the `apply_immediately` flag to instruct the service to apply the change immediately (see documentation below). Using `apply_immediately` can result in a brief downtime as the broker reboots.


{{% examples %}}
## Example Usage
{{% example %}}
### Basic Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.mq.Broker("example", {
    configuration: {
        id: aws_mq_configuration.test.id,
        revision: aws_mq_configuration.test.latest_revision,
    },
    engineType: "ActiveMQ",
    engineVersion: "5.17.6",
    hostInstanceType: "mq.t2.micro",
    securityGroups: [aws_security_group.test.id],
    users: [{
        username: "ExampleUser",
        password: "MindTheGap",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.mq.Broker("example",
    configuration=aws.mq.BrokerConfigurationArgs(
        id=aws_mq_configuration["test"]["id"],
        revision=aws_mq_configuration["test"]["latest_revision"],
    ),
    engine_type="ActiveMQ",
    engine_version="5.17.6",
    host_instance_type="mq.t2.micro",
    security_groups=[aws_security_group["test"]["id"]],
    users=[aws.mq.BrokerUserArgs(
        username="ExampleUser",
        password="MindTheGap",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Mq.Broker("example", new()
    {
        Configuration = new Aws.Mq.Inputs.BrokerConfigurationArgs
        {
            Id = aws_mq_configuration.Test.Id,
            Revision = aws_mq_configuration.Test.Latest_revision,
        },
        EngineType = "ActiveMQ",
        EngineVersion = "5.17.6",
        HostInstanceType = "mq.t2.micro",
        SecurityGroups = new[]
        {
            aws_security_group.Test.Id,
        },
        Users = new[]
        {
            new Aws.Mq.Inputs.BrokerUserArgs
            {
                Username = "ExampleUser",
                Password = "MindTheGap",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/mq"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mq.NewBroker(ctx, "example", &mq.BrokerArgs{
			Configuration: &mq.BrokerConfigurationArgs{
				Id:       pulumi.Any(aws_mq_configuration.Test.Id),
				Revision: pulumi.Any(aws_mq_configuration.Test.Latest_revision),
			},
			EngineType:       pulumi.String("ActiveMQ"),
			EngineVersion:    pulumi.String("5.17.6"),
			HostInstanceType: pulumi.String("mq.t2.micro"),
			SecurityGroups: pulumi.StringArray{
				aws_security_group.Test.Id,
			},
			Users: mq.BrokerUserArray{
				&mq.BrokerUserArgs{
					Username: pulumi.String("ExampleUser"),
					Password: pulumi.String("MindTheGap"),
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
import com.pulumi.aws.mq.Broker;
import com.pulumi.aws.mq.BrokerArgs;
import com.pulumi.aws.mq.inputs.BrokerConfigurationArgs;
import com.pulumi.aws.mq.inputs.BrokerUserArgs;
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
        var example = new Broker("example", BrokerArgs.builder()        
            .configuration(BrokerConfigurationArgs.builder()
                .id(aws_mq_configuration.test().id())
                .revision(aws_mq_configuration.test().latest_revision())
                .build())
            .engineType("ActiveMQ")
            .engineVersion("5.17.6")
            .hostInstanceType("mq.t2.micro")
            .securityGroups(aws_security_group.test().id())
            .users(BrokerUserArgs.builder()
                .username("ExampleUser")
                .password("MindTheGap")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:mq:Broker
    properties:
      configuration:
        id: ${aws_mq_configuration.test.id}
        revision: ${aws_mq_configuration.test.latest_revision}
      engineType: ActiveMQ
      engineVersion: 5.17.6
      hostInstanceType: mq.t2.micro
      securityGroups:
        - ${aws_security_group.test.id}
      users:
        - username: ExampleUser
          password: MindTheGap
```
{{% /example %}}
{{% example %}}
### High-throughput Optimized Example

This example shows the use of EBS storage for high-throughput optimized performance.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.mq.Broker("example", {
    configuration: {
        id: aws_mq_configuration.test.id,
        revision: aws_mq_configuration.test.latest_revision,
    },
    engineType: "ActiveMQ",
    engineVersion: "5.17.6",
    storageType: "ebs",
    hostInstanceType: "mq.m5.large",
    securityGroups: [aws_security_group.test.id],
    users: [{
        username: "ExampleUser",
        password: "MindTheGap",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.mq.Broker("example",
    configuration=aws.mq.BrokerConfigurationArgs(
        id=aws_mq_configuration["test"]["id"],
        revision=aws_mq_configuration["test"]["latest_revision"],
    ),
    engine_type="ActiveMQ",
    engine_version="5.17.6",
    storage_type="ebs",
    host_instance_type="mq.m5.large",
    security_groups=[aws_security_group["test"]["id"]],
    users=[aws.mq.BrokerUserArgs(
        username="ExampleUser",
        password="MindTheGap",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Mq.Broker("example", new()
    {
        Configuration = new Aws.Mq.Inputs.BrokerConfigurationArgs
        {
            Id = aws_mq_configuration.Test.Id,
            Revision = aws_mq_configuration.Test.Latest_revision,
        },
        EngineType = "ActiveMQ",
        EngineVersion = "5.17.6",
        StorageType = "ebs",
        HostInstanceType = "mq.m5.large",
        SecurityGroups = new[]
        {
            aws_security_group.Test.Id,
        },
        Users = new[]
        {
            new Aws.Mq.Inputs.BrokerUserArgs
            {
                Username = "ExampleUser",
                Password = "MindTheGap",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/mq"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mq.NewBroker(ctx, "example", &mq.BrokerArgs{
			Configuration: &mq.BrokerConfigurationArgs{
				Id:       pulumi.Any(aws_mq_configuration.Test.Id),
				Revision: pulumi.Any(aws_mq_configuration.Test.Latest_revision),
			},
			EngineType:       pulumi.String("ActiveMQ"),
			EngineVersion:    pulumi.String("5.17.6"),
			StorageType:      pulumi.String("ebs"),
			HostInstanceType: pulumi.String("mq.m5.large"),
			SecurityGroups: pulumi.StringArray{
				aws_security_group.Test.Id,
			},
			Users: mq.BrokerUserArray{
				&mq.BrokerUserArgs{
					Username: pulumi.String("ExampleUser"),
					Password: pulumi.String("MindTheGap"),
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
import com.pulumi.aws.mq.Broker;
import com.pulumi.aws.mq.BrokerArgs;
import com.pulumi.aws.mq.inputs.BrokerConfigurationArgs;
import com.pulumi.aws.mq.inputs.BrokerUserArgs;
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
        var example = new Broker("example", BrokerArgs.builder()        
            .configuration(BrokerConfigurationArgs.builder()
                .id(aws_mq_configuration.test().id())
                .revision(aws_mq_configuration.test().latest_revision())
                .build())
            .engineType("ActiveMQ")
            .engineVersion("5.17.6")
            .storageType("ebs")
            .hostInstanceType("mq.m5.large")
            .securityGroups(aws_security_group.test().id())
            .users(BrokerUserArgs.builder()
                .username("ExampleUser")
                .password("MindTheGap")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:mq:Broker
    properties:
      configuration:
        id: ${aws_mq_configuration.test.id}
        revision: ${aws_mq_configuration.test.latest_revision}
      engineType: ActiveMQ
      engineVersion: 5.17.6
      storageType: ebs
      hostInstanceType: mq.m5.large
      securityGroups:
        - ${aws_security_group.test.id}
      users:
        - username: ExampleUser
          password: MindTheGap
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import MQ Brokers using their broker id. For example:

```sh
 $ pulumi import aws:mq/broker:Broker example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
```
 