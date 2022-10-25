# cumulus

A library and command line tool for concurrent bulk investigation and occasional manipulation of
multiple clouds.

# TL;DR

Do you

* have multiple AWS accounts or multiple regions and need to list resources across them?
* get tired of the aws cli baroque command syntax?
* get annoyed by how slow and awkward the aws cli is?

This tool might be for you (at least eventually)

Futures:  many more resource types, DigitalOcean, GCP, and Kubernetes

## Development Status

Becoming Useful. Only handles AWS, only handles a few resource types, but handles many
accounts/regions concurrently.

# Installation

Pick one:

* MacOS: `brew install deweysasser/tap/Cumulus`
* Pre-built binaries from [GitHub release page](https://github.com/deweysasser/cumulus/releases).
* Other platforms (deb, rpm) TBD as I get a curiosity bump or someone needs them.
* You can always `go install github.com/deweysasser/cumulus@latest` (currently the
only thing you'll miss out on is the Makefile embedding version information in the binary.)

# Running

## Quickstart

There's nothing to set up. The command will pick up all your profiles from `~/.aws/credentials`,
deduplicate them (i.e. you may have more than 1 pointing to an account -- it will take care of it),
and use them.

If you use AWS creds in your environment -- file a bug and I'll make that work. It's not really
important to me using multiple clouds.

## Examples

```commandline
 $ time bin/cumulus instance list > /dev/null

real	0m5.585s
user	0m0.718s
sys	    0m0.141s
```

```text
8:19PM INF discovered instance count=259
8:19PM INF timer average_latency=195.194586 count=47 name="AWS API calls" rate=8.566558814468866
8:19PM INF timer totals TotalCalls=47 TotalDuration=9174.145546 name="AWS API calls"
8:19PM INF run time duration=5484.798748
```

```commandline
$ time ./bin/cumulus  instance list -I '^type$' -X . | sort | uniq -c | sort -rn 
  59 c5.2xlarge
   4 c5.xlarge
   3 c5.4xlarge
   1 c5.18xlarge
 ...

real	0m5.935s
user	0m0.738s
sys	0m0.150s
```

(Breakdown:  List the instances, include the field matching regexp `^type$` in output, exclude all
other fields. You can figure out the rest of the pipeline.)

Other possibilities:

```commandline
# Show the "environment" tag in a column
$ cumulus instance list -I tag:environment

# show only instances where the "environment" tag is "production"
$ cumulus instance list -l tag:environment=production

# show all instances where *any* column contains the word "test"
$ cumulus instance list -l .=test

# show all available fields for snapshots.  Warning:  there are a *LOT* of them.
$ cumulus snapshot list -v

# show snapshots tagged with "Release" containing "v1.0" (it's a regexp) *AND* "Status" containing "alpha"
$ cumulus snapshot list -l tag:Release=v1.0,tag:Status=alpha
```

It has similar performance for listing instances, zones...everything else (except, oddly, the API
for S3 is pretty slow)

## Command Overview

```commandline
$ bin/cumulus  -h
Usage: cumulus <command>

Flags:
  -h, --help       Show context-sensitive help.
      --version    Show program version

Output
      --debug                   Show debugging information
      --output-format="auto"    How to show program output (auto|terminal|jsonl)
  -q, --quiet                   Be less verbose than usual
  -v, --verbose                 Be more verbose than usual

Profile
  --profile.cpu       profile the CPU
  --profile.memory    profile the Memory usage

Commands:
  account list

  instance list

  snapshot list

  machine-image list

  volume list

  dns zone list

  dns record list
```

# More detailed Development Status

Becoming Useful. Only handles AWS, only handles a few resource types, but handles many
accounts/regions concurrently.

Allows column selection and filtering.

Very fast

CLI might be awkward for now -- I haven't fully exercises all the options and made sure that
everything behaves senibly all the time (e.g. when running the "--version" flag, you still have to
give a subcommand)

Currently handles:

* AWS
    * listing
        * accounts
        * instances
        * snapshots
        * AMIs
        * EBS volumes
        * Route53 Zones
        * Route53 Recordsets

# Plans and Intentions

* Codgen as much as possible
* Handle many more AWS types (all the ones I use regularly)
* Add kubernetes (discover clusters across accounts, *then* discover resources inside them)
* Add Digital Ocean
* Add GCP
* Add Azure?  (I don't currently use Azure, but maybe by this time I'll find someone who does to add
  to it)

# Why am I writing this?

The AWS command line is slow. S....l....o....w.

It's also...awkwardly laid out, in the format of "service - endpoint". GCP and Digital Ocean CLIs
are much more intuitively designed in the form of "subject - object - verb - modifiers".

And, last, the AWS CLI works with a single account, and a single region, at a time. Even wehn I'm
exclusively using AWS, this is not the right abstraction.

Oh, and very last...I have hit some non-linearity in the AWS command line, where large result sets
are incredibly delayed. I'm not that patient.

So the goal of this tool is to produce a cloud library that's much better structured, faster,
concurrent, and not tied strictly to AWS.

That last part is the most experimental. I have tried better structures for library and found them
good. In this tool, I'm attempting a "common denominator" API approach to multiple clouds
and...well...we're going to see how well that works.

Also, I'm using that library to create a better CLI.

# Features and Bugs

Please use [Github Issues](https://github.com/deweysasser/cumulus/issues) for all bugs and feature
requests. 