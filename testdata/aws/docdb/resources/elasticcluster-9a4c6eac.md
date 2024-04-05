Manages an AWS DocDB (DocumentDB) Elastic Cluster.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.docdb.ElasticCluster;
import com.pulumi.aws.docdb.ElasticClusterArgs;
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
        var example = new ElasticCluster("example", ElasticClusterArgs.builder()        
            .adminUserName("foo")
            .adminUserPassword("mustbeeightchars")
            .authType("PLAIN_TEXT")
            .clusterName("my-docdb-cluster")
            .shardCapacity(2)
            .shardCount(1)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:docdb:ElasticCluster
    properties:
      adminUserName: foo
      adminUserPassword: mustbeeightchars
      authType: PLAIN_TEXT
      clusterName: my-docdb-cluster
      shardCapacity: 2
      shardCount: 1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DocDB (DocumentDB) Elastic Cluster using the `arn` argument. For example,

```sh
 $ pulumi import aws:docdb/elasticCluster:ElasticCluster example arn:aws:docdb-elastic:us-east-1:000011112222:cluster/12345678-7abc-def0-1234-56789abcdef
```
 