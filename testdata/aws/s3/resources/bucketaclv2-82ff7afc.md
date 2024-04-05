Provides an S3 bucket ACL resource.

> **Note:** destroy does not delete the S3 Bucket ACL but does remove the resource from state.

> This resource cannot be used with S3 directory buckets.

{{% examples %}}
## Example Usage
{{% example %}}
### With `private` ACL

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleBucketOwnershipControls = new aws.s3.BucketOwnershipControls("exampleBucketOwnershipControls", {
    bucket: exampleBucketV2.id,
    rule: {
        objectOwnership: "BucketOwnerPreferred",
    },
});
const exampleBucketAclV2 = new aws.s3.BucketAclV2("exampleBucketAclV2", {
    bucket: exampleBucketV2.id,
    acl: "private",
}, {
    dependsOn: [exampleBucketOwnershipControls],
});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_bucket_ownership_controls = aws.s3.BucketOwnershipControls("exampleBucketOwnershipControls",
    bucket=example_bucket_v2.id,
    rule=aws.s3.BucketOwnershipControlsRuleArgs(
        object_ownership="BucketOwnerPreferred",
    ))
example_bucket_acl_v2 = aws.s3.BucketAclV2("exampleBucketAclV2",
    bucket=example_bucket_v2.id,
    acl="private",
    opts=pulumi.ResourceOptions(depends_on=[example_bucket_ownership_controls]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleBucketOwnershipControls = new Aws.S3.BucketOwnershipControls("exampleBucketOwnershipControls", new()
    {
        Bucket = exampleBucketV2.Id,
        Rule = new Aws.S3.Inputs.BucketOwnershipControlsRuleArgs
        {
            ObjectOwnership = "BucketOwnerPreferred",
        },
    });

    var exampleBucketAclV2 = new Aws.S3.BucketAclV2("exampleBucketAclV2", new()
    {
        Bucket = exampleBucketV2.Id,
        Acl = "private",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleBucketOwnershipControls,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
		if err != nil {
			return err
		}
		exampleBucketOwnershipControls, err := s3.NewBucketOwnershipControls(ctx, "exampleBucketOwnershipControls", &s3.BucketOwnershipControlsArgs{
			Bucket: exampleBucketV2.ID(),
			Rule: &s3.BucketOwnershipControlsRuleArgs{
				ObjectOwnership: pulumi.String("BucketOwnerPreferred"),
			},
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAclV2(ctx, "exampleBucketAclV2", &s3.BucketAclV2Args{
			Bucket: exampleBucketV2.ID(),
			Acl:    pulumi.String("private"),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleBucketOwnershipControls,
		}))
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketOwnershipControls;
import com.pulumi.aws.s3.BucketOwnershipControlsArgs;
import com.pulumi.aws.s3.inputs.BucketOwnershipControlsRuleArgs;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.resources.CustomResourceOptions;
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
        var exampleBucketV2 = new BucketV2("exampleBucketV2");

        var exampleBucketOwnershipControls = new BucketOwnershipControls("exampleBucketOwnershipControls", BucketOwnershipControlsArgs.builder()        
            .bucket(exampleBucketV2.id())
            .rule(BucketOwnershipControlsRuleArgs.builder()
                .objectOwnership("BucketOwnerPreferred")
                .build())
            .build());

        var exampleBucketAclV2 = new BucketAclV2("exampleBucketAclV2", BucketAclV2Args.builder()        
            .bucket(exampleBucketV2.id())
            .acl("private")
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleBucketOwnershipControls)
                .build());

    }
}
```
```yaml
resources:
  exampleBucketV2:
    type: aws:s3:BucketV2
  exampleBucketOwnershipControls:
    type: aws:s3:BucketOwnershipControls
    properties:
      bucket: ${exampleBucketV2.id}
      rule:
        objectOwnership: BucketOwnerPreferred
  exampleBucketAclV2:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${exampleBucketV2.id}
      acl: private
    options:
      dependson:
        - ${exampleBucketOwnershipControls}
```
{{% /example %}}
{{% example %}}
### With `public-read` ACL

> This example explicitly disables the default S3 bucket security settings. This
should be done with caution, as all bucket objects become publicly exposed.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleBucketOwnershipControls = new aws.s3.BucketOwnershipControls("exampleBucketOwnershipControls", {
    bucket: exampleBucketV2.id,
    rule: {
        objectOwnership: "BucketOwnerPreferred",
    },
});
const exampleBucketPublicAccessBlock = new aws.s3.BucketPublicAccessBlock("exampleBucketPublicAccessBlock", {
    bucket: exampleBucketV2.id,
    blockPublicAcls: false,
    blockPublicPolicy: false,
    ignorePublicAcls: false,
    restrictPublicBuckets: false,
});
const exampleBucketAclV2 = new aws.s3.BucketAclV2("exampleBucketAclV2", {
    bucket: exampleBucketV2.id,
    acl: "public-read",
}, {
    dependsOn: [
        exampleBucketOwnershipControls,
        exampleBucketPublicAccessBlock,
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_bucket_ownership_controls = aws.s3.BucketOwnershipControls("exampleBucketOwnershipControls",
    bucket=example_bucket_v2.id,
    rule=aws.s3.BucketOwnershipControlsRuleArgs(
        object_ownership="BucketOwnerPreferred",
    ))
example_bucket_public_access_block = aws.s3.BucketPublicAccessBlock("exampleBucketPublicAccessBlock",
    bucket=example_bucket_v2.id,
    block_public_acls=False,
    block_public_policy=False,
    ignore_public_acls=False,
    restrict_public_buckets=False)
example_bucket_acl_v2 = aws.s3.BucketAclV2("exampleBucketAclV2",
    bucket=example_bucket_v2.id,
    acl="public-read",
    opts=pulumi.ResourceOptions(depends_on=[
            example_bucket_ownership_controls,
            example_bucket_public_access_block,
        ]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleBucketOwnershipControls = new Aws.S3.BucketOwnershipControls("exampleBucketOwnershipControls", new()
    {
        Bucket = exampleBucketV2.Id,
        Rule = new Aws.S3.Inputs.BucketOwnershipControlsRuleArgs
        {
            ObjectOwnership = "BucketOwnerPreferred",
        },
    });

    var exampleBucketPublicAccessBlock = new Aws.S3.BucketPublicAccessBlock("exampleBucketPublicAccessBlock", new()
    {
        Bucket = exampleBucketV2.Id,
        BlockPublicAcls = false,
        BlockPublicPolicy = false,
        IgnorePublicAcls = false,
        RestrictPublicBuckets = false,
    });

    var exampleBucketAclV2 = new Aws.S3.BucketAclV2("exampleBucketAclV2", new()
    {
        Bucket = exampleBucketV2.Id,
        Acl = "public-read",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleBucketOwnershipControls,
            exampleBucketPublicAccessBlock,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
		if err != nil {
			return err
		}
		exampleBucketOwnershipControls, err := s3.NewBucketOwnershipControls(ctx, "exampleBucketOwnershipControls", &s3.BucketOwnershipControlsArgs{
			Bucket: exampleBucketV2.ID(),
			Rule: &s3.BucketOwnershipControlsRuleArgs{
				ObjectOwnership: pulumi.String("BucketOwnerPreferred"),
			},
		})
		if err != nil {
			return err
		}
		exampleBucketPublicAccessBlock, err := s3.NewBucketPublicAccessBlock(ctx, "exampleBucketPublicAccessBlock", &s3.BucketPublicAccessBlockArgs{
			Bucket:                exampleBucketV2.ID(),
			BlockPublicAcls:       pulumi.Bool(false),
			BlockPublicPolicy:     pulumi.Bool(false),
			IgnorePublicAcls:      pulumi.Bool(false),
			RestrictPublicBuckets: pulumi.Bool(false),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAclV2(ctx, "exampleBucketAclV2", &s3.BucketAclV2Args{
			Bucket: exampleBucketV2.ID(),
			Acl:    pulumi.String("public-read"),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleBucketOwnershipControls,
			exampleBucketPublicAccessBlock,
		}))
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketOwnershipControls;
import com.pulumi.aws.s3.BucketOwnershipControlsArgs;
import com.pulumi.aws.s3.inputs.BucketOwnershipControlsRuleArgs;
import com.pulumi.aws.s3.BucketPublicAccessBlock;
import com.pulumi.aws.s3.BucketPublicAccessBlockArgs;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.resources.CustomResourceOptions;
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
        var exampleBucketV2 = new BucketV2("exampleBucketV2");

        var exampleBucketOwnershipControls = new BucketOwnershipControls("exampleBucketOwnershipControls", BucketOwnershipControlsArgs.builder()        
            .bucket(exampleBucketV2.id())
            .rule(BucketOwnershipControlsRuleArgs.builder()
                .objectOwnership("BucketOwnerPreferred")
                .build())
            .build());

        var exampleBucketPublicAccessBlock = new BucketPublicAccessBlock("exampleBucketPublicAccessBlock", BucketPublicAccessBlockArgs.builder()        
            .bucket(exampleBucketV2.id())
            .blockPublicAcls(false)
            .blockPublicPolicy(false)
            .ignorePublicAcls(false)
            .restrictPublicBuckets(false)
            .build());

        var exampleBucketAclV2 = new BucketAclV2("exampleBucketAclV2", BucketAclV2Args.builder()        
            .bucket(exampleBucketV2.id())
            .acl("public-read")
            .build(), CustomResourceOptions.builder()
                .dependsOn(                
                    exampleBucketOwnershipControls,
                    exampleBucketPublicAccessBlock)
                .build());

    }
}
```
```yaml
resources:
  exampleBucketV2:
    type: aws:s3:BucketV2
  exampleBucketOwnershipControls:
    type: aws:s3:BucketOwnershipControls
    properties:
      bucket: ${exampleBucketV2.id}
      rule:
        objectOwnership: BucketOwnerPreferred
  exampleBucketPublicAccessBlock:
    type: aws:s3:BucketPublicAccessBlock
    properties:
      bucket: ${exampleBucketV2.id}
      blockPublicAcls: false
      blockPublicPolicy: false
      ignorePublicAcls: false
      restrictPublicBuckets: false
  exampleBucketAclV2:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${exampleBucketV2.id}
      acl: public-read
    options:
      dependson:
        - ${exampleBucketOwnershipControls}
        - ${exampleBucketPublicAccessBlock}
```
{{% /example %}}
{{% example %}}
### With Grants

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.s3.getCanonicalUserId({});
const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleBucketOwnershipControls = new aws.s3.BucketOwnershipControls("exampleBucketOwnershipControls", {
    bucket: exampleBucketV2.id,
    rule: {
        objectOwnership: "BucketOwnerPreferred",
    },
});
const exampleBucketAclV2 = new aws.s3.BucketAclV2("exampleBucketAclV2", {
    bucket: exampleBucketV2.id,
    accessControlPolicy: {
        grants: [
            {
                grantee: {
                    id: current.then(current => current.id),
                    type: "CanonicalUser",
                },
                permission: "READ",
            },
            {
                grantee: {
                    type: "Group",
                    uri: "http://acs.amazonaws.com/groups/s3/LogDelivery",
                },
                permission: "READ_ACP",
            },
        ],
        owner: {
            id: current.then(current => current.id),
        },
    },
}, {
    dependsOn: [exampleBucketOwnershipControls],
});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.s3.get_canonical_user_id()
example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_bucket_ownership_controls = aws.s3.BucketOwnershipControls("exampleBucketOwnershipControls",
    bucket=example_bucket_v2.id,
    rule=aws.s3.BucketOwnershipControlsRuleArgs(
        object_ownership="BucketOwnerPreferred",
    ))
example_bucket_acl_v2 = aws.s3.BucketAclV2("exampleBucketAclV2",
    bucket=example_bucket_v2.id,
    access_control_policy=aws.s3.BucketAclV2AccessControlPolicyArgs(
        grants=[
            aws.s3.BucketAclV2AccessControlPolicyGrantArgs(
                grantee=aws.s3.BucketAclV2AccessControlPolicyGrantGranteeArgs(
                    id=current.id,
                    type="CanonicalUser",
                ),
                permission="READ",
            ),
            aws.s3.BucketAclV2AccessControlPolicyGrantArgs(
                grantee=aws.s3.BucketAclV2AccessControlPolicyGrantGranteeArgs(
                    type="Group",
                    uri="http://acs.amazonaws.com/groups/s3/LogDelivery",
                ),
                permission="READ_ACP",
            ),
        ],
        owner=aws.s3.BucketAclV2AccessControlPolicyOwnerArgs(
            id=current.id,
        ),
    ),
    opts=pulumi.ResourceOptions(depends_on=[example_bucket_ownership_controls]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.S3.GetCanonicalUserId.Invoke();

    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleBucketOwnershipControls = new Aws.S3.BucketOwnershipControls("exampleBucketOwnershipControls", new()
    {
        Bucket = exampleBucketV2.Id,
        Rule = new Aws.S3.Inputs.BucketOwnershipControlsRuleArgs
        {
            ObjectOwnership = "BucketOwnerPreferred",
        },
    });

    var exampleBucketAclV2 = new Aws.S3.BucketAclV2("exampleBucketAclV2", new()
    {
        Bucket = exampleBucketV2.Id,
        AccessControlPolicy = new Aws.S3.Inputs.BucketAclV2AccessControlPolicyArgs
        {
            Grants = new[]
            {
                new Aws.S3.Inputs.BucketAclV2AccessControlPolicyGrantArgs
                {
                    Grantee = new Aws.S3.Inputs.BucketAclV2AccessControlPolicyGrantGranteeArgs
                    {
                        Id = current.Apply(getCanonicalUserIdResult => getCanonicalUserIdResult.Id),
                        Type = "CanonicalUser",
                    },
                    Permission = "READ",
                },
                new Aws.S3.Inputs.BucketAclV2AccessControlPolicyGrantArgs
                {
                    Grantee = new Aws.S3.Inputs.BucketAclV2AccessControlPolicyGrantGranteeArgs
                    {
                        Type = "Group",
                        Uri = "http://acs.amazonaws.com/groups/s3/LogDelivery",
                    },
                    Permission = "READ_ACP",
                },
            },
            Owner = new Aws.S3.Inputs.BucketAclV2AccessControlPolicyOwnerArgs
            {
                Id = current.Apply(getCanonicalUserIdResult => getCanonicalUserIdResult.Id),
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleBucketOwnershipControls,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := s3.GetCanonicalUserId(ctx, nil, nil)
		if err != nil {
			return err
		}
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
		if err != nil {
			return err
		}
		exampleBucketOwnershipControls, err := s3.NewBucketOwnershipControls(ctx, "exampleBucketOwnershipControls", &s3.BucketOwnershipControlsArgs{
			Bucket: exampleBucketV2.ID(),
			Rule: &s3.BucketOwnershipControlsRuleArgs{
				ObjectOwnership: pulumi.String("BucketOwnerPreferred"),
			},
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAclV2(ctx, "exampleBucketAclV2", &s3.BucketAclV2Args{
			Bucket: exampleBucketV2.ID(),
			AccessControlPolicy: &s3.BucketAclV2AccessControlPolicyArgs{
				Grants: s3.BucketAclV2AccessControlPolicyGrantArray{
					&s3.BucketAclV2AccessControlPolicyGrantArgs{
						Grantee: &s3.BucketAclV2AccessControlPolicyGrantGranteeArgs{
							Id:   *pulumi.String(current.Id),
							Type: pulumi.String("CanonicalUser"),
						},
						Permission: pulumi.String("READ"),
					},
					&s3.BucketAclV2AccessControlPolicyGrantArgs{
						Grantee: &s3.BucketAclV2AccessControlPolicyGrantGranteeArgs{
							Type: pulumi.String("Group"),
							Uri:  pulumi.String("http://acs.amazonaws.com/groups/s3/LogDelivery"),
						},
						Permission: pulumi.String("READ_ACP"),
					},
				},
				Owner: &s3.BucketAclV2AccessControlPolicyOwnerArgs{
					Id: *pulumi.String(current.Id),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleBucketOwnershipControls,
		}))
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
import com.pulumi.aws.s3.S3Functions;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketOwnershipControls;
import com.pulumi.aws.s3.BucketOwnershipControlsArgs;
import com.pulumi.aws.s3.inputs.BucketOwnershipControlsRuleArgs;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.aws.s3.inputs.BucketAclV2AccessControlPolicyArgs;
import com.pulumi.aws.s3.inputs.BucketAclV2AccessControlPolicyOwnerArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        final var current = S3Functions.getCanonicalUserId();

        var exampleBucketV2 = new BucketV2("exampleBucketV2");

        var exampleBucketOwnershipControls = new BucketOwnershipControls("exampleBucketOwnershipControls", BucketOwnershipControlsArgs.builder()        
            .bucket(exampleBucketV2.id())
            .rule(BucketOwnershipControlsRuleArgs.builder()
                .objectOwnership("BucketOwnerPreferred")
                .build())
            .build());

        var exampleBucketAclV2 = new BucketAclV2("exampleBucketAclV2", BucketAclV2Args.builder()        
            .bucket(exampleBucketV2.id())
            .accessControlPolicy(BucketAclV2AccessControlPolicyArgs.builder()
                .grants(                
                    BucketAclV2AccessControlPolicyGrantArgs.builder()
                        .grantee(BucketAclV2AccessControlPolicyGrantGranteeArgs.builder()
                            .id(current.applyValue(getCanonicalUserIdResult -> getCanonicalUserIdResult.id()))
                            .type("CanonicalUser")
                            .build())
                        .permission("READ")
                        .build(),
                    BucketAclV2AccessControlPolicyGrantArgs.builder()
                        .grantee(BucketAclV2AccessControlPolicyGrantGranteeArgs.builder()
                            .type("Group")
                            .uri("http://acs.amazonaws.com/groups/s3/LogDelivery")
                            .build())
                        .permission("READ_ACP")
                        .build())
                .owner(BucketAclV2AccessControlPolicyOwnerArgs.builder()
                    .id(current.applyValue(getCanonicalUserIdResult -> getCanonicalUserIdResult.id()))
                    .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleBucketOwnershipControls)
                .build());

    }
}
```
```yaml
resources:
  exampleBucketV2:
    type: aws:s3:BucketV2
  exampleBucketOwnershipControls:
    type: aws:s3:BucketOwnershipControls
    properties:
      bucket: ${exampleBucketV2.id}
      rule:
        objectOwnership: BucketOwnerPreferred
  exampleBucketAclV2:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${exampleBucketV2.id}
      accessControlPolicy:
        grants:
          - grantee:
              id: ${current.id}
              type: CanonicalUser
            permission: READ
          - grantee:
              type: Group
              uri: http://acs.amazonaws.com/groups/s3/LogDelivery
            permission: READ_ACP
        owner:
          id: ${current.id}
    options:
      dependson:
        - ${exampleBucketOwnershipControls}
variables:
  current:
    fn::invoke:
      Function: aws:s3:getCanonicalUserId
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

If the owner (account ID) of the source bucket is the _same_ account used to configure the AWS Provider, and the source bucket is __configured__ with a [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#canned-acl) (i.e. predefined grant), import using the `bucket` and `acl` separated by a comma (`,`):

If the owner (account ID) of the source bucket _differs_ from the account used to configure the AWS Provider, and the source bucket is __not configured__ with a [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#canned-acl) (i.e. predefined grant), imported using the `bucket` and `expected_bucket_owner` separated by a comma (`,`):

If the owner (account ID) of the source bucket _differs_ from the account used to configure the AWS Provider, and the source bucket is __configured__ with a [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#canned-acl) (i.e. predefined grant), imported using the `bucket`, `expected_bucket_owner`, and `acl` separated by commas (`,`):

__Using `pulumi import` to import__ using `bucket`, `expected_bucket_owner`, and/or `acl`, depending on your situation. For example:

If the owner (account ID) of the source bucket is the _same_ account used to configure the AWS Provider, and the source bucket is __not configured__ with a [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#canned-acl) (i.e. predefined grant), import using the `bucket`:

```sh
 $ pulumi import aws:s3/bucketAclV2:BucketAclV2 example bucket-name
```
 If the owner (account ID) of the source bucket is the _same_ account used to configure the AWS Provider, and the source bucket is __configured__ with a [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#canned-acl) (i.e. predefined grant), import using the `bucket` and `acl` separated by a comma (`,`):

```sh
 $ pulumi import aws:s3/bucketAclV2:BucketAclV2 example bucket-name,private
```
 If the owner (account ID) of the source bucket _differs_ from the account used to configure the AWS Provider, and the source bucket is __not configured__ with a [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#canned-acl) (i.e. predefined grant), imported using the `bucket` and `expected_bucket_owner` separated by a comma (`,`):

```sh
 $ pulumi import aws:s3/bucketAclV2:BucketAclV2 example bucket-name,123456789012
```
 If the owner (account ID) of the source bucket _differs_ from the account used to configure the AWS Provider, and the source bucket is __configured__ with a [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#canned-acl) (i.e. predefined grant), imported using the `bucket`, `expected_bucket_owner`, and `acl` separated by commas (`,`):

```sh
 $ pulumi import aws:s3/bucketAclV2:BucketAclV2 example bucket-name,123456789012,private
```
 