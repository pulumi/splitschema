Resource for managing an AWS QuickSight Folder Membership.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.FolderMembership("example", {
    folderId: aws_quicksight_folder.example.folder_id,
    memberType: "DATASET",
    memberId: aws_quicksight_data_set.example.data_set_id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.quicksight.FolderMembership("example",
    folder_id=aws_quicksight_folder["example"]["folder_id"],
    member_type="DATASET",
    member_id=aws_quicksight_data_set["example"]["data_set_id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Quicksight.FolderMembership("example", new()
    {
        FolderId = aws_quicksight_folder.Example.Folder_id,
        MemberType = "DATASET",
        MemberId = aws_quicksight_data_set.Example.Data_set_id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.NewFolderMembership(ctx, "example", &quicksight.FolderMembershipArgs{
			FolderId:   pulumi.Any(aws_quicksight_folder.Example.Folder_id),
			MemberType: pulumi.String("DATASET"),
			MemberId:   pulumi.Any(aws_quicksight_data_set.Example.Data_set_id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.quicksight.FolderMembership;
import com.pulumi.aws.quicksight.FolderMembershipArgs;
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
        var example = new FolderMembership("example", FolderMembershipArgs.builder()        
            .folderId(aws_quicksight_folder.example().folder_id())
            .memberType("DATASET")
            .memberId(aws_quicksight_data_set.example().data_set_id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:FolderMembership
    properties:
      folderId: ${aws_quicksight_folder.example.folder_id}
      memberType: DATASET
      memberId: ${aws_quicksight_data_set.example.data_set_id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import QuickSight Folder Membership using the AWS account ID, folder ID, member type, and member ID separated by commas (`,`). For example:

```sh
 $ pulumi import aws:quicksight/folderMembership:FolderMembership example 123456789012,example-folder,DATASET,example-dataset
```
 