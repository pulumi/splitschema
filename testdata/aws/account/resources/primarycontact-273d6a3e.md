Manages the specified primary contact information associated with an AWS Account.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.account.PrimaryContact("test", {
    addressLine1: "123 Any Street",
    city: "Seattle",
    companyName: "Example Corp, Inc.",
    countryCode: "US",
    districtOrCounty: "King",
    fullName: "My Name",
    phoneNumber: "+64211111111",
    postalCode: "98101",
    stateOrRegion: "WA",
    websiteUrl: "https://www.examplecorp.com",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.account.PrimaryContact("test",
    address_line1="123 Any Street",
    city="Seattle",
    company_name="Example Corp, Inc.",
    country_code="US",
    district_or_county="King",
    full_name="My Name",
    phone_number="+64211111111",
    postal_code="98101",
    state_or_region="WA",
    website_url="https://www.examplecorp.com")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Account.PrimaryContact("test", new()
    {
        AddressLine1 = "123 Any Street",
        City = "Seattle",
        CompanyName = "Example Corp, Inc.",
        CountryCode = "US",
        DistrictOrCounty = "King",
        FullName = "My Name",
        PhoneNumber = "+64211111111",
        PostalCode = "98101",
        StateOrRegion = "WA",
        WebsiteUrl = "https://www.examplecorp.com",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/account"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := account.NewPrimaryContact(ctx, "test", &account.PrimaryContactArgs{
			AddressLine1:     pulumi.String("123 Any Street"),
			City:             pulumi.String("Seattle"),
			CompanyName:      pulumi.String("Example Corp, Inc."),
			CountryCode:      pulumi.String("US"),
			DistrictOrCounty: pulumi.String("King"),
			FullName:         pulumi.String("My Name"),
			PhoneNumber:      pulumi.String("+64211111111"),
			PostalCode:       pulumi.String("98101"),
			StateOrRegion:    pulumi.String("WA"),
			WebsiteUrl:       pulumi.String("https://www.examplecorp.com"),
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
import com.pulumi.aws.account.PrimaryContact;
import com.pulumi.aws.account.PrimaryContactArgs;
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
        var test = new PrimaryContact("test", PrimaryContactArgs.builder()        
            .addressLine1("123 Any Street")
            .city("Seattle")
            .companyName("Example Corp, Inc.")
            .countryCode("US")
            .districtOrCounty("King")
            .fullName("My Name")
            .phoneNumber("+64211111111")
            .postalCode("98101")
            .stateOrRegion("WA")
            .websiteUrl("https://www.examplecorp.com")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:account:PrimaryContact
    properties:
      addressLine1: 123 Any Street
      city: Seattle
      companyName: Example Corp, Inc.
      countryCode: US
      districtOrCounty: King
      fullName: My Name
      phoneNumber: '+64211111111'
      postalCode: '98101'
      stateOrRegion: WA
      websiteUrl: https://www.examplecorp.com
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import the Primary Contact using the `account_id`. For example:

```sh
 $ pulumi import aws:account/primaryContact:PrimaryContact test 1234567890
```
 