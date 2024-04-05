Provides a Cognito User Pool UI Customization resource.

> **Note:** To use this resource, the user pool must have a domain associated with it. For more information, see the Amazon Cognito Developer Guide on [Customizing the Built-in Sign-In and Sign-up Webpages](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-app-ui-customization.html).

{{% examples %}}
## Example Usage
{{% example %}}
### UI customization settings for a single client

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const exampleUserPool = new aws.cognito.UserPool("exampleUserPool", {});
const exampleUserPoolDomain = new aws.cognito.UserPoolDomain("exampleUserPoolDomain", {
    domain: "example",
    userPoolId: exampleUserPool.id,
});
const exampleUserPoolClient = new aws.cognito.UserPoolClient("exampleUserPoolClient", {userPoolId: exampleUserPool.id});
const exampleUserPoolUICustomization = new aws.cognito.UserPoolUICustomization("exampleUserPoolUICustomization", {
    clientId: exampleUserPoolClient.id,
    css: ".label-customizable {font-weight: 400;}",
    imageFile: Buffer.from(fs.readFileSync("logo.png", 'binary')).toString('base64'),
    userPoolId: exampleUserPoolDomain.userPoolId,
});
```
```python
import pulumi
import base64
import pulumi_aws as aws

example_user_pool = aws.cognito.UserPool("exampleUserPool")
example_user_pool_domain = aws.cognito.UserPoolDomain("exampleUserPoolDomain",
    domain="example",
    user_pool_id=example_user_pool.id)
example_user_pool_client = aws.cognito.UserPoolClient("exampleUserPoolClient", user_pool_id=example_user_pool.id)
example_user_pool_ui_customization = aws.cognito.UserPoolUICustomization("exampleUserPoolUICustomization",
    client_id=example_user_pool_client.id,
    css=".label-customizable {font-weight: 400;}",
    image_file=(lambda path: base64.b64encode(open(path).read().encode()).decode())("logo.png"),
    user_pool_id=example_user_pool_domain.user_pool_id)
```
```csharp
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

	
string ReadFileBase64(string path) 
{
    return Convert.ToBase64String(Encoding.UTF8.GetBytes(File.ReadAllText(path)));
}

return await Deployment.RunAsync(() => 
{
    var exampleUserPool = new Aws.Cognito.UserPool("exampleUserPool");

    var exampleUserPoolDomain = new Aws.Cognito.UserPoolDomain("exampleUserPoolDomain", new()
    {
        Domain = "example",
        UserPoolId = exampleUserPool.Id,
    });

    var exampleUserPoolClient = new Aws.Cognito.UserPoolClient("exampleUserPoolClient", new()
    {
        UserPoolId = exampleUserPool.Id,
    });

    var exampleUserPoolUICustomization = new Aws.Cognito.UserPoolUICustomization("exampleUserPoolUICustomization", new()
    {
        ClientId = exampleUserPoolClient.Id,
        Css = ".label-customizable {font-weight: 400;}",
        ImageFile = ReadFileBase64("logo.png"),
        UserPoolId = exampleUserPoolDomain.UserPoolId,
    });

});
```
```go
package main

import (
	"encoding/base64"
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func filebase64OrPanic(path string) string {
	if fileData, err := os.ReadFile(path); err == nil {
		return base64.StdEncoding.EncodeToString(fileData[:])
	} else {
		panic(err.Error())
	}
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleUserPool, err := cognito.NewUserPool(ctx, "exampleUserPool", nil)
		if err != nil {
			return err
		}
		exampleUserPoolDomain, err := cognito.NewUserPoolDomain(ctx, "exampleUserPoolDomain", &cognito.UserPoolDomainArgs{
			Domain:     pulumi.String("example"),
			UserPoolId: exampleUserPool.ID(),
		})
		if err != nil {
			return err
		}
		exampleUserPoolClient, err := cognito.NewUserPoolClient(ctx, "exampleUserPoolClient", &cognito.UserPoolClientArgs{
			UserPoolId: exampleUserPool.ID(),
		})
		if err != nil {
			return err
		}
		_, err = cognito.NewUserPoolUICustomization(ctx, "exampleUserPoolUICustomization", &cognito.UserPoolUICustomizationArgs{
			ClientId:   exampleUserPoolClient.ID(),
			Css:        pulumi.String(".label-customizable {font-weight: 400;}"),
			ImageFile:  filebase64OrPanic("logo.png"),
			UserPoolId: exampleUserPoolDomain.UserPoolId,
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
import com.pulumi.aws.cognito.UserPool;
import com.pulumi.aws.cognito.UserPoolDomain;
import com.pulumi.aws.cognito.UserPoolDomainArgs;
import com.pulumi.aws.cognito.UserPoolClient;
import com.pulumi.aws.cognito.UserPoolClientArgs;
import com.pulumi.aws.cognito.UserPoolUICustomization;
import com.pulumi.aws.cognito.UserPoolUICustomizationArgs;
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
        var exampleUserPool = new UserPool("exampleUserPool");

        var exampleUserPoolDomain = new UserPoolDomain("exampleUserPoolDomain", UserPoolDomainArgs.builder()        
            .domain("example")
            .userPoolId(exampleUserPool.id())
            .build());

        var exampleUserPoolClient = new UserPoolClient("exampleUserPoolClient", UserPoolClientArgs.builder()        
            .userPoolId(exampleUserPool.id())
            .build());

        var exampleUserPoolUICustomization = new UserPoolUICustomization("exampleUserPoolUICustomization", UserPoolUICustomizationArgs.builder()        
            .clientId(exampleUserPoolClient.id())
            .css(".label-customizable {font-weight: 400;}")
            .imageFile(Base64.getEncoder().encodeToString(Files.readAllBytes(Paths.get("logo.png"))))
            .userPoolId(exampleUserPoolDomain.userPoolId())
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### UI customization settings for all clients

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const exampleUserPool = new aws.cognito.UserPool("exampleUserPool", {});
const exampleUserPoolDomain = new aws.cognito.UserPoolDomain("exampleUserPoolDomain", {
    domain: "example",
    userPoolId: exampleUserPool.id,
});
const exampleUserPoolUICustomization = new aws.cognito.UserPoolUICustomization("exampleUserPoolUICustomization", {
    css: ".label-customizable {font-weight: 400;}",
    imageFile: Buffer.from(fs.readFileSync("logo.png", 'binary')).toString('base64'),
    userPoolId: exampleUserPoolDomain.userPoolId,
});
```
```python
import pulumi
import base64
import pulumi_aws as aws

example_user_pool = aws.cognito.UserPool("exampleUserPool")
example_user_pool_domain = aws.cognito.UserPoolDomain("exampleUserPoolDomain",
    domain="example",
    user_pool_id=example_user_pool.id)
example_user_pool_ui_customization = aws.cognito.UserPoolUICustomization("exampleUserPoolUICustomization",
    css=".label-customizable {font-weight: 400;}",
    image_file=(lambda path: base64.b64encode(open(path).read().encode()).decode())("logo.png"),
    user_pool_id=example_user_pool_domain.user_pool_id)
```
```csharp
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

	
string ReadFileBase64(string path) 
{
    return Convert.ToBase64String(Encoding.UTF8.GetBytes(File.ReadAllText(path)));
}

return await Deployment.RunAsync(() => 
{
    var exampleUserPool = new Aws.Cognito.UserPool("exampleUserPool");

    var exampleUserPoolDomain = new Aws.Cognito.UserPoolDomain("exampleUserPoolDomain", new()
    {
        Domain = "example",
        UserPoolId = exampleUserPool.Id,
    });

    var exampleUserPoolUICustomization = new Aws.Cognito.UserPoolUICustomization("exampleUserPoolUICustomization", new()
    {
        Css = ".label-customizable {font-weight: 400;}",
        ImageFile = ReadFileBase64("logo.png"),
        UserPoolId = exampleUserPoolDomain.UserPoolId,
    });

});
```
```go
package main

import (
	"encoding/base64"
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func filebase64OrPanic(path string) string {
	if fileData, err := os.ReadFile(path); err == nil {
		return base64.StdEncoding.EncodeToString(fileData[:])
	} else {
		panic(err.Error())
	}
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleUserPool, err := cognito.NewUserPool(ctx, "exampleUserPool", nil)
		if err != nil {
			return err
		}
		exampleUserPoolDomain, err := cognito.NewUserPoolDomain(ctx, "exampleUserPoolDomain", &cognito.UserPoolDomainArgs{
			Domain:     pulumi.String("example"),
			UserPoolId: exampleUserPool.ID(),
		})
		if err != nil {
			return err
		}
		_, err = cognito.NewUserPoolUICustomization(ctx, "exampleUserPoolUICustomization", &cognito.UserPoolUICustomizationArgs{
			Css:        pulumi.String(".label-customizable {font-weight: 400;}"),
			ImageFile:  filebase64OrPanic("logo.png"),
			UserPoolId: exampleUserPoolDomain.UserPoolId,
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
import com.pulumi.aws.cognito.UserPool;
import com.pulumi.aws.cognito.UserPoolDomain;
import com.pulumi.aws.cognito.UserPoolDomainArgs;
import com.pulumi.aws.cognito.UserPoolUICustomization;
import com.pulumi.aws.cognito.UserPoolUICustomizationArgs;
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
        var exampleUserPool = new UserPool("exampleUserPool");

        var exampleUserPoolDomain = new UserPoolDomain("exampleUserPoolDomain", UserPoolDomainArgs.builder()        
            .domain("example")
            .userPoolId(exampleUserPool.id())
            .build());

        var exampleUserPoolUICustomization = new UserPoolUICustomization("exampleUserPoolUICustomization", UserPoolUICustomizationArgs.builder()        
            .css(".label-customizable {font-weight: 400;}")
            .imageFile(Base64.getEncoder().encodeToString(Files.readAllBytes(Paths.get("logo.png"))))
            .userPoolId(exampleUserPoolDomain.userPoolId())
            .build());

    }
}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Cognito User Pool UI Customizations using the `user_pool_id` and `client_id` separated by `,`. For example:

```sh
 $ pulumi import aws:cognito/userPoolUICustomization:UserPoolUICustomization example us-west-2_ZCTarbt5C,12bu4fuk3mlgqa2rtrujgp6egq
```
 