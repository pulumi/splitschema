The App Mesh Route data source allows details of an App Mesh Route to be retrieved by its name, mesh_name, virtual_router_name, and optionally the mesh_owner.

{{% examples %}}
## Example Usage
{{% example %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.appmesh.AppmeshFunctions;
import com.pulumi.aws.appmesh.inputs.GetVirtualServiceArgs;
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
        final var test = AppmeshFunctions.getVirtualService(GetVirtualServiceArgs.builder()
            .meshName("test-mesh")
            .name("test-route")
            .virtualRouterName("test-router")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:appmesh:getVirtualService
      Arguments:
        meshName: test-mesh
        name: test-route
        virtualRouterName: test-router
```
{{% /example %}}
{{% /examples %}}