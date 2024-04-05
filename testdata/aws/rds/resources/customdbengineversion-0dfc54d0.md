Provides an custom engine version (CEV) resource for Amazon RDS Custom. For additional information, see [Working with CEVs for RDS Custom for Oracle](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.html) and [Working with CEVs for RDS Custom for SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev-sqlserver.html) in the the [RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Welcome.html).

{{% examples %}}
## Example Usage
{{% example %}}
### RDS Custom for Oracle Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleKey = new aws.kms.Key("exampleKey", {description: "KMS symmetric key for RDS Custom for Oracle"});
const exampleCustomDbEngineVersion = new aws.rds.CustomDbEngineVersion("exampleCustomDbEngineVersion", {
    databaseInstallationFilesS3BucketName: "DOC-EXAMPLE-BUCKET",
    databaseInstallationFilesS3Prefix: "1915_GI/",
    engine: "custom-oracle-ee-cdb",
    engineVersion: "19.cdb_cev1",
    kmsKeyId: exampleKey.arn,
    manifest: `  {
	"databaseInstallationFileNames":["V982063-01.zip"]
  }
`,
    tags: {
        Name: "example",
        Key: "value",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_key = aws.kms.Key("exampleKey", description="KMS symmetric key for RDS Custom for Oracle")
example_custom_db_engine_version = aws.rds.CustomDbEngineVersion("exampleCustomDbEngineVersion",
    database_installation_files_s3_bucket_name="DOC-EXAMPLE-BUCKET",
    database_installation_files_s3_prefix="1915_GI/",
    engine="custom-oracle-ee-cdb",
    engine_version="19.cdb_cev1",
    kms_key_id=example_key.arn,
    manifest="""  {
	"databaseInstallationFileNames":["V982063-01.zip"]
  }
""",
    tags={
        "Name": "example",
        "Key": "value",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleKey = new Aws.Kms.Key("exampleKey", new()
    {
        Description = "KMS symmetric key for RDS Custom for Oracle",
    });

    var exampleCustomDbEngineVersion = new Aws.Rds.CustomDbEngineVersion("exampleCustomDbEngineVersion", new()
    {
        DatabaseInstallationFilesS3BucketName = "DOC-EXAMPLE-BUCKET",
        DatabaseInstallationFilesS3Prefix = "1915_GI/",
        Engine = "custom-oracle-ee-cdb",
        EngineVersion = "19.cdb_cev1",
        KmsKeyId = exampleKey.Arn,
        Manifest = @"  {
	""databaseInstallationFileNames"":[""V982063-01.zip""]
  }
",
        Tags = 
        {
            { "Name", "example" },
            { "Key", "value" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
			Description: pulumi.String("KMS symmetric key for RDS Custom for Oracle"),
		})
		if err != nil {
			return err
		}
		_, err = rds.NewCustomDbEngineVersion(ctx, "exampleCustomDbEngineVersion", &rds.CustomDbEngineVersionArgs{
			DatabaseInstallationFilesS3BucketName: pulumi.String("DOC-EXAMPLE-BUCKET"),
			DatabaseInstallationFilesS3Prefix:     pulumi.String("1915_GI/"),
			Engine:                                pulumi.String("custom-oracle-ee-cdb"),
			EngineVersion:                         pulumi.String("19.cdb_cev1"),
			KmsKeyId:                              exampleKey.Arn,
			Manifest:                              pulumi.String("  {\n	\"databaseInstallationFileNames\":[\"V982063-01.zip\"]\n  }\n"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example"),
				"Key":  pulumi.String("value"),
			},
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
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.rds.CustomDbEngineVersion;
import com.pulumi.aws.rds.CustomDbEngineVersionArgs;
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
        var exampleKey = new Key("exampleKey", KeyArgs.builder()        
            .description("KMS symmetric key for RDS Custom for Oracle")
            .build());

        var exampleCustomDbEngineVersion = new CustomDbEngineVersion("exampleCustomDbEngineVersion", CustomDbEngineVersionArgs.builder()        
            .databaseInstallationFilesS3BucketName("DOC-EXAMPLE-BUCKET")
            .databaseInstallationFilesS3Prefix("1915_GI/")
            .engine("custom-oracle-ee-cdb")
            .engineVersion("19.cdb_cev1")
            .kmsKeyId(exampleKey.arn())
            .manifest("""
  {
	"databaseInstallationFileNames":["V982063-01.zip"]
  }
            """)
            .tags(Map.ofEntries(
                Map.entry("Name", "example"),
                Map.entry("Key", "value")
            ))
            .build());

    }
}
```
```yaml
resources:
  exampleKey:
    type: aws:kms:Key
    properties:
      description: KMS symmetric key for RDS Custom for Oracle
  exampleCustomDbEngineVersion:
    type: aws:rds:CustomDbEngineVersion
    properties:
      databaseInstallationFilesS3BucketName: DOC-EXAMPLE-BUCKET
      databaseInstallationFilesS3Prefix: 1915_GI/
      engine: custom-oracle-ee-cdb
      engineVersion: 19.cdb_cev1
      kmsKeyId: ${exampleKey.arn}
      manifest: |2
          {
        	"databaseInstallationFileNames":["V982063-01.zip"]
          }
      tags:
        Name: example
        Key: value
```
{{% /example %}}
{{% example %}}
### RDS Custom for Oracle External Manifest Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as crypto from "crypto";
import * as fs from "fs";

function computeFilebase64sha256(path: string): string {
	const fileData = Buffer.from(fs.readFileSync(path, 'binary'))
	return crypto.createHash('sha256').update(fileData).digest('hex')
}

const exampleKey = new aws.kms.Key("exampleKey", {description: "KMS symmetric key for RDS Custom for Oracle"});
const exampleCustomDbEngineVersion = new aws.rds.CustomDbEngineVersion("exampleCustomDbEngineVersion", {
    databaseInstallationFilesS3BucketName: "DOC-EXAMPLE-BUCKET",
    databaseInstallationFilesS3Prefix: "1915_GI/",
    engine: "custom-oracle-ee-cdb",
    engineVersion: "19.cdb_cev1",
    kmsKeyId: exampleKey.arn,
    filename: "manifest_1915_GI.json",
    manifestHash: computeFilebase64sha256(manifest_1915_GI.json),
    tags: {
        Name: "example",
        Key: "value",
    },
});
```
```python
import pulumi
import base64
import hashlib
import pulumi_aws as aws

def computeFilebase64sha256(path):
	fileData = open(path).read().encode()
	hashedData = hashlib.sha256(fileData.encode()).digest()
	return base64.b64encode(hashedData).decode()

example_key = aws.kms.Key("exampleKey", description="KMS symmetric key for RDS Custom for Oracle")
example_custom_db_engine_version = aws.rds.CustomDbEngineVersion("exampleCustomDbEngineVersion",
    database_installation_files_s3_bucket_name="DOC-EXAMPLE-BUCKET",
    database_installation_files_s3_prefix="1915_GI/",
    engine="custom-oracle-ee-cdb",
    engine_version="19.cdb_cev1",
    kms_key_id=example_key.arn,
    filename="manifest_1915_GI.json",
    manifest_hash=computeFilebase64sha256(manifest_1915__gi["json"]),
    tags={
        "Name": "example",
        "Key": "value",
    })
```
```csharp
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Security.Cryptography;
using System.Text;
using Pulumi;
using Aws = Pulumi.Aws;

	
string ComputeFileBase64Sha256(string path) 
{
    var fileData = Encoding.UTF8.GetBytes(File.ReadAllText(path));
    var hashData = SHA256.Create().ComputeHash(fileData);
    return Convert.ToBase64String(hashData);
}

return await Deployment.RunAsync(() => 
{
    var exampleKey = new Aws.Kms.Key("exampleKey", new()
    {
        Description = "KMS symmetric key for RDS Custom for Oracle",
    });

    var exampleCustomDbEngineVersion = new Aws.Rds.CustomDbEngineVersion("exampleCustomDbEngineVersion", new()
    {
        DatabaseInstallationFilesS3BucketName = "DOC-EXAMPLE-BUCKET",
        DatabaseInstallationFilesS3Prefix = "1915_GI/",
        Engine = "custom-oracle-ee-cdb",
        EngineVersion = "19.cdb_cev1",
        KmsKeyId = exampleKey.Arn,
        Filename = "manifest_1915_GI.json",
        ManifestHash = ComputeFileBase64Sha256(manifest_1915_GI.Json),
        Tags = 
        {
            { "Name", "example" },
            { "Key", "value" },
        },
    });

});
```
```go
package main

import (
	"crypto/sha256"
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func filebase64sha256OrPanic(path string) string {
	if fileData, err := os.ReadFile(path); err == nil {
		hashedData := sha256.Sum256([]byte(fileData))
		return base64.StdEncoding.EncodeToString(hashedData[:])
	} else {
		panic(err.Error())
	}
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
			Description: pulumi.String("KMS symmetric key for RDS Custom for Oracle"),
		})
		if err != nil {
			return err
		}
		_, err = rds.NewCustomDbEngineVersion(ctx, "exampleCustomDbEngineVersion", &rds.CustomDbEngineVersionArgs{
			DatabaseInstallationFilesS3BucketName: pulumi.String("DOC-EXAMPLE-BUCKET"),
			DatabaseInstallationFilesS3Prefix:     pulumi.String("1915_GI/"),
			Engine:                                pulumi.String("custom-oracle-ee-cdb"),
			EngineVersion:                         pulumi.String("19.cdb_cev1"),
			KmsKeyId:                              exampleKey.Arn,
			Filename:                              pulumi.String("manifest_1915_GI.json"),
			ManifestHash:                          filebase64sha256OrPanic(manifest_1915_GI.Json),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example"),
				"Key":  pulumi.String("value"),
			},
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
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.rds.CustomDbEngineVersion;
import com.pulumi.aws.rds.CustomDbEngineVersionArgs;
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
        var exampleKey = new Key("exampleKey", KeyArgs.builder()        
            .description("KMS symmetric key for RDS Custom for Oracle")
            .build());

        var exampleCustomDbEngineVersion = new CustomDbEngineVersion("exampleCustomDbEngineVersion", CustomDbEngineVersionArgs.builder()        
            .databaseInstallationFilesS3BucketName("DOC-EXAMPLE-BUCKET")
            .databaseInstallationFilesS3Prefix("1915_GI/")
            .engine("custom-oracle-ee-cdb")
            .engineVersion("19.cdb_cev1")
            .kmsKeyId(exampleKey.arn())
            .filename("manifest_1915_GI.json")
            .manifestHash(computeFileBase64Sha256(manifest_1915_GI.json()))
            .tags(Map.ofEntries(
                Map.entry("Name", "example"),
                Map.entry("Key", "value")
            ))
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### RDS Custom for SQL Server Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// CEV creation requires an AMI owned by the operator
const test = new aws.rds.CustomDbEngineVersion("test", {
    engine: "custom-sqlserver-se",
    engineVersion: "15.00.4249.2.cev-1",
    sourceImageId: "ami-0aa12345678a12ab1",
});
```
```python
import pulumi
import pulumi_aws as aws

# CEV creation requires an AMI owned by the operator
test = aws.rds.CustomDbEngineVersion("test",
    engine="custom-sqlserver-se",
    engine_version="15.00.4249.2.cev-1",
    source_image_id="ami-0aa12345678a12ab1")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // CEV creation requires an AMI owned by the operator
    var test = new Aws.Rds.CustomDbEngineVersion("test", new()
    {
        Engine = "custom-sqlserver-se",
        EngineVersion = "15.00.4249.2.cev-1",
        SourceImageId = "ami-0aa12345678a12ab1",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewCustomDbEngineVersion(ctx, "test", &rds.CustomDbEngineVersionArgs{
			Engine:        pulumi.String("custom-sqlserver-se"),
			EngineVersion: pulumi.String("15.00.4249.2.cev-1"),
			SourceImageId: pulumi.String("ami-0aa12345678a12ab1"),
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
import com.pulumi.aws.rds.CustomDbEngineVersion;
import com.pulumi.aws.rds.CustomDbEngineVersionArgs;
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
        var test = new CustomDbEngineVersion("test", CustomDbEngineVersionArgs.builder()        
            .engine("custom-sqlserver-se")
            .engineVersion("15.00.4249.2.cev-1")
            .sourceImageId("ami-0aa12345678a12ab1")
            .build());

    }
}
```
```yaml
resources:
  # CEV creation requires an AMI owned by the operator
  test:
    type: aws:rds:CustomDbEngineVersion
    properties:
      engine: custom-sqlserver-se
      engineVersion: 15.00.4249.2.cev-1
      sourceImageId: ami-0aa12345678a12ab1
```
{{% /example %}}
{{% example %}}
### RDS Custom for SQL Server Usage with AMI from another region

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.AmiCopy("example", {
    description: "A copy of ami-xxxxxxxx",
    sourceAmiId: "ami-xxxxxxxx",
    sourceAmiRegion: "us-east-1",
});
// CEV creation requires an AMI owned by the operator
const test = new aws.rds.CustomDbEngineVersion("test", {
    engine: "custom-sqlserver-se",
    engineVersion: "15.00.4249.2.cev-1",
    sourceImageId: example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.AmiCopy("example",
    description="A copy of ami-xxxxxxxx",
    source_ami_id="ami-xxxxxxxx",
    source_ami_region="us-east-1")
# CEV creation requires an AMI owned by the operator
test = aws.rds.CustomDbEngineVersion("test",
    engine="custom-sqlserver-se",
    engine_version="15.00.4249.2.cev-1",
    source_image_id=example.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.AmiCopy("example", new()
    {
        Description = "A copy of ami-xxxxxxxx",
        SourceAmiId = "ami-xxxxxxxx",
        SourceAmiRegion = "us-east-1",
    });

    // CEV creation requires an AMI owned by the operator
    var test = new Aws.Rds.CustomDbEngineVersion("test", new()
    {
        Engine = "custom-sqlserver-se",
        EngineVersion = "15.00.4249.2.cev-1",
        SourceImageId = example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := ec2.NewAmiCopy(ctx, "example", &ec2.AmiCopyArgs{
			Description:     pulumi.String("A copy of ami-xxxxxxxx"),
			SourceAmiId:     pulumi.String("ami-xxxxxxxx"),
			SourceAmiRegion: pulumi.String("us-east-1"),
		})
		if err != nil {
			return err
		}
		_, err = rds.NewCustomDbEngineVersion(ctx, "test", &rds.CustomDbEngineVersionArgs{
			Engine:        pulumi.String("custom-sqlserver-se"),
			EngineVersion: pulumi.String("15.00.4249.2.cev-1"),
			SourceImageId: example.ID(),
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
import com.pulumi.aws.ec2.AmiCopy;
import com.pulumi.aws.ec2.AmiCopyArgs;
import com.pulumi.aws.rds.CustomDbEngineVersion;
import com.pulumi.aws.rds.CustomDbEngineVersionArgs;
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
        var example = new AmiCopy("example", AmiCopyArgs.builder()        
            .description("A copy of ami-xxxxxxxx")
            .sourceAmiId("ami-xxxxxxxx")
            .sourceAmiRegion("us-east-1")
            .build());

        var test = new CustomDbEngineVersion("test", CustomDbEngineVersionArgs.builder()        
            .engine("custom-sqlserver-se")
            .engineVersion("15.00.4249.2.cev-1")
            .sourceImageId(example.id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:AmiCopy
    properties:
      description: A copy of ami-xxxxxxxx
      sourceAmiId: ami-xxxxxxxx
      sourceAmiRegion: us-east-1
  # CEV creation requires an AMI owned by the operator
  test:
    type: aws:rds:CustomDbEngineVersion
    properties:
      engine: custom-sqlserver-se
      engineVersion: 15.00.4249.2.cev-1
      sourceImageId: ${example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import custom engine versions for Amazon RDS custom using the `engine` and `engine_version` separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:rds/customDbEngineVersion:CustomDbEngineVersion example custom-oracle-ee-cdb:19.cdb_cev1
```
 