Provides an SSM Document resource

> **NOTE on updating SSM documents:** Only documents with a schema version of 2.0
or greater can update their content once created, see [SSM Schema Features](http://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-ssm-docs.html#document-schemas-features). To update a document with an older schema version you must recreate the resource. Not all document types support a schema version of 2.0 or greater. Refer to [SSM document schema features and examples](https://docs.aws.amazon.com/systems-manager/latest/userguide/document-schemas-features.html) for information about which schema versions are supported for the respective `document_type`.

{{% examples %}}
## Example Usage
{{% example %}}
### Create an ssm document in JSON format

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.ssm.Document("foo", {
    content: `  {
    "schemaVersion": "1.2",
    "description": "Check ip configuration of a Linux instance.",
    "parameters": {

    },
    "runtimeConfig": {
      "aws:runShellScript": {
        "properties": [
          {
            "id": "0.aws:runShellScript",
            "runCommand": ["ifconfig"]
          }
        ]
      }
    }
  }

`,
    documentType: "Command",
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.ssm.Document("foo",
    content="""  {
    "schemaVersion": "1.2",
    "description": "Check ip configuration of a Linux instance.",
    "parameters": {

    },
    "runtimeConfig": {
      "aws:runShellScript": {
        "properties": [
          {
            "id": "0.aws:runShellScript",
            "runCommand": ["ifconfig"]
          }
        ]
      }
    }
  }

""",
    document_type="Command")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Ssm.Document("foo", new()
    {
        Content = @"  {
    ""schemaVersion"": ""1.2"",
    ""description"": ""Check ip configuration of a Linux instance."",
    ""parameters"": {

    },
    ""runtimeConfig"": {
      ""aws:runShellScript"": {
        ""properties"": [
          {
            ""id"": ""0.aws:runShellScript"",
            ""runCommand"": [""ifconfig""]
          }
        ]
      }
    }
  }

",
        DocumentType = "Command",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.NewDocument(ctx, "foo", &ssm.DocumentArgs{
			Content: pulumi.String(`  {
    "schemaVersion": "1.2",
    "description": "Check ip configuration of a Linux instance.",
    "parameters": {

    },
    "runtimeConfig": {
      "aws:runShellScript": {
        "properties": [
          {
            "id": "0.aws:runShellScript",
            "runCommand": ["ifconfig"]
          }
        ]
      }
    }
  }

`),
			DocumentType: pulumi.String("Command"),
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
import com.pulumi.aws.ssm.Document;
import com.pulumi.aws.ssm.DocumentArgs;
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
        var foo = new Document("foo", DocumentArgs.builder()        
            .content("""
  {
    "schemaVersion": "1.2",
    "description": "Check ip configuration of a Linux instance.",
    "parameters": {

    },
    "runtimeConfig": {
      "aws:runShellScript": {
        "properties": [
          {
            "id": "0.aws:runShellScript",
            "runCommand": ["ifconfig"]
          }
        ]
      }
    }
  }

            """)
            .documentType("Command")
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:ssm:Document
    properties:
      content: |2+
          {
            "schemaVersion": "1.2",
            "description": "Check ip configuration of a Linux instance.",
            "parameters": {

            },
            "runtimeConfig": {
              "aws:runShellScript": {
                "properties": [
                  {
                    "id": "0.aws:runShellScript",
                    "runCommand": ["ifconfig"]
                  }
                ]
              }
            }
          }

      documentType: Command
```
{{% /example %}}
{{% example %}}
### Create an ssm document in YAML format

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.ssm.Document("foo", {
    content: `schemaVersion: '1.2'
description: Check ip configuration of a Linux instance.
parameters: {}
runtimeConfig:
  'aws:runShellScript':
    properties:
      - id: '0.aws:runShellScript'
        runCommand:
          - ifconfig

`,
    documentFormat: "YAML",
    documentType: "Command",
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.ssm.Document("foo",
    content="""schemaVersion: '1.2'
description: Check ip configuration of a Linux instance.
parameters: {}
runtimeConfig:
  'aws:runShellScript':
    properties:
      - id: '0.aws:runShellScript'
        runCommand:
          - ifconfig

""",
    document_format="YAML",
    document_type="Command")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Ssm.Document("foo", new()
    {
        Content = @"schemaVersion: '1.2'
description: Check ip configuration of a Linux instance.
parameters: {}
runtimeConfig:
  'aws:runShellScript':
    properties:
      - id: '0.aws:runShellScript'
        runCommand:
          - ifconfig

",
        DocumentFormat = "YAML",
        DocumentType = "Command",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.NewDocument(ctx, "foo", &ssm.DocumentArgs{
			Content: pulumi.String(`schemaVersion: '1.2'
description: Check ip configuration of a Linux instance.
parameters: {}
runtimeConfig:
  'aws:runShellScript':
    properties:
      - id: '0.aws:runShellScript'
        runCommand:
          - ifconfig

`),
			DocumentFormat: pulumi.String("YAML"),
			DocumentType:   pulumi.String("Command"),
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
import com.pulumi.aws.ssm.Document;
import com.pulumi.aws.ssm.DocumentArgs;
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
        var foo = new Document("foo", DocumentArgs.builder()        
            .content("""
schemaVersion: '1.2'
description: Check ip configuration of a Linux instance.
parameters: {}
runtimeConfig:
  'aws:runShellScript':
    properties:
      - id: '0.aws:runShellScript'
        runCommand:
          - ifconfig

            """)
            .documentFormat("YAML")
            .documentType("Command")
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:ssm:Document
    properties:
      content: |+
        schemaVersion: '1.2'
        description: Check ip configuration of a Linux instance.
        parameters: {}
        runtimeConfig:
          'aws:runShellScript':
            properties:
              - id: '0.aws:runShellScript'
                runCommand:
                  - ifconfig

      documentFormat: YAML
      documentType: Command
```
{{% /example %}}
{{% /examples %}}
## Permissions

The permissions attribute specifies how you want to share the document. If you share a document privately,
you must specify the AWS user account IDs for those people who can use the document. If you share a document
publicly, you must specify All as the account ID.

The permissions mapping supports the following:

* `type` - The permission type for the document. The permission type can be `Share`.
* `account_ids` - The AWS user accounts that should have access to the document. The account IDs can either be a group of account IDs or `All`.


## Import

Using `pulumi import`, import SSM Documents using the name. For example:

```sh
 $ pulumi import aws:ssm/document:Document example example
```
 The `attachments_source` argument does not have an SSM API method for reading the attachment information detail after creation. If the argument is set in the Pulumi program on an imported resource, Pulumi will always show a difference. To workaround this behavior, either omit the argument from the Pulumi program or use `ignore_changes` to hide the difference. For example:

