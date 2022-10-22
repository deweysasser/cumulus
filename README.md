# cumulus

A library and command line tool for concurrent bulk investigation and occational manipulation of
multiple clouds.

# Status

HIGHLY EXPERIMENTAL

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

## Examples

```shell
$ time aws --region us-east-1 ec2 describe-snapshots --owner-id 536746477680 | grep SnapshotId | wc -l
      20

real	0m1.135s
user	0m0.695s
sys	0m0.246s
```

vs

checking 4 regions in each of 5 accounts

```shell
$ time ./bulk-cloud snapshot list | wc -l
     446

real	0m0.968s
user	0m0.292s
sys	0m0.096s
```

It has similar performance for listing instances, zones...everything else (except, oddly, the API
for S3 is pretty slow)

