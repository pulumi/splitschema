Data source for managing an AWS Polly Voices.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.polly.getVoices({});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.polly.get_voices()
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Polly.GetVoices.Invoke();

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/polly"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := polly.GetVoices(ctx, nil, nil)
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
import com.pulumi.aws.polly.PollyFunctions;
import com.pulumi.aws.polly.inputs.GetVoicesArgs;
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
        final var example = PollyFunctions.getVoices();

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:polly:getVoices
      Arguments: {}
```
{{% /example %}}
{{% example %}}
### With Language Code

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.polly.getVoices({
    languageCode: "en-GB",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.polly.get_voices(language_code="en-GB")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Polly.GetVoices.Invoke(new()
    {
        LanguageCode = "en-GB",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/polly"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := polly.GetVoices(ctx, &polly.GetVoicesArgs{
			LanguageCode: pulumi.StringRef("en-GB"),
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
import com.pulumi.aws.polly.PollyFunctions;
import com.pulumi.aws.polly.inputs.GetVoicesArgs;
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
        final var example = PollyFunctions.getVoices(GetVoicesArgs.builder()
            .languageCode("en-GB")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:polly:getVoices
      Arguments:
        languageCode: en-GB
```
{{% /example %}}
{{% /examples %}}