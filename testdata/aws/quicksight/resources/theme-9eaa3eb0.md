Resource for managing a QuickSight Theme.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.quicksight.Theme;
import com.pulumi.aws.quicksight.ThemeArgs;
import com.pulumi.aws.quicksight.inputs.ThemeConfigurationArgs;
import com.pulumi.aws.quicksight.inputs.ThemeConfigurationDataColorPaletteArgs;
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
        var example = new Theme("example", ThemeArgs.builder()        
            .baseThemeId("MIDNIGHT")
            .configuration(ThemeConfigurationArgs.builder()
                .dataColorPalette(ThemeConfigurationDataColorPaletteArgs.builder()
                    .colors(                    
                        "#FFFFFF",
                        "#111111",
                        "#222222",
                        "#333333",
                        "#444444",
                        "#555555",
                        "#666666",
                        "#777777",
                        "#888888",
                        "#999999")
                    .emptyFillColor("#FFFFFF")
                    .minMaxGradient(                    
                        "#FFFFFF",
                        "#111111")
                    .build())
                .build())
            .themeId("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:Theme
    properties:
      baseThemeId: MIDNIGHT
      configuration:
        dataColorPalette:
          colors:
            - '#FFFFFF'
            - '#111111'
            - '#222222'
            - '#333333'
            - '#444444'
            - '#555555'
            - '#666666'
            - '#777777'
            - '#888888'
            - '#999999'
          emptyFillColor: '#FFFFFF'
          minMaxGradient:
            - '#FFFFFF'
            - '#111111'
      themeId: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a QuickSight Theme using the AWS account ID and theme ID separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:quicksight/theme:Theme example 123456789012,example-id
```
 