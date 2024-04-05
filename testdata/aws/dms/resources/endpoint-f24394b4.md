Provides a DMS (Data Migration Service) endpoint resource. DMS endpoints can be created, updated, deleted, and imported.

> **Note:** All arguments including the password will be stored in the raw state as plain-text. > **Note:** The `s3_settings` argument is deprecated, may not be maintained, and will be removed in a future version. Use the `aws.dms.S3Endpoint` resource instead.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new endpoint
const test = new aws.dms.Endpoint("test", {
    certificateArn: "arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012",
    databaseName: "test",
    endpointId: "test-dms-endpoint-tf",
    endpointType: "source",
    engineName: "aurora",
    extraConnectionAttributes: "",
    kmsKeyArn: "arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012",
    password: "test",
    port: 3306,
    serverName: "test",
    sslMode: "none",
    tags: {
        Name: "test",
    },
    username: "test",
});
```
```python
import pulumi
import pulumi_aws as aws

# Create a new endpoint
test = aws.dms.Endpoint("test",
    certificate_arn="arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012",
    database_name="test",
    endpoint_id="test-dms-endpoint-tf",
    endpoint_type="source",
    engine_name="aurora",
    extra_connection_attributes="",
    kms_key_arn="arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012",
    password="test",
    port=3306,
    server_name="test",
    ssl_mode="none",
    tags={
        "Name": "test",
    },
    username="test")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // Create a new endpoint
    var test = new Aws.Dms.Endpoint("test", new()
    {
        CertificateArn = "arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012",
        DatabaseName = "test",
        EndpointId = "test-dms-endpoint-tf",
        EndpointType = "source",
        EngineName = "aurora",
        ExtraConnectionAttributes = "",
        KmsKeyArn = "arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012",
        Password = "test",
        Port = 3306,
        ServerName = "test",
        SslMode = "none",
        Tags = 
        {
            { "Name", "test" },
        },
        Username = "test",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dms.NewEndpoint(ctx, "test", &dms.EndpointArgs{
			CertificateArn:            pulumi.String("arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012"),
			DatabaseName:              pulumi.String("test"),
			EndpointId:                pulumi.String("test-dms-endpoint-tf"),
			EndpointType:              pulumi.String("source"),
			EngineName:                pulumi.String("aurora"),
			ExtraConnectionAttributes: pulumi.String(""),
			KmsKeyArn:                 pulumi.String("arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012"),
			Password:                  pulumi.String("test"),
			Port:                      pulumi.Int(3306),
			ServerName:                pulumi.String("test"),
			SslMode:                   pulumi.String("none"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("test"),
			},
			Username: pulumi.String("test"),
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
import com.pulumi.aws.dms.Endpoint;
import com.pulumi.aws.dms.EndpointArgs;
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
        var test = new Endpoint("test", EndpointArgs.builder()        
            .certificateArn("arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012")
            .databaseName("test")
            .endpointId("test-dms-endpoint-tf")
            .endpointType("source")
            .engineName("aurora")
            .extraConnectionAttributes("")
            .kmsKeyArn("arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012")
            .password("test")
            .port(3306)
            .serverName("test")
            .sslMode("none")
            .tags(Map.of("Name", "test"))
            .username("test")
            .build());

    }
}
```
```yaml
resources:
  # Create a new endpoint
  test:
    type: aws:dms:Endpoint
    properties:
      certificateArn: arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012
      databaseName: test
      endpointId: test-dms-endpoint-tf
      endpointType: source
      engineName: aurora
      extraConnectionAttributes:
      kmsKeyArn: arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012
      password: test
      port: 3306
      serverName: test
      sslMode: none
      tags:
        Name: test
      username: test
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import endpoints using the `endpoint_id`. For example:

```sh
 $ pulumi import aws:dms/endpoint:Endpoint test test-dms-endpoint-tf
```
 