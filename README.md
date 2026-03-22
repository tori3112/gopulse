# gopulse
GoPulse is a lightweight distributed task processing system designed to explore real-world backend patterns using Go and gRPC.

It allows clients to submit jobs, distributes them across multiple workers, and streams results back efficiently. The system demonstrates how modern microservices communicate using Protocol Buffers and streaming RPCs.

## features
- [ ] submit and process jobs asynchronously
- [ ] real-time job updates via gRPC streaming
- [ ] concurrent worker execution using goroutines
- [ ] bidirectional communication between server and workers
- [ ] pluggable job types
- [ ] job status tracking (pending, running, completed, failed

## architecture

``` mermaid
flowchart LR
  a[**coordinator**
    - manages job queue
    - assigns tasks to workers
    - tracks job lifecycle]
  b[**workers**
    - register with the server
    - receive and execute jobs
    - return results]
  c[**client**
    - submits jobs
    - queries status
    - listens to updates]

  c --> a --> b
```
