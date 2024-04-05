Use this data source to manage a response plan in AWS Systems Manager Incident Manager.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.ssmincidents.ResponsePlan;
import com.pulumi.aws.ssmincidents.ResponsePlanArgs;
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
        var example = new ResponsePlan("example", ResponsePlanArgs.builder()        
            .arn("exampleARN")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ssmincidents:ResponsePlan
    properties:
      arn: exampleARN
```
{{% /example %}}
{{% /examples %}}