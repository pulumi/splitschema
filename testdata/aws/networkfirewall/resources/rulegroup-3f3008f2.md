Provides an AWS Network Firewall Rule Group Resource

{{% examples %}}
## Example Usage
{{% example %}}
### Stateful Inspection for denying access to a domain

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.networkfirewall.RuleGroup("example", {
    capacity: 100,
    ruleGroup: {
        rulesSource: {
            rulesSourceList: {
                generatedRulesType: "DENYLIST",
                targetTypes: ["HTTP_HOST"],
                targets: ["test.example.com"],
            },
        },
    },
    tags: {
        Tag1: "Value1",
        Tag2: "Value2",
    },
    type: "STATEFUL",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkfirewall.RuleGroup("example",
    capacity=100,
    rule_group=aws.networkfirewall.RuleGroupRuleGroupArgs(
        rules_source=aws.networkfirewall.RuleGroupRuleGroupRulesSourceArgs(
            rules_source_list=aws.networkfirewall.RuleGroupRuleGroupRulesSourceRulesSourceListArgs(
                generated_rules_type="DENYLIST",
                target_types=["HTTP_HOST"],
                targets=["test.example.com"],
            ),
        ),
    ),
    tags={
        "Tag1": "Value1",
        "Tag2": "Value2",
    },
    type="STATEFUL")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.NetworkFirewall.RuleGroup("example", new()
    {
        Capacity = 100,
        RuleGroupConfiguration = new Aws.NetworkFirewall.Inputs.RuleGroupRuleGroupArgs
        {
            RulesSource = new Aws.NetworkFirewall.Inputs.RuleGroupRuleGroupRulesSourceArgs
            {
                RulesSourceList = new Aws.NetworkFirewall.Inputs.RuleGroupRuleGroupRulesSourceRulesSourceListArgs
                {
                    GeneratedRulesType = "DENYLIST",
                    TargetTypes = new[]
                    {
                        "HTTP_HOST",
                    },
                    Targets = new[]
                    {
                        "test.example.com",
                    },
                },
            },
        },
        Tags = 
        {
            { "Tag1", "Value1" },
            { "Tag2", "Value2" },
        },
        Type = "STATEFUL",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkfirewall"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkfirewall.NewRuleGroup(ctx, "example", &networkfirewall.RuleGroupArgs{
			Capacity: pulumi.Int(100),
			RuleGroup: &networkfirewall.RuleGroupRuleGroupArgs{
				RulesSource: &networkfirewall.RuleGroupRuleGroupRulesSourceArgs{
					RulesSourceList: &networkfirewall.RuleGroupRuleGroupRulesSourceRulesSourceListArgs{
						GeneratedRulesType: pulumi.String("DENYLIST"),
						TargetTypes: pulumi.StringArray{
							pulumi.String("HTTP_HOST"),
						},
						Targets: pulumi.StringArray{
							pulumi.String("test.example.com"),
						},
					},
				},
			},
			Tags: pulumi.StringMap{
				"Tag1": pulumi.String("Value1"),
				"Tag2": pulumi.String("Value2"),
			},
			Type: pulumi.String("STATEFUL"),
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
import com.pulumi.aws.networkfirewall.RuleGroup;
import com.pulumi.aws.networkfirewall.RuleGroupArgs;
import com.pulumi.aws.networkfirewall.inputs.RuleGroupRuleGroupArgs;
import com.pulumi.aws.networkfirewall.inputs.RuleGroupRuleGroupRulesSourceArgs;
import com.pulumi.aws.networkfirewall.inputs.RuleGroupRuleGroupRulesSourceRulesSourceListArgs;
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
        var example = new RuleGroup("example", RuleGroupArgs.builder()        
            .capacity(100)
            .ruleGroup(RuleGroupRuleGroupArgs.builder()
                .rulesSource(RuleGroupRuleGroupRulesSourceArgs.builder()
                    .rulesSourceList(RuleGroupRuleGroupRulesSourceRulesSourceListArgs.builder()
                        .generatedRulesType("DENYLIST")
                        .targetTypes("HTTP_HOST")
                        .targets("test.example.com")
                        .build())
                    .build())
                .build())
            .tags(Map.ofEntries(
                Map.entry("Tag1", "Value1"),
                Map.entry("Tag2", "Value2")
            ))
            .type("STATEFUL")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkfirewall:RuleGroup
    properties:
      capacity: 100
      ruleGroup:
        rulesSource:
          rulesSourceList:
            generatedRulesType: DENYLIST
            targetTypes:
              - HTTP_HOST
            targets:
              - test.example.com
      tags:
        Tag1: Value1
        Tag2: Value2
      type: STATEFUL
```
{{% /example %}}
{{% example %}}
### Stateful Inspection for permitting packets from a source IP address

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.networkfirewall.RuleGroup;
import com.pulumi.aws.networkfirewall.RuleGroupArgs;
import com.pulumi.aws.networkfirewall.inputs.RuleGroupRuleGroupArgs;
import com.pulumi.aws.networkfirewall.inputs.RuleGroupRuleGroupRulesSourceArgs;
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
        final var ips =         
            "1.1.1.1/32",
            "1.0.0.1/32";

        var example = new RuleGroup("example", RuleGroupArgs.builder()        
            .capacity(50)
            .description("Permits http traffic from source")
            .type("STATEFUL")
            .ruleGroup(RuleGroupRuleGroupArgs.builder()
                .rulesSource(RuleGroupRuleGroupRulesSourceArgs.builder()
                    .dynamic(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
                    .build())
                .build())
            .tags(Map.of("Name", "permit HTTP from source"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkfirewall:RuleGroup
    properties:
      capacity: 50
      description: Permits http traffic from source
      type: STATEFUL
      ruleGroup:
        rulesSource:
          dynamic:
            - forEach: ${ips}
              content:
                - action: PASS
                  header:
                    - destination: ANY
                      destinationPort: ANY
                      protocol: HTTP
                      direction: ANY
                      sourcePort: ANY
                      source: ${stateful_rule.value}
                  ruleOption:
                    - keyword: sid
                      settings:
                        - '1'
      tags:
        Name: permit HTTP from source
variables:
  ips:
    - 1.1.1.1/32
    - 1.0.0.1/32
```
{{% /example %}}
{{% example %}}
### Stateful Inspection for blocking packets from going to an intended destination

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.networkfirewall.RuleGroup;
import com.pulumi.aws.networkfirewall.RuleGroupArgs;
import com.pulumi.aws.networkfirewall.inputs.RuleGroupRuleGroupArgs;
import com.pulumi.aws.networkfirewall.inputs.RuleGroupRuleGroupRulesSourceArgs;
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
        var example = new RuleGroup("example", RuleGroupArgs.builder()        
            .capacity(100)
            .ruleGroup(RuleGroupRuleGroupArgs.builder()
                .rulesSource(RuleGroupRuleGroupRulesSourceArgs.builder()
                    .statefulRule(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
                    .build())
                .build())
            .tags(Map.ofEntries(
                Map.entry("Tag1", "Value1"),
                Map.entry("Tag2", "Value2")
            ))
            .type("STATEFUL")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkfirewall:RuleGroup
    properties:
      capacity: 100
      ruleGroup:
        rulesSource:
          statefulRule:
            - action: DROP
              header:
                destination: 124.1.1.24/32
                destinationPort: 53
                direction: ANY
                protocol: TCP
                source: 1.2.3.4/32
                sourcePort: 53
              ruleOption:
                - keyword: sid
                  settings:
                    - '1'
      tags:
        Tag1: Value1
        Tag2: Value2
      type: STATEFUL
```
{{% /example %}}
{{% example %}}
### Stateful Inspection from rules specifications defined in Suricata flat format

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const example = new aws.networkfirewall.RuleGroup("example", {
    capacity: 100,
    type: "STATEFUL",
    rules: fs.readFileSync("example.rules", "utf8"),
    tags: {
        Tag1: "Value1",
        Tag2: "Value2",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkfirewall.RuleGroup("example",
    capacity=100,
    type="STATEFUL",
    rules=(lambda path: open(path).read())("example.rules"),
    tags={
        "Tag1": "Value1",
        "Tag2": "Value2",
    })
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.NetworkFirewall.RuleGroup("example", new()
    {
        Capacity = 100,
        Type = "STATEFUL",
        Rules = File.ReadAllText("example.rules"),
        Tags = 
        {
            { "Tag1", "Value1" },
            { "Tag2", "Value2" },
        },
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkfirewall"
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
		_, err := networkfirewall.NewRuleGroup(ctx, "example", &networkfirewall.RuleGroupArgs{
			Capacity: pulumi.Int(100),
			Type:     pulumi.String("STATEFUL"),
			Rules:    readFileOrPanic("example.rules"),
			Tags: pulumi.StringMap{
				"Tag1": pulumi.String("Value1"),
				"Tag2": pulumi.String("Value2"),
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
import com.pulumi.aws.networkfirewall.RuleGroup;
import com.pulumi.aws.networkfirewall.RuleGroupArgs;
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
        var example = new RuleGroup("example", RuleGroupArgs.builder()        
            .capacity(100)
            .type("STATEFUL")
            .rules(Files.readString(Paths.get("example.rules")))
            .tags(Map.ofEntries(
                Map.entry("Tag1", "Value1"),
                Map.entry("Tag2", "Value2")
            ))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkfirewall:RuleGroup
    properties:
      capacity: 100
      type: STATEFUL
      rules:
        fn::readFile: example.rules
      tags:
        Tag1: Value1
        Tag2: Value2
```
{{% /example %}}
{{% example %}}
### Stateful Inspection from rule group specifications using rule variables and Suricata format rules

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const example = new aws.networkfirewall.RuleGroup("example", {
    capacity: 100,
    type: "STATEFUL",
    ruleGroup: {
        ruleVariables: {
            ipSets: [
                {
                    key: "WEBSERVERS_HOSTS",
                    ipSet: {
                        definitions: [
                            "10.0.0.0/16",
                            "10.0.1.0/24",
                            "192.168.0.0/16",
                        ],
                    },
                },
                {
                    key: "EXTERNAL_HOST",
                    ipSet: {
                        definitions: ["1.2.3.4/32"],
                    },
                },
            ],
            portSets: [{
                key: "HTTP_PORTS",
                portSet: {
                    definitions: [
                        "443",
                        "80",
                    ],
                },
            }],
        },
        rulesSource: {
            rulesString: fs.readFileSync("suricata_rules_file", "utf8"),
        },
    },
    tags: {
        Tag1: "Value1",
        Tag2: "Value2",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkfirewall.RuleGroup("example",
    capacity=100,
    type="STATEFUL",
    rule_group=aws.networkfirewall.RuleGroupRuleGroupArgs(
        rule_variables=aws.networkfirewall.RuleGroupRuleGroupRuleVariablesArgs(
            ip_sets=[
                aws.networkfirewall.RuleGroupRuleGroupRuleVariablesIpSetArgs(
                    key="WEBSERVERS_HOSTS",
                    ip_set=aws.networkfirewall.RuleGroupRuleGroupRuleVariablesIpSetIpSetArgs(
                        definitions=[
                            "10.0.0.0/16",
                            "10.0.1.0/24",
                            "192.168.0.0/16",
                        ],
                    ),
                ),
                aws.networkfirewall.RuleGroupRuleGroupRuleVariablesIpSetArgs(
                    key="EXTERNAL_HOST",
                    ip_set=aws.networkfirewall.RuleGroupRuleGroupRuleVariablesIpSetIpSetArgs(
                        definitions=["1.2.3.4/32"],
                    ),
                ),
            ],
            port_sets=[aws.networkfirewall.RuleGroupRuleGroupRuleVariablesPortSetArgs(
                key="HTTP_PORTS",
                port_set=aws.networkfirewall.RuleGroupRuleGroupRuleVariablesPortSetPortSetArgs(
                    definitions=[
                        "443",
                        "80",
                    ],
                ),
            )],
        ),
        rules_source=aws.networkfirewall.RuleGroupRuleGroupRulesSourceArgs(
            rules_string=(lambda path: open(path).read())("suricata_rules_file"),
        ),
    ),
    tags={
        "Tag1": "Value1",
        "Tag2": "Value2",
    })
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.NetworkFirewall.RuleGroup("example", new()
    {
        Capacity = 100,
        Type = "STATEFUL",
        RuleGroupConfiguration = new Aws.NetworkFirewall.Inputs.RuleGroupRuleGroupArgs
        {
            RuleVariables = new Aws.NetworkFirewall.Inputs.RuleGroupRuleGroupRuleVariablesArgs
            {
                IpSets = new[]
                {
                    new Aws.NetworkFirewall.Inputs.RuleGroupRuleGroupRuleVariablesIpSetArgs
                    {
                        Key = "WEBSERVERS_HOSTS",
                        IpSet = new Aws.NetworkFirewall.Inputs.RuleGroupRuleGroupRuleVariablesIpSetIpSetArgs
                        {
                            Definitions = new[]
                            {
                                "10.0.0.0/16",
                                "10.0.1.0/24",
                                "192.168.0.0/16",
                            },
                        },
                    },
                    new Aws.NetworkFirewall.Inputs.RuleGroupRuleGroupRuleVariablesIpSetArgs
                    {
                        Key = "EXTERNAL_HOST",
                        IpSet = new Aws.NetworkFirewall.Inputs.RuleGroupRuleGroupRuleVariablesIpSetIpSetArgs
                        {
                            Definitions = new[]
                            {
                                "1.2.3.4/32",
                            },
                        },
                    },
                },
                PortSets = new[]
                {
                    new Aws.NetworkFirewall.Inputs.RuleGroupRuleGroupRuleVariablesPortSetArgs
                    {
                        Key = "HTTP_PORTS",
                        PortSet = new Aws.NetworkFirewall.Inputs.RuleGroupRuleGroupRuleVariablesPortSetPortSetArgs
                        {
                            Definitions = new[]
                            {
                                "443",
                                "80",
                            },
                        },
                    },
                },
            },
            RulesSource = new Aws.NetworkFirewall.Inputs.RuleGroupRuleGroupRulesSourceArgs
            {
                RulesString = File.ReadAllText("suricata_rules_file"),
            },
        },
        Tags = 
        {
            { "Tag1", "Value1" },
            { "Tag2", "Value2" },
        },
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkfirewall"
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
		_, err := networkfirewall.NewRuleGroup(ctx, "example", &networkfirewall.RuleGroupArgs{
			Capacity: pulumi.Int(100),
			Type:     pulumi.String("STATEFUL"),
			RuleGroup: &networkfirewall.RuleGroupRuleGroupArgs{
				RuleVariables: &networkfirewall.RuleGroupRuleGroupRuleVariablesArgs{
					IpSets: networkfirewall.RuleGroupRuleGroupRuleVariablesIpSetArray{
						&networkfirewall.RuleGroupRuleGroupRuleVariablesIpSetArgs{
							Key: pulumi.String("WEBSERVERS_HOSTS"),
							IpSet: &networkfirewall.RuleGroupRuleGroupRuleVariablesIpSetIpSetArgs{
								Definitions: pulumi.StringArray{
									pulumi.String("10.0.0.0/16"),
									pulumi.String("10.0.1.0/24"),
									pulumi.String("192.168.0.0/16"),
								},
							},
						},
						&networkfirewall.RuleGroupRuleGroupRuleVariablesIpSetArgs{
							Key: pulumi.String("EXTERNAL_HOST"),
							IpSet: &networkfirewall.RuleGroupRuleGroupRuleVariablesIpSetIpSetArgs{
								Definitions: pulumi.StringArray{
									pulumi.String("1.2.3.4/32"),
								},
							},
						},
					},
					PortSets: networkfirewall.RuleGroupRuleGroupRuleVariablesPortSetArray{
						&networkfirewall.RuleGroupRuleGroupRuleVariablesPortSetArgs{
							Key: pulumi.String("HTTP_PORTS"),
							PortSet: &networkfirewall.RuleGroupRuleGroupRuleVariablesPortSetPortSetArgs{
								Definitions: pulumi.StringArray{
									pulumi.String("443"),
									pulumi.String("80"),
								},
							},
						},
					},
				},
				RulesSource: &networkfirewall.RuleGroupRuleGroupRulesSourceArgs{
					RulesString: readFileOrPanic("suricata_rules_file"),
				},
			},
			Tags: pulumi.StringMap{
				"Tag1": pulumi.String("Value1"),
				"Tag2": pulumi.String("Value2"),
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
import com.pulumi.aws.networkfirewall.RuleGroup;
import com.pulumi.aws.networkfirewall.RuleGroupArgs;
import com.pulumi.aws.networkfirewall.inputs.RuleGroupRuleGroupArgs;
import com.pulumi.aws.networkfirewall.inputs.RuleGroupRuleGroupRuleVariablesArgs;
import com.pulumi.aws.networkfirewall.inputs.RuleGroupRuleGroupRulesSourceArgs;
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
        var example = new RuleGroup("example", RuleGroupArgs.builder()        
            .capacity(100)
            .type("STATEFUL")
            .ruleGroup(RuleGroupRuleGroupArgs.builder()
                .ruleVariables(RuleGroupRuleGroupRuleVariablesArgs.builder()
                    .ipSets(                    
                        RuleGroupRuleGroupRuleVariablesIpSetArgs.builder()
                            .key("WEBSERVERS_HOSTS")
                            .ipSet(RuleGroupRuleGroupRuleVariablesIpSetIpSetArgs.builder()
                                .definitions(                                
                                    "10.0.0.0/16",
                                    "10.0.1.0/24",
                                    "192.168.0.0/16")
                                .build())
                            .build(),
                        RuleGroupRuleGroupRuleVariablesIpSetArgs.builder()
                            .key("EXTERNAL_HOST")
                            .ipSet(RuleGroupRuleGroupRuleVariablesIpSetIpSetArgs.builder()
                                .definitions("1.2.3.4/32")
                                .build())
                            .build())
                    .portSets(RuleGroupRuleGroupRuleVariablesPortSetArgs.builder()
                        .key("HTTP_PORTS")
                        .portSet(RuleGroupRuleGroupRuleVariablesPortSetPortSetArgs.builder()
                            .definitions(                            
                                "443",
                                "80")
                            .build())
                        .build())
                    .build())
                .rulesSource(RuleGroupRuleGroupRulesSourceArgs.builder()
                    .rulesString(Files.readString(Paths.get("suricata_rules_file")))
                    .build())
                .build())
            .tags(Map.ofEntries(
                Map.entry("Tag1", "Value1"),
                Map.entry("Tag2", "Value2")
            ))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkfirewall:RuleGroup
    properties:
      capacity: 100
      type: STATEFUL
      ruleGroup:
        ruleVariables:
          ipSets:
            - key: WEBSERVERS_HOSTS
              ipSet:
                definitions:
                  - 10.0.0.0/16
                  - 10.0.1.0/24
                  - 192.168.0.0/16
            - key: EXTERNAL_HOST
              ipSet:
                definitions:
                  - 1.2.3.4/32
          portSets:
            - key: HTTP_PORTS
              portSet:
                definitions:
                  - '443'
                  - '80'
        rulesSource:
          rulesString:
            fn::readFile: suricata_rules_file
      tags:
        Tag1: Value1
        Tag2: Value2
```
{{% /example %}}
{{% example %}}
### Stateless Inspection with a Custom Action

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.networkfirewall.RuleGroup;
import com.pulumi.aws.networkfirewall.RuleGroupArgs;
import com.pulumi.aws.networkfirewall.inputs.RuleGroupRuleGroupArgs;
import com.pulumi.aws.networkfirewall.inputs.RuleGroupRuleGroupRulesSourceArgs;
import com.pulumi.aws.networkfirewall.inputs.RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsArgs;
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
        var example = new RuleGroup("example", RuleGroupArgs.builder()        
            .capacity(100)
            .description("Stateless Rate Limiting Rule")
            .ruleGroup(RuleGroupRuleGroupArgs.builder()
                .rulesSource(RuleGroupRuleGroupRulesSourceArgs.builder()
                    .statelessRulesAndCustomActions(RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsArgs.builder()
                        .customAction(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
                        .statelessRule(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
                        .build())
                    .build())
                .build())
            .tags(Map.ofEntries(
                Map.entry("Tag1", "Value1"),
                Map.entry("Tag2", "Value2")
            ))
            .type("STATELESS")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkfirewall:RuleGroup
    properties:
      capacity: 100
      description: Stateless Rate Limiting Rule
      ruleGroup:
        rulesSource:
          statelessRulesAndCustomActions:
            customAction:
              - actionDefinition:
                  publishMetricAction:
                    dimension:
                      - value: '2'
                actionName: ExampleMetricsAction
            statelessRule:
              - priority: 1
                ruleDefinition:
                  actions:
                    - aws:pass
                    - ExampleMetricsAction
                  matchAttributes:
                    destination:
                      - addressDefinition: 124.1.1.5/32
                    destinationPort:
                      - fromPort: 443
                        toPort: 443
                    protocols:
                      - 6
                    source:
                      - addressDefinition: 1.2.3.4/32
                    sourcePort:
                      - fromPort: 443
                        toPort: 443
                    tcpFlag:
                      - flags:
                          - SYN
                        masks:
                          - SYN
                          - ACK
      tags:
        Tag1: Value1
        Tag2: Value2
      type: STATELESS
```
{{% /example %}}
{{% example %}}
### IP Set References to the Rule Group

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.networkfirewall.RuleGroup("example", {
    capacity: 100,
    type: "STATEFUL",
    ruleGroup: {
        rulesSource: {
            rulesSourceList: {
                generatedRulesType: "DENYLIST",
                targetTypes: ["HTTP_HOST"],
                targets: ["test.example.com"],
            },
        },
        referenceSets: {
            ipSetReferences: [{
                key: "example",
                ipSetReferences: [{
                    referenceArn: aws_ec2_managed_prefix_list["this"].arn,
                }],
            }],
        },
    },
    tags: {
        Tag1: "Value1",
        Tag2: "Value2",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkfirewall.RuleGroup("example",
    capacity=100,
    type="STATEFUL",
    rule_group=aws.networkfirewall.RuleGroupRuleGroupArgs(
        rules_source=aws.networkfirewall.RuleGroupRuleGroupRulesSourceArgs(
            rules_source_list=aws.networkfirewall.RuleGroupRuleGroupRulesSourceRulesSourceListArgs(
                generated_rules_type="DENYLIST",
                target_types=["HTTP_HOST"],
                targets=["test.example.com"],
            ),
        ),
        reference_sets=aws.networkfirewall.RuleGroupRuleGroupReferenceSetsArgs(
            ip_set_references=[aws.networkfirewall.RuleGroupRuleGroupReferenceSetsIpSetReferenceArgs(
                key="example",
                ip_set_references=[aws.networkfirewall.RuleGroupRuleGroupReferenceSetsIpSetReferenceIpSetReferenceArgs(
                    reference_arn=aws_ec2_managed_prefix_list["this"]["arn"],
                )],
            )],
        ),
    ),
    tags={
        "Tag1": "Value1",
        "Tag2": "Value2",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.NetworkFirewall.RuleGroup("example", new()
    {
        Capacity = 100,
        Type = "STATEFUL",
        RuleGroupConfiguration = new Aws.NetworkFirewall.Inputs.RuleGroupRuleGroupArgs
        {
            RulesSource = new Aws.NetworkFirewall.Inputs.RuleGroupRuleGroupRulesSourceArgs
            {
                RulesSourceList = new Aws.NetworkFirewall.Inputs.RuleGroupRuleGroupRulesSourceRulesSourceListArgs
                {
                    GeneratedRulesType = "DENYLIST",
                    TargetTypes = new[]
                    {
                        "HTTP_HOST",
                    },
                    Targets = new[]
                    {
                        "test.example.com",
                    },
                },
            },
            ReferenceSets = new Aws.NetworkFirewall.Inputs.RuleGroupRuleGroupReferenceSetsArgs
            {
                IpSetReferences = new[]
                {
                    new Aws.NetworkFirewall.Inputs.RuleGroupRuleGroupReferenceSetsIpSetReferenceArgs
                    {
                        Key = "example",
                        IpSetReferences = new[]
                        {
                            new Aws.NetworkFirewall.Inputs.RuleGroupRuleGroupReferenceSetsIpSetReferenceIpSetReferenceArgs
                            {
                                ReferenceArn = aws_ec2_managed_prefix_list.This.Arn,
                            },
                        },
                    },
                },
            },
        },
        Tags = 
        {
            { "Tag1", "Value1" },
            { "Tag2", "Value2" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkfirewall"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkfirewall.NewRuleGroup(ctx, "example", &networkfirewall.RuleGroupArgs{
			Capacity: pulumi.Int(100),
			Type:     pulumi.String("STATEFUL"),
			RuleGroup: &networkfirewall.RuleGroupRuleGroupArgs{
				RulesSource: &networkfirewall.RuleGroupRuleGroupRulesSourceArgs{
					RulesSourceList: &networkfirewall.RuleGroupRuleGroupRulesSourceRulesSourceListArgs{
						GeneratedRulesType: pulumi.String("DENYLIST"),
						TargetTypes: pulumi.StringArray{
							pulumi.String("HTTP_HOST"),
						},
						Targets: pulumi.StringArray{
							pulumi.String("test.example.com"),
						},
					},
				},
				ReferenceSets: &networkfirewall.RuleGroupRuleGroupReferenceSetsArgs{
					IpSetReferences: networkfirewall.RuleGroupRuleGroupReferenceSetsIpSetReferenceArray{
						&networkfirewall.RuleGroupRuleGroupReferenceSetsIpSetReferenceArgs{
							Key: pulumi.String("example"),
							IpSetReferences: networkfirewall.RuleGroupRuleGroupReferenceSetsIpSetReferenceIpSetReferenceArray{
								&networkfirewall.RuleGroupRuleGroupReferenceSetsIpSetReferenceIpSetReferenceArgs{
									ReferenceArn: pulumi.Any(aws_ec2_managed_prefix_list.This.Arn),
								},
							},
						},
					},
				},
			},
			Tags: pulumi.StringMap{
				"Tag1": pulumi.String("Value1"),
				"Tag2": pulumi.String("Value2"),
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
import com.pulumi.aws.networkfirewall.RuleGroup;
import com.pulumi.aws.networkfirewall.RuleGroupArgs;
import com.pulumi.aws.networkfirewall.inputs.RuleGroupRuleGroupArgs;
import com.pulumi.aws.networkfirewall.inputs.RuleGroupRuleGroupRulesSourceArgs;
import com.pulumi.aws.networkfirewall.inputs.RuleGroupRuleGroupRulesSourceRulesSourceListArgs;
import com.pulumi.aws.networkfirewall.inputs.RuleGroupRuleGroupReferenceSetsArgs;
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
        var example = new RuleGroup("example", RuleGroupArgs.builder()        
            .capacity(100)
            .type("STATEFUL")
            .ruleGroup(RuleGroupRuleGroupArgs.builder()
                .rulesSource(RuleGroupRuleGroupRulesSourceArgs.builder()
                    .rulesSourceList(RuleGroupRuleGroupRulesSourceRulesSourceListArgs.builder()
                        .generatedRulesType("DENYLIST")
                        .targetTypes("HTTP_HOST")
                        .targets("test.example.com")
                        .build())
                    .build())
                .referenceSets(RuleGroupRuleGroupReferenceSetsArgs.builder()
                    .ipSetReferences(RuleGroupRuleGroupReferenceSetsIpSetReferenceArgs.builder()
                        .key("example")
                        .ipSetReferences(RuleGroupRuleGroupReferenceSetsIpSetReferenceIpSetReferenceArgs.builder()
                            .referenceArn(aws_ec2_managed_prefix_list.this().arn())
                            .build())
                        .build())
                    .build())
                .build())
            .tags(Map.ofEntries(
                Map.entry("Tag1", "Value1"),
                Map.entry("Tag2", "Value2")
            ))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkfirewall:RuleGroup
    properties:
      capacity: 100
      type: STATEFUL
      ruleGroup:
        rulesSource:
          rulesSourceList:
            generatedRulesType: DENYLIST
            targetTypes:
              - HTTP_HOST
            targets:
              - test.example.com
        referenceSets:
          ipSetReferences:
            - key: example
              ipSetReferences:
                - referenceArn: ${aws_ec2_managed_prefix_list.this.arn}
      tags:
        Tag1: Value1
        Tag2: Value2
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Network Firewall Rule Groups using their `arn`. For example:

```sh
 $ pulumi import aws:networkfirewall/ruleGroup:RuleGroup example arn:aws:network-firewall:us-west-1:123456789012:stateful-rulegroup/example
```
 