Creates an Amazon CloudHSM v2 cluster.

For information about CloudHSM v2, see the
[AWS CloudHSM User Guide](https://docs.aws.amazon.com/cloudhsm/latest/userguide/introduction.html) and the [Amazon
CloudHSM API Reference][2].

> **NOTE:** A CloudHSM Cluster can take several minutes to set up.
Practically no single attribute can be updated, except for `tags`.
If you need to delete a cluster, you have to remove its HSM modules first.
To initialize cluster, you have to add an HSM instance to the cluster, then sign CSR and upload it.


## Import

Using `pulumi import`, import CloudHSM v2 Clusters using the cluster `id`. For example:

```sh
 $ pulumi import aws:cloudhsmv2/cluster:Cluster test_cluster cluster-aeb282a201
```
 