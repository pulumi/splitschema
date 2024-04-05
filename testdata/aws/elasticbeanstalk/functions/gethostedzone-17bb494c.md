Use this data source to get the ID of an [elastic beanstalk hosted zone](http://docs.aws.amazon.com/general/latest/gr/rande.html#elasticbeanstalk_region).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.elasticbeanstalk.getHostedZone({});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.elasticbeanstalk.get_hosted_zone()
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.ElasticBeanstalk.GetHostedZone.Invoke();

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticbeanstalk"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticbeanstalk.GetHostedZone(ctx, nil, nil)
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
import com.pulumi.aws.elasticbeanstalk.ElasticbeanstalkFunctions;
import com.pulumi.aws.elasticbeanstalk.inputs.GetHostedZoneArgs;
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
        final var current = ElasticbeanstalkFunctions.getHostedZone();

    }
}
```
```yaml
variables:
  current:
    fn::invoke:
      Function: aws:elasticbeanstalk:getHostedZone
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}