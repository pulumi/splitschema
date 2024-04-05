Provides an Amazon Lex Intent resource. For more information see
[Amazon Lex: How It Works](https://docs.aws.amazon.com/lex/latest/dg/how-it-works.html)

{{% examples %}}
## Example Usage
{{% example %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.lex.Intent;
import com.pulumi.aws.lex.IntentArgs;
import com.pulumi.aws.lex.inputs.IntentConfirmationPromptArgs;
import com.pulumi.aws.lex.inputs.IntentFulfillmentActivityArgs;
import com.pulumi.aws.lex.inputs.IntentRejectionStatementArgs;
import com.pulumi.aws.lex.inputs.IntentSlotArgs;
import com.pulumi.aws.lex.inputs.IntentSlotValueElicitationPromptArgs;
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
        var orderFlowersIntent = new Intent("orderFlowersIntent", IntentArgs.builder()        
            .confirmationPrompt(IntentConfirmationPromptArgs.builder()
                .maxAttempts(2)
                .messages(IntentConfirmationPromptMessageArgs.builder()
                    .content("Okay, your {FlowerType} will be ready for pickup by {PickupTime} on {PickupDate}.  Does this sound okay?")
                    .contentType("PlainText")
                    .build())
                .build())
            .createVersion(false)
            .description("Intent to order a bouquet of flowers for pick up")
            .fulfillmentActivity(IntentFulfillmentActivityArgs.builder()
                .type("ReturnIntent")
                .build())
            .name("OrderFlowers")
            .rejectionStatement(IntentRejectionStatementArgs.builder()
                .messages(IntentRejectionStatementMessageArgs.builder()
                    .content("Okay, I will not place your order.")
                    .contentType("PlainText")
                    .build())
                .build())
            .sampleUtterances(            
                "I would like to order some flowers",
                "I would like to pick up flowers")
            .slots(            
                IntentSlotArgs.builder()
                    .description("The type of flowers to pick up")
                    .name("FlowerType")
                    .priority(1)
                    .sampleUtterances("I would like to order {FlowerType}")
                    .slotConstraint("Required")
                    .slotType("FlowerTypes")
                    .slotTypeVersion("$LATEST")
                    .valueElicitationPrompt(IntentSlotValueElicitationPromptArgs.builder()
                        .maxAttempts(2)
                        .message(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
                        .build())
                    .build(),
                IntentSlotArgs.builder()
                    .description("The date to pick up the flowers")
                    .name("PickupDate")
                    .priority(2)
                    .sampleUtterances("I would like to order {FlowerType}")
                    .slotConstraint("Required")
                    .slotType("AMAZON.DATE")
                    .slotTypeVersion("$LATEST")
                    .valueElicitationPrompt(IntentSlotValueElicitationPromptArgs.builder()
                        .maxAttempts(2)
                        .message(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
                        .build())
                    .build(),
                IntentSlotArgs.builder()
                    .description("The time to pick up the flowers")
                    .name("PickupTime")
                    .priority(3)
                    .sampleUtterances("I would like to order {FlowerType}")
                    .slotConstraint("Required")
                    .slotType("AMAZON.TIME")
                    .slotTypeVersion("$LATEST")
                    .valueElicitationPrompt(IntentSlotValueElicitationPromptArgs.builder()
                        .maxAttempts(2)
                        .message(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
                        .build())
                    .build())
            .build());

    }
}
```
```yaml
resources:
  orderFlowersIntent:
    type: aws:lex:Intent
    properties:
      confirmationPrompt:
        maxAttempts: 2
        messages:
          - content: Okay, your {FlowerType} will be ready for pickup by {PickupTime} on {PickupDate}.  Does this sound okay?
            contentType: PlainText
      createVersion: false
      description: Intent to order a bouquet of flowers for pick up
      fulfillmentActivity:
        type: ReturnIntent
      name: OrderFlowers
      rejectionStatement:
        messages:
          - content: Okay, I will not place your order.
            contentType: PlainText
      sampleUtterances:
        - I would like to order some flowers
        - I would like to pick up flowers
      slots:
        - description: The type of flowers to pick up
          name: FlowerType
          priority: 1
          sampleUtterances:
            - I would like to order {FlowerType}
          slotConstraint: Required
          slotType: FlowerTypes
          slotTypeVersion: $LATEST
          valueElicitationPrompt:
            maxAttempts: 2
            message:
              - content: What type of flowers would you like to order?
                contentType: PlainText
        - description: The date to pick up the flowers
          name: PickupDate
          priority: 2
          sampleUtterances:
            - I would like to order {FlowerType}
          slotConstraint: Required
          slotType: AMAZON.DATE
          slotTypeVersion: $LATEST
          valueElicitationPrompt:
            maxAttempts: 2
            message:
              - content: What day do you want the {FlowerType} to be picked up?
                contentType: PlainText
        - description: The time to pick up the flowers
          name: PickupTime
          priority: 3
          sampleUtterances:
            - I would like to order {FlowerType}
          slotConstraint: Required
          slotType: AMAZON.TIME
          slotTypeVersion: $LATEST
          valueElicitationPrompt:
            maxAttempts: 2
            message:
              - content: Pick up the {FlowerType} at what time on {PickupDate}?
                contentType: PlainText
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import intents using their name. For example:

```sh
 $ pulumi import aws:lex/intent:Intent order_flowers_intent OrderFlowers
```
 