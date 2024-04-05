Provides a Glue Partition Resource.

{{% examples %}}
## Example Usage
{{% example %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.glue.Partition;
import com.pulumi.aws.glue.PartitionArgs;
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
        var example = new Partition("example", PartitionArgs.builder()        
            .databaseName("some-database")
            .tableName("some-table")
            .values("some-value")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Partition
    properties:
      databaseName: some-database
      tableName: some-table
      values:
        - some-value
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Glue Partitions using the catalog ID (usually AWS account ID), database name, table name and partition values. For example:

```sh
 $ pulumi import aws:glue/partition:Partition part 123456789012:MyDatabase:MyTable:val1#val2
```
 