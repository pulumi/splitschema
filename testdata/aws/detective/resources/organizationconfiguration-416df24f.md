{{% examples %}}
## Example Usage
{{% example %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.detective.Graph;
import com.pulumi.aws.detective.GraphArgs;
import com.pulumi.aws.detective.OrganizationConfiguration;
import com.pulumi.aws.detective.OrganizationConfigurationArgs;
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
        var exampleGraph = new Graph("exampleGraph", GraphArgs.builder()        
            .enable(true)
            .build());

        var exampleOrganizationConfiguration = new OrganizationConfiguration("exampleOrganizationConfiguration", OrganizationConfigurationArgs.builder()        
            .autoEnable(true)
            .graphArn(exampleGraph.id())
            .build());

    }
}
```
```yaml
resources:
  exampleGraph:
    type: aws:detective:Graph
    properties:
      enable: true
  exampleOrganizationConfiguration:
    type: aws:detective:OrganizationConfiguration
    properties:
      autoEnable: true
      graphArn: ${exampleGraph.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_detective_organization_admin_account` using the Detective Graph ID. For example:

```sh
 $ pulumi import aws:detective/organizationConfiguration:OrganizationConfiguration example 00b00fd5aecc0ab60a708659477e9617
```
 