Provides a Simple or Managed Microsoft directory in AWS Directory Service.

{{% examples %}}
## Example Usage
{{% example %}}
### SimpleAD

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ec2.Vpc("main", {cidrBlock: "10.0.0.0/16"});
const foo = new aws.ec2.Subnet("foo", {
    vpcId: main.id,
    availabilityZone: "us-west-2a",
    cidrBlock: "10.0.1.0/24",
});
const barSubnet = new aws.ec2.Subnet("barSubnet", {
    vpcId: main.id,
    availabilityZone: "us-west-2b",
    cidrBlock: "10.0.2.0/24",
});
const barDirectory = new aws.directoryservice.Directory("barDirectory", {
    name: "corp.notexample.com",
    password: "SuperSecretPassw0rd",
    size: "Small",
    vpcSettings: {
        vpcId: main.id,
        subnetIds: [
            foo.id,
            barSubnet.id,
        ],
    },
    tags: {
        Project: "foo",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
foo = aws.ec2.Subnet("foo",
    vpc_id=main.id,
    availability_zone="us-west-2a",
    cidr_block="10.0.1.0/24")
bar_subnet = aws.ec2.Subnet("barSubnet",
    vpc_id=main.id,
    availability_zone="us-west-2b",
    cidr_block="10.0.2.0/24")
bar_directory = aws.directoryservice.Directory("barDirectory",
    name="corp.notexample.com",
    password="SuperSecretPassw0rd",
    size="Small",
    vpc_settings=aws.directoryservice.DirectoryVpcSettingsArgs(
        vpc_id=main.id,
        subnet_ids=[
            foo.id,
            bar_subnet.id,
        ],
    ),
    tags={
        "Project": "foo",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = new Aws.Ec2.Vpc("main", new()
    {
        CidrBlock = "10.0.0.0/16",
    });

    var foo = new Aws.Ec2.Subnet("foo", new()
    {
        VpcId = main.Id,
        AvailabilityZone = "us-west-2a",
        CidrBlock = "10.0.1.0/24",
    });

    var barSubnet = new Aws.Ec2.Subnet("barSubnet", new()
    {
        VpcId = main.Id,
        AvailabilityZone = "us-west-2b",
        CidrBlock = "10.0.2.0/24",
    });

    var barDirectory = new Aws.DirectoryService.Directory("barDirectory", new()
    {
        Name = "corp.notexample.com",
        Password = "SuperSecretPassw0rd",
        Size = "Small",
        VpcSettings = new Aws.DirectoryService.Inputs.DirectoryVpcSettingsArgs
        {
            VpcId = main.Id,
            SubnetIds = new[]
            {
                foo.Id,
                barSubnet.Id,
            },
        },
        Tags = 
        {
            { "Project", "foo" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directoryservice"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		foo, err := ec2.NewSubnet(ctx, "foo", &ec2.SubnetArgs{
			VpcId:            main.ID(),
			AvailabilityZone: pulumi.String("us-west-2a"),
			CidrBlock:        pulumi.String("10.0.1.0/24"),
		})
		if err != nil {
			return err
		}
		barSubnet, err := ec2.NewSubnet(ctx, "barSubnet", &ec2.SubnetArgs{
			VpcId:            main.ID(),
			AvailabilityZone: pulumi.String("us-west-2b"),
			CidrBlock:        pulumi.String("10.0.2.0/24"),
		})
		if err != nil {
			return err
		}
		_, err = directoryservice.NewDirectory(ctx, "barDirectory", &directoryservice.DirectoryArgs{
			Name:     pulumi.String("corp.notexample.com"),
			Password: pulumi.String("SuperSecretPassw0rd"),
			Size:     pulumi.String("Small"),
			VpcSettings: &directoryservice.DirectoryVpcSettingsArgs{
				VpcId: main.ID(),
				SubnetIds: pulumi.StringArray{
					foo.ID(),
					barSubnet.ID(),
				},
			},
			Tags: pulumi.StringMap{
				"Project": pulumi.String("foo"),
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.ec2.Subnet;
import com.pulumi.aws.ec2.SubnetArgs;
import com.pulumi.aws.directoryservice.Directory;
import com.pulumi.aws.directoryservice.DirectoryArgs;
import com.pulumi.aws.directoryservice.inputs.DirectoryVpcSettingsArgs;
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
        var main = new Vpc("main", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .build());

        var foo = new Subnet("foo", SubnetArgs.builder()        
            .vpcId(main.id())
            .availabilityZone("us-west-2a")
            .cidrBlock("10.0.1.0/24")
            .build());

        var barSubnet = new Subnet("barSubnet", SubnetArgs.builder()        
            .vpcId(main.id())
            .availabilityZone("us-west-2b")
            .cidrBlock("10.0.2.0/24")
            .build());

        var barDirectory = new Directory("barDirectory", DirectoryArgs.builder()        
            .name("corp.notexample.com")
            .password("SuperSecretPassw0rd")
            .size("Small")
            .vpcSettings(DirectoryVpcSettingsArgs.builder()
                .vpcId(main.id())
                .subnetIds(                
                    foo.id(),
                    barSubnet.id())
                .build())
            .tags(Map.of("Project", "foo"))
            .build());

    }
}
```
```yaml
resources:
  barDirectory:
    type: aws:directoryservice:Directory
    properties:
      name: corp.notexample.com
      password: SuperSecretPassw0rd
      size: Small
      vpcSettings:
        vpcId: ${main.id}
        subnetIds:
          - ${foo.id}
          - ${barSubnet.id}
      tags:
        Project: foo
  main:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
  foo:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${main.id}
      availabilityZone: us-west-2a
      cidrBlock: 10.0.1.0/24
  barSubnet:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${main.id}
      availabilityZone: us-west-2b
      cidrBlock: 10.0.2.0/24
```
{{% /example %}}
{{% example %}}
### Microsoft Active Directory (MicrosoftAD)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ec2.Vpc("main", {cidrBlock: "10.0.0.0/16"});
const foo = new aws.ec2.Subnet("foo", {
    vpcId: main.id,
    availabilityZone: "us-west-2a",
    cidrBlock: "10.0.1.0/24",
});
const barSubnet = new aws.ec2.Subnet("barSubnet", {
    vpcId: main.id,
    availabilityZone: "us-west-2b",
    cidrBlock: "10.0.2.0/24",
});
const barDirectory = new aws.directoryservice.Directory("barDirectory", {
    name: "corp.notexample.com",
    password: "SuperSecretPassw0rd",
    edition: "Standard",
    type: "MicrosoftAD",
    vpcSettings: {
        vpcId: main.id,
        subnetIds: [
            foo.id,
            barSubnet.id,
        ],
    },
    tags: {
        Project: "foo",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
foo = aws.ec2.Subnet("foo",
    vpc_id=main.id,
    availability_zone="us-west-2a",
    cidr_block="10.0.1.0/24")
bar_subnet = aws.ec2.Subnet("barSubnet",
    vpc_id=main.id,
    availability_zone="us-west-2b",
    cidr_block="10.0.2.0/24")
bar_directory = aws.directoryservice.Directory("barDirectory",
    name="corp.notexample.com",
    password="SuperSecretPassw0rd",
    edition="Standard",
    type="MicrosoftAD",
    vpc_settings=aws.directoryservice.DirectoryVpcSettingsArgs(
        vpc_id=main.id,
        subnet_ids=[
            foo.id,
            bar_subnet.id,
        ],
    ),
    tags={
        "Project": "foo",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = new Aws.Ec2.Vpc("main", new()
    {
        CidrBlock = "10.0.0.0/16",
    });

    var foo = new Aws.Ec2.Subnet("foo", new()
    {
        VpcId = main.Id,
        AvailabilityZone = "us-west-2a",
        CidrBlock = "10.0.1.0/24",
    });

    var barSubnet = new Aws.Ec2.Subnet("barSubnet", new()
    {
        VpcId = main.Id,
        AvailabilityZone = "us-west-2b",
        CidrBlock = "10.0.2.0/24",
    });

    var barDirectory = new Aws.DirectoryService.Directory("barDirectory", new()
    {
        Name = "corp.notexample.com",
        Password = "SuperSecretPassw0rd",
        Edition = "Standard",
        Type = "MicrosoftAD",
        VpcSettings = new Aws.DirectoryService.Inputs.DirectoryVpcSettingsArgs
        {
            VpcId = main.Id,
            SubnetIds = new[]
            {
                foo.Id,
                barSubnet.Id,
            },
        },
        Tags = 
        {
            { "Project", "foo" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directoryservice"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		foo, err := ec2.NewSubnet(ctx, "foo", &ec2.SubnetArgs{
			VpcId:            main.ID(),
			AvailabilityZone: pulumi.String("us-west-2a"),
			CidrBlock:        pulumi.String("10.0.1.0/24"),
		})
		if err != nil {
			return err
		}
		barSubnet, err := ec2.NewSubnet(ctx, "barSubnet", &ec2.SubnetArgs{
			VpcId:            main.ID(),
			AvailabilityZone: pulumi.String("us-west-2b"),
			CidrBlock:        pulumi.String("10.0.2.0/24"),
		})
		if err != nil {
			return err
		}
		_, err = directoryservice.NewDirectory(ctx, "barDirectory", &directoryservice.DirectoryArgs{
			Name:     pulumi.String("corp.notexample.com"),
			Password: pulumi.String("SuperSecretPassw0rd"),
			Edition:  pulumi.String("Standard"),
			Type:     pulumi.String("MicrosoftAD"),
			VpcSettings: &directoryservice.DirectoryVpcSettingsArgs{
				VpcId: main.ID(),
				SubnetIds: pulumi.StringArray{
					foo.ID(),
					barSubnet.ID(),
				},
			},
			Tags: pulumi.StringMap{
				"Project": pulumi.String("foo"),
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.ec2.Subnet;
import com.pulumi.aws.ec2.SubnetArgs;
import com.pulumi.aws.directoryservice.Directory;
import com.pulumi.aws.directoryservice.DirectoryArgs;
import com.pulumi.aws.directoryservice.inputs.DirectoryVpcSettingsArgs;
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
        var main = new Vpc("main", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .build());

        var foo = new Subnet("foo", SubnetArgs.builder()        
            .vpcId(main.id())
            .availabilityZone("us-west-2a")
            .cidrBlock("10.0.1.0/24")
            .build());

        var barSubnet = new Subnet("barSubnet", SubnetArgs.builder()        
            .vpcId(main.id())
            .availabilityZone("us-west-2b")
            .cidrBlock("10.0.2.0/24")
            .build());

        var barDirectory = new Directory("barDirectory", DirectoryArgs.builder()        
            .name("corp.notexample.com")
            .password("SuperSecretPassw0rd")
            .edition("Standard")
            .type("MicrosoftAD")
            .vpcSettings(DirectoryVpcSettingsArgs.builder()
                .vpcId(main.id())
                .subnetIds(                
                    foo.id(),
                    barSubnet.id())
                .build())
            .tags(Map.of("Project", "foo"))
            .build());

    }
}
```
```yaml
resources:
  barDirectory:
    type: aws:directoryservice:Directory
    properties:
      name: corp.notexample.com
      password: SuperSecretPassw0rd
      edition: Standard
      type: MicrosoftAD
      vpcSettings:
        vpcId: ${main.id}
        subnetIds:
          - ${foo.id}
          - ${barSubnet.id}
      tags:
        Project: foo
  main:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
  foo:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${main.id}
      availabilityZone: us-west-2a
      cidrBlock: 10.0.1.0/24
  barSubnet:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${main.id}
      availabilityZone: us-west-2b
      cidrBlock: 10.0.2.0/24
```
{{% /example %}}
{{% example %}}
### Microsoft Active Directory Connector (ADConnector)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ec2.Vpc("main", {cidrBlock: "10.0.0.0/16"});
const foo = new aws.ec2.Subnet("foo", {
    vpcId: main.id,
    availabilityZone: "us-west-2a",
    cidrBlock: "10.0.1.0/24",
});
const bar = new aws.ec2.Subnet("bar", {
    vpcId: main.id,
    availabilityZone: "us-west-2b",
    cidrBlock: "10.0.2.0/24",
});
const connector = new aws.directoryservice.Directory("connector", {
    name: "corp.notexample.com",
    password: "SuperSecretPassw0rd",
    size: "Small",
    type: "ADConnector",
    connectSettings: {
        customerDnsIps: ["A.B.C.D"],
        customerUsername: "Admin",
        subnetIds: [
            foo.id,
            bar.id,
        ],
        vpcId: main.id,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
foo = aws.ec2.Subnet("foo",
    vpc_id=main.id,
    availability_zone="us-west-2a",
    cidr_block="10.0.1.0/24")
bar = aws.ec2.Subnet("bar",
    vpc_id=main.id,
    availability_zone="us-west-2b",
    cidr_block="10.0.2.0/24")
connector = aws.directoryservice.Directory("connector",
    name="corp.notexample.com",
    password="SuperSecretPassw0rd",
    size="Small",
    type="ADConnector",
    connect_settings=aws.directoryservice.DirectoryConnectSettingsArgs(
        customer_dns_ips=["A.B.C.D"],
        customer_username="Admin",
        subnet_ids=[
            foo.id,
            bar.id,
        ],
        vpc_id=main.id,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = new Aws.Ec2.Vpc("main", new()
    {
        CidrBlock = "10.0.0.0/16",
    });

    var foo = new Aws.Ec2.Subnet("foo", new()
    {
        VpcId = main.Id,
        AvailabilityZone = "us-west-2a",
        CidrBlock = "10.0.1.0/24",
    });

    var bar = new Aws.Ec2.Subnet("bar", new()
    {
        VpcId = main.Id,
        AvailabilityZone = "us-west-2b",
        CidrBlock = "10.0.2.0/24",
    });

    var connector = new Aws.DirectoryService.Directory("connector", new()
    {
        Name = "corp.notexample.com",
        Password = "SuperSecretPassw0rd",
        Size = "Small",
        Type = "ADConnector",
        ConnectSettings = new Aws.DirectoryService.Inputs.DirectoryConnectSettingsArgs
        {
            CustomerDnsIps = new[]
            {
                "A.B.C.D",
            },
            CustomerUsername = "Admin",
            SubnetIds = new[]
            {
                foo.Id,
                bar.Id,
            },
            VpcId = main.Id,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directoryservice"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		foo, err := ec2.NewSubnet(ctx, "foo", &ec2.SubnetArgs{
			VpcId:            main.ID(),
			AvailabilityZone: pulumi.String("us-west-2a"),
			CidrBlock:        pulumi.String("10.0.1.0/24"),
		})
		if err != nil {
			return err
		}
		bar, err := ec2.NewSubnet(ctx, "bar", &ec2.SubnetArgs{
			VpcId:            main.ID(),
			AvailabilityZone: pulumi.String("us-west-2b"),
			CidrBlock:        pulumi.String("10.0.2.0/24"),
		})
		if err != nil {
			return err
		}
		_, err = directoryservice.NewDirectory(ctx, "connector", &directoryservice.DirectoryArgs{
			Name:     pulumi.String("corp.notexample.com"),
			Password: pulumi.String("SuperSecretPassw0rd"),
			Size:     pulumi.String("Small"),
			Type:     pulumi.String("ADConnector"),
			ConnectSettings: &directoryservice.DirectoryConnectSettingsArgs{
				CustomerDnsIps: pulumi.StringArray{
					pulumi.String("A.B.C.D"),
				},
				CustomerUsername: pulumi.String("Admin"),
				SubnetIds: pulumi.StringArray{
					foo.ID(),
					bar.ID(),
				},
				VpcId: main.ID(),
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.ec2.Subnet;
import com.pulumi.aws.ec2.SubnetArgs;
import com.pulumi.aws.directoryservice.Directory;
import com.pulumi.aws.directoryservice.DirectoryArgs;
import com.pulumi.aws.directoryservice.inputs.DirectoryConnectSettingsArgs;
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
        var main = new Vpc("main", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .build());

        var foo = new Subnet("foo", SubnetArgs.builder()        
            .vpcId(main.id())
            .availabilityZone("us-west-2a")
            .cidrBlock("10.0.1.0/24")
            .build());

        var bar = new Subnet("bar", SubnetArgs.builder()        
            .vpcId(main.id())
            .availabilityZone("us-west-2b")
            .cidrBlock("10.0.2.0/24")
            .build());

        var connector = new Directory("connector", DirectoryArgs.builder()        
            .name("corp.notexample.com")
            .password("SuperSecretPassw0rd")
            .size("Small")
            .type("ADConnector")
            .connectSettings(DirectoryConnectSettingsArgs.builder()
                .customerDnsIps("A.B.C.D")
                .customerUsername("Admin")
                .subnetIds(                
                    foo.id(),
                    bar.id())
                .vpcId(main.id())
                .build())
            .build());

    }
}
```
```yaml
resources:
  connector:
    type: aws:directoryservice:Directory
    properties:
      name: corp.notexample.com
      password: SuperSecretPassw0rd
      size: Small
      type: ADConnector
      connectSettings:
        customerDnsIps:
          - A.B.C.D
        customerUsername: Admin
        subnetIds:
          - ${foo.id}
          - ${bar.id}
        vpcId: ${main.id}
  main:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
  foo:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${main.id}
      availabilityZone: us-west-2a
      cidrBlock: 10.0.1.0/24
  bar:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${main.id}
      availabilityZone: us-west-2b
      cidrBlock: 10.0.2.0/24
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DirectoryService directories using the directory `id`. For example:

```sh
 $ pulumi import aws:directoryservice/directory:Directory sample d-926724cf57
```
 