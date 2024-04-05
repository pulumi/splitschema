Provides details about a specific redshift cluster.

{{% examples %}}
## Example Usage
{{% example %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.redshift.RedshiftFunctions;
import com.pulumi.aws.redshift.inputs.GetClusterArgs;
import com.pulumi.aws.kinesis.FirehoseDeliveryStream;
import com.pulumi.aws.kinesis.FirehoseDeliveryStreamArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamRedshiftConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamRedshiftConfigurationS3ConfigurationArgs;
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
        final var example = RedshiftFunctions.getCluster(GetClusterArgs.builder()
            .clusterIdentifier("example-cluster")
            .build());

        var exampleStream = new FirehoseDeliveryStream("exampleStream", FirehoseDeliveryStreamArgs.builder()        
            .destination("redshift")
            .redshiftConfiguration(FirehoseDeliveryStreamRedshiftConfigurationArgs.builder()
                .roleArn(aws_iam_role.firehose_role().arn())
                .clusterJdbcurl(String.format("jdbc:redshift://%s/%s", example.applyValue(getClusterResult -> getClusterResult.endpoint()),example.applyValue(getClusterResult -> getClusterResult.databaseName())))
                .username("exampleuser")
                .password("Exampl3Pass")
                .dataTableName("example-table")
                .copyOptions("delimiter '|'")
                .dataTableColumns("example-col")
                .s3Configuration(FirehoseDeliveryStreamRedshiftConfigurationS3ConfigurationArgs.builder()
                    .roleArn(aws_iam_role.firehose_role().arn())
                    .bucketArn(aws_s3_bucket.bucket().arn())
                    .bufferSize(10)
                    .bufferInterval(400)
                    .compressionFormat("GZIP")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleStream:
    type: aws:kinesis:FirehoseDeliveryStream
    properties:
      destination: redshift
      redshiftConfiguration:
        roleArn: ${aws_iam_role.firehose_role.arn}
        clusterJdbcurl: jdbc:redshift://${example.endpoint}/${example.databaseName}
        username: exampleuser
        password: Exampl3Pass
        dataTableName: example-table
        copyOptions: delimiter '|'
        dataTableColumns: example-col
        s3Configuration:
          roleArn: ${aws_iam_role.firehose_role.arn}
          bucketArn: ${aws_s3_bucket.bucket.arn}
          bufferSize: 10
          bufferInterval: 400
          compressionFormat: GZIP
variables:
  example:
    fn::invoke:
      Function: aws:redshift:getCluster
      Arguments:
        clusterIdentifier: example-cluster
```
{{% /example %}}
{{% /examples %}}