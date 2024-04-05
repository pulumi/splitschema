Provides a Lightsail Instance. Amazon Lightsail is a service to provide easy virtual private servers
with custom software already setup. See [What is Amazon Lightsail?](https://lightsail.aws.amazon.com/ls/docs/getting-started/article/what-is-amazon-lightsail)
for more information.

> **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new GitLab Lightsail Instance
const gitlabTest = new aws.lightsail.Instance("gitlabTest", {
    availabilityZone: "us-east-1b",
    blueprintId: "amazon_linux_2",
    bundleId: "nano_1_0",
    keyPairName: "some_key_name",
    tags: {
        foo: "bar",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

# Create a new GitLab Lightsail Instance
gitlab_test = aws.lightsail.Instance("gitlabTest",
    availability_zone="us-east-1b",
    blueprint_id="amazon_linux_2",
    bundle_id="nano_1_0",
    key_pair_name="some_key_name",
    tags={
        "foo": "bar",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // Create a new GitLab Lightsail Instance
    var gitlabTest = new Aws.LightSail.Instance("gitlabTest", new()
    {
        AvailabilityZone = "us-east-1b",
        BlueprintId = "amazon_linux_2",
        BundleId = "nano_1_0",
        KeyPairName = "some_key_name",
        Tags = 
        {
            { "foo", "bar" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lightsail.NewInstance(ctx, "gitlabTest", &lightsail.InstanceArgs{
			AvailabilityZone: pulumi.String("us-east-1b"),
			BlueprintId:      pulumi.String("amazon_linux_2"),
			BundleId:         pulumi.String("nano_1_0"),
			KeyPairName:      pulumi.String("some_key_name"),
			Tags: pulumi.StringMap{
				"foo": pulumi.String("bar"),
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
import com.pulumi.aws.lightsail.Instance;
import com.pulumi.aws.lightsail.InstanceArgs;
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
        var gitlabTest = new Instance("gitlabTest", InstanceArgs.builder()        
            .availabilityZone("us-east-1b")
            .blueprintId("amazon_linux_2")
            .bundleId("nano_1_0")
            .keyPairName("some_key_name")
            .tags(Map.of("foo", "bar"))
            .build());

    }
}
```
```yaml
resources:
  # Create a new GitLab Lightsail Instance
  gitlabTest:
    type: aws:lightsail:Instance
    properties:
      availabilityZone: us-east-1b
      blueprintId: amazon_linux_2
      bundleId: nano_1_0
      keyPairName: some_key_name
      tags:
        foo: bar
```
{{% /example %}}
{{% example %}}
### Example With User Data

Lightsail user data is handled differently than ec2 user data. Lightsail user data only accepts a single lined string. The below example shows installing apache and creating the index page.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const custom = new aws.lightsail.Instance("custom", {
    availabilityZone: "us-east-1b",
    blueprintId: "amazon_linux_2",
    bundleId: "nano_1_0",
    userData: "sudo yum install -y httpd && sudo systemctl start httpd && sudo systemctl enable httpd && echo '<h1>Deployed via Pulumi</h1>' | sudo tee /var/www/html/index.html",
});
```
```python
import pulumi
import pulumi_aws as aws

custom = aws.lightsail.Instance("custom",
    availability_zone="us-east-1b",
    blueprint_id="amazon_linux_2",
    bundle_id="nano_1_0",
    user_data="sudo yum install -y httpd && sudo systemctl start httpd && sudo systemctl enable httpd && echo '<h1>Deployed via Pulumi</h1>' | sudo tee /var/www/html/index.html")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var custom = new Aws.LightSail.Instance("custom", new()
    {
        AvailabilityZone = "us-east-1b",
        BlueprintId = "amazon_linux_2",
        BundleId = "nano_1_0",
        UserData = "sudo yum install -y httpd && sudo systemctl start httpd && sudo systemctl enable httpd && echo '<h1>Deployed via Pulumi</h1>' | sudo tee /var/www/html/index.html",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lightsail.NewInstance(ctx, "custom", &lightsail.InstanceArgs{
			AvailabilityZone: pulumi.String("us-east-1b"),
			BlueprintId:      pulumi.String("amazon_linux_2"),
			BundleId:         pulumi.String("nano_1_0"),
			UserData:         pulumi.String("sudo yum install -y httpd && sudo systemctl start httpd && sudo systemctl enable httpd && echo '<h1>Deployed via Pulumi</h1>' | sudo tee /var/www/html/index.html"),
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
import com.pulumi.aws.lightsail.Instance;
import com.pulumi.aws.lightsail.InstanceArgs;
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
        var custom = new Instance("custom", InstanceArgs.builder()        
            .availabilityZone("us-east-1b")
            .blueprintId("amazon_linux_2")
            .bundleId("nano_1_0")
            .userData("sudo yum install -y httpd && sudo systemctl start httpd && sudo systemctl enable httpd && echo '<h1>Deployed via Pulumi</h1>' | sudo tee /var/www/html/index.html")
            .build());

    }
}
```
```yaml
resources:
  custom:
    type: aws:lightsail:Instance
    properties:
      availabilityZone: us-east-1b
      blueprintId: amazon_linux_2
      bundleId: nano_1_0
      userData: sudo yum install -y httpd && sudo systemctl start httpd && sudo systemctl enable httpd && echo '<h1>Deployed via Pulumi</h1>' | sudo tee /var/www/html/index.html
```
{{% /example %}}
{{% example %}}
### Enable Auto Snapshots

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.lightsail.Instance("test", {
    addOn: {
        snapshotTime: "06:00",
        status: "Enabled",
        type: "AutoSnapshot",
    },
    availabilityZone: "us-east-1b",
    blueprintId: "amazon_linux_2",
    bundleId: "nano_1_0",
    tags: {
        foo: "bar",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.lightsail.Instance("test",
    add_on=aws.lightsail.InstanceAddOnArgs(
        snapshot_time="06:00",
        status="Enabled",
        type="AutoSnapshot",
    ),
    availability_zone="us-east-1b",
    blueprint_id="amazon_linux_2",
    bundle_id="nano_1_0",
    tags={
        "foo": "bar",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.LightSail.Instance("test", new()
    {
        AddOn = new Aws.LightSail.Inputs.InstanceAddOnArgs
        {
            SnapshotTime = "06:00",
            Status = "Enabled",
            Type = "AutoSnapshot",
        },
        AvailabilityZone = "us-east-1b",
        BlueprintId = "amazon_linux_2",
        BundleId = "nano_1_0",
        Tags = 
        {
            { "foo", "bar" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lightsail.NewInstance(ctx, "test", &lightsail.InstanceArgs{
			AddOn: &lightsail.InstanceAddOnArgs{
				SnapshotTime: pulumi.String("06:00"),
				Status:       pulumi.String("Enabled"),
				Type:         pulumi.String("AutoSnapshot"),
			},
			AvailabilityZone: pulumi.String("us-east-1b"),
			BlueprintId:      pulumi.String("amazon_linux_2"),
			BundleId:         pulumi.String("nano_1_0"),
			Tags: pulumi.StringMap{
				"foo": pulumi.String("bar"),
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
import com.pulumi.aws.lightsail.Instance;
import com.pulumi.aws.lightsail.InstanceArgs;
import com.pulumi.aws.lightsail.inputs.InstanceAddOnArgs;
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
        var test = new Instance("test", InstanceArgs.builder()        
            .addOn(InstanceAddOnArgs.builder()
                .snapshotTime("06:00")
                .status("Enabled")
                .type("AutoSnapshot")
                .build())
            .availabilityZone("us-east-1b")
            .blueprintId("amazon_linux_2")
            .bundleId("nano_1_0")
            .tags(Map.of("foo", "bar"))
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:lightsail:Instance
    properties:
      addOn:
        snapshotTime: 06:00
        status: Enabled
        type: AutoSnapshot
      availabilityZone: us-east-1b
      blueprintId: amazon_linux_2
      bundleId: nano_1_0
      tags:
        foo: bar
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Lightsail Instances using their name. For example:

```sh
 $ pulumi import aws:lightsail/instance:Instance gitlab_test 'custom_gitlab'
```
 