Resource for managing an AWS Audit Manager Assessment Report.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.auditmanager.AssessmentReport("test", {assessmentId: aws_auditmanager_assessment.test.id});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.auditmanager.AssessmentReport("test", assessment_id=aws_auditmanager_assessment["test"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Auditmanager.AssessmentReport("test", new()
    {
        AssessmentId = aws_auditmanager_assessment.Test.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/auditmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := auditmanager.NewAssessmentReport(ctx, "test", &auditmanager.AssessmentReportArgs{
			AssessmentId: pulumi.Any(aws_auditmanager_assessment.Test.Id),
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
import com.pulumi.aws.auditmanager.AssessmentReport;
import com.pulumi.aws.auditmanager.AssessmentReportArgs;
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
        var test = new AssessmentReport("test", AssessmentReportArgs.builder()        
            .assessmentId(aws_auditmanager_assessment.test().id())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:auditmanager:AssessmentReport
    properties:
      assessmentId: ${aws_auditmanager_assessment.test.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Audit Manager Assessment Reports using the assessment report `id`. For example:

```sh
 $ pulumi import aws:auditmanager/assessmentReport:AssessmentReport example abc123-de45
```
 