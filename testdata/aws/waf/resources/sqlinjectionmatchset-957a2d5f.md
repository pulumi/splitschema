Provides a WAF SQL Injection Match Set Resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const sqlInjectionMatchSet = new aws.waf.SqlInjectionMatchSet("sqlInjectionMatchSet", {sqlInjectionMatchTuples: [{
    fieldToMatch: {
        type: "QUERY_STRING",
    },
    textTransformation: "URL_DECODE",
}]});
```
```python
import pulumi
import pulumi_aws as aws

sql_injection_match_set = aws.waf.SqlInjectionMatchSet("sqlInjectionMatchSet", sql_injection_match_tuples=[aws.waf.SqlInjectionMatchSetSqlInjectionMatchTupleArgs(
    field_to_match=aws.waf.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatchArgs(
        type="QUERY_STRING",
    ),
    text_transformation="URL_DECODE",
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var sqlInjectionMatchSet = new Aws.Waf.SqlInjectionMatchSet("sqlInjectionMatchSet", new()
    {
        SqlInjectionMatchTuples = new[]
        {
            new Aws.Waf.Inputs.SqlInjectionMatchSetSqlInjectionMatchTupleArgs
            {
                FieldToMatch = new Aws.Waf.Inputs.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatchArgs
                {
                    Type = "QUERY_STRING",
                },
                TextTransformation = "URL_DECODE",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/waf"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := waf.NewSqlInjectionMatchSet(ctx, "sqlInjectionMatchSet", &waf.SqlInjectionMatchSetArgs{
			SqlInjectionMatchTuples: waf.SqlInjectionMatchSetSqlInjectionMatchTupleArray{
				&waf.SqlInjectionMatchSetSqlInjectionMatchTupleArgs{
					FieldToMatch: &waf.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatchArgs{
						Type: pulumi.String("QUERY_STRING"),
					},
					TextTransformation: pulumi.String("URL_DECODE"),
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
import com.pulumi.aws.waf.SqlInjectionMatchSet;
import com.pulumi.aws.waf.SqlInjectionMatchSetArgs;
import com.pulumi.aws.waf.inputs.SqlInjectionMatchSetSqlInjectionMatchTupleArgs;
import com.pulumi.aws.waf.inputs.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatchArgs;
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
        var sqlInjectionMatchSet = new SqlInjectionMatchSet("sqlInjectionMatchSet", SqlInjectionMatchSetArgs.builder()        
            .sqlInjectionMatchTuples(SqlInjectionMatchSetSqlInjectionMatchTupleArgs.builder()
                .fieldToMatch(SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatchArgs.builder()
                    .type("QUERY_STRING")
                    .build())
                .textTransformation("URL_DECODE")
                .build())
            .build());

    }
}
```
```yaml
resources:
  sqlInjectionMatchSet:
    type: aws:waf:SqlInjectionMatchSet
    properties:
      sqlInjectionMatchTuples:
        - fieldToMatch:
            type: QUERY_STRING
          textTransformation: URL_DECODE
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AWS WAF SQL Injection Match Set using their ID. For example:

```sh
 $ pulumi import aws:waf/sqlInjectionMatchSet:SqlInjectionMatchSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
```
 