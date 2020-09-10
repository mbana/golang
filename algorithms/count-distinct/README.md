# `countdistinct`

[![CircleCI](https://circleci.com/gh/banaio/countdistinct.svg?style=svg)](https://circleci.com/gh/banaio/countdistinct)

Implementations of the [count-distinct problem](https://en.wikipedia.org/wiki/Count-distinct_problem) algorithms in Go.

## Install

```bash
$ go get -u github.com/banaio/countdistinct/cmd/...
$ countdistinct --help
In computer science, the count-distinct problem (also
known in applied mathematics as the cardinality estimation
problem) is the problem of finding the number of distinct
elements in a data stream with repeated elements. This is
a well-known problem with numerous applications. The elements
might represent IP addresses of packets passing through a router,
unique visitors to a web site, elements in a large database,
motifs in a DNA sequence, or elements of RFID/sensor networks.

See: https://en.wikipedia.org/wiki/Count-distinct_problem

Usage:
  countdistinct [flags]

Flags:
  -a, --algorithm string   The algorithm to use (default "pcsa")
  -f, --file string        File containing the elements to add (default "/Users/mbana/dev/banaio/github/countdistinct/countdistinct-elements.txt")
  -h, --help               help for countdistinct
```

## Reading

### Probabilistic Counting with Stochastic Averaging (PCSA)

* [Probabilistic Counting with Stochastic Averaging](https://bana.io/blog/pcsa/)
* [Sketch of the Day: Probabilistic Counting with Stochastic Averaging (PCSA)](https://research.neustar.biz/2013/04/02/sketch-of-the-day-probabilistic-counting-with-stochastic-averaging-pcsa/)
* [Flajolet–Martin algorithm](https://en.wikipedia.org/wiki/Flajolet%E2%80%93Martin_algorithm)
* [Probabilistic Counting Algorithms for Data Base Applications](http://algo.inria.fr/flajolet/Publications/src/FlMa85.pdf)
