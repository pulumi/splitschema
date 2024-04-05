Provides a License Manager association.

> **Note:** License configurations can also be associated with launch templates by specifying the `license_specifications` block for an `aws.ec2.LaunchTemplate`.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAmi = aws.ec2.getAmi({
    mostRecent: true,
    owners: ["amazon"],
    filters: [{
        name: "name",
        values: ["amzn-ami-vpc-nat*"],
    }],
});
const exampleInstance = new aws.ec2.Instance("exampleInstance", {
    ami: exampleAmi.then(exampleAmi => exampleAmi.id),
    instanceType: "t2.micro",
});
const exampleLicenseConfiguration = new aws.licensemanager.LicenseConfiguration("exampleLicenseConfiguration", {licenseCountingType: "Instance"});
const exampleAssociation = new aws.licensemanager.Association("exampleAssociation", {
    licenseConfigurationArn: exampleLicenseConfiguration.arn,
    resourceArn: exampleInstance.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example_ami = aws.ec2.get_ami(most_recent=True,
    owners=["amazon"],
    filters=[aws.ec2.GetAmiFilterArgs(
        name="name",
        values=["amzn-ami-vpc-nat*"],
    )])
example_instance = aws.ec2.Instance("exampleInstance",
    ami=example_ami.id,
    instance_type="t2.micro")
example_license_configuration = aws.licensemanager.LicenseConfiguration("exampleLicenseConfiguration", license_counting_type="Instance")
example_association = aws.licensemanager.Association("exampleAssociation",
    license_configuration_arn=example_license_configuration.arn,
    resource_arn=example_instance.arn)
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
                    "amzn-ami-vpc-nat*",
                },
            },
        },
    });

    var exampleInstance = new Aws.Ec2.Instance("exampleInstance", new()
    {
        Ami = exampleAmi.Apply(getAmiResult => getAmiResult.Id),
        InstanceType = "t2.micro",
    });

    var exampleLicenseConfiguration = new Aws.LicenseManager.LicenseConfiguration("exampleLicenseConfiguration", new()
    {
        LicenseCountingType = "Instance",
    });

    var exampleAssociation = new Aws.LicenseManager.Association("exampleAssociation", new()
    {
        LicenseConfigurationArn = exampleLicenseConfiguration.Arn,
        ResourceArn = exampleInstance.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/licensemanager"
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
						"amzn-ami-vpc-nat*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		exampleInstance, err := ec2.NewInstance(ctx, "exampleInstance", &ec2.InstanceArgs{
			Ami:          *pulumi.String(exampleAmi.Id),
			InstanceType: pulumi.String("t2.micro"),
		})
		if err != nil {
			return err
		}
		exampleLicenseConfiguration, err := licensemanager.NewLicenseConfiguration(ctx, "exampleLicenseConfiguration", &licensemanager.LicenseConfigurationArgs{
			LicenseCountingType: pulumi.String("Instance"),
		})
		if err != nil {
			return err
		}
		_, err = licensemanager.NewAssociation(ctx, "exampleAssociation", &licensemanager.AssociationArgs{
			LicenseConfigurationArn: exampleLicenseConfiguration.Arn,
			ResourceArn:             exampleInstance.Arn,
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
import com.pulumi.aws.ec2.Instance;
import com.pulumi.aws.ec2.InstanceArgs;
import com.pulumi.aws.licensemanager.LicenseConfiguration;
import com.pulumi.aws.licensemanager.LicenseConfigurationArgs;
import com.pulumi.aws.licensemanager.Association;
import com.pulumi.aws.licensemanager.AssociationArgs;
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
                .values("amzn-ami-vpc-nat*")
                .build())
            .build());

        var exampleInstance = new Instance("exampleInstance", InstanceArgs.builder()        
            .ami(exampleAmi.applyValue(getAmiResult -> getAmiResult.id()))
            .instanceType("t2.micro")
            .build());

        var exampleLicenseConfiguration = new LicenseConfiguration("exampleLicenseConfiguration", LicenseConfigurationArgs.builder()        
            .licenseCountingType("Instance")
            .build());

        var exampleAssociation = new Association("exampleAssociation", AssociationArgs.builder()        
            .licenseConfigurationArn(exampleLicenseConfiguration.arn())
            .resourceArn(exampleInstance.arn())
            .build());

    }
}
```
```yaml
resources:
  exampleInstance:
    type: aws:ec2:Instance
    properties:
      ami: ${exampleAmi.id}
      instanceType: t2.micro
  exampleLicenseConfiguration:
    type: aws:licensemanager:LicenseConfiguration
    properties:
      licenseCountingType: Instance
  exampleAssociation:
    type: aws:licensemanager:Association
    properties:
      licenseConfigurationArn: ${exampleLicenseConfiguration.arn}
      resourceArn: ${exampleInstance.arn}
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
              - amzn-ami-vpc-nat*
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import license configurations using `resource_arn,license_configuration_arn`. For example:

```sh
 $ pulumi import aws:licensemanager/association:Association example arn:aws:ec2:eu-west-1:123456789012:image/ami-123456789abcdef01,arn:aws:license-manager:eu-west-1:123456789012:license-configuration:lic-0123456789abcdef0123456789abcdef
```
 