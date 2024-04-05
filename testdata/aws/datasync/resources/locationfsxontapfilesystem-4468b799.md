Resource for managing an AWS DataSync Location FSx Ontap File System.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.datasync.LocationFsxOntapFileSystem;
import com.pulumi.aws.datasync.LocationFsxOntapFileSystemArgs;
import com.pulumi.aws.datasync.inputs.LocationFsxOntapFileSystemProtocolArgs;
import com.pulumi.aws.datasync.inputs.LocationFsxOntapFileSystemProtocolNfsArgs;
import com.pulumi.aws.datasync.inputs.LocationFsxOntapFileSystemProtocolNfsMountOptionsArgs;
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
        var test = new LocationFsxOntapFileSystem("test", LocationFsxOntapFileSystemArgs.builder()        
            .fsxFilesystemArn(aws_fsx_ontap_file_system.test().arn())
            .securityGroupArns(aws_security_group.test().arn())
            .storageVirtualMachineArn(aws_fsx_ontap_storage_virtual_machine.test().arn())
            .protocol(LocationFsxOntapFileSystemProtocolArgs.builder()
                .nfs(LocationFsxOntapFileSystemProtocolNfsArgs.builder()
                    .mountOptions(LocationFsxOntapFileSystemProtocolNfsMountOptionsArgs.builder()
                        .version("NFS3")
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:datasync:LocationFsxOntapFileSystem
    properties:
      fsxFilesystemArn: ${aws_fsx_ontap_file_system.test.arn}
      securityGroupArns:
        - ${aws_security_group.test.arn}
      storageVirtualMachineArn: ${aws_fsx_ontap_storage_virtual_machine.test.arn}
      protocol:
        nfs:
          mountOptions:
            version: NFS3
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_datasync_location_fsx_ontap_file_system` using the `DataSync-ARN#FSx-ontap-svm-ARN`. For example:

```sh
 $ pulumi import aws:datasync/locationFsxOntapFileSystem:LocationFsxOntapFileSystem example arn:aws:datasync:us-west-2:123456789012:location/loc-12345678901234567#arn:aws:fsx:us-west-2:123456789012:storage-virtual-machine/svm-12345678abcdef123
```
 