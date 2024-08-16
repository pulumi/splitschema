# Pulumi Schema Splitting

_Efficient and human-friendly schema serialisation._

Historically, schema and metadata has always been stored and embedded within providers as a single json file. For non-trivial provider, the process of deserializing the whole file can take a noticeable amount of time. This library helps with this problem in two ways:

1. Only the segment of the schema that's needed can be quickly accessed without parsing the whole schema.
2. When parsing the whole schema, we can perform parsing of each segment in parallel.

## Usage

### CLI

Install:

```bash
go install github.com/pulumi/splitschema/cmd/splitschema
```

Usage:

```bash
splitschema split -s schema.json -d schema-dir
splitschema merge -s schema-dir -d schema.json
```

### In code

Writing:

```go
var pkg *schema.PackageSpec
err = json.Unmarshal(bytes, &pkg)
// Write to the "schema" folder
err := WritePackageSpec("schema", pkg)
```

Embedding & reading:

```go
//go:embed schema/*
var schema embed.FS

pkg := NewPartialPackage(awsEmbeddedSplit, "schema")
// Read a single resource
instanceSpec, err := pkg.GetResource("aws:ec2/instance:Instance")
// Put the whole package back together
pkgSpec, err := pkg.ReadPackageSpec()
```

## File Structure

- `core.json`: the original schema excluding resources, functions and types.
- `resources.json`, `functions.json`, `types.json`: map of all available tokens for iteration.
- `{module}/[resources|functions|types]/{name}-{hash}.json`: A single resource/function/type specification.
- `{module}/[resources|functions|types]/{name}-{hash}.md`: Optional standalone description for any multi-line descriptions.

### Example

```
- core.json                                      Core schema fields
- functions.json                                 List of function tokens
- resources.json                                 List of resource tokens
- types.json                                     List of type tokens
- vpc/                                           vpc module
  - functions/                                   Functions in the vpc module
    - getsecuritygrouprule-de68e9b0.json         getSecurityGroupRule function spec
    - getsecuritygrouprule-de68e9b0.md           getSecurityGroupRule description
    - getsecuritygrouprule-de68e9b0.meta.json    Additional provider-specific metadata for the function
  - resources/                                   Resources in the vpc module
    - securitygroupegressrule-d60ba7ae.json      SecurityGroupEgressRule resource spec
    - securitygroupegressrule-d60ba7ae.md        SecurityGroupEgressRule description
    - securitygroupegressrule-d60ba7ae.meta.json Additional provider-specific metadata for the resource
  - types/                                       Types in the vpc module
    - getsecuritygrouprulefilter-2c6dddd7.json   GetSecurityGroupRuleFilter type spec
```

## Key Features

- **Lazy Loading**: Only the parts of the package which are requested are read, then cached.
- **Thread safe**: Multiple goroutines can request different or the same elements at the same time. This uses lock-free strategies.
- **Fast**: Batch loading is performed in parallel.
- **Human Readable**: The structure of the filesystem is designed to be easily navigable & good for showing diffs.
- **Custom Metadata**: Add extra, typed metadata for the provider's internal use (e.g. mappings to API endpoints)

## Potential Applications

- Improving YAML language support responsiveness by reducing start-up time.
- Allow the engine to interact with provider's schemas. This is a foundation for the engine to be able to request specific parts of the schema from a provider.
- Reduce our dependency on committing SDKs. One reason we commit all SDKs is that they provide a better diff of how the schema has changed than the monolithic schema file.

Prior art: [github.com/pulumi/pulumi-schema-explode](https://github.com/pulumi/pulumi-schema-explode).

## Performance

| Reading | Time |
| -- | -- |
| AWS Embedded Monolith - Schema | 290ms |
| AWS Embedded Split - Schema | 62ms |
| AWS Embedded Split - Single resource | 0.4ms |

## Developing Locally

1. Run `go generate` to populate the `testdata` directory.
2. Run `go test -bench .` to run tests including benchmarks.
