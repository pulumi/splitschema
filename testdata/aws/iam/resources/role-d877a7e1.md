Provides an IAM role.

> **NOTE:** If policies are attached to the role via the `aws.iam.PolicyAttachment` resource and you are modifying the role `name` or `path`, the `force_detach_policies` argument must be set to `true` and applied before attempting the operation otherwise you will encounter a `DeleteConflict` error. The `aws.iam.RolePolicyAttachment` resource (recommended) does not have this requirement.

> **NOTE:** If you use this resource's `managed_policy_arns` argument or `inline_policy` configuration blocks, this resource will take over exclusive management of the role's respective policy types (e.g., both policy types if both arguments are used). These arguments are incompatible with other ways of managing a role's policies, such as `aws.iam.PolicyAttachment`, `aws.iam.RolePolicyAttachment`, and `aws.iam.RolePolicy`. If you attempt to manage a role's policies by multiple means, you will get resource cycling and/or errors.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testRole = new aws.iam.Role("testRole", {
    assumeRolePolicy: JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Action: "sts:AssumeRole",
            Effect: "Allow",
            Sid: "",
            Principal: {
                Service: "ec2.amazonaws.com",
            },
        }],
    }),
    tags: {
        "tag-key": "tag-value",
    },
});
```
```python
import pulumi
import json
import pulumi_aws as aws

test_role = aws.iam.Role("testRole",
    assume_role_policy=json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Action": "sts:AssumeRole",
            "Effect": "Allow",
            "Sid": "",
            "Principal": {
                "Service": "ec2.amazonaws.com",
            },
        }],
    }),
    tags={
        "tag-key": "tag-value",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testRole = new Aws.Iam.Role("testRole", new()
    {
        AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = "sts:AssumeRole",
                    ["Effect"] = "Allow",
                    ["Sid"] = "",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["Service"] = "ec2.amazonaws.com",
                    },
                },
            },
        }),
        Tags = 
        {
            { "tag-key", "tag-value" },
        },
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": "sts:AssumeRole",
					"Effect": "Allow",
					"Sid":    "",
					"Principal": map[string]interface{}{
						"Service": "ec2.amazonaws.com",
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = iam.NewRole(ctx, "testRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(json0),
			Tags: pulumi.StringMap{
				"tag-key": pulumi.String("tag-value"),
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
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
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
        var testRole = new Role("testRole", RoleArgs.builder()        
            .assumeRolePolicy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", "sts:AssumeRole"),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Sid", ""),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("Service", "ec2.amazonaws.com")
                        ))
                    )))
                )))
            .tags(Map.of("tag-key", "tag-value"))
            .build());

    }
}
```
```yaml
resources:
  testRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Action: sts:AssumeRole
              Effect: Allow
              Sid:
              Principal:
                Service: ec2.amazonaws.com
      tags:
        tag-key: tag-value
```
{{% /example %}}
{{% example %}}
### Example of Using Data Source for Assume Role Policy

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const instanceAssumeRolePolicy = aws.iam.getPolicyDocument({
    statements: [{
        actions: ["sts:AssumeRole"],
        principals: [{
            type: "Service",
            identifiers: ["ec2.amazonaws.com"],
        }],
    }],
});
const instance = new aws.iam.Role("instance", {
    path: "/system/",
    assumeRolePolicy: instanceAssumeRolePolicy.then(instanceAssumeRolePolicy => instanceAssumeRolePolicy.json),
});
```
```python
import pulumi
import pulumi_aws as aws

instance_assume_role_policy = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["sts:AssumeRole"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["ec2.amazonaws.com"],
    )],
)])
instance = aws.iam.Role("instance",
    path="/system/",
    assume_role_policy=instance_assume_role_policy.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var instanceAssumeRolePolicy = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "ec2.amazonaws.com",
                        },
                    },
                },
            },
        },
    });

    var instance = new Aws.Iam.Role("instance", new()
    {
        Path = "/system/",
        AssumeRolePolicy = instanceAssumeRolePolicy.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		instanceAssumeRolePolicy, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Actions: []string{
						"sts:AssumeRole",
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"ec2.amazonaws.com",
							},
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = iam.NewRole(ctx, "instance", &iam.RoleArgs{
			Path:             pulumi.String("/system/"),
			AssumeRolePolicy: *pulumi.String(instanceAssumeRolePolicy.Json),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
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
        final var instanceAssumeRolePolicy = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("sts:AssumeRole")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("ec2.amazonaws.com")
                    .build())
                .build())
            .build());

        var instance = new Role("instance", RoleArgs.builder()        
            .path("/system/")
            .assumeRolePolicy(instanceAssumeRolePolicy.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

    }
}
```
```yaml
resources:
  instance:
    type: aws:iam:Role
    properties:
      path: /system/
      assumeRolePolicy: ${instanceAssumeRolePolicy.json}
variables:
  instanceAssumeRolePolicy:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - sts:AssumeRole
            principals:
              - type: Service
                identifiers:
                  - ec2.amazonaws.com
```
{{% /example %}}
{{% example %}}
### Example of Exclusive Inline Policies

This example creates an IAM role with two inline IAM policies. If someone adds another inline policy out-of-band, on the next apply, this provider will remove that policy. If someone deletes these policies out-of-band, this provider will recreate them.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const inlinePolicy = aws.iam.getPolicyDocument({
    statements: [{
        actions: ["ec2:DescribeAccountAttributes"],
        resources: ["*"],
    }],
});
const example = new aws.iam.Role("example", {
    assumeRolePolicy: data.aws_iam_policy_document.instance_assume_role_policy.json,
    inlinePolicies: [
        {
            name: "my_inline_policy",
            policy: JSON.stringify({
                Version: "2012-10-17",
                Statement: [{
                    Action: ["ec2:Describe*"],
                    Effect: "Allow",
                    Resource: "*",
                }],
            }),
        },
        {
            name: "policy-8675309",
            policy: inlinePolicy.then(inlinePolicy => inlinePolicy.json),
        },
    ],
});
```
```python
import pulumi
import json
import pulumi_aws as aws

inline_policy = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["ec2:DescribeAccountAttributes"],
    resources=["*"],
)])
example = aws.iam.Role("example",
    assume_role_policy=data["aws_iam_policy_document"]["instance_assume_role_policy"]["json"],
    inline_policies=[
        aws.iam.RoleInlinePolicyArgs(
            name="my_inline_policy",
            policy=json.dumps({
                "Version": "2012-10-17",
                "Statement": [{
                    "Action": ["ec2:Describe*"],
                    "Effect": "Allow",
                    "Resource": "*",
                }],
            }),
        ),
        aws.iam.RoleInlinePolicyArgs(
            name="policy-8675309",
            policy=inline_policy.json,
        ),
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var inlinePolicy = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "ec2:DescribeAccountAttributes",
                },
                Resources = new[]
                {
                    "*",
                },
            },
        },
    });

    var example = new Aws.Iam.Role("example", new()
    {
        AssumeRolePolicy = data.Aws_iam_policy_document.Instance_assume_role_policy.Json,
        InlinePolicies = new[]
        {
            new Aws.Iam.Inputs.RoleInlinePolicyArgs
            {
                Name = "my_inline_policy",
                Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
                {
                    ["Version"] = "2012-10-17",
                    ["Statement"] = new[]
                    {
                        new Dictionary<string, object?>
                        {
                            ["Action"] = new[]
                            {
                                "ec2:Describe*",
                            },
                            ["Effect"] = "Allow",
                            ["Resource"] = "*",
                        },
                    },
                }),
            },
            new Aws.Iam.Inputs.RoleInlinePolicyArgs
            {
                Name = "policy-8675309",
                Policy = inlinePolicy.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
            },
        },
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		inlinePolicy, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Actions: []string{
						"ec2:DescribeAccountAttributes",
					},
					Resources: []string{
						"*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": []string{
						"ec2:Describe*",
					},
					"Effect":   "Allow",
					"Resource": "*",
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = iam.NewRole(ctx, "example", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.Any(data.Aws_iam_policy_document.Instance_assume_role_policy.Json),
			InlinePolicies: iam.RoleInlinePolicyArray{
				&iam.RoleInlinePolicyArgs{
					Name:   pulumi.String("my_inline_policy"),
					Policy: pulumi.String(json0),
				},
				&iam.RoleInlinePolicyArgs{
					Name:   pulumi.String("policy-8675309"),
					Policy: *pulumi.String(inlinePolicy.Json),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.inputs.RoleInlinePolicyArgs;
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
        final var inlinePolicy = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("ec2:DescribeAccountAttributes")
                .resources("*")
                .build())
            .build());

        var example = new Role("example", RoleArgs.builder()        
            .assumeRolePolicy(data.aws_iam_policy_document().instance_assume_role_policy().json())
            .inlinePolicies(            
                RoleInlinePolicyArgs.builder()
                    .name("my_inline_policy")
                    .policy(serializeJson(
                        jsonObject(
                            jsonProperty("Version", "2012-10-17"),
                            jsonProperty("Statement", jsonArray(jsonObject(
                                jsonProperty("Action", jsonArray("ec2:Describe*")),
                                jsonProperty("Effect", "Allow"),
                                jsonProperty("Resource", "*")
                            )))
                        )))
                    .build(),
                RoleInlinePolicyArgs.builder()
                    .name("policy-8675309")
                    .policy(inlinePolicy.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
                    .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${data.aws_iam_policy_document.instance_assume_role_policy.json} # (not shown)
      inlinePolicies:
        - name: my_inline_policy
          policy:
            fn::toJSON:
              Version: 2012-10-17
              Statement:
                - Action:
                    - ec2:Describe*
                  Effect: Allow
                  Resource: '*'
        - name: policy-8675309
          policy: ${inlinePolicy.json}
variables:
  inlinePolicy:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - ec2:DescribeAccountAttributes
            resources:
              - '*'
```
{{% /example %}}
{{% example %}}
### Example of Removing Inline Policies

This example creates an IAM role with what appears to be empty IAM `inline_policy` argument instead of using `inline_policy` as a configuration block. The result is that if someone were to add an inline policy out-of-band, on the next apply, this provider will remove that policy.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.iam.Role("example", {
    assumeRolePolicy: data.aws_iam_policy_document.instance_assume_role_policy.json,
    inlinePolicies: [{}],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.iam.Role("example",
    assume_role_policy=data["aws_iam_policy_document"]["instance_assume_role_policy"]["json"],
    inline_policies=[aws.iam.RoleInlinePolicyArgs()])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Iam.Role("example", new()
    {
        AssumeRolePolicy = data.Aws_iam_policy_document.Instance_assume_role_policy.Json,
        InlinePolicies = new[]
        {
            null,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.NewRole(ctx, "example", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.Any(data.Aws_iam_policy_document.Instance_assume_role_policy.Json),
			InlinePolicies: iam.RoleInlinePolicyArray{
				nil,
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
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.inputs.RoleInlinePolicyArgs;
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
        var example = new Role("example", RoleArgs.builder()        
            .assumeRolePolicy(data.aws_iam_policy_document().instance_assume_role_policy().json())
            .inlinePolicies()
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${data.aws_iam_policy_document.instance_assume_role_policy.json} # (not shown)
      inlinePolicies:
        - {}
```
{{% /example %}}
{{% example %}}
### Example of Exclusive Managed Policies

This example creates an IAM role and attaches two managed IAM policies. If someone attaches another managed policy out-of-band, on the next apply, this provider will detach that policy. If someone detaches these policies out-of-band, this provider will attach them again.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const policyOne = new aws.iam.Policy("policyOne", {policy: JSON.stringify({
    Version: "2012-10-17",
    Statement: [{
        Action: ["ec2:Describe*"],
        Effect: "Allow",
        Resource: "*",
    }],
})});
const policyTwo = new aws.iam.Policy("policyTwo", {policy: JSON.stringify({
    Version: "2012-10-17",
    Statement: [{
        Action: [
            "s3:ListAllMyBuckets",
            "s3:ListBucket",
            "s3:HeadBucket",
        ],
        Effect: "Allow",
        Resource: "*",
    }],
})});
const example = new aws.iam.Role("example", {
    assumeRolePolicy: data.aws_iam_policy_document.instance_assume_role_policy.json,
    managedPolicyArns: [
        policyOne.arn,
        policyTwo.arn,
    ],
});
```
```python
import pulumi
import json
import pulumi_aws as aws

policy_one = aws.iam.Policy("policyOne", policy=json.dumps({
    "Version": "2012-10-17",
    "Statement": [{
        "Action": ["ec2:Describe*"],
        "Effect": "Allow",
        "Resource": "*",
    }],
}))
policy_two = aws.iam.Policy("policyTwo", policy=json.dumps({
    "Version": "2012-10-17",
    "Statement": [{
        "Action": [
            "s3:ListAllMyBuckets",
            "s3:ListBucket",
            "s3:HeadBucket",
        ],
        "Effect": "Allow",
        "Resource": "*",
    }],
}))
example = aws.iam.Role("example",
    assume_role_policy=data["aws_iam_policy_document"]["instance_assume_role_policy"]["json"],
    managed_policy_arns=[
        policy_one.arn,
        policy_two.arn,
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var policyOne = new Aws.Iam.Policy("policyOne", new()
    {
        PolicyDocument = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = new[]
                    {
                        "ec2:Describe*",
                    },
                    ["Effect"] = "Allow",
                    ["Resource"] = "*",
                },
            },
        }),
    });

    var policyTwo = new Aws.Iam.Policy("policyTwo", new()
    {
        PolicyDocument = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = new[]
                    {
                        "s3:ListAllMyBuckets",
                        "s3:ListBucket",
                        "s3:HeadBucket",
                    },
                    ["Effect"] = "Allow",
                    ["Resource"] = "*",
                },
            },
        }),
    });

    var example = new Aws.Iam.Role("example", new()
    {
        AssumeRolePolicy = data.Aws_iam_policy_document.Instance_assume_role_policy.Json,
        ManagedPolicyArns = new[]
        {
            policyOne.Arn,
            policyTwo.Arn,
        },
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": []string{
						"ec2:Describe*",
					},
					"Effect":   "Allow",
					"Resource": "*",
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		policyOne, err := iam.NewPolicy(ctx, "policyOne", &iam.PolicyArgs{
			Policy: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		tmpJSON1, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": []string{
						"s3:ListAllMyBuckets",
						"s3:ListBucket",
						"s3:HeadBucket",
					},
					"Effect":   "Allow",
					"Resource": "*",
				},
			},
		})
		if err != nil {
			return err
		}
		json1 := string(tmpJSON1)
		policyTwo, err := iam.NewPolicy(ctx, "policyTwo", &iam.PolicyArgs{
			Policy: pulumi.String(json1),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRole(ctx, "example", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.Any(data.Aws_iam_policy_document.Instance_assume_role_policy.Json),
			ManagedPolicyArns: pulumi.StringArray{
				policyOne.Arn,
				policyTwo.Arn,
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
import com.pulumi.aws.iam.Policy;
import com.pulumi.aws.iam.PolicyArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
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
        var policyOne = new Policy("policyOne", PolicyArgs.builder()        
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", jsonArray("ec2:Describe*")),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Resource", "*")
                    )))
                )))
            .build());

        var policyTwo = new Policy("policyTwo", PolicyArgs.builder()        
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", jsonArray(
                            "s3:ListAllMyBuckets", 
                            "s3:ListBucket", 
                            "s3:HeadBucket"
                        )),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Resource", "*")
                    )))
                )))
            .build());

        var example = new Role("example", RoleArgs.builder()        
            .assumeRolePolicy(data.aws_iam_policy_document().instance_assume_role_policy().json())
            .managedPolicyArns(            
                policyOne.arn(),
                policyTwo.arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${data.aws_iam_policy_document.instance_assume_role_policy.json}
      # (not shown)
      managedPolicyArns:
        - ${policyOne.arn}
        - ${policyTwo.arn}
  policyOne:
    type: aws:iam:Policy
    properties:
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Action:
                - ec2:Describe*
              Effect: Allow
              Resource: '*'
  policyTwo:
    type: aws:iam:Policy
    properties:
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Action:
                - s3:ListAllMyBuckets
                - s3:ListBucket
                - s3:HeadBucket
              Effect: Allow
              Resource: '*'
```
{{% /example %}}
{{% example %}}
### Example of Removing Managed Policies

This example creates an IAM role with an empty `managed_policy_arns` argument. If someone attaches a policy out-of-band, on the next apply, this provider will detach that policy.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.iam.Role("example", {
    assumeRolePolicy: data.aws_iam_policy_document.instance_assume_role_policy.json,
    managedPolicyArns: [],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.iam.Role("example",
    assume_role_policy=data["aws_iam_policy_document"]["instance_assume_role_policy"]["json"],
    managed_policy_arns=[])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Iam.Role("example", new()
    {
        AssumeRolePolicy = data.Aws_iam_policy_document.Instance_assume_role_policy.Json,
        ManagedPolicyArns = new[] {},
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.NewRole(ctx, "example", &iam.RoleArgs{
			AssumeRolePolicy:  pulumi.Any(data.Aws_iam_policy_document.Instance_assume_role_policy.Json),
			ManagedPolicyArns: pulumi.StringArray{},
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
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
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
        var example = new Role("example", RoleArgs.builder()        
            .assumeRolePolicy(data.aws_iam_policy_document().instance_assume_role_policy().json())
            .managedPolicyArns()
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${data.aws_iam_policy_document.instance_assume_role_policy.json}
      # (not shown)
      managedPolicyArns: []
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IAM Roles using the `name`. For example:

```sh
 $ pulumi import aws:iam/role:Role developer developer_name
```
 