Provides details about a specific EC2 Customer-Owned IP Pool.

This data source can prove useful when a module accepts a coip pool id as
an input variable and needs to, for example, determine the CIDR block of that
COIP Pool.

{{% examples %}}
## Example Usage
{{% example %}}

The following example returns a specific coip pool ID

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetCoipPoolArgs;
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
        final var config = ctx.config();
        final var coipPoolId = config.get("coipPoolId");
        final var selected = Ec2Functions.getCoipPool(GetCoipPoolArgs.builder()
            .id(coipPoolId)
            .build());

    }
}
```
```yaml
configuration:
  coipPoolId:
    type: dynamic
variables:
  selected:
    fn::invoke:
      Function: aws:ec2:getCoipPool
      Arguments:
        id: ${coipPoolId}
```
{{% /example %}}
{{% /examples %}}