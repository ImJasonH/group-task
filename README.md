# GroupTask for Tekton

This repo provides an experimental [Tekton Custom
Task](https://github.com/tektoncd/community/pull/128) that specifies multiple
`Task`s to run in the same underlying Pod. This can enable faster execution of
multiple Tasks that share underlying resources (e.g., workspaces), since they
can avoid scheduling PVCs.

The intention is to demonstrate the kinds of things a Custom Task can do, and
to demonstrate how to write a Custom Task.

## Install

Install and configure `ko`.

```
ko apply -f controller.yaml
```

This will build and install the controller on your cluster, in the namespace
`group-task`.

## TODO: Define a GroupTask

## TODO: Run the GroupTask.
