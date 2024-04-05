Resource for managing an AWS FinSpace Kx Cluster.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.finspace.KxCluster;
import com.pulumi.aws.finspace.KxClusterArgs;
import com.pulumi.aws.finspace.inputs.KxClusterCapacityConfigurationArgs;
import com.pulumi.aws.finspace.inputs.KxClusterVpcConfigurationArgs;
import com.pulumi.aws.finspace.inputs.KxClusterCacheStorageConfigurationArgs;
import com.pulumi.aws.finspace.inputs.KxClusterDatabaseArgs;
import com.pulumi.aws.finspace.inputs.KxClusterCodeArgs;
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
        var example = new KxCluster("example", KxClusterArgs.builder()        
            .environmentId(aws_finspace_kx_environment.example().id())
            .type("HDB")
            .releaseLabel("1.0")
            .azMode("SINGLE")
            .availabilityZoneId("use1-az2")
            .capacityConfiguration(KxClusterCapacityConfigurationArgs.builder()
                .nodeType("kx.s.2xlarge")
                .nodeCount(2)
                .build())
            .vpcConfiguration(KxClusterVpcConfigurationArgs.builder()
                .vpcId(aws_vpc.test().id())
                .securityGroupIds(aws_security_group.example().id())
                .subnetIds(aws_subnet.example().id())
                .ipAddressType("IP_V4")
                .build())
            .cacheStorageConfigurations(KxClusterCacheStorageConfigurationArgs.builder()
                .type("CACHE_1000")
                .size(1200)
                .build())
            .databases(KxClusterDatabaseArgs.builder()
                .databaseName(aws_finspace_kx_database.example().name())
                .cacheConfiguration(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
                .build())
            .code(KxClusterCodeArgs.builder()
                .s3Bucket(aws_s3_bucket.test().id())
                .s3Key(aws_s3_object.object().key())
                .build())
            .timeouts(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:finspace:KxCluster
    properties:
      environmentId: ${aws_finspace_kx_environment.example.id}
      type: HDB
      releaseLabel: '1.0'
      azMode: SINGLE
      availabilityZoneId: use1-az2
      capacityConfiguration:
        nodeType: kx.s.2xlarge
        nodeCount: 2
      vpcConfiguration:
        vpcId: ${aws_vpc.test.id}
        securityGroupIds:
          - ${aws_security_group.example.id}
        subnetIds:
          - ${aws_subnet.example.id}
        ipAddressType: IP_V4
      cacheStorageConfigurations:
        - type: CACHE_1000
          size: 1200
      databases:
        - databaseName: ${aws_finspace_kx_database.example.name}
          cacheConfiguration:
            - cacheType: CACHE_1000
              dbPaths: /
      code:
        s3Bucket: ${aws_s3_bucket.test.id}
        s3Key: ${aws_s3_object.object.key}
      # Depending on the amount of data cached, create/update timeouts 
      #   # may need to be increased up to a potential maximum of 18 hours.
      timeouts:
        - create: 18h
          update: 18h
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import an AWS FinSpace Kx Cluster using the `id` (environment ID and cluster name, comma-delimited). For example:

```sh
 $ pulumi import aws:finspace/kxCluster:KxCluster example n3ceo7wqxoxcti5tujqwzs,my-tf-kx-cluster
```
 