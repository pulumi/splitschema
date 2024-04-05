Use this data source to get the Account ID of the [AWS Elastic Load Balancing Service Account](http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-access-logs.html#attach-bucket-policy)
in a given region for the purpose of permitting in S3 bucket policy.

> **Note:** For AWS Regions opened since Jakarta (`ap-southeast-3`) in December 2021, AWS [documents that](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-access-logs.html#attach-bucket-policy) a [service principal name](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html#principal-services) should be used instead of an AWS account ID in any relevant IAM policy.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = aws.elb.getServiceAccount({});
const elbLogs = new aws.s3.BucketV2("elbLogs", {});
const elbLogsAcl = new aws.s3.BucketAclV2("elbLogsAcl", {
    bucket: elbLogs.id,
    acl: "private",
});
const allowElbLoggingPolicyDocument = pulumi.all([main, elbLogs.arn]).apply(([main, arn]) => aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "AWS",
            identifiers: [main.arn],
        }],
        actions: ["s3:PutObject"],
        resources: [`${arn}/AWSLogs/*`],
    }],
}));
const allowElbLoggingBucketPolicy = new aws.s3.BucketPolicy("allowElbLoggingBucketPolicy", {
    bucket: elbLogs.id,
    policy: allowElbLoggingPolicyDocument.apply(allowElbLoggingPolicyDocument => allowElbLoggingPolicyDocument.json),
});
const bar = new aws.elb.LoadBalancer("bar", {
    availabilityZones: ["us-west-2a"],
    accessLogs: {
        bucket: elbLogs.id,
        interval: 5,
    },
    listeners: [{
        instancePort: 8000,
        instanceProtocol: "http",
        lbPort: 80,
        lbProtocol: "http",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.elb.get_service_account()
elb_logs = aws.s3.BucketV2("elbLogs")
elb_logs_acl = aws.s3.BucketAclV2("elbLogsAcl",
    bucket=elb_logs.id,
    acl="private")
allow_elb_logging_policy_document = elb_logs.arn.apply(lambda arn: aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="AWS",
        identifiers=[main.arn],
    )],
    actions=["s3:PutObject"],
    resources=[f"{arn}/AWSLogs/*"],
)]))
allow_elb_logging_bucket_policy = aws.s3.BucketPolicy("allowElbLoggingBucketPolicy",
    bucket=elb_logs.id,
    policy=allow_elb_logging_policy_document.json)
bar = aws.elb.LoadBalancer("bar",
    availability_zones=["us-west-2a"],
    access_logs=aws.elb.LoadBalancerAccessLogsArgs(
        bucket=elb_logs.id,
        interval=5,
    ),
    listeners=[aws.elb.LoadBalancerListenerArgs(
        instance_port=8000,
        instance_protocol="http",
        lb_port=80,
        lb_protocol="http",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = Aws.Elb.GetServiceAccount.Invoke();

    var elbLogs = new Aws.S3.BucketV2("elbLogs");

    var elbLogsAcl = new Aws.S3.BucketAclV2("elbLogsAcl", new()
    {
        Bucket = elbLogs.Id,
        Acl = "private",
    });

    var allowElbLoggingPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                        Type = "AWS",
                        Identifiers = new[]
                        {
                            main.Apply(getServiceAccountResult => getServiceAccountResult.Arn),
                        },
                    },
                },
                Actions = new[]
                {
                    "s3:PutObject",
                },
                Resources = new[]
                {
                    $"{elbLogs.Arn}/AWSLogs/*",
                },
            },
        },
    });

    var allowElbLoggingBucketPolicy = new Aws.S3.BucketPolicy("allowElbLoggingBucketPolicy", new()
    {
        Bucket = elbLogs.Id,
        Policy = allowElbLoggingPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var bar = new Aws.Elb.LoadBalancer("bar", new()
    {
        AvailabilityZones = new[]
        {
            "us-west-2a",
        },
        AccessLogs = new Aws.Elb.Inputs.LoadBalancerAccessLogsArgs
        {
            Bucket = elbLogs.Id,
            Interval = 5,
        },
        Listeners = new[]
        {
            new Aws.Elb.Inputs.LoadBalancerListenerArgs
            {
                InstancePort = 8000,
                InstanceProtocol = "http",
                LbPort = 80,
                LbProtocol = "http",
            },
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elb"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
main, err := elb.GetServiceAccount(ctx, nil, nil);
if err != nil {
return err
}
elbLogs, err := s3.NewBucketV2(ctx, "elbLogs", nil)
if err != nil {
return err
}
_, err = s3.NewBucketAclV2(ctx, "elbLogsAcl", &s3.BucketAclV2Args{
Bucket: elbLogs.ID(),
Acl: pulumi.String("private"),
})
if err != nil {
return err
}
allowElbLoggingPolicyDocument := elbLogs.Arn.ApplyT(func(arn string) (iam.GetPolicyDocumentResult, error) {
return iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
Statements: []iam.GetPolicyDocumentStatement{
{
Effect: "Allow",
Principals: []iam.GetPolicyDocumentStatementPrincipal{
{
Type: "AWS",
Identifiers: interface{}{
main.Arn,
},
},
},
Actions: []string{
"s3:PutObject",
},
Resources: []string{
fmt.Sprintf("%v/AWSLogs/*", arn),
},
},
},
}, nil), nil
}).(iam.GetPolicyDocumentResultOutput)
_, err = s3.NewBucketPolicy(ctx, "allowElbLoggingBucketPolicy", &s3.BucketPolicyArgs{
Bucket: elbLogs.ID(),
Policy: allowElbLoggingPolicyDocument.ApplyT(func(allowElbLoggingPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
return &allowElbLoggingPolicyDocument.Json, nil
}).(pulumi.StringPtrOutput),
})
if err != nil {
return err
}
_, err = elb.NewLoadBalancer(ctx, "bar", &elb.LoadBalancerArgs{
AvailabilityZones: pulumi.StringArray{
pulumi.String("us-west-2a"),
},
AccessLogs: &elb.LoadBalancerAccessLogsArgs{
Bucket: elbLogs.ID(),
Interval: pulumi.Int(5),
},
Listeners: elb.LoadBalancerListenerArray{
&elb.LoadBalancerListenerArgs{
InstancePort: pulumi.Int(8000),
InstanceProtocol: pulumi.String("http"),
LbPort: pulumi.Int(80),
LbProtocol: pulumi.String("http"),
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
import com.pulumi.aws.elb.ElbFunctions;
import com.pulumi.aws.elb.inputs.GetServiceAccountArgs;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.s3.BucketPolicy;
import com.pulumi.aws.s3.BucketPolicyArgs;
import com.pulumi.aws.elb.LoadBalancer;
import com.pulumi.aws.elb.LoadBalancerArgs;
import com.pulumi.aws.elb.inputs.LoadBalancerAccessLogsArgs;
import com.pulumi.aws.elb.inputs.LoadBalancerListenerArgs;
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
        final var main = ElbFunctions.getServiceAccount();

        var elbLogs = new BucketV2("elbLogs");

        var elbLogsAcl = new BucketAclV2("elbLogsAcl", BucketAclV2Args.builder()        
            .bucket(elbLogs.id())
            .acl("private")
            .build());

        final var allowElbLoggingPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("AWS")
                    .identifiers(main.applyValue(getServiceAccountResult -> getServiceAccountResult.arn()))
                    .build())
                .actions("s3:PutObject")
                .resources(elbLogs.arn().applyValue(arn -> String.format("%s/AWSLogs/*", arn)))
                .build())
            .build());

        var allowElbLoggingBucketPolicy = new BucketPolicy("allowElbLoggingBucketPolicy", BucketPolicyArgs.builder()        
            .bucket(elbLogs.id())
            .policy(allowElbLoggingPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(allowElbLoggingPolicyDocument -> allowElbLoggingPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

        var bar = new LoadBalancer("bar", LoadBalancerArgs.builder()        
            .availabilityZones("us-west-2a")
            .accessLogs(LoadBalancerAccessLogsArgs.builder()
                .bucket(elbLogs.id())
                .interval(5)
                .build())
            .listeners(LoadBalancerListenerArgs.builder()
                .instancePort(8000)
                .instanceProtocol("http")
                .lbPort(80)
                .lbProtocol("http")
                .build())
            .build());

    }
}
```
```yaml
resources:
  elbLogs:
    type: aws:s3:BucketV2
  elbLogsAcl:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${elbLogs.id}
      acl: private
  allowElbLoggingBucketPolicy:
    type: aws:s3:BucketPolicy
    properties:
      bucket: ${elbLogs.id}
      policy: ${allowElbLoggingPolicyDocument.json}
  bar:
    type: aws:elb:LoadBalancer
    properties:
      availabilityZones:
        - us-west-2a
      accessLogs:
        bucket: ${elbLogs.id}
        interval: 5
      listeners:
        - instancePort: 8000
          instanceProtocol: http
          lbPort: 80
          lbProtocol: http
variables:
  main:
    fn::invoke:
      Function: aws:elb:getServiceAccount
      Arguments: {}
  allowElbLoggingPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: AWS
                identifiers:
                  - ${main.arn}
            actions:
              - s3:PutObject
            resources:
              - ${elbLogs.arn}/AWSLogs/*
```
{{% /example %}}
{{% /examples %}}