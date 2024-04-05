Provides an AppSync Function.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleGraphQLApi = new aws.appsync.GraphQLApi("exampleGraphQLApi", {
    authenticationType: "API_KEY",
    schema: `type Mutation {
  putPost(id: ID!, title: String!): Post
}

type Post {
  id: ID!
  title: String!
}

type Query {
  singlePost(id: ID!): Post
}

schema {
  query: Query
  mutation: Mutation
}
`,
});
const exampleDataSource = new aws.appsync.DataSource("exampleDataSource", {
    apiId: exampleGraphQLApi.id,
    name: "example",
    type: "HTTP",
    httpConfig: {
        endpoint: "http://example.com",
    },
});
const exampleFunction = new aws.appsync.Function("exampleFunction", {
    apiId: exampleGraphQLApi.id,
    dataSource: exampleDataSource.name,
    name: "example",
    requestMappingTemplate: `{
    "version": "2018-05-29",
    "method": "GET",
    "resourcePath": "/",
    "params":{
        "headers": $utils.http.copyheaders($ctx.request.headers)
    }
}
`,
    responseMappingTemplate: `#if($ctx.result.statusCode == 200)
    $ctx.result.body
#else
    $utils.appendError($ctx.result.body, $ctx.result.statusCode)
#end
`,
});
```
```python
import pulumi
import pulumi_aws as aws

example_graph_ql_api = aws.appsync.GraphQLApi("exampleGraphQLApi",
    authentication_type="API_KEY",
    schema="""type Mutation {
  putPost(id: ID!, title: String!): Post
}

type Post {
  id: ID!
  title: String!
}

type Query {
  singlePost(id: ID!): Post
}

schema {
  query: Query
  mutation: Mutation
}
""")
example_data_source = aws.appsync.DataSource("exampleDataSource",
    api_id=example_graph_ql_api.id,
    name="example",
    type="HTTP",
    http_config=aws.appsync.DataSourceHttpConfigArgs(
        endpoint="http://example.com",
    ))
example_function = aws.appsync.Function("exampleFunction",
    api_id=example_graph_ql_api.id,
    data_source=example_data_source.name,
    name="example",
    request_mapping_template="""{
    "version": "2018-05-29",
    "method": "GET",
    "resourcePath": "/",
    "params":{
        "headers": $utils.http.copyheaders($ctx.request.headers)
    }
}
""",
    response_mapping_template="""#if($ctx.result.statusCode == 200)
    $ctx.result.body
#else
    $utils.appendError($ctx.result.body, $ctx.result.statusCode)
#end
""")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleGraphQLApi = new Aws.AppSync.GraphQLApi("exampleGraphQLApi", new()
    {
        AuthenticationType = "API_KEY",
        Schema = @"type Mutation {
  putPost(id: ID!, title: String!): Post
}

type Post {
  id: ID!
  title: String!
}

type Query {
  singlePost(id: ID!): Post
}

schema {
  query: Query
  mutation: Mutation
}
",
    });

    var exampleDataSource = new Aws.AppSync.DataSource("exampleDataSource", new()
    {
        ApiId = exampleGraphQLApi.Id,
        Name = "example",
        Type = "HTTP",
        HttpConfig = new Aws.AppSync.Inputs.DataSourceHttpConfigArgs
        {
            Endpoint = "http://example.com",
        },
    });

    var exampleFunction = new Aws.AppSync.Function("exampleFunction", new()
    {
        ApiId = exampleGraphQLApi.Id,
        DataSource = exampleDataSource.Name,
        Name = "example",
        RequestMappingTemplate = @"{
    ""version"": ""2018-05-29"",
    ""method"": ""GET"",
    ""resourcePath"": ""/"",
    ""params"":{
        ""headers"": $utils.http.copyheaders($ctx.request.headers)
    }
}
",
        ResponseMappingTemplate = @"#if($ctx.result.statusCode == 200)
    $ctx.result.body
#else
    $utils.appendError($ctx.result.body, $ctx.result.statusCode)
#end
",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appsync"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleGraphQLApi, err := appsync.NewGraphQLApi(ctx, "exampleGraphQLApi", &appsync.GraphQLApiArgs{
			AuthenticationType: pulumi.String("API_KEY"),
			Schema: pulumi.String(`type Mutation {
  putPost(id: ID!, title: String!): Post
}

type Post {
  id: ID!
  title: String!
}

type Query {
  singlePost(id: ID!): Post
}

schema {
  query: Query
  mutation: Mutation
}
`),
		})
		if err != nil {
			return err
		}
		exampleDataSource, err := appsync.NewDataSource(ctx, "exampleDataSource", &appsync.DataSourceArgs{
			ApiId: exampleGraphQLApi.ID(),
			Name:  pulumi.String("example"),
			Type:  pulumi.String("HTTP"),
			HttpConfig: &appsync.DataSourceHttpConfigArgs{
				Endpoint: pulumi.String("http://example.com"),
			},
		})
		if err != nil {
			return err
		}
		_, err = appsync.NewFunction(ctx, "exampleFunction", &appsync.FunctionArgs{
			ApiId:      exampleGraphQLApi.ID(),
			DataSource: exampleDataSource.Name,
			Name:       pulumi.String("example"),
			RequestMappingTemplate: pulumi.String(`{
    "version": "2018-05-29",
    "method": "GET",
    "resourcePath": "/",
    "params":{
        "headers": $utils.http.copyheaders($ctx.request.headers)
    }
}
`),
			ResponseMappingTemplate: pulumi.String(`#if($ctx.result.statusCode == 200)
    $ctx.result.body
#else
    $utils.appendError($ctx.result.body, $ctx.result.statusCode)
#end
`),
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
import com.pulumi.aws.appsync.GraphQLApi;
import com.pulumi.aws.appsync.GraphQLApiArgs;
import com.pulumi.aws.appsync.DataSource;
import com.pulumi.aws.appsync.DataSourceArgs;
import com.pulumi.aws.appsync.inputs.DataSourceHttpConfigArgs;
import com.pulumi.aws.appsync.Function;
import com.pulumi.aws.appsync.FunctionArgs;
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
        var exampleGraphQLApi = new GraphQLApi("exampleGraphQLApi", GraphQLApiArgs.builder()        
            .authenticationType("API_KEY")
            .schema("""
type Mutation {
  putPost(id: ID!, title: String!): Post
}

type Post {
  id: ID!
  title: String!
}

type Query {
  singlePost(id: ID!): Post
}

schema {
  query: Query
  mutation: Mutation
}
            """)
            .build());

        var exampleDataSource = new DataSource("exampleDataSource", DataSourceArgs.builder()        
            .apiId(exampleGraphQLApi.id())
            .name("example")
            .type("HTTP")
            .httpConfig(DataSourceHttpConfigArgs.builder()
                .endpoint("http://example.com")
                .build())
            .build());

        var exampleFunction = new Function("exampleFunction", FunctionArgs.builder()        
            .apiId(exampleGraphQLApi.id())
            .dataSource(exampleDataSource.name())
            .name("example")
            .requestMappingTemplate("""
{
    "version": "2018-05-29",
    "method": "GET",
    "resourcePath": "/",
    "params":{
        "headers": $utils.http.copyheaders($ctx.request.headers)
    }
}
            """)
            .responseMappingTemplate("""
#if($ctx.result.statusCode == 200)
    $ctx.result.body
#else
    $utils.appendError($ctx.result.body, $ctx.result.statusCode)
#end
            """)
            .build());

    }
}
```
```yaml
resources:
  exampleGraphQLApi:
    type: aws:appsync:GraphQLApi
    properties:
      authenticationType: API_KEY
      schema: |
        type Mutation {
          putPost(id: ID!, title: String!): Post
        }

        type Post {
          id: ID!
          title: String!
        }

        type Query {
          singlePost(id: ID!): Post
        }

        schema {
          query: Query
          mutation: Mutation
        }
  exampleDataSource:
    type: aws:appsync:DataSource
    properties:
      apiId: ${exampleGraphQLApi.id}
      name: example
      type: HTTP
      httpConfig:
        endpoint: http://example.com
  exampleFunction:
    type: aws:appsync:Function
    properties:
      apiId: ${exampleGraphQLApi.id}
      dataSource: ${exampleDataSource.name}
      name: example
      requestMappingTemplate: |
        {
            "version": "2018-05-29",
            "method": "GET",
            "resourcePath": "/",
            "params":{
                "headers": $utils.http.copyheaders($ctx.request.headers)
            }
        }
      responseMappingTemplate: |
        #if($ctx.result.statusCode == 200)
            $ctx.result.body
        #else
            $utils.appendError($ctx.result.body, $ctx.result.statusCode)
        #end
```

{{% /example %}}
{{% example %}}
### With Code

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const example = new aws.appsync.Function("example", {
    apiId: aws_appsync_graphql_api.example.id,
    dataSource: aws_appsync_datasource.example.name,
    name: "example",
    code: fs.readFileSync("some-code-dir", "utf8"),
    runtime: {
        name: "APPSYNC_JS",
        runtimeVersion: "1.0.0",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.appsync.Function("example",
    api_id=aws_appsync_graphql_api["example"]["id"],
    data_source=aws_appsync_datasource["example"]["name"],
    name="example",
    code=(lambda path: open(path).read())("some-code-dir"),
    runtime=aws.appsync.FunctionRuntimeArgs(
        name="APPSYNC_JS",
        runtime_version="1.0.0",
    ))
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AppSync.Function("example", new()
    {
        ApiId = aws_appsync_graphql_api.Example.Id,
        DataSource = aws_appsync_datasource.Example.Name,
        Name = "example",
        Code = File.ReadAllText("some-code-dir"),
        Runtime = new Aws.AppSync.Inputs.FunctionRuntimeArgs
        {
            Name = "APPSYNC_JS",
            RuntimeVersion = "1.0.0",
        },
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appsync"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func readFileOrPanic(path string) pulumi.StringPtrInput {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return pulumi.String(string(data))
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appsync.NewFunction(ctx, "example", &appsync.FunctionArgs{
			ApiId:      pulumi.Any(aws_appsync_graphql_api.Example.Id),
			DataSource: pulumi.Any(aws_appsync_datasource.Example.Name),
			Name:       pulumi.String("example"),
			Code:       readFileOrPanic("some-code-dir"),
			Runtime: &appsync.FunctionRuntimeArgs{
				Name:           pulumi.String("APPSYNC_JS"),
				RuntimeVersion: pulumi.String("1.0.0"),
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
import com.pulumi.aws.appsync.Function;
import com.pulumi.aws.appsync.FunctionArgs;
import com.pulumi.aws.appsync.inputs.FunctionRuntimeArgs;
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
        var example = new Function("example", FunctionArgs.builder()        
            .apiId(aws_appsync_graphql_api.example().id())
            .dataSource(aws_appsync_datasource.example().name())
            .name("example")
            .code(Files.readString(Paths.get("some-code-dir")))
            .runtime(FunctionRuntimeArgs.builder()
                .name("APPSYNC_JS")
                .runtimeVersion("1.0.0")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:appsync:Function
    properties:
      apiId: ${aws_appsync_graphql_api.example.id}
      dataSource: ${aws_appsync_datasource.example.name}
      name: example
      code:
        fn::readFile: some-code-dir
      runtime:
        name: APPSYNC_JS
        runtimeVersion: 1.0.0
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_appsync_function` using the AppSync API ID and Function ID separated by `-`. For example:

```sh
 $ pulumi import aws:appsync/function:Function example xxxxx-yyyyy
```
 