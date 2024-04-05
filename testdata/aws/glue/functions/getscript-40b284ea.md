Use this data source to generate a Glue script from a Directed Acyclic Graph (DAG).

{{% examples %}}
## Example Usage
{{% example %}}
### Generate Python Script

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.glue.getScript({
    language: "PYTHON",
    dagEdges: [
        {
            source: "datasource0",
            target: "applymapping1",
        },
        {
            source: "applymapping1",
            target: "selectfields2",
        },
        {
            source: "selectfields2",
            target: "resolvechoice3",
        },
        {
            source: "resolvechoice3",
            target: "datasink4",
        },
    ],
    dagNodes: [
        {
            id: "datasource0",
            nodeType: "DataSource",
            args: [
                {
                    name: "database",
                    value: `"${aws_glue_catalog_database.source.name}"`,
                },
                {
                    name: "table_name",
                    value: `"${aws_glue_catalog_table.source.name}"`,
                },
            ],
        },
        {
            id: "applymapping1",
            nodeType: "ApplyMapping",
            args: [{
                name: "mapping",
                value: "[(\"column1\", \"string\", \"column1\", \"string\")]",
            }],
        },
        {
            id: "selectfields2",
            nodeType: "SelectFields",
            args: [{
                name: "paths",
                value: "[\"column1\"]",
            }],
        },
        {
            id: "resolvechoice3",
            nodeType: "ResolveChoice",
            args: [
                {
                    name: "choice",
                    value: "\"MATCH_CATALOG\"",
                },
                {
                    name: "database",
                    value: `"${aws_glue_catalog_database.destination.name}"`,
                },
                {
                    name: "table_name",
                    value: `"${aws_glue_catalog_table.destination.name}"`,
                },
            ],
        },
        {
            id: "datasink4",
            nodeType: "DataSink",
            args: [
                {
                    name: "database",
                    value: `"${aws_glue_catalog_database.destination.name}"`,
                },
                {
                    name: "table_name",
                    value: `"${aws_glue_catalog_table.destination.name}"`,
                },
            ],
        },
    ],
});
export const pythonScript = example.then(example => example.pythonScript);
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.get_script(language="PYTHON",
    dag_edges=[
        aws.glue.GetScriptDagEdgeArgs(
            source="datasource0",
            target="applymapping1",
        ),
        aws.glue.GetScriptDagEdgeArgs(
            source="applymapping1",
            target="selectfields2",
        ),
        aws.glue.GetScriptDagEdgeArgs(
            source="selectfields2",
            target="resolvechoice3",
        ),
        aws.glue.GetScriptDagEdgeArgs(
            source="resolvechoice3",
            target="datasink4",
        ),
    ],
    dag_nodes=[
        aws.glue.GetScriptDagNodeArgs(
            id="datasource0",
            node_type="DataSource",
            args=[
                aws.glue.GetScriptDagNodeArgArgs(
                    name="database",
                    value=f"\"{aws_glue_catalog_database['source']['name']}\"",
                ),
                aws.glue.GetScriptDagNodeArgArgs(
                    name="table_name",
                    value=f"\"{aws_glue_catalog_table['source']['name']}\"",
                ),
            ],
        ),
        aws.glue.GetScriptDagNodeArgs(
            id="applymapping1",
            node_type="ApplyMapping",
            args=[aws.glue.GetScriptDagNodeArgArgs(
                name="mapping",
                value="[(\"column1\", \"string\", \"column1\", \"string\")]",
            )],
        ),
        aws.glue.GetScriptDagNodeArgs(
            id="selectfields2",
            node_type="SelectFields",
            args=[aws.glue.GetScriptDagNodeArgArgs(
                name="paths",
                value="[\"column1\"]",
            )],
        ),
        aws.glue.GetScriptDagNodeArgs(
            id="resolvechoice3",
            node_type="ResolveChoice",
            args=[
                aws.glue.GetScriptDagNodeArgArgs(
                    name="choice",
                    value="\"MATCH_CATALOG\"",
                ),
                aws.glue.GetScriptDagNodeArgArgs(
                    name="database",
                    value=f"\"{aws_glue_catalog_database['destination']['name']}\"",
                ),
                aws.glue.GetScriptDagNodeArgArgs(
                    name="table_name",
                    value=f"\"{aws_glue_catalog_table['destination']['name']}\"",
                ),
            ],
        ),
        aws.glue.GetScriptDagNodeArgs(
            id="datasink4",
            node_type="DataSink",
            args=[
                aws.glue.GetScriptDagNodeArgArgs(
                    name="database",
                    value=f"\"{aws_glue_catalog_database['destination']['name']}\"",
                ),
                aws.glue.GetScriptDagNodeArgArgs(
                    name="table_name",
                    value=f"\"{aws_glue_catalog_table['destination']['name']}\"",
                ),
            ],
        ),
    ])
pulumi.export("pythonScript", example.python_script)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Glue.GetScript.Invoke(new()
    {
        Language = "PYTHON",
        DagEdges = new[]
        {
            new Aws.Glue.Inputs.GetScriptDagEdgeInputArgs
            {
                Source = "datasource0",
                Target = "applymapping1",
            },
            new Aws.Glue.Inputs.GetScriptDagEdgeInputArgs
            {
                Source = "applymapping1",
                Target = "selectfields2",
            },
            new Aws.Glue.Inputs.GetScriptDagEdgeInputArgs
            {
                Source = "selectfields2",
                Target = "resolvechoice3",
            },
            new Aws.Glue.Inputs.GetScriptDagEdgeInputArgs
            {
                Source = "resolvechoice3",
                Target = "datasink4",
            },
        },
        DagNodes = new[]
        {
            new Aws.Glue.Inputs.GetScriptDagNodeInputArgs
            {
                Id = "datasource0",
                NodeType = "DataSource",
                Args = new[]
                {
                    new Aws.Glue.Inputs.GetScriptDagNodeArgInputArgs
                    {
                        Name = "database",
                        Value = $"\"{aws_glue_catalog_database.Source.Name}\"",
                    },
                    new Aws.Glue.Inputs.GetScriptDagNodeArgInputArgs
                    {
                        Name = "table_name",
                        Value = $"\"{aws_glue_catalog_table.Source.Name}\"",
                    },
                },
            },
            new Aws.Glue.Inputs.GetScriptDagNodeInputArgs
            {
                Id = "applymapping1",
                NodeType = "ApplyMapping",
                Args = new[]
                {
                    new Aws.Glue.Inputs.GetScriptDagNodeArgInputArgs
                    {
                        Name = "mapping",
                        Value = "[(\"column1\", \"string\", \"column1\", \"string\")]",
                    },
                },
            },
            new Aws.Glue.Inputs.GetScriptDagNodeInputArgs
            {
                Id = "selectfields2",
                NodeType = "SelectFields",
                Args = new[]
                {
                    new Aws.Glue.Inputs.GetScriptDagNodeArgInputArgs
                    {
                        Name = "paths",
                        Value = "[\"column1\"]",
                    },
                },
            },
            new Aws.Glue.Inputs.GetScriptDagNodeInputArgs
            {
                Id = "resolvechoice3",
                NodeType = "ResolveChoice",
                Args = new[]
                {
                    new Aws.Glue.Inputs.GetScriptDagNodeArgInputArgs
                    {
                        Name = "choice",
                        Value = "\"MATCH_CATALOG\"",
                    },
                    new Aws.Glue.Inputs.GetScriptDagNodeArgInputArgs
                    {
                        Name = "database",
                        Value = $"\"{aws_glue_catalog_database.Destination.Name}\"",
                    },
                    new Aws.Glue.Inputs.GetScriptDagNodeArgInputArgs
                    {
                        Name = "table_name",
                        Value = $"\"{aws_glue_catalog_table.Destination.Name}\"",
                    },
                },
            },
            new Aws.Glue.Inputs.GetScriptDagNodeInputArgs
            {
                Id = "datasink4",
                NodeType = "DataSink",
                Args = new[]
                {
                    new Aws.Glue.Inputs.GetScriptDagNodeArgInputArgs
                    {
                        Name = "database",
                        Value = $"\"{aws_glue_catalog_database.Destination.Name}\"",
                    },
                    new Aws.Glue.Inputs.GetScriptDagNodeArgInputArgs
                    {
                        Name = "table_name",
                        Value = $"\"{aws_glue_catalog_table.Destination.Name}\"",
                    },
                },
            },
        },
    });

    return new Dictionary<string, object?>
    {
        ["pythonScript"] = example.Apply(getScriptResult => getScriptResult.PythonScript),
    };
});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := glue.GetScript(ctx, &glue.GetScriptArgs{
			Language: pulumi.StringRef("PYTHON"),
			DagEdges: []glue.GetScriptDagEdge{
				{
					Source: "datasource0",
					Target: "applymapping1",
				},
				{
					Source: "applymapping1",
					Target: "selectfields2",
				},
				{
					Source: "selectfields2",
					Target: "resolvechoice3",
				},
				{
					Source: "resolvechoice3",
					Target: "datasink4",
				},
			},
			DagNodes: []glue.GetScriptDagNode{
				{
					Id:       "datasource0",
					NodeType: "DataSource",
					Args: []glue.GetScriptDagNodeArg{
						{
							Name:  "database",
							Value: fmt.Sprintf("\"%v\"", aws_glue_catalog_database.Source.Name),
						},
						{
							Name:  "table_name",
							Value: fmt.Sprintf("\"%v\"", aws_glue_catalog_table.Source.Name),
						},
					},
				},
				{
					Id:       "applymapping1",
					NodeType: "ApplyMapping",
					Args: []glue.GetScriptDagNodeArg{
						{
							Name:  "mapping",
							Value: "[(\"column1\", \"string\", \"column1\", \"string\")]",
						},
					},
				},
				{
					Id:       "selectfields2",
					NodeType: "SelectFields",
					Args: []glue.GetScriptDagNodeArg{
						{
							Name:  "paths",
							Value: "[\"column1\"]",
						},
					},
				},
				{
					Id:       "resolvechoice3",
					NodeType: "ResolveChoice",
					Args: []glue.GetScriptDagNodeArg{
						{
							Name:  "choice",
							Value: "\"MATCH_CATALOG\"",
						},
						{
							Name:  "database",
							Value: fmt.Sprintf("\"%v\"", aws_glue_catalog_database.Destination.Name),
						},
						{
							Name:  "table_name",
							Value: fmt.Sprintf("\"%v\"", aws_glue_catalog_table.Destination.Name),
						},
					},
				},
				{
					Id:       "datasink4",
					NodeType: "DataSink",
					Args: []glue.GetScriptDagNodeArg{
						{
							Name:  "database",
							Value: fmt.Sprintf("\"%v\"", aws_glue_catalog_database.Destination.Name),
						},
						{
							Name:  "table_name",
							Value: fmt.Sprintf("\"%v\"", aws_glue_catalog_table.Destination.Name),
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("pythonScript", example.PythonScript)
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.glue.GlueFunctions;
import com.pulumi.aws.glue.inputs.GetScriptArgs;
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
        final var example = GlueFunctions.getScript(GetScriptArgs.builder()
            .language("PYTHON")
            .dagEdges(            
                GetScriptDagEdgeArgs.builder()
                    .source("datasource0")
                    .target("applymapping1")
                    .build(),
                GetScriptDagEdgeArgs.builder()
                    .source("applymapping1")
                    .target("selectfields2")
                    .build(),
                GetScriptDagEdgeArgs.builder()
                    .source("selectfields2")
                    .target("resolvechoice3")
                    .build(),
                GetScriptDagEdgeArgs.builder()
                    .source("resolvechoice3")
                    .target("datasink4")
                    .build())
            .dagNodes(            
                GetScriptDagNodeArgs.builder()
                    .id("datasource0")
                    .nodeType("DataSource")
                    .args(                    
                        GetScriptDagNodeArgArgs.builder()
                            .name("database")
                            .value(String.format("\"%s\"", aws_glue_catalog_database.source().name()))
                            .build(),
                        GetScriptDagNodeArgArgs.builder()
                            .name("table_name")
                            .value(String.format("\"%s\"", aws_glue_catalog_table.source().name()))
                            .build())
                    .build(),
                GetScriptDagNodeArgs.builder()
                    .id("applymapping1")
                    .nodeType("ApplyMapping")
                    .args(GetScriptDagNodeArgArgs.builder()
                        .name("mapping")
                        .value("[(\"column1\", \"string\", \"column1\", \"string\")]")
                        .build())
                    .build(),
                GetScriptDagNodeArgs.builder()
                    .id("selectfields2")
                    .nodeType("SelectFields")
                    .args(GetScriptDagNodeArgArgs.builder()
                        .name("paths")
                        .value("[\"column1\"]")
                        .build())
                    .build(),
                GetScriptDagNodeArgs.builder()
                    .id("resolvechoice3")
                    .nodeType("ResolveChoice")
                    .args(                    
                        GetScriptDagNodeArgArgs.builder()
                            .name("choice")
                            .value("\"MATCH_CATALOG\"")
                            .build(),
                        GetScriptDagNodeArgArgs.builder()
                            .name("database")
                            .value(String.format("\"%s\"", aws_glue_catalog_database.destination().name()))
                            .build(),
                        GetScriptDagNodeArgArgs.builder()
                            .name("table_name")
                            .value(String.format("\"%s\"", aws_glue_catalog_table.destination().name()))
                            .build())
                    .build(),
                GetScriptDagNodeArgs.builder()
                    .id("datasink4")
                    .nodeType("DataSink")
                    .args(                    
                        GetScriptDagNodeArgArgs.builder()
                            .name("database")
                            .value(String.format("\"%s\"", aws_glue_catalog_database.destination().name()))
                            .build(),
                        GetScriptDagNodeArgArgs.builder()
                            .name("table_name")
                            .value(String.format("\"%s\"", aws_glue_catalog_table.destination().name()))
                            .build())
                    .build())
            .build());

        ctx.export("pythonScript", example.applyValue(getScriptResult -> getScriptResult.pythonScript()));
    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:glue:getScript
      Arguments:
        language: PYTHON
        dagEdges:
          - source: datasource0
            target: applymapping1
          - source: applymapping1
            target: selectfields2
          - source: selectfields2
            target: resolvechoice3
          - source: resolvechoice3
            target: datasink4
        dagNodes:
          - id: datasource0
            nodeType: DataSource
            args:
              - name: database
                value: '"${aws_glue_catalog_database.source.name}"'
              - name: table_name
                value: '"${aws_glue_catalog_table.source.name}"'
          - id: applymapping1
            nodeType: ApplyMapping
            args:
              - name: mapping
                value: '[("column1", "string", "column1", "string")]'
          - id: selectfields2
            nodeType: SelectFields
            args:
              - name: paths
                value: '["column1"]'
          - id: resolvechoice3
            nodeType: ResolveChoice
            args:
              - name: choice
                value: '"MATCH_CATALOG"'
              - name: database
                value: '"${aws_glue_catalog_database.destination.name}"'
              - name: table_name
                value: '"${aws_glue_catalog_table.destination.name}"'
          - id: datasink4
            nodeType: DataSink
            args:
              - name: database
                value: '"${aws_glue_catalog_database.destination.name}"'
              - name: table_name
                value: '"${aws_glue_catalog_table.destination.name}"'
outputs:
  pythonScript: ${example.pythonScript}
```
{{% /example %}}
{{% example %}}
### Generate Scala Code

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.glue.getScript({
    language: "SCALA",
    dagEdges: [
        {
            source: "datasource0",
            target: "applymapping1",
        },
        {
            source: "applymapping1",
            target: "selectfields2",
        },
        {
            source: "selectfields2",
            target: "resolvechoice3",
        },
        {
            source: "resolvechoice3",
            target: "datasink4",
        },
    ],
    dagNodes: [
        {
            id: "datasource0",
            nodeType: "DataSource",
            args: [
                {
                    name: "database",
                    value: `"${aws_glue_catalog_database.source.name}"`,
                },
                {
                    name: "table_name",
                    value: `"${aws_glue_catalog_table.source.name}"`,
                },
            ],
        },
        {
            id: "applymapping1",
            nodeType: "ApplyMapping",
            args: [{
                name: "mappings",
                value: "[(\"column1\", \"string\", \"column1\", \"string\")]",
            }],
        },
        {
            id: "selectfields2",
            nodeType: "SelectFields",
            args: [{
                name: "paths",
                value: "[\"column1\"]",
            }],
        },
        {
            id: "resolvechoice3",
            nodeType: "ResolveChoice",
            args: [
                {
                    name: "choice",
                    value: "\"MATCH_CATALOG\"",
                },
                {
                    name: "database",
                    value: `"${aws_glue_catalog_database.destination.name}"`,
                },
                {
                    name: "table_name",
                    value: `"${aws_glue_catalog_table.destination.name}"`,
                },
            ],
        },
        {
            id: "datasink4",
            nodeType: "DataSink",
            args: [
                {
                    name: "database",
                    value: `"${aws_glue_catalog_database.destination.name}"`,
                },
                {
                    name: "table_name",
                    value: `"${aws_glue_catalog_table.destination.name}"`,
                },
            ],
        },
    ],
});
export const scalaCode = example.then(example => example.scalaCode);
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.get_script(language="SCALA",
    dag_edges=[
        aws.glue.GetScriptDagEdgeArgs(
            source="datasource0",
            target="applymapping1",
        ),
        aws.glue.GetScriptDagEdgeArgs(
            source="applymapping1",
            target="selectfields2",
        ),
        aws.glue.GetScriptDagEdgeArgs(
            source="selectfields2",
            target="resolvechoice3",
        ),
        aws.glue.GetScriptDagEdgeArgs(
            source="resolvechoice3",
            target="datasink4",
        ),
    ],
    dag_nodes=[
        aws.glue.GetScriptDagNodeArgs(
            id="datasource0",
            node_type="DataSource",
            args=[
                aws.glue.GetScriptDagNodeArgArgs(
                    name="database",
                    value=f"\"{aws_glue_catalog_database['source']['name']}\"",
                ),
                aws.glue.GetScriptDagNodeArgArgs(
                    name="table_name",
                    value=f"\"{aws_glue_catalog_table['source']['name']}\"",
                ),
            ],
        ),
        aws.glue.GetScriptDagNodeArgs(
            id="applymapping1",
            node_type="ApplyMapping",
            args=[aws.glue.GetScriptDagNodeArgArgs(
                name="mappings",
                value="[(\"column1\", \"string\", \"column1\", \"string\")]",
            )],
        ),
        aws.glue.GetScriptDagNodeArgs(
            id="selectfields2",
            node_type="SelectFields",
            args=[aws.glue.GetScriptDagNodeArgArgs(
                name="paths",
                value="[\"column1\"]",
            )],
        ),
        aws.glue.GetScriptDagNodeArgs(
            id="resolvechoice3",
            node_type="ResolveChoice",
            args=[
                aws.glue.GetScriptDagNodeArgArgs(
                    name="choice",
                    value="\"MATCH_CATALOG\"",
                ),
                aws.glue.GetScriptDagNodeArgArgs(
                    name="database",
                    value=f"\"{aws_glue_catalog_database['destination']['name']}\"",
                ),
                aws.glue.GetScriptDagNodeArgArgs(
                    name="table_name",
                    value=f"\"{aws_glue_catalog_table['destination']['name']}\"",
                ),
            ],
        ),
        aws.glue.GetScriptDagNodeArgs(
            id="datasink4",
            node_type="DataSink",
            args=[
                aws.glue.GetScriptDagNodeArgArgs(
                    name="database",
                    value=f"\"{aws_glue_catalog_database['destination']['name']}\"",
                ),
                aws.glue.GetScriptDagNodeArgArgs(
                    name="table_name",
                    value=f"\"{aws_glue_catalog_table['destination']['name']}\"",
                ),
            ],
        ),
    ])
pulumi.export("scalaCode", example.scala_code)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Glue.GetScript.Invoke(new()
    {
        Language = "SCALA",
        DagEdges = new[]
        {
            new Aws.Glue.Inputs.GetScriptDagEdgeInputArgs
            {
                Source = "datasource0",
                Target = "applymapping1",
            },
            new Aws.Glue.Inputs.GetScriptDagEdgeInputArgs
            {
                Source = "applymapping1",
                Target = "selectfields2",
            },
            new Aws.Glue.Inputs.GetScriptDagEdgeInputArgs
            {
                Source = "selectfields2",
                Target = "resolvechoice3",
            },
            new Aws.Glue.Inputs.GetScriptDagEdgeInputArgs
            {
                Source = "resolvechoice3",
                Target = "datasink4",
            },
        },
        DagNodes = new[]
        {
            new Aws.Glue.Inputs.GetScriptDagNodeInputArgs
            {
                Id = "datasource0",
                NodeType = "DataSource",
                Args = new[]
                {
                    new Aws.Glue.Inputs.GetScriptDagNodeArgInputArgs
                    {
                        Name = "database",
                        Value = $"\"{aws_glue_catalog_database.Source.Name}\"",
                    },
                    new Aws.Glue.Inputs.GetScriptDagNodeArgInputArgs
                    {
                        Name = "table_name",
                        Value = $"\"{aws_glue_catalog_table.Source.Name}\"",
                    },
                },
            },
            new Aws.Glue.Inputs.GetScriptDagNodeInputArgs
            {
                Id = "applymapping1",
                NodeType = "ApplyMapping",
                Args = new[]
                {
                    new Aws.Glue.Inputs.GetScriptDagNodeArgInputArgs
                    {
                        Name = "mappings",
                        Value = "[(\"column1\", \"string\", \"column1\", \"string\")]",
                    },
                },
            },
            new Aws.Glue.Inputs.GetScriptDagNodeInputArgs
            {
                Id = "selectfields2",
                NodeType = "SelectFields",
                Args = new[]
                {
                    new Aws.Glue.Inputs.GetScriptDagNodeArgInputArgs
                    {
                        Name = "paths",
                        Value = "[\"column1\"]",
                    },
                },
            },
            new Aws.Glue.Inputs.GetScriptDagNodeInputArgs
            {
                Id = "resolvechoice3",
                NodeType = "ResolveChoice",
                Args = new[]
                {
                    new Aws.Glue.Inputs.GetScriptDagNodeArgInputArgs
                    {
                        Name = "choice",
                        Value = "\"MATCH_CATALOG\"",
                    },
                    new Aws.Glue.Inputs.GetScriptDagNodeArgInputArgs
                    {
                        Name = "database",
                        Value = $"\"{aws_glue_catalog_database.Destination.Name}\"",
                    },
                    new Aws.Glue.Inputs.GetScriptDagNodeArgInputArgs
                    {
                        Name = "table_name",
                        Value = $"\"{aws_glue_catalog_table.Destination.Name}\"",
                    },
                },
            },
            new Aws.Glue.Inputs.GetScriptDagNodeInputArgs
            {
                Id = "datasink4",
                NodeType = "DataSink",
                Args = new[]
                {
                    new Aws.Glue.Inputs.GetScriptDagNodeArgInputArgs
                    {
                        Name = "database",
                        Value = $"\"{aws_glue_catalog_database.Destination.Name}\"",
                    },
                    new Aws.Glue.Inputs.GetScriptDagNodeArgInputArgs
                    {
                        Name = "table_name",
                        Value = $"\"{aws_glue_catalog_table.Destination.Name}\"",
                    },
                },
            },
        },
    });

    return new Dictionary<string, object?>
    {
        ["scalaCode"] = example.Apply(getScriptResult => getScriptResult.ScalaCode),
    };
});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := glue.GetScript(ctx, &glue.GetScriptArgs{
			Language: pulumi.StringRef("SCALA"),
			DagEdges: []glue.GetScriptDagEdge{
				{
					Source: "datasource0",
					Target: "applymapping1",
				},
				{
					Source: "applymapping1",
					Target: "selectfields2",
				},
				{
					Source: "selectfields2",
					Target: "resolvechoice3",
				},
				{
					Source: "resolvechoice3",
					Target: "datasink4",
				},
			},
			DagNodes: []glue.GetScriptDagNode{
				{
					Id:       "datasource0",
					NodeType: "DataSource",
					Args: []glue.GetScriptDagNodeArg{
						{
							Name:  "database",
							Value: fmt.Sprintf("\"%v\"", aws_glue_catalog_database.Source.Name),
						},
						{
							Name:  "table_name",
							Value: fmt.Sprintf("\"%v\"", aws_glue_catalog_table.Source.Name),
						},
					},
				},
				{
					Id:       "applymapping1",
					NodeType: "ApplyMapping",
					Args: []glue.GetScriptDagNodeArg{
						{
							Name:  "mappings",
							Value: "[(\"column1\", \"string\", \"column1\", \"string\")]",
						},
					},
				},
				{
					Id:       "selectfields2",
					NodeType: "SelectFields",
					Args: []glue.GetScriptDagNodeArg{
						{
							Name:  "paths",
							Value: "[\"column1\"]",
						},
					},
				},
				{
					Id:       "resolvechoice3",
					NodeType: "ResolveChoice",
					Args: []glue.GetScriptDagNodeArg{
						{
							Name:  "choice",
							Value: "\"MATCH_CATALOG\"",
						},
						{
							Name:  "database",
							Value: fmt.Sprintf("\"%v\"", aws_glue_catalog_database.Destination.Name),
						},
						{
							Name:  "table_name",
							Value: fmt.Sprintf("\"%v\"", aws_glue_catalog_table.Destination.Name),
						},
					},
				},
				{
					Id:       "datasink4",
					NodeType: "DataSink",
					Args: []glue.GetScriptDagNodeArg{
						{
							Name:  "database",
							Value: fmt.Sprintf("\"%v\"", aws_glue_catalog_database.Destination.Name),
						},
						{
							Name:  "table_name",
							Value: fmt.Sprintf("\"%v\"", aws_glue_catalog_table.Destination.Name),
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("scalaCode", example.ScalaCode)
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.glue.GlueFunctions;
import com.pulumi.aws.glue.inputs.GetScriptArgs;
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
        final var example = GlueFunctions.getScript(GetScriptArgs.builder()
            .language("SCALA")
            .dagEdges(            
                GetScriptDagEdgeArgs.builder()
                    .source("datasource0")
                    .target("applymapping1")
                    .build(),
                GetScriptDagEdgeArgs.builder()
                    .source("applymapping1")
                    .target("selectfields2")
                    .build(),
                GetScriptDagEdgeArgs.builder()
                    .source("selectfields2")
                    .target("resolvechoice3")
                    .build(),
                GetScriptDagEdgeArgs.builder()
                    .source("resolvechoice3")
                    .target("datasink4")
                    .build())
            .dagNodes(            
                GetScriptDagNodeArgs.builder()
                    .id("datasource0")
                    .nodeType("DataSource")
                    .args(                    
                        GetScriptDagNodeArgArgs.builder()
                            .name("database")
                            .value(String.format("\"%s\"", aws_glue_catalog_database.source().name()))
                            .build(),
                        GetScriptDagNodeArgArgs.builder()
                            .name("table_name")
                            .value(String.format("\"%s\"", aws_glue_catalog_table.source().name()))
                            .build())
                    .build(),
                GetScriptDagNodeArgs.builder()
                    .id("applymapping1")
                    .nodeType("ApplyMapping")
                    .args(GetScriptDagNodeArgArgs.builder()
                        .name("mappings")
                        .value("[(\"column1\", \"string\", \"column1\", \"string\")]")
                        .build())
                    .build(),
                GetScriptDagNodeArgs.builder()
                    .id("selectfields2")
                    .nodeType("SelectFields")
                    .args(GetScriptDagNodeArgArgs.builder()
                        .name("paths")
                        .value("[\"column1\"]")
                        .build())
                    .build(),
                GetScriptDagNodeArgs.builder()
                    .id("resolvechoice3")
                    .nodeType("ResolveChoice")
                    .args(                    
                        GetScriptDagNodeArgArgs.builder()
                            .name("choice")
                            .value("\"MATCH_CATALOG\"")
                            .build(),
                        GetScriptDagNodeArgArgs.builder()
                            .name("database")
                            .value(String.format("\"%s\"", aws_glue_catalog_database.destination().name()))
                            .build(),
                        GetScriptDagNodeArgArgs.builder()
                            .name("table_name")
                            .value(String.format("\"%s\"", aws_glue_catalog_table.destination().name()))
                            .build())
                    .build(),
                GetScriptDagNodeArgs.builder()
                    .id("datasink4")
                    .nodeType("DataSink")
                    .args(                    
                        GetScriptDagNodeArgArgs.builder()
                            .name("database")
                            .value(String.format("\"%s\"", aws_glue_catalog_database.destination().name()))
                            .build(),
                        GetScriptDagNodeArgArgs.builder()
                            .name("table_name")
                            .value(String.format("\"%s\"", aws_glue_catalog_table.destination().name()))
                            .build())
                    .build())
            .build());

        ctx.export("scalaCode", example.applyValue(getScriptResult -> getScriptResult.scalaCode()));
    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:glue:getScript
      Arguments:
        language: SCALA
        dagEdges:
          - source: datasource0
            target: applymapping1
          - source: applymapping1
            target: selectfields2
          - source: selectfields2
            target: resolvechoice3
          - source: resolvechoice3
            target: datasink4
        dagNodes:
          - id: datasource0
            nodeType: DataSource
            args:
              - name: database
                value: '"${aws_glue_catalog_database.source.name}"'
              - name: table_name
                value: '"${aws_glue_catalog_table.source.name}"'
          - id: applymapping1
            nodeType: ApplyMapping
            args:
              - name: mappings
                value: '[("column1", "string", "column1", "string")]'
          - id: selectfields2
            nodeType: SelectFields
            args:
              - name: paths
                value: '["column1"]'
          - id: resolvechoice3
            nodeType: ResolveChoice
            args:
              - name: choice
                value: '"MATCH_CATALOG"'
              - name: database
                value: '"${aws_glue_catalog_database.destination.name}"'
              - name: table_name
                value: '"${aws_glue_catalog_table.destination.name}"'
          - id: datasink4
            nodeType: DataSink
            args:
              - name: database
                value: '"${aws_glue_catalog_database.destination.name}"'
              - name: table_name
                value: '"${aws_glue_catalog_table.destination.name}"'
outputs:
  scalaCode: ${example.scalaCode}
```
{{% /example %}}
{{% /examples %}}