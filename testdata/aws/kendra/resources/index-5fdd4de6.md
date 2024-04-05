Provides an Amazon Kendra Index resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.Index("example", {
    description: "example",
    edition: "DEVELOPER_EDITION",
    roleArn: aws_iam_role["this"].arn,
    tags: {
        Key1: "Value1",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.Index("example",
    description="example",
    edition="DEVELOPER_EDITION",
    role_arn=aws_iam_role["this"]["arn"],
    tags={
        "Key1": "Value1",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Kendra.Index("example", new()
    {
        Description = "example",
        Edition = "DEVELOPER_EDITION",
        RoleArn = aws_iam_role.This.Arn,
        Tags = 
        {
            { "Key1", "Value1" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kendra"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kendra.NewIndex(ctx, "example", &kendra.IndexArgs{
			Description: pulumi.String("example"),
			Edition:     pulumi.String("DEVELOPER_EDITION"),
			RoleArn:     pulumi.Any(aws_iam_role.This.Arn),
			Tags: pulumi.StringMap{
				"Key1": pulumi.String("Value1"),
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
import com.pulumi.aws.kendra.Index;
import com.pulumi.aws.kendra.IndexArgs;
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
        var example = new Index("example", IndexArgs.builder()        
            .description("example")
            .edition("DEVELOPER_EDITION")
            .roleArn(aws_iam_role.this().arn())
            .tags(Map.of("Key1", "Value1"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:Index
    properties:
      description: example
      edition: DEVELOPER_EDITION
      roleArn: ${aws_iam_role.this.arn}
      tags:
        Key1: Value1
```
{{% /example %}}
{{% example %}}
### With capacity units

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.Index("example", {
    edition: "DEVELOPER_EDITION",
    roleArn: aws_iam_role["this"].arn,
    capacityUnits: {
        queryCapacityUnits: 2,
        storageCapacityUnits: 2,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.Index("example",
    edition="DEVELOPER_EDITION",
    role_arn=aws_iam_role["this"]["arn"],
    capacity_units=aws.kendra.IndexCapacityUnitsArgs(
        query_capacity_units=2,
        storage_capacity_units=2,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Kendra.Index("example", new()
    {
        Edition = "DEVELOPER_EDITION",
        RoleArn = aws_iam_role.This.Arn,
        CapacityUnits = new Aws.Kendra.Inputs.IndexCapacityUnitsArgs
        {
            QueryCapacityUnits = 2,
            StorageCapacityUnits = 2,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kendra"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kendra.NewIndex(ctx, "example", &kendra.IndexArgs{
			Edition: pulumi.String("DEVELOPER_EDITION"),
			RoleArn: pulumi.Any(aws_iam_role.This.Arn),
			CapacityUnits: &kendra.IndexCapacityUnitsArgs{
				QueryCapacityUnits:   pulumi.Int(2),
				StorageCapacityUnits: pulumi.Int(2),
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
import com.pulumi.aws.kendra.Index;
import com.pulumi.aws.kendra.IndexArgs;
import com.pulumi.aws.kendra.inputs.IndexCapacityUnitsArgs;
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
        var example = new Index("example", IndexArgs.builder()        
            .edition("DEVELOPER_EDITION")
            .roleArn(aws_iam_role.this().arn())
            .capacityUnits(IndexCapacityUnitsArgs.builder()
                .queryCapacityUnits(2)
                .storageCapacityUnits(2)
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:Index
    properties:
      edition: DEVELOPER_EDITION
      roleArn: ${aws_iam_role.this.arn}
      capacityUnits:
        queryCapacityUnits: 2
        storageCapacityUnits: 2
```
{{% /example %}}
{{% example %}}
### With server side encryption configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.Index("example", {
    roleArn: aws_iam_role["this"].arn,
    serverSideEncryptionConfiguration: {
        kmsKeyId: data.aws_kms_key["this"].arn,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.Index("example",
    role_arn=aws_iam_role["this"]["arn"],
    server_side_encryption_configuration=aws.kendra.IndexServerSideEncryptionConfigurationArgs(
        kms_key_id=data["aws_kms_key"]["this"]["arn"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Kendra.Index("example", new()
    {
        RoleArn = aws_iam_role.This.Arn,
        ServerSideEncryptionConfiguration = new Aws.Kendra.Inputs.IndexServerSideEncryptionConfigurationArgs
        {
            KmsKeyId = data.Aws_kms_key.This.Arn,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kendra"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kendra.NewIndex(ctx, "example", &kendra.IndexArgs{
			RoleArn: pulumi.Any(aws_iam_role.This.Arn),
			ServerSideEncryptionConfiguration: &kendra.IndexServerSideEncryptionConfigurationArgs{
				KmsKeyId: pulumi.Any(data.Aws_kms_key.This.Arn),
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
import com.pulumi.aws.kendra.Index;
import com.pulumi.aws.kendra.IndexArgs;
import com.pulumi.aws.kendra.inputs.IndexServerSideEncryptionConfigurationArgs;
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
        var example = new Index("example", IndexArgs.builder()        
            .roleArn(aws_iam_role.this().arn())
            .serverSideEncryptionConfiguration(IndexServerSideEncryptionConfigurationArgs.builder()
                .kmsKeyId(data.aws_kms_key().this().arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:Index
    properties:
      roleArn: ${aws_iam_role.this.arn}
      serverSideEncryptionConfiguration:
        kmsKeyId: ${data.aws_kms_key.this.arn}
```
{{% /example %}}
{{% example %}}
### With user group resolution configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.Index("example", {
    roleArn: aws_iam_role["this"].arn,
    userGroupResolutionConfiguration: {
        userGroupResolutionMode: "AWS_SSO",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.Index("example",
    role_arn=aws_iam_role["this"]["arn"],
    user_group_resolution_configuration=aws.kendra.IndexUserGroupResolutionConfigurationArgs(
        user_group_resolution_mode="AWS_SSO",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Kendra.Index("example", new()
    {
        RoleArn = aws_iam_role.This.Arn,
        UserGroupResolutionConfiguration = new Aws.Kendra.Inputs.IndexUserGroupResolutionConfigurationArgs
        {
            UserGroupResolutionMode = "AWS_SSO",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kendra"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kendra.NewIndex(ctx, "example", &kendra.IndexArgs{
			RoleArn: pulumi.Any(aws_iam_role.This.Arn),
			UserGroupResolutionConfiguration: &kendra.IndexUserGroupResolutionConfigurationArgs{
				UserGroupResolutionMode: pulumi.String("AWS_SSO"),
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
import com.pulumi.aws.kendra.Index;
import com.pulumi.aws.kendra.IndexArgs;
import com.pulumi.aws.kendra.inputs.IndexUserGroupResolutionConfigurationArgs;
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
        var example = new Index("example", IndexArgs.builder()        
            .roleArn(aws_iam_role.this().arn())
            .userGroupResolutionConfiguration(IndexUserGroupResolutionConfigurationArgs.builder()
                .userGroupResolutionMode("AWS_SSO")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:Index
    properties:
      roleArn: ${aws_iam_role.this.arn}
      userGroupResolutionConfiguration:
        userGroupResolutionMode: AWS_SSO
```
{{% /example %}}
### With Document Metadata Configuration Updates
{{% example %}}
### Specifying the predefined elements

Refer to [Amazon Kendra documentation on built-in document fields](https://docs.aws.amazon.com/kendra/latest/dg/hiw-index.html#index-reserved-fields) for more information.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.Index("example", {
    roleArn: aws_iam_role["this"].arn,
    documentMetadataConfigurationUpdates: [
        {
            name: "_authors",
            type: "STRING_LIST_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: false,
            },
            relevance: {
                importance: 1,
            },
        },
        {
            name: "_category",
            type: "STRING_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                importance: 1,
                valuesImportanceMap: {},
            },
        },
        {
            name: "_created_at",
            type: "DATE_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                freshness: false,
                importance: 1,
                duration: "25920000s",
                rankOrder: "ASCENDING",
            },
        },
        {
            name: "_data_source_id",
            type: "STRING_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                importance: 1,
                valuesImportanceMap: {},
            },
        },
        {
            name: "_document_title",
            type: "STRING_VALUE",
            search: {
                displayable: true,
                facetable: false,
                searchable: true,
                sortable: true,
            },
            relevance: {
                importance: 2,
                valuesImportanceMap: {},
            },
        },
        {
            name: "_excerpt_page_number",
            type: "LONG_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: false,
            },
            relevance: {
                importance: 2,
                rankOrder: "ASCENDING",
            },
        },
        {
            name: "_faq_id",
            type: "STRING_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                importance: 1,
                valuesImportanceMap: {},
            },
        },
        {
            name: "_file_type",
            type: "STRING_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                importance: 1,
                valuesImportanceMap: {},
            },
        },
        {
            name: "_language_code",
            type: "STRING_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                importance: 1,
                valuesImportanceMap: {},
            },
        },
        {
            name: "_last_updated_at",
            type: "DATE_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                freshness: false,
                importance: 1,
                duration: "25920000s",
                rankOrder: "ASCENDING",
            },
        },
        {
            name: "_source_uri",
            type: "STRING_VALUE",
            search: {
                displayable: true,
                facetable: false,
                searchable: false,
                sortable: false,
            },
            relevance: {
                importance: 1,
                valuesImportanceMap: {},
            },
        },
        {
            name: "_tenant_id",
            type: "STRING_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                importance: 1,
                valuesImportanceMap: {},
            },
        },
        {
            name: "_version",
            type: "STRING_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                importance: 1,
                valuesImportanceMap: {},
            },
        },
        {
            name: "_view_count",
            type: "LONG_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                importance: 1,
                rankOrder: "ASCENDING",
            },
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.Index("example",
    role_arn=aws_iam_role["this"]["arn"],
    document_metadata_configuration_updates=[
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_authors",
            type="STRING_LIST_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=False,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_category",
            type="STRING_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                values_importance_map={},
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_created_at",
            type="DATE_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                freshness=False,
                importance=1,
                duration="25920000s",
                rank_order="ASCENDING",
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_data_source_id",
            type="STRING_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                values_importance_map={},
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_document_title",
            type="STRING_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=True,
                facetable=False,
                searchable=True,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=2,
                values_importance_map={},
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_excerpt_page_number",
            type="LONG_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=False,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=2,
                rank_order="ASCENDING",
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_faq_id",
            type="STRING_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                values_importance_map={},
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_file_type",
            type="STRING_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                values_importance_map={},
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_language_code",
            type="STRING_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                values_importance_map={},
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_last_updated_at",
            type="DATE_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                freshness=False,
                importance=1,
                duration="25920000s",
                rank_order="ASCENDING",
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_source_uri",
            type="STRING_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=True,
                facetable=False,
                searchable=False,
                sortable=False,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                values_importance_map={},
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_tenant_id",
            type="STRING_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                values_importance_map={},
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_version",
            type="STRING_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                values_importance_map={},
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_view_count",
            type="LONG_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                rank_order="ASCENDING",
            ),
        ),
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Kendra.Index("example", new()
    {
        RoleArn = aws_iam_role.This.Arn,
        DocumentMetadataConfigurationUpdates = new[]
        {
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_authors",
                Type = "STRING_LIST_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = false,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_category",
                Type = "STRING_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    ValuesImportanceMap = null,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_created_at",
                Type = "DATE_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Freshness = false,
                    Importance = 1,
                    Duration = "25920000s",
                    RankOrder = "ASCENDING",
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_data_source_id",
                Type = "STRING_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    ValuesImportanceMap = null,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_document_title",
                Type = "STRING_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = true,
                    Facetable = false,
                    Searchable = true,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 2,
                    ValuesImportanceMap = null,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_excerpt_page_number",
                Type = "LONG_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = false,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 2,
                    RankOrder = "ASCENDING",
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_faq_id",
                Type = "STRING_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    ValuesImportanceMap = null,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_file_type",
                Type = "STRING_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    ValuesImportanceMap = null,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_language_code",
                Type = "STRING_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    ValuesImportanceMap = null,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_last_updated_at",
                Type = "DATE_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Freshness = false,
                    Importance = 1,
                    Duration = "25920000s",
                    RankOrder = "ASCENDING",
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_source_uri",
                Type = "STRING_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = true,
                    Facetable = false,
                    Searchable = false,
                    Sortable = false,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    ValuesImportanceMap = null,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_tenant_id",
                Type = "STRING_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    ValuesImportanceMap = null,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_version",
                Type = "STRING_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    ValuesImportanceMap = null,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_view_count",
                Type = "LONG_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    RankOrder = "ASCENDING",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kendra"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kendra.NewIndex(ctx, "example", &kendra.IndexArgs{
			RoleArn: pulumi.Any(aws_iam_role.This.Arn),
			DocumentMetadataConfigurationUpdates: kendra.IndexDocumentMetadataConfigurationUpdateArray{
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_authors"),
					Type: pulumi.String("STRING_LIST_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(false),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance: pulumi.Int(1),
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_category"),
					Type: pulumi.String("STRING_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance:          pulumi.Int(1),
						ValuesImportanceMap: nil,
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_created_at"),
					Type: pulumi.String("DATE_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Freshness:  pulumi.Bool(false),
						Importance: pulumi.Int(1),
						Duration:   pulumi.String("25920000s"),
						RankOrder:  pulumi.String("ASCENDING"),
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_data_source_id"),
					Type: pulumi.String("STRING_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance:          pulumi.Int(1),
						ValuesImportanceMap: nil,
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_document_title"),
					Type: pulumi.String("STRING_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(true),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(true),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance:          pulumi.Int(2),
						ValuesImportanceMap: nil,
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_excerpt_page_number"),
					Type: pulumi.String("LONG_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(false),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance: pulumi.Int(2),
						RankOrder:  pulumi.String("ASCENDING"),
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_faq_id"),
					Type: pulumi.String("STRING_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance:          pulumi.Int(1),
						ValuesImportanceMap: nil,
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_file_type"),
					Type: pulumi.String("STRING_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance:          pulumi.Int(1),
						ValuesImportanceMap: nil,
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_language_code"),
					Type: pulumi.String("STRING_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance:          pulumi.Int(1),
						ValuesImportanceMap: nil,
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_last_updated_at"),
					Type: pulumi.String("DATE_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Freshness:  pulumi.Bool(false),
						Importance: pulumi.Int(1),
						Duration:   pulumi.String("25920000s"),
						RankOrder:  pulumi.String("ASCENDING"),
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_source_uri"),
					Type: pulumi.String("STRING_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(true),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(false),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance:          pulumi.Int(1),
						ValuesImportanceMap: nil,
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_tenant_id"),
					Type: pulumi.String("STRING_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance:          pulumi.Int(1),
						ValuesImportanceMap: nil,
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_version"),
					Type: pulumi.String("STRING_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance:          pulumi.Int(1),
						ValuesImportanceMap: nil,
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_view_count"),
					Type: pulumi.String("LONG_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance: pulumi.Int(1),
						RankOrder:  pulumi.String("ASCENDING"),
					},
				},
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
import com.pulumi.aws.kendra.Index;
import com.pulumi.aws.kendra.IndexArgs;
import com.pulumi.aws.kendra.inputs.IndexDocumentMetadataConfigurationUpdateArgs;
import com.pulumi.aws.kendra.inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs;
import com.pulumi.aws.kendra.inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs;
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
        var example = new Index("example", IndexArgs.builder()        
            .roleArn(aws_iam_role.this().arn())
            .documentMetadataConfigurationUpdates(            
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_authors")
                    .type("STRING_LIST_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(false)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_category")
                    .type("STRING_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .valuesImportanceMap()
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_created_at")
                    .type("DATE_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .freshness(false)
                        .importance(1)
                        .duration("25920000s")
                        .rankOrder("ASCENDING")
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_data_source_id")
                    .type("STRING_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .valuesImportanceMap()
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_document_title")
                    .type("STRING_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(true)
                        .facetable(false)
                        .searchable(true)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(2)
                        .valuesImportanceMap()
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_excerpt_page_number")
                    .type("LONG_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(false)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(2)
                        .rankOrder("ASCENDING")
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_faq_id")
                    .type("STRING_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .valuesImportanceMap()
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_file_type")
                    .type("STRING_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .valuesImportanceMap()
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_language_code")
                    .type("STRING_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .valuesImportanceMap()
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_last_updated_at")
                    .type("DATE_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .freshness(false)
                        .importance(1)
                        .duration("25920000s")
                        .rankOrder("ASCENDING")
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_source_uri")
                    .type("STRING_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(true)
                        .facetable(false)
                        .searchable(false)
                        .sortable(false)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .valuesImportanceMap()
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_tenant_id")
                    .type("STRING_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .valuesImportanceMap()
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_version")
                    .type("STRING_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .valuesImportanceMap()
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_view_count")
                    .type("LONG_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .rankOrder("ASCENDING")
                        .build())
                    .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:Index
    properties:
      roleArn: ${aws_iam_role.this.arn}
      documentMetadataConfigurationUpdates:
        - name: _authors
          type: STRING_LIST_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: false
          relevance:
            importance: 1
        - name: _category
          type: STRING_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            importance: 1
            valuesImportanceMap: {}
        - name: _created_at
          type: DATE_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            freshness: false
            importance: 1
            duration: 25920000s
            rankOrder: ASCENDING
        - name: _data_source_id
          type: STRING_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            importance: 1
            valuesImportanceMap: {}
        - name: _document_title
          type: STRING_VALUE
          search:
            displayable: true
            facetable: false
            searchable: true
            sortable: true
          relevance:
            importance: 2
            valuesImportanceMap: {}
        - name: _excerpt_page_number
          type: LONG_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: false
          relevance:
            importance: 2
            rankOrder: ASCENDING
        - name: _faq_id
          type: STRING_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            importance: 1
            valuesImportanceMap: {}
        - name: _file_type
          type: STRING_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            importance: 1
            valuesImportanceMap: {}
        - name: _language_code
          type: STRING_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            importance: 1
            valuesImportanceMap: {}
        - name: _last_updated_at
          type: DATE_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            freshness: false
            importance: 1
            duration: 25920000s
            rankOrder: ASCENDING
        - name: _source_uri
          type: STRING_VALUE
          search:
            displayable: true
            facetable: false
            searchable: false
            sortable: false
          relevance:
            importance: 1
            valuesImportanceMap: {}
        - name: _tenant_id
          type: STRING_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            importance: 1
            valuesImportanceMap: {}
        - name: _version
          type: STRING_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            importance: 1
            valuesImportanceMap: {}
        - name: _view_count
          type: LONG_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            importance: 1
            rankOrder: ASCENDING
```
{{% /example %}}
{{% example %}}
### Appending additional elements

The example below shows additional elements with names, `example-string-value`, `example-long-value`, `example-string-list-value`, `example-date-value` representing the 4 types of `STRING_VALUE`, `LONG_VALUE`, `STRING_LIST_VALUE`, `DATE_VALUE` respectively.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.Index("example", {
    roleArn: aws_iam_role["this"].arn,
    documentMetadataConfigurationUpdates: [
        {
            name: "_authors",
            type: "STRING_LIST_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: false,
            },
            relevance: {
                importance: 1,
            },
        },
        {
            name: "_category",
            type: "STRING_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                importance: 1,
                valuesImportanceMap: {},
            },
        },
        {
            name: "_created_at",
            type: "DATE_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                freshness: false,
                importance: 1,
                duration: "25920000s",
                rankOrder: "ASCENDING",
            },
        },
        {
            name: "_data_source_id",
            type: "STRING_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                importance: 1,
                valuesImportanceMap: {},
            },
        },
        {
            name: "_document_title",
            type: "STRING_VALUE",
            search: {
                displayable: true,
                facetable: false,
                searchable: true,
                sortable: true,
            },
            relevance: {
                importance: 2,
                valuesImportanceMap: {},
            },
        },
        {
            name: "_excerpt_page_number",
            type: "LONG_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: false,
            },
            relevance: {
                importance: 2,
                rankOrder: "ASCENDING",
            },
        },
        {
            name: "_faq_id",
            type: "STRING_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                importance: 1,
                valuesImportanceMap: {},
            },
        },
        {
            name: "_file_type",
            type: "STRING_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                importance: 1,
                valuesImportanceMap: {},
            },
        },
        {
            name: "_language_code",
            type: "STRING_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                importance: 1,
                valuesImportanceMap: {},
            },
        },
        {
            name: "_last_updated_at",
            type: "DATE_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                freshness: false,
                importance: 1,
                duration: "25920000s",
                rankOrder: "ASCENDING",
            },
        },
        {
            name: "_source_uri",
            type: "STRING_VALUE",
            search: {
                displayable: true,
                facetable: false,
                searchable: false,
                sortable: false,
            },
            relevance: {
                importance: 1,
                valuesImportanceMap: {},
            },
        },
        {
            name: "_tenant_id",
            type: "STRING_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                importance: 1,
                valuesImportanceMap: {},
            },
        },
        {
            name: "_version",
            type: "STRING_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                importance: 1,
                valuesImportanceMap: {},
            },
        },
        {
            name: "_view_count",
            type: "LONG_VALUE",
            search: {
                displayable: false,
                facetable: false,
                searchable: false,
                sortable: true,
            },
            relevance: {
                importance: 1,
                rankOrder: "ASCENDING",
            },
        },
        {
            name: "example-string-value",
            type: "STRING_VALUE",
            search: {
                displayable: true,
                facetable: true,
                searchable: true,
                sortable: true,
            },
            relevance: {
                importance: 1,
                valuesImportanceMap: {},
            },
        },
        {
            name: "example-long-value",
            type: "LONG_VALUE",
            search: {
                displayable: true,
                facetable: true,
                searchable: false,
                sortable: true,
            },
            relevance: {
                importance: 1,
                rankOrder: "ASCENDING",
            },
        },
        {
            name: "example-string-list-value",
            type: "STRING_LIST_VALUE",
            search: {
                displayable: true,
                facetable: true,
                searchable: true,
                sortable: false,
            },
            relevance: {
                importance: 1,
            },
        },
        {
            name: "example-date-value",
            type: "DATE_VALUE",
            search: {
                displayable: true,
                facetable: true,
                searchable: false,
                sortable: false,
            },
            relevance: {
                freshness: false,
                importance: 1,
                duration: "25920000s",
                rankOrder: "ASCENDING",
            },
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.Index("example",
    role_arn=aws_iam_role["this"]["arn"],
    document_metadata_configuration_updates=[
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_authors",
            type="STRING_LIST_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=False,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_category",
            type="STRING_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                values_importance_map={},
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_created_at",
            type="DATE_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                freshness=False,
                importance=1,
                duration="25920000s",
                rank_order="ASCENDING",
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_data_source_id",
            type="STRING_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                values_importance_map={},
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_document_title",
            type="STRING_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=True,
                facetable=False,
                searchable=True,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=2,
                values_importance_map={},
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_excerpt_page_number",
            type="LONG_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=False,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=2,
                rank_order="ASCENDING",
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_faq_id",
            type="STRING_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                values_importance_map={},
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_file_type",
            type="STRING_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                values_importance_map={},
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_language_code",
            type="STRING_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                values_importance_map={},
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_last_updated_at",
            type="DATE_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                freshness=False,
                importance=1,
                duration="25920000s",
                rank_order="ASCENDING",
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_source_uri",
            type="STRING_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=True,
                facetable=False,
                searchable=False,
                sortable=False,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                values_importance_map={},
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_tenant_id",
            type="STRING_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                values_importance_map={},
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_version",
            type="STRING_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                values_importance_map={},
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="_view_count",
            type="LONG_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=False,
                facetable=False,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                rank_order="ASCENDING",
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="example-string-value",
            type="STRING_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=True,
                facetable=True,
                searchable=True,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                values_importance_map={},
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="example-long-value",
            type="LONG_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=True,
                facetable=True,
                searchable=False,
                sortable=True,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
                rank_order="ASCENDING",
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="example-string-list-value",
            type="STRING_LIST_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=True,
                facetable=True,
                searchable=True,
                sortable=False,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                importance=1,
            ),
        ),
        aws.kendra.IndexDocumentMetadataConfigurationUpdateArgs(
            name="example-date-value",
            type="DATE_VALUE",
            search=aws.kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs(
                displayable=True,
                facetable=True,
                searchable=False,
                sortable=False,
            ),
            relevance=aws.kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs(
                freshness=False,
                importance=1,
                duration="25920000s",
                rank_order="ASCENDING",
            ),
        ),
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Kendra.Index("example", new()
    {
        RoleArn = aws_iam_role.This.Arn,
        DocumentMetadataConfigurationUpdates = new[]
        {
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_authors",
                Type = "STRING_LIST_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = false,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_category",
                Type = "STRING_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    ValuesImportanceMap = null,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_created_at",
                Type = "DATE_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Freshness = false,
                    Importance = 1,
                    Duration = "25920000s",
                    RankOrder = "ASCENDING",
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_data_source_id",
                Type = "STRING_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    ValuesImportanceMap = null,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_document_title",
                Type = "STRING_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = true,
                    Facetable = false,
                    Searchable = true,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 2,
                    ValuesImportanceMap = null,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_excerpt_page_number",
                Type = "LONG_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = false,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 2,
                    RankOrder = "ASCENDING",
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_faq_id",
                Type = "STRING_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    ValuesImportanceMap = null,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_file_type",
                Type = "STRING_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    ValuesImportanceMap = null,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_language_code",
                Type = "STRING_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    ValuesImportanceMap = null,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_last_updated_at",
                Type = "DATE_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Freshness = false,
                    Importance = 1,
                    Duration = "25920000s",
                    RankOrder = "ASCENDING",
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_source_uri",
                Type = "STRING_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = true,
                    Facetable = false,
                    Searchable = false,
                    Sortable = false,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    ValuesImportanceMap = null,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_tenant_id",
                Type = "STRING_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    ValuesImportanceMap = null,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_version",
                Type = "STRING_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    ValuesImportanceMap = null,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "_view_count",
                Type = "LONG_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = false,
                    Facetable = false,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    RankOrder = "ASCENDING",
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "example-string-value",
                Type = "STRING_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = true,
                    Facetable = true,
                    Searchable = true,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    ValuesImportanceMap = null,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "example-long-value",
                Type = "LONG_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = true,
                    Facetable = true,
                    Searchable = false,
                    Sortable = true,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                    RankOrder = "ASCENDING",
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "example-string-list-value",
                Type = "STRING_LIST_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = true,
                    Facetable = true,
                    Searchable = true,
                    Sortable = false,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Importance = 1,
                },
            },
            new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateArgs
            {
                Name = "example-date-value",
                Type = "DATE_VALUE",
                Search = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs
                {
                    Displayable = true,
                    Facetable = true,
                    Searchable = false,
                    Sortable = false,
                },
                Relevance = new Aws.Kendra.Inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs
                {
                    Freshness = false,
                    Importance = 1,
                    Duration = "25920000s",
                    RankOrder = "ASCENDING",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kendra"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kendra.NewIndex(ctx, "example", &kendra.IndexArgs{
			RoleArn: pulumi.Any(aws_iam_role.This.Arn),
			DocumentMetadataConfigurationUpdates: kendra.IndexDocumentMetadataConfigurationUpdateArray{
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_authors"),
					Type: pulumi.String("STRING_LIST_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(false),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance: pulumi.Int(1),
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_category"),
					Type: pulumi.String("STRING_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance:          pulumi.Int(1),
						ValuesImportanceMap: nil,
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_created_at"),
					Type: pulumi.String("DATE_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Freshness:  pulumi.Bool(false),
						Importance: pulumi.Int(1),
						Duration:   pulumi.String("25920000s"),
						RankOrder:  pulumi.String("ASCENDING"),
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_data_source_id"),
					Type: pulumi.String("STRING_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance:          pulumi.Int(1),
						ValuesImportanceMap: nil,
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_document_title"),
					Type: pulumi.String("STRING_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(true),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(true),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance:          pulumi.Int(2),
						ValuesImportanceMap: nil,
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_excerpt_page_number"),
					Type: pulumi.String("LONG_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(false),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance: pulumi.Int(2),
						RankOrder:  pulumi.String("ASCENDING"),
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_faq_id"),
					Type: pulumi.String("STRING_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance:          pulumi.Int(1),
						ValuesImportanceMap: nil,
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_file_type"),
					Type: pulumi.String("STRING_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance:          pulumi.Int(1),
						ValuesImportanceMap: nil,
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_language_code"),
					Type: pulumi.String("STRING_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance:          pulumi.Int(1),
						ValuesImportanceMap: nil,
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_last_updated_at"),
					Type: pulumi.String("DATE_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Freshness:  pulumi.Bool(false),
						Importance: pulumi.Int(1),
						Duration:   pulumi.String("25920000s"),
						RankOrder:  pulumi.String("ASCENDING"),
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_source_uri"),
					Type: pulumi.String("STRING_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(true),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(false),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance:          pulumi.Int(1),
						ValuesImportanceMap: nil,
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_tenant_id"),
					Type: pulumi.String("STRING_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance:          pulumi.Int(1),
						ValuesImportanceMap: nil,
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_version"),
					Type: pulumi.String("STRING_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance:          pulumi.Int(1),
						ValuesImportanceMap: nil,
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("_view_count"),
					Type: pulumi.String("LONG_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(false),
						Facetable:   pulumi.Bool(false),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance: pulumi.Int(1),
						RankOrder:  pulumi.String("ASCENDING"),
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("example-string-value"),
					Type: pulumi.String("STRING_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(true),
						Facetable:   pulumi.Bool(true),
						Searchable:  pulumi.Bool(true),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance:          pulumi.Int(1),
						ValuesImportanceMap: nil,
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("example-long-value"),
					Type: pulumi.String("LONG_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(true),
						Facetable:   pulumi.Bool(true),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(true),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance: pulumi.Int(1),
						RankOrder:  pulumi.String("ASCENDING"),
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("example-string-list-value"),
					Type: pulumi.String("STRING_LIST_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(true),
						Facetable:   pulumi.Bool(true),
						Searchable:  pulumi.Bool(true),
						Sortable:    pulumi.Bool(false),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Importance: pulumi.Int(1),
					},
				},
				&kendra.IndexDocumentMetadataConfigurationUpdateArgs{
					Name: pulumi.String("example-date-value"),
					Type: pulumi.String("DATE_VALUE"),
					Search: &kendra.IndexDocumentMetadataConfigurationUpdateSearchArgs{
						Displayable: pulumi.Bool(true),
						Facetable:   pulumi.Bool(true),
						Searchable:  pulumi.Bool(false),
						Sortable:    pulumi.Bool(false),
					},
					Relevance: &kendra.IndexDocumentMetadataConfigurationUpdateRelevanceArgs{
						Freshness:  pulumi.Bool(false),
						Importance: pulumi.Int(1),
						Duration:   pulumi.String("25920000s"),
						RankOrder:  pulumi.String("ASCENDING"),
					},
				},
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
import com.pulumi.aws.kendra.Index;
import com.pulumi.aws.kendra.IndexArgs;
import com.pulumi.aws.kendra.inputs.IndexDocumentMetadataConfigurationUpdateArgs;
import com.pulumi.aws.kendra.inputs.IndexDocumentMetadataConfigurationUpdateSearchArgs;
import com.pulumi.aws.kendra.inputs.IndexDocumentMetadataConfigurationUpdateRelevanceArgs;
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
        var example = new Index("example", IndexArgs.builder()        
            .roleArn(aws_iam_role.this().arn())
            .documentMetadataConfigurationUpdates(            
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_authors")
                    .type("STRING_LIST_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(false)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_category")
                    .type("STRING_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .valuesImportanceMap()
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_created_at")
                    .type("DATE_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .freshness(false)
                        .importance(1)
                        .duration("25920000s")
                        .rankOrder("ASCENDING")
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_data_source_id")
                    .type("STRING_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .valuesImportanceMap()
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_document_title")
                    .type("STRING_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(true)
                        .facetable(false)
                        .searchable(true)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(2)
                        .valuesImportanceMap()
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_excerpt_page_number")
                    .type("LONG_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(false)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(2)
                        .rankOrder("ASCENDING")
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_faq_id")
                    .type("STRING_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .valuesImportanceMap()
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_file_type")
                    .type("STRING_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .valuesImportanceMap()
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_language_code")
                    .type("STRING_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .valuesImportanceMap()
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_last_updated_at")
                    .type("DATE_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .freshness(false)
                        .importance(1)
                        .duration("25920000s")
                        .rankOrder("ASCENDING")
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_source_uri")
                    .type("STRING_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(true)
                        .facetable(false)
                        .searchable(false)
                        .sortable(false)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .valuesImportanceMap()
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_tenant_id")
                    .type("STRING_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .valuesImportanceMap()
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_version")
                    .type("STRING_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .valuesImportanceMap()
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("_view_count")
                    .type("LONG_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(false)
                        .facetable(false)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .rankOrder("ASCENDING")
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("example-string-value")
                    .type("STRING_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(true)
                        .facetable(true)
                        .searchable(true)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .valuesImportanceMap()
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("example-long-value")
                    .type("LONG_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(true)
                        .facetable(true)
                        .searchable(false)
                        .sortable(true)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .rankOrder("ASCENDING")
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("example-string-list-value")
                    .type("STRING_LIST_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(true)
                        .facetable(true)
                        .searchable(true)
                        .sortable(false)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .importance(1)
                        .build())
                    .build(),
                IndexDocumentMetadataConfigurationUpdateArgs.builder()
                    .name("example-date-value")
                    .type("DATE_VALUE")
                    .search(IndexDocumentMetadataConfigurationUpdateSearchArgs.builder()
                        .displayable(true)
                        .facetable(true)
                        .searchable(false)
                        .sortable(false)
                        .build())
                    .relevance(IndexDocumentMetadataConfigurationUpdateRelevanceArgs.builder()
                        .freshness(false)
                        .importance(1)
                        .duration("25920000s")
                        .rankOrder("ASCENDING")
                        .build())
                    .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:Index
    properties:
      roleArn: ${aws_iam_role.this.arn}
      documentMetadataConfigurationUpdates:
        - name: _authors
          type: STRING_LIST_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: false
          relevance:
            importance: 1
        - name: _category
          type: STRING_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            importance: 1
            valuesImportanceMap: {}
        - name: _created_at
          type: DATE_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            freshness: false
            importance: 1
            duration: 25920000s
            rankOrder: ASCENDING
        - name: _data_source_id
          type: STRING_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            importance: 1
            valuesImportanceMap: {}
        - name: _document_title
          type: STRING_VALUE
          search:
            displayable: true
            facetable: false
            searchable: true
            sortable: true
          relevance:
            importance: 2
            valuesImportanceMap: {}
        - name: _excerpt_page_number
          type: LONG_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: false
          relevance:
            importance: 2
            rankOrder: ASCENDING
        - name: _faq_id
          type: STRING_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            importance: 1
            valuesImportanceMap: {}
        - name: _file_type
          type: STRING_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            importance: 1
            valuesImportanceMap: {}
        - name: _language_code
          type: STRING_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            importance: 1
            valuesImportanceMap: {}
        - name: _last_updated_at
          type: DATE_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            freshness: false
            importance: 1
            duration: 25920000s
            rankOrder: ASCENDING
        - name: _source_uri
          type: STRING_VALUE
          search:
            displayable: true
            facetable: false
            searchable: false
            sortable: false
          relevance:
            importance: 1
            valuesImportanceMap: {}
        - name: _tenant_id
          type: STRING_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            importance: 1
            valuesImportanceMap: {}
        - name: _version
          type: STRING_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            importance: 1
            valuesImportanceMap: {}
        - name: _view_count
          type: LONG_VALUE
          search:
            displayable: false
            facetable: false
            searchable: false
            sortable: true
          relevance:
            importance: 1
            rankOrder: ASCENDING
        - name: example-string-value
          type: STRING_VALUE
          search:
            displayable: true
            facetable: true
            searchable: true
            sortable: true
          relevance:
            importance: 1
            valuesImportanceMap: {}
        - name: example-long-value
          type: LONG_VALUE
          search:
            displayable: true
            facetable: true
            searchable: false
            sortable: true
          relevance:
            importance: 1
            rankOrder: ASCENDING
        - name: example-string-list-value
          type: STRING_LIST_VALUE
          search:
            displayable: true
            facetable: true
            searchable: true
            sortable: false
          relevance:
            importance: 1
        - name: example-date-value
          type: DATE_VALUE
          search:
            displayable: true
            facetable: true
            searchable: false
            sortable: false
          relevance:
            freshness: false
            importance: 1
            duration: 25920000s
            rankOrder: ASCENDING
```
{{% /example %}}
{{% example %}}
### With JSON token type configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.Index("example", {
    roleArn: aws_iam_role["this"].arn,
    userTokenConfigurations: {
        jsonTokenTypeConfiguration: {
            groupAttributeField: "groups",
            userNameAttributeField: "username",
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.Index("example",
    role_arn=aws_iam_role["this"]["arn"],
    user_token_configurations=aws.kendra.IndexUserTokenConfigurationsArgs(
        json_token_type_configuration=aws.kendra.IndexUserTokenConfigurationsJsonTokenTypeConfigurationArgs(
            group_attribute_field="groups",
            user_name_attribute_field="username",
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Kendra.Index("example", new()
    {
        RoleArn = aws_iam_role.This.Arn,
        UserTokenConfigurations = new Aws.Kendra.Inputs.IndexUserTokenConfigurationsArgs
        {
            JsonTokenTypeConfiguration = new Aws.Kendra.Inputs.IndexUserTokenConfigurationsJsonTokenTypeConfigurationArgs
            {
                GroupAttributeField = "groups",
                UserNameAttributeField = "username",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kendra"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kendra.NewIndex(ctx, "example", &kendra.IndexArgs{
			RoleArn: pulumi.Any(aws_iam_role.This.Arn),
			UserTokenConfigurations: &kendra.IndexUserTokenConfigurationsArgs{
				JsonTokenTypeConfiguration: &kendra.IndexUserTokenConfigurationsJsonTokenTypeConfigurationArgs{
					GroupAttributeField:    pulumi.String("groups"),
					UserNameAttributeField: pulumi.String("username"),
				},
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
import com.pulumi.aws.kendra.Index;
import com.pulumi.aws.kendra.IndexArgs;
import com.pulumi.aws.kendra.inputs.IndexUserTokenConfigurationsArgs;
import com.pulumi.aws.kendra.inputs.IndexUserTokenConfigurationsJsonTokenTypeConfigurationArgs;
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
        var example = new Index("example", IndexArgs.builder()        
            .roleArn(aws_iam_role.this().arn())
            .userTokenConfigurations(IndexUserTokenConfigurationsArgs.builder()
                .jsonTokenTypeConfiguration(IndexUserTokenConfigurationsJsonTokenTypeConfigurationArgs.builder()
                    .groupAttributeField("groups")
                    .userNameAttributeField("username")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:Index
    properties:
      roleArn: ${aws_iam_role.this.arn}
      userTokenConfigurations:
        jsonTokenTypeConfiguration:
          groupAttributeField: groups
          userNameAttributeField: username
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amazon Kendra Indexes using its `id`. For example:

```sh
 $ pulumi import aws:kendra/index:Index example 12345678-1234-5678-9123-123456789123
```
 