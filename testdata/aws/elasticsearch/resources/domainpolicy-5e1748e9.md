Allows setting policy to an Elasticsearch domain while referencing domain attributes (e.g., ARN)

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.elasticsearch.Domain("example", {elasticsearchVersion: "2.3"});
const main = new aws.elasticsearch.DomainPolicy("main", {
    domainName: example.domainName,
    accessPolicies: pulumi.interpolate`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "es:*",
            "Principal": "*",
            "Effect": "Allow",
            "Condition": {
                "IpAddress": {"aws:SourceIp": "127.0.0.1/32"}
            },
            "Resource": "${example.arn}/*"
        }
    ]
}
`,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.elasticsearch.Domain("example", elasticsearch_version="2.3")
main = aws.elasticsearch.DomainPolicy("main",
    domain_name=example.domain_name,
    access_policies=example.arn.apply(lambda arn: f"""{{
    "Version": "2012-10-17",
    "Statement": [
        {{
            "Action": "es:*",
            "Principal": "*",
            "Effect": "Allow",
            "Condition": {{
                "IpAddress": {{"aws:SourceIp": "127.0.0.1/32"}}
            }},
            "Resource": "{arn}/*"
        }}
    ]
}}
"""))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ElasticSearch.Domain("example", new()
    {
        ElasticsearchVersion = "2.3",
    });

    var main = new Aws.ElasticSearch.DomainPolicy("main", new()
    {
        DomainName = example.DomainName,
        AccessPolicies = example.Arn.Apply(arn => @$"{{
    ""Version"": ""2012-10-17"",
    ""Statement"": [
        {{
            ""Action"": ""es:*"",
            ""Principal"": ""*"",
            ""Effect"": ""Allow"",
            ""Condition"": {{
                ""IpAddress"": {{""aws:SourceIp"": ""127.0.0.1/32""}}
            }},
            ""Resource"": ""{arn}/*""
        }}
    ]
}}
"),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticsearch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := elasticsearch.NewDomain(ctx, "example", &elasticsearch.DomainArgs{
			ElasticsearchVersion: pulumi.String("2.3"),
		})
		if err != nil {
			return err
		}
		_, err = elasticsearch.NewDomainPolicy(ctx, "main", &elasticsearch.DomainPolicyArgs{
			DomainName: example.DomainName,
			AccessPolicies: example.Arn.ApplyT(func(arn string) (string, error) {
				return fmt.Sprintf(`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "es:*",
            "Principal": "*",
            "Effect": "Allow",
            "Condition": {
                "IpAddress": {"aws:SourceIp": "127.0.0.1/32"}
            },
            "Resource": "%v/*"
        }
    ]
}
`, arn), nil
			}).(pulumi.StringOutput),
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
import com.pulumi.aws.elasticsearch.Domain;
import com.pulumi.aws.elasticsearch.DomainArgs;
import com.pulumi.aws.elasticsearch.DomainPolicy;
import com.pulumi.aws.elasticsearch.DomainPolicyArgs;
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
        var example = new Domain("example", DomainArgs.builder()        
            .elasticsearchVersion("2.3")
            .build());

        var main = new DomainPolicy("main", DomainPolicyArgs.builder()        
            .domainName(example.domainName())
            .accessPolicies(example.arn().applyValue(arn -> """
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "es:*",
            "Principal": "*",
            "Effect": "Allow",
            "Condition": {
                "IpAddress": {"aws:SourceIp": "127.0.0.1/32"}
            },
            "Resource": "%s/*"
        }
    ]
}
", arn)))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:elasticsearch:Domain
    properties:
      elasticsearchVersion: '2.3'
  main:
    type: aws:elasticsearch:DomainPolicy
    properties:
      domainName: ${example.domainName}
      accessPolicies: |
        {
            "Version": "2012-10-17",
            "Statement": [
                {
                    "Action": "es:*",
                    "Principal": "*",
                    "Effect": "Allow",
                    "Condition": {
                        "IpAddress": {"aws:SourceIp": "127.0.0.1/32"}
                    },
                    "Resource": "${example.arn}/*"
                }
            ]
        }
```
{{% /example %}}
{{% /examples %}}