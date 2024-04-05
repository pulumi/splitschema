Provides a License Manager grant. This allows for sharing licenses with other AWS accounts.

{{% examples %}}
## Example Usage
{{% example %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.licensemanager.LicenseGrant;
import com.pulumi.aws.licensemanager.LicenseGrantArgs;
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
        var test = new LicenseGrant("test", LicenseGrantArgs.builder()        
            .allowedOperations(            
                "ListPurchasedLicenses",
                "CheckoutLicense",
                "CheckInLicense",
                "ExtendConsumptionLicense",
                "CreateToken")
            .homeRegion("us-east-1")
            .licenseArn("arn:aws:license-manager::111111111111:license:l-exampleARN")
            .principal("arn:aws:iam::111111111112:root")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:licensemanager:LicenseGrant
    properties:
      allowedOperations:
        - ListPurchasedLicenses
        - CheckoutLicense
        - CheckInLicense
        - ExtendConsumptionLicense
        - CreateToken
      homeRegion: us-east-1
      licenseArn: arn:aws:license-manager::111111111111:license:l-exampleARN
      principal: arn:aws:iam::111111111112:root
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_licensemanager_grant` using the grant arn. For example:

```sh
 $ pulumi import aws:licensemanager/licenseGrant:LicenseGrant test arn:aws:license-manager::123456789011:grant:g-01d313393d9e443d8664cc054db1e089
```
 