# KNE Binding

The KNE Binding is an implementation of the Ondatra binding interface that runs
on a network topology created with [openconfig/kne](https://github.com/openconfig/kne).

## Config File

The KNE Binding requires a YAML config file that lets the binding know how to
connect to your KNE topology. The YAML must be specified in a `-config` flag
passed to the Ondatra test. The file supports the following keys:

Key        | Required? | Description
---------- | :-------: | ----------------------------------
`username` | yes       | username to log into the KNE nodes
`password` | yes       | password to log into the KNE nodes
`topology` | yes       | path to a KNE topology text proto
`cli`      | no        | path to the kne binary
`kubecfg`  | no        | path to your kubeconfig file

If `cli` and `kubecfg` are not specified, they will be inferred from the `PATH`
environment.

An example YAML config file:

```
username: tester
password: hunter2
topology: /home/tester/topo.textproto
```

## Running the Integration Test

This repo includes an
[example integration test](integration/integration_test.go) that uses the KNE
binding, as well as a [testbed file](integration/testbed.textproto) for that
test. To execute the test, you must:

1.  create a local KNE topology with at least two linked nodes, as the testbed
    requires
1.  create a config file for your topology, as specified above
1.  run the test passing both the testbed and config file flags:

```
go test -testbed=testbed.textproto -config=path/to/config.yaml
```

One example KNE topology that satisfies this testbed is [2node-ixia-ceos.pb.txt](https://github.com/openconfig/kne/blob/main/examples/2node-ixia-ceos.pb.txt)
provided by the openconfig/kne project. A YAML config file that works with that
topology is:

```
username: admin
password: admin
topology: /[YOUR GIT CLONE PATH]/kne/examples/2node-ixia-ceos.pb.txt
```
