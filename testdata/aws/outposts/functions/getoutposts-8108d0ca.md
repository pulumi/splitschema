Provides details about multiple Outposts.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.outposts.getOutposts({
    siteId: data.aws_outposts_site.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.outposts.get_outposts(site_id=data["aws_outposts_site"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Outposts.GetOutposts.Invoke(new()
    {
        SiteId = data.Aws_outposts_site.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/outposts"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := outposts.GetOutposts(ctx, &outposts.GetOutpostsArgs{
			SiteId: pulumi.StringRef(data.Aws_outposts_site.Id),
		}, nil)
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
import com.pulumi.aws.outposts.OutpostsFunctions;
import com.pulumi.aws.outposts.inputs.GetOutpostsArgs;
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
        final var example = OutpostsFunctions.getOutposts(GetOutpostsArgs.builder()
            .siteId(data.aws_outposts_site().id())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:outposts:getOutposts
      Arguments:
        siteId: ${data.aws_outposts_site.id}
```
{{% /example %}}
{{% /examples %}}