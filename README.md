# nearmap.com Go Test

Welcome to the Nearmap Go test. The purpose of this assignment is to test your
familiarity with Go, distributed systems concepts and unit testing.

## Background

The source code that you are given is a very simple imitation of a key/value
store:

* `Database` represents a client to the central store that takes a long time
  (500ms) to store and retrieve data.
* `DistributedCache` represents a client to the distributed cache (Redis for
  example) that takes much less time to turn around (100ms to store or retrieve).

This scenario is a simplified example of a typical high performance server
cluster with a database, a distributed cache and multiple worker nodes.

## Assumptions

After startup:

* Data in `Database` never changes and can be cached forever.
* If `Database.Value()` returns `nil` for a key, the requested data item does
  not exist and will never exist.
* `DistributedCache` is initially empty.

## Task

Complete the 2 parts below and submit the solution.
If the solution is incomplete, please state what hasn't been finished and
outline how you are planning on solving it.

* Provided code can be modified at will.
* The whole solution must build with no errors.

### Part 1

Create an implementation for the `DataSource` interface to create a mechanism to
retrieve data from `Database` with lowest possible latency. An example
`LocalDataSource` struct has been provided as a start.

For a frequently-requested item your `DataSource.Value()` implementation should
have a better response time than the distributed cache store (ie < 100ms).

* The user of the `DataSource` interface must not have to deal with thread
  synchronisation.
* Write unit tests for the new `DataSource` implementation (only), and ensure
  all tests pass.
* The solution should aim to minimise calls to the database.
* Limit your use of libraries to the Go standard library only.

### Part 2

Complete `main()` to test your `DataSource` implementation; it must:

* Populate `Database` with the following data at startup:
<pre>
    | key          | value         |
    --------------------------------
    | key0         | value0        |
    | key1         | value1        |
    | key2         | value2        |
    | key3         | value3        |
    | key4         | value4        |
    | key5         | value5        |
    | key6         | value6        |
    | key7         | value7        |
    | key8         | value8        |
    | key9         | value9        |
</pre>
* Use 10 goroutines (simulating separate threads on a single worker node) each
  making 50 consecutive requests for a random key in the range (key0-key9).
  I.e. there should be a total of 500 requests.
* For each request, print the requested key name, returned value, time to
* complete that request; similar to the following example:
<pre>
    [1] Request 'key1', response 'value1', time: 50.05 ms
    [2] Request 'key2', response 'value2', time: 50.05 ms
</pre>

## Submission

* **DO NOT** publicly fork this repository or create pull requests on it as we
  don't want other candidates to see your solution.
* Please submit your solution as per instructions provided in the correspondence.
