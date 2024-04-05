Provides an Amplify Domain Association resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleApp = new aws.amplify.App("exampleApp", {customRules: [{
    source: "https://example.com",
    status: "302",
    target: "https://www.example.com",
}]});
const master = new aws.amplify.Branch("master", {
    appId: exampleApp.id,
    branchName: "master",
});
const exampleDomainAssociation = new aws.amplify.DomainAssociation("exampleDomainAssociation", {
    appId: exampleApp.id,
    domainName: "example.com",
    subDomains: [
        {
            branchName: master.branchName,
            prefix: "",
        },
        {
            branchName: master.branchName,
            prefix: "www",
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example_app = aws.amplify.App("exampleApp", custom_rules=[aws.amplify.AppCustomRuleArgs(
    source="https://example.com",
    status="302",
    target="https://www.example.com",
)])
master = aws.amplify.Branch("master",
    app_id=example_app.id,
    branch_name="master")
example_domain_association = aws.amplify.DomainAssociation("exampleDomainAssociation",
    app_id=example_app.id,
    domain_name="example.com",
    sub_domains=[
        aws.amplify.DomainAssociationSubDomainArgs(
            branch_name=master.branch_name,
            prefix="",
        ),
        aws.amplify.DomainAssociationSubDomainArgs(
            branch_name=master.branch_name,
            prefix="www",
        ),
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleApp = new Aws.Amplify.App("exampleApp", new()
    {
        CustomRules = new[]
        {
            new Aws.Amplify.Inputs.AppCustomRuleArgs
            {
                Source = "https://example.com",
                Status = "302",
                Target = "https://www.example.com",
            },
        },
    });

    var master = new Aws.Amplify.Branch("master", new()
    {
        AppId = exampleApp.Id,
        BranchName = "master",
    });

    var exampleDomainAssociation = new Aws.Amplify.DomainAssociation("exampleDomainAssociation", new()
    {
        AppId = exampleApp.Id,
        DomainName = "example.com",
        SubDomains = new[]
        {
            new Aws.Amplify.Inputs.DomainAssociationSubDomainArgs
            {
                BranchName = master.BranchName,
                Prefix = "",
            },
            new Aws.Amplify.Inputs.DomainAssociationSubDomainArgs
            {
                BranchName = master.BranchName,
                Prefix = "www",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/amplify"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleApp, err := amplify.NewApp(ctx, "exampleApp", &amplify.AppArgs{
			CustomRules: amplify.AppCustomRuleArray{
				&amplify.AppCustomRuleArgs{
					Source: pulumi.String("https://example.com"),
					Status: pulumi.String("302"),
					Target: pulumi.String("https://www.example.com"),
				},
			},
		})
		if err != nil {
			return err
		}
		master, err := amplify.NewBranch(ctx, "master", &amplify.BranchArgs{
			AppId:      exampleApp.ID(),
			BranchName: pulumi.String("master"),
		})
		if err != nil {
			return err
		}
		_, err = amplify.NewDomainAssociation(ctx, "exampleDomainAssociation", &amplify.DomainAssociationArgs{
			AppId:      exampleApp.ID(),
			DomainName: pulumi.String("example.com"),
			SubDomains: amplify.DomainAssociationSubDomainArray{
				&amplify.DomainAssociationSubDomainArgs{
					BranchName: master.BranchName,
					Prefix:     pulumi.String(""),
				},
				&amplify.DomainAssociationSubDomainArgs{
					BranchName: master.BranchName,
					Prefix:     pulumi.String("www"),
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
import com.pulumi.aws.amplify.App;
import com.pulumi.aws.amplify.AppArgs;
import com.pulumi.aws.amplify.inputs.AppCustomRuleArgs;
import com.pulumi.aws.amplify.Branch;
import com.pulumi.aws.amplify.BranchArgs;
import com.pulumi.aws.amplify.DomainAssociation;
import com.pulumi.aws.amplify.DomainAssociationArgs;
import com.pulumi.aws.amplify.inputs.DomainAssociationSubDomainArgs;
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
        var exampleApp = new App("exampleApp", AppArgs.builder()        
            .customRules(AppCustomRuleArgs.builder()
                .source("https://example.com")
                .status("302")
                .target("https://www.example.com")
                .build())
            .build());

        var master = new Branch("master", BranchArgs.builder()        
            .appId(exampleApp.id())
            .branchName("master")
            .build());

        var exampleDomainAssociation = new DomainAssociation("exampleDomainAssociation", DomainAssociationArgs.builder()        
            .appId(exampleApp.id())
            .domainName("example.com")
            .subDomains(            
                DomainAssociationSubDomainArgs.builder()
                    .branchName(master.branchName())
                    .prefix("")
                    .build(),
                DomainAssociationSubDomainArgs.builder()
                    .branchName(master.branchName())
                    .prefix("www")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  exampleApp:
    type: aws:amplify:App
    properties:
      # Setup redirect from https://example.com to https://www.example.com
      customRules:
        - source: https://example.com
          status: '302'
          target: https://www.example.com
  master:
    type: aws:amplify:Branch
    properties:
      appId: ${exampleApp.id}
      branchName: master
  exampleDomainAssociation:
    type: aws:amplify:DomainAssociation
    properties:
      appId: ${exampleApp.id}
      domainName: example.com
      # https://example.com
      subDomains:
        - branchName: ${master.branchName}
          prefix:
        - branchName: ${master.branchName}
          prefix: www
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amplify domain association using `app_id` and `domain_name`. For example:

```sh
 $ pulumi import aws:amplify/domainAssociation:DomainAssociation app d2ypk4k47z8u6/example.com
```
 