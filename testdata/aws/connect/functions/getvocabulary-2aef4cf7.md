Provides details about a specific Amazon Connect Vocabulary.

{{% examples %}}
## Example Usage
{{% example %}}

By `name`

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.connect.getVocabulary({
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
    name: "Example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.get_vocabulary(instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
    name="Example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Connect.GetVocabulary.Invoke(new()
    {
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
        Name = "Example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.LookupVocabulary(ctx, &connect.LookupVocabularyArgs{
			InstanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
			Name:       pulumi.StringRef("Example"),
		}, nil)
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
import com.pulumi.aws.connect.ConnectFunctions;
import com.pulumi.aws.connect.inputs.GetVocabularyArgs;
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
        final var example = ConnectFunctions.getVocabulary(GetVocabularyArgs.builder()
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .name("Example")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:connect:getVocabulary
      Arguments:
        instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
        name: Example
```

By `vocabulary_id`

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.connect.getVocabulary({
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
    vocabularyId: "cccccccc-bbbb-cccc-dddd-111111111111",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.get_vocabulary(instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
    vocabulary_id="cccccccc-bbbb-cccc-dddd-111111111111")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Connect.GetVocabulary.Invoke(new()
    {
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
        VocabularyId = "cccccccc-bbbb-cccc-dddd-111111111111",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.LookupVocabulary(ctx, &connect.LookupVocabularyArgs{
			InstanceId:   "aaaaaaaa-bbbb-cccc-dddd-111111111111",
			VocabularyId: pulumi.StringRef("cccccccc-bbbb-cccc-dddd-111111111111"),
		}, nil)
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
import com.pulumi.aws.connect.ConnectFunctions;
import com.pulumi.aws.connect.inputs.GetVocabularyArgs;
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
        final var example = ConnectFunctions.getVocabulary(GetVocabularyArgs.builder()
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .vocabularyId("cccccccc-bbbb-cccc-dddd-111111111111")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:connect:getVocabulary
      Arguments:
        instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
        vocabularyId: cccccccc-bbbb-cccc-dddd-111111111111
```
{{% /example %}}
{{% /examples %}}