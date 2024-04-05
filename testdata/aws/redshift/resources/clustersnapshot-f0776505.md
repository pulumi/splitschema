Creates a Redshift cluster snapshot

{{% examples %}}
## Example Usage
{{% example %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.redshift.ClusterSnapshot;
import com.pulumi.aws.redshift.ClusterSnapshotArgs;
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
        var example = new ClusterSnapshot("example", ClusterSnapshotArgs.builder()        
            .clusterSnapshotName("example")
            .clusterSnapshotContent(serializeJson(
                jsonObject(
                    jsonProperty("AllowDBUserOverride", "1"),
                    jsonProperty("Client_ID", "ExampleClientID"),
                    jsonProperty("App_ID", "example")
                )))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:redshift:ClusterSnapshot
    properties:
      clusterSnapshotName: example
      clusterSnapshotContent:
        fn::toJSON:
          AllowDBUserOverride: '1'
          Client_ID: ExampleClientID
          App_ID: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift Cluster Snapshots using `snapshot_identifier`. For example:

```sh
 $ pulumi import aws:redshift/clusterSnapshot:ClusterSnapshot test example
```
 