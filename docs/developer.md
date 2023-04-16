# Flux Go Development

These are notes for Go developers of Flux!

## Development Environment

We recommended that you use the provided (VSCode) DevContainers environment.
This will provide Flux and Go already installed, so you can jump right in
to development! You can follow the [tutorial](https://code.visualstudio.com/docs/remote/containers-tutorial) where you'll basically
need to:

1. Install Docker, or compatible engine
2. Install the [Development Containers](vscode:extension/ms-vscode-remote.remote-containers) extension

Then you can go to the command palette (View -> Command Palette) and select `Dev Containers: Open Workspace in Container.`
This will build a development environment from [fluxrm/flux-sched](https://hub.docker.com/r/fluxrm/flux-sched/tags).

You are free to change the base image and rebuild if you need to test on another operating system!
When your container is built, when you open `Terminal -> New Terminal`, surprise! You're
in the container! 

**Important** the development container assumes you are on a system with uid 1000 and gid 1000. If this isn't the case,
edit the [.devcontainer/Dockerfile](.devcontainer/Dockerfile) to be your user and group id. This will ensure
changes written inside the container are owned by your user. It's recommended that you commit on your system
(not inside the container) because if you need to sign your commits, the container doesn't
have access and won't be able to. If you find that you accidentally muck up permissions
and need to fix, you can run this from your terminal outside of VSCode:

```bash
$ sudo chown -R $USER .git/
# and then commit
```

## Building Bindings

You can build the examples using the [Makefile](../Makefile). You can inspect it 
to see how simple it is - we target the Flux install in the container (and you
would want to customize this based on your install of Flux).

```bash
$ make
```

This will create a `bin` in the root directory, and each executable demonstrates 
an example that you can match to a command file in `cmd`.  For any of the examples
below, since the container starts running a flux instance, you won't need to start
one. If you are developing (and an instance isn't started) you can do:

```bash
$ flux start --test-size=4
```

### Submit Example

Here is an example of running submit. You can target this executable:

```bash
$ ./bin/fluxgo-submit 
```
```console
⭐️ Testing flux-core in Go! ⭐️
Submitting a Sleep Job: sleep 10
Flux Future: &{{{}}}
```

We likely will add more output from the job (e.g., the identifier) and metadata,
but for now you can use `flux jobs` to see that the job ran and completed:

```bash
vscode@3da6fb865c22:/workspaces/flux-go$ flux jobs -a
       JOBID USER     NAME       ST NTASKS NNODES     TIME INFO
   ƒ3Mop5847 vscode   sleep      CD      1      1   10.04s 3da6fb865c22
```

Of course you would likely use this in your own library, and do something more
interesting! More examples coming soon!
