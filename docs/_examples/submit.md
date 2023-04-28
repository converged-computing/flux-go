---
title: Submit
layout: snippet-page
description: How to submit a job to Flux from Go
---

# Submit Example

Here is an example of how to submit a job. Each of the files below is provided in
completion, meaning that they include all of the headers that are necessary
to build with Flux.

## 1. Create a Flux Handle

You might first want to start with
creating a function to instantiate a Flux Handle.

{% include snippet.html language="go" id="submitgo" path="pkg/core/handle.go" %}

While a lot of this isn't necessarily needed, this will give you control over
determing if the user isn't running in an environment where Flux is available,
and taking the appropriate action.

## 2. Create a JobSpec

The specification for a job in flux is called a "jobspec." It comes down to a json
object with a series of expected fields describing the job you want to submit
(e.g., command, environment) and then the resources that you need. The example
below shows creating your own `JobSpec` type to interact with.

{% include snippet.html language="go" id="jobgo" path="pkg/core/job.go" %}

## 3. Write a Submit Function

To submit a job, you'll want to create a new JobSpec with your parameters of
interest, ensure it is encoded, and then call the C function to submit the job.
This is illustrated in the `core.go` file of the repository here:

{% include snippet.html language="go" id="corego" path="pkg/core/core.go" %}

## 4. Put it all together!

For the complete example, we will want to create the handle (checking that we have Flux)
and then call the function that we wrote above.

{% include snippet.html language="go" id="maingo" path="cmd/submit/main.go" %}


On the command line (after compiling this entrypoint) the interaction might look
like the following:

```bash
$ ./bin/fluxgo-submit 
```
```console
⭐️ Testing flux submit in Go! ⭐️
Submitting a Sleep Job: sleep 10
...
```

And you can see the job is running!

```bash
$ flux jobs -a
```
```console
       JOBID USER     NAME       ST NTASKS NNODES     TIME INFO
    ƒdcsg8Nj vscode   sleep       R      1      1   2.728s 5e3ccb811e04
```

We likely will add more output from the job (e.g., the identifier) and metadata,
but for now you can use `flux jobs` to see that the job ran (above) and completed:

```bash
$ flux jobs -a
       JOBID USER     NAME       ST NTASKS NNODES     TIME INFO
   ƒ3Mop5847 vscode   sleep      CD      1      1   10.04s 3da6fb865c22
```

Of course you would likely use this in your own library, and do something more
interesting! More examples coming soon!

