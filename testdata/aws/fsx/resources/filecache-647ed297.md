Resource for managing an Amazon File Cache cache.
See the [Create File Cache](https://docs.aws.amazon.com/fsx/latest/APIReference/API_CreateFileCache.html) for more information.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.fsx.FileCache("example", {
    dataRepositoryAssociations: [{
        dataRepositoryPath: "nfs://filer.domain.com",
        dataRepositorySubdirectories: [
            "test",
            "test2",
        ],
        fileCachePath: "/ns1",
        nfs: [{
            dnsIps: [
                "192.168.0.1",
                "192.168.0.2",
            ],
            version: "NFS3",
        }],
    }],
    fileCacheType: "LUSTRE",
    fileCacheTypeVersion: "2.12",
    lustreConfigurations: [{
        deploymentType: "CACHE_1",
        metadataConfigurations: [{
            storageCapacity: 2400,
        }],
        perUnitStorageThroughput: 1000,
        weeklyMaintenanceStartTime: "2:05:00",
    }],
    subnetIds: [aws_subnet.test1.id],
    storageCapacity: 1200,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.fsx.FileCache("example",
    data_repository_associations=[aws.fsx.FileCacheDataRepositoryAssociationArgs(
        data_repository_path="nfs://filer.domain.com",
        data_repository_subdirectories=[
            "test",
            "test2",
        ],
        file_cache_path="/ns1",
        nfs=[aws.fsx.FileCacheDataRepositoryAssociationNfArgs(
            dns_ips=[
                "192.168.0.1",
                "192.168.0.2",
            ],
            version="NFS3",
        )],
    )],
    file_cache_type="LUSTRE",
    file_cache_type_version="2.12",
    lustre_configurations=[aws.fsx.FileCacheLustreConfigurationArgs(
        deployment_type="CACHE_1",
        metadata_configurations=[aws.fsx.FileCacheLustreConfigurationMetadataConfigurationArgs(
            storage_capacity=2400,
        )],
        per_unit_storage_throughput=1000,
        weekly_maintenance_start_time="2:05:00",
    )],
    subnet_ids=[aws_subnet["test1"]["id"]],
    storage_capacity=1200)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Fsx.FileCache("example", new()
    {
        DataRepositoryAssociations = new[]
        {
            new Aws.Fsx.Inputs.FileCacheDataRepositoryAssociationArgs
            {
                DataRepositoryPath = "nfs://filer.domain.com",
                DataRepositorySubdirectories = new[]
                {
                    "test",
                    "test2",
                },
                FileCachePath = "/ns1",
                Nfs = new[]
                {
                    new Aws.Fsx.Inputs.FileCacheDataRepositoryAssociationNfArgs
                    {
                        DnsIps = new[]
                        {
                            "192.168.0.1",
                            "192.168.0.2",
                        },
                        Version = "NFS3",
                    },
                },
            },
        },
        FileCacheType = "LUSTRE",
        FileCacheTypeVersion = "2.12",
        LustreConfigurations = new[]
        {
            new Aws.Fsx.Inputs.FileCacheLustreConfigurationArgs
            {
                DeploymentType = "CACHE_1",
                MetadataConfigurations = new[]
                {
                    new Aws.Fsx.Inputs.FileCacheLustreConfigurationMetadataConfigurationArgs
                    {
                        StorageCapacity = 2400,
                    },
                },
                PerUnitStorageThroughput = 1000,
                WeeklyMaintenanceStartTime = "2:05:00",
            },
        },
        SubnetIds = new[]
        {
            aws_subnet.Test1.Id,
        },
        StorageCapacity = 1200,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/fsx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := fsx.NewFileCache(ctx, "example", &fsx.FileCacheArgs{
			DataRepositoryAssociations: fsx.FileCacheDataRepositoryAssociationArray{
				&fsx.FileCacheDataRepositoryAssociationArgs{
					DataRepositoryPath: pulumi.String("nfs://filer.domain.com"),
					DataRepositorySubdirectories: pulumi.StringArray{
						pulumi.String("test"),
						pulumi.String("test2"),
					},
					FileCachePath: pulumi.String("/ns1"),
					Nfs: fsx.FileCacheDataRepositoryAssociationNfArray{
						&fsx.FileCacheDataRepositoryAssociationNfArgs{
							DnsIps: pulumi.StringArray{
								pulumi.String("192.168.0.1"),
								pulumi.String("192.168.0.2"),
							},
							Version: pulumi.String("NFS3"),
						},
					},
				},
			},
			FileCacheType:        pulumi.String("LUSTRE"),
			FileCacheTypeVersion: pulumi.String("2.12"),
			LustreConfigurations: fsx.FileCacheLustreConfigurationArray{
				&fsx.FileCacheLustreConfigurationArgs{
					DeploymentType: pulumi.String("CACHE_1"),
					MetadataConfigurations: fsx.FileCacheLustreConfigurationMetadataConfigurationArray{
						&fsx.FileCacheLustreConfigurationMetadataConfigurationArgs{
							StorageCapacity: pulumi.Int(2400),
						},
					},
					PerUnitStorageThroughput:   pulumi.Int(1000),
					WeeklyMaintenanceStartTime: pulumi.String("2:05:00"),
				},
			},
			SubnetIds: pulumi.StringArray{
				aws_subnet.Test1.Id,
			},
			StorageCapacity: pulumi.Int(1200),
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
import com.pulumi.aws.fsx.FileCache;
import com.pulumi.aws.fsx.FileCacheArgs;
import com.pulumi.aws.fsx.inputs.FileCacheDataRepositoryAssociationArgs;
import com.pulumi.aws.fsx.inputs.FileCacheLustreConfigurationArgs;
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
        var example = new FileCache("example", FileCacheArgs.builder()        
            .dataRepositoryAssociations(FileCacheDataRepositoryAssociationArgs.builder()
                .dataRepositoryPath("nfs://filer.domain.com")
                .dataRepositorySubdirectories(                
                    "test",
                    "test2")
                .fileCachePath("/ns1")
                .nfs(FileCacheDataRepositoryAssociationNfArgs.builder()
                    .dnsIps(                    
                        "192.168.0.1",
                        "192.168.0.2")
                    .version("NFS3")
                    .build())
                .build())
            .fileCacheType("LUSTRE")
            .fileCacheTypeVersion("2.12")
            .lustreConfigurations(FileCacheLustreConfigurationArgs.builder()
                .deploymentType("CACHE_1")
                .metadataConfigurations(FileCacheLustreConfigurationMetadataConfigurationArgs.builder()
                    .storageCapacity(2400)
                    .build())
                .perUnitStorageThroughput(1000)
                .weeklyMaintenanceStartTime("2:05:00")
                .build())
            .subnetIds(aws_subnet.test1().id())
            .storageCapacity(1200)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:fsx:FileCache
    properties:
      dataRepositoryAssociations:
        - dataRepositoryPath: nfs://filer.domain.com
          dataRepositorySubdirectories:
            - test
            - test2
          fileCachePath: /ns1
          nfs:
            - dnsIps:
                - 192.168.0.1
                - 192.168.0.2
              version: NFS3
      fileCacheType: LUSTRE
      fileCacheTypeVersion: '2.12'
      lustreConfigurations:
        - deploymentType: CACHE_1
          metadataConfigurations:
            - storageCapacity: 2400
          perUnitStorageThroughput: 1000
          weeklyMaintenanceStartTime: 2:05:00
      subnetIds:
        - ${aws_subnet.test1.id}
      storageCapacity: 1200
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amazon File Cache cache using the resource `id`. For example:

```sh
 $ pulumi import aws:fsx/fileCache:FileCache example fc-8012925589
```
 