Provides an Athena data catalog.

More information about Athena and Athena data catalogs can be found in the [Athena User Guide](https://docs.aws.amazon.com/athena/latest/ug/what-is.html).

> **Tip:** for a more detailed explanation on the usage of `parameters`, see the [DataCatalog API documentation](https://docs.aws.amazon.com/athena/latest/APIReference/API_DataCatalog.html)

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.athena.DataCatalog("example", {
    description: "Example Athena data catalog",
    parameters: {
        "function": "arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function",
    },
    tags: {
        Name: "example-athena-data-catalog",
    },
    type: "LAMBDA",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.athena.DataCatalog("example",
    description="Example Athena data catalog",
    parameters={
        "function": "arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function",
    },
    tags={
        "Name": "example-athena-data-catalog",
    },
    type="LAMBDA")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Athena.DataCatalog("example", new()
    {
        Description = "Example Athena data catalog",
        Parameters = 
        {
            { "function", "arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function" },
        },
        Tags = 
        {
            { "Name", "example-athena-data-catalog" },
        },
        Type = "LAMBDA",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/athena"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := athena.NewDataCatalog(ctx, "example", &athena.DataCatalogArgs{
			Description: pulumi.String("Example Athena data catalog"),
			Parameters: pulumi.StringMap{
				"function": pulumi.String("arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function"),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example-athena-data-catalog"),
			},
			Type: pulumi.String("LAMBDA"),
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
import com.pulumi.aws.athena.DataCatalog;
import com.pulumi.aws.athena.DataCatalogArgs;
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
        var example = new DataCatalog("example", DataCatalogArgs.builder()        
            .description("Example Athena data catalog")
            .parameters(Map.of("function", "arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function"))
            .tags(Map.of("Name", "example-athena-data-catalog"))
            .type("LAMBDA")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:athena:DataCatalog
    properties:
      description: Example Athena data catalog
      parameters:
        function: arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function
      tags:
        Name: example-athena-data-catalog
      type: LAMBDA
```
{{% /example %}}
{{% example %}}
### Hive based Data Catalog

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.athena.DataCatalog("example", {
    description: "Hive based Data Catalog",
    parameters: {
        "metadata-function": "arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function",
    },
    type: "HIVE",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.athena.DataCatalog("example",
    description="Hive based Data Catalog",
    parameters={
        "metadata-function": "arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function",
    },
    type="HIVE")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Athena.DataCatalog("example", new()
    {
        Description = "Hive based Data Catalog",
        Parameters = 
        {
            { "metadata-function", "arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function" },
        },
        Type = "HIVE",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/athena"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := athena.NewDataCatalog(ctx, "example", &athena.DataCatalogArgs{
			Description: pulumi.String("Hive based Data Catalog"),
			Parameters: pulumi.StringMap{
				"metadata-function": pulumi.String("arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function"),
			},
			Type: pulumi.String("HIVE"),
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
import com.pulumi.aws.athena.DataCatalog;
import com.pulumi.aws.athena.DataCatalogArgs;
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
        var example = new DataCatalog("example", DataCatalogArgs.builder()        
            .description("Hive based Data Catalog")
            .parameters(Map.of("metadata-function", "arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function"))
            .type("HIVE")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:athena:DataCatalog
    properties:
      description: Hive based Data Catalog
      parameters:
        metadata-function: arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function
      type: HIVE
```
{{% /example %}}
{{% example %}}
### Glue based Data Catalog

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.athena.DataCatalog("example", {
    description: "Glue based Data Catalog",
    parameters: {
        "catalog-id": "123456789012",
    },
    type: "GLUE",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.athena.DataCatalog("example",
    description="Glue based Data Catalog",
    parameters={
        "catalog-id": "123456789012",
    },
    type="GLUE")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Athena.DataCatalog("example", new()
    {
        Description = "Glue based Data Catalog",
        Parameters = 
        {
            { "catalog-id", "123456789012" },
        },
        Type = "GLUE",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/athena"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := athena.NewDataCatalog(ctx, "example", &athena.DataCatalogArgs{
			Description: pulumi.String("Glue based Data Catalog"),
			Parameters: pulumi.StringMap{
				"catalog-id": pulumi.String("123456789012"),
			},
			Type: pulumi.String("GLUE"),
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
import com.pulumi.aws.athena.DataCatalog;
import com.pulumi.aws.athena.DataCatalogArgs;
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
        var example = new DataCatalog("example", DataCatalogArgs.builder()        
            .description("Glue based Data Catalog")
            .parameters(Map.of("catalog-id", "123456789012"))
            .type("GLUE")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:athena:DataCatalog
    properties:
      description: Glue based Data Catalog
      parameters:
        catalog-id: '123456789012'
      type: GLUE
```
{{% /example %}}
{{% example %}}
### Lambda based Data Catalog

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.athena.DataCatalog("example", {
    description: "Lambda based Data Catalog",
    parameters: {
        "metadata-function": "arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function-1",
        "record-function": "arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function-2",
    },
    type: "LAMBDA",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.athena.DataCatalog("example",
    description="Lambda based Data Catalog",
    parameters={
        "metadata-function": "arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function-1",
        "record-function": "arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function-2",
    },
    type="LAMBDA")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Athena.DataCatalog("example", new()
    {
        Description = "Lambda based Data Catalog",
        Parameters = 
        {
            { "metadata-function", "arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function-1" },
            { "record-function", "arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function-2" },
        },
        Type = "LAMBDA",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/athena"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := athena.NewDataCatalog(ctx, "example", &athena.DataCatalogArgs{
			Description: pulumi.String("Lambda based Data Catalog"),
			Parameters: pulumi.StringMap{
				"metadata-function": pulumi.String("arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function-1"),
				"record-function":   pulumi.String("arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function-2"),
			},
			Type: pulumi.String("LAMBDA"),
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
import com.pulumi.aws.athena.DataCatalog;
import com.pulumi.aws.athena.DataCatalogArgs;
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
        var example = new DataCatalog("example", DataCatalogArgs.builder()        
            .description("Lambda based Data Catalog")
            .parameters(Map.ofEntries(
                Map.entry("metadata-function", "arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function-1"),
                Map.entry("record-function", "arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function-2")
            ))
            .type("LAMBDA")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:athena:DataCatalog
    properties:
      description: Lambda based Data Catalog
      parameters:
        metadata-function: arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function-1
        record-function: arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function-2
      type: LAMBDA
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import data catalogs using their `name`. For example:

```sh
 $ pulumi import aws:athena/dataCatalog:DataCatalog example example-data-catalog
```
 