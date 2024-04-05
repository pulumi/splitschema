> **NOTE:** The `aws.s3.getBucketObjects` data source is DEPRECATED and will be removed in a future version! Use `aws.s3.getObjects` instead, where new features and fixes will be added.

> **NOTE on `max_keys`:** Retrieving very large numbers of keys can adversely affect this provider's performance.

The objects data source returns keys (i.e., file names) and other metadata about objects in an S3 bucket.
