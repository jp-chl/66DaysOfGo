# Day 6 of #66DaysOfGo

_Last update:  Jul 18, 2023_.

---

Today, I've continued with the Concurrency series, this time understanding a more realistic scenario with a parallel approach.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6
- pre-commit: 3.3.3

---

## Setup

The following examples are available in the official [Go Blog](https://go.dev/blog/), article "[Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)", specifically in the files [serial.go](https://go.dev/blog/pipelines/serial.go), [parallel.go](https://go.dev/blog/pipelines/parallel.go) and [bounded.go](https://go.dev/blog/pipelines/bounded.go) (also available in the current directory of this repo).

## Real example: MD5 digest (sequential version)

The aforementioned three files/approach do the following:

- The [serial.go](./serial.go) version computes the MD5 checksum for each file in a directory sequentially. It uses the `filepath.Walk` function to navigate the file structure and `ioutil.ReadFile` to read each file's data. It then applies the `md5.Sum` function to compute the hash.
- The [parallel.go](./parallel.go) version optimizes this process by computing the MD5 checksums concurrently using goroutines. The `sumFiles` function walks the file tree, spinning up a new goroutine to read and compute the MD5 sum for each file. It uses channels to communicate the results and any errors that occur.
- However, the previous (parallel) version can run into memory issues if the directory contains many large files. To solve this, the [bounded.go](./bounded.go) version introduces _bounded parallelism_, limiting the number of files read in parallel by creating a fixed number of goroutines for reading files (adding a limited number to a `sync.WaitGroup`). This pipeline version consists of three stages: walking the file tree, reading and digesting the files, and collecting the digests.

All versions result in a map of file paths to their corresponding MD5 checksums, which is then sorted and printed by the main function.

## Sequential version

The [serial.go](./serial.go) works similarly to the `md5sum` command line utility. It computes the MD5 checksum for all files in a given directory, and then prints the sorted digest values associated with each file's path.

The main function calls a helper function `MD5All` that returns a map linking each file path to its MD5 checksum. `MD5All`, defined in _serial.go_, reads and sums each file's contents while walking the file tree without concurrency. It returns a map from the file path to the MD5 sum of the file's contents. If a directory walk fails, or any read operation fails, `MD5All` returns an error. This approach works sequentially, processing one file at a time.

Main parts of the go code:

> Reads every file and puts the path as a map key, and its md5 sum as its value

```go
func MD5All(root string) (map[string][md5.Size]byte, error) {
    m := make(map[string][md5.Size]byte)
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        // shortened for brevity...
        data, err := ioutil.ReadFile(path)
        // shortened for brevity...
        m[path] = md5.Sum(data)
        return nil
    })
    // shortened for brevity...
    return m, nil
}
```

```go
func main() {
    m, err := MD5All(os.Args[1])
    // shortened for brevity...
    var paths []string
    for path := range m {
        paths = append(paths, path)
    }
    sort.Strings(paths)
    // shortened for brevity...
}
```

## Parallel version

The [parallel.go](./parallel.go) it's an enhanced version of the MD5 checksum program, where the computation process is split into a two-stage pipeline using goroutines.

The first stage, `sumFiles`, walks the file tree, spawning a new goroutine to read and compute the MD5 sum for each file, sending the results over a channel. The `MD5All` function, responsible for initiating this process, receives the digest values from the channel and returns early in case of error, closing the `done` channel. By computing the MD5 checksums in parallel, the program is optimized for increased efficiency. However, this approach might run into memory issues with a large directory of sizeable files.

The main function section stays the same as the previous version.

Main parts of the go code:

> Results to be sent in a channel

```go
type result struct {
    path string
    sum  [md5.Size]byte
    err  error
}
```

> The MD5All function now uses a done channel and leverage the digest retrieval concurrently via a result channel ("c")

```go
func MD5All(root string) (map[string][md5.Size]byte, error) {
    done := make(chan struct{})
    defer close(done)

    // See the next function
    c, errc := sumFiles(done, root)

    m := make(map[string][md5.Size]byte)
    for r := range c {
        // shortened for brevity...
        m[r.path] = r.sum
    }
    // shortened for brevity...
}
```

> For each file, a new goroutine sends data to the results channel (`c`)

```go
func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
    c := make(chan result)
    errc := make(chan error, 1)
    go func() {
        var wg sync.WaitGroup
        err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
            // shortened for brevity...
            wg.Add(1)
            go func() {
                data, err := ioutil.ReadFile(path)
                select {
                case c <- result{path, md5.Sum(data), err}:
                case <-done:
                }
                wg.Done()
            }()
            // Abort the walk if done is closed.
            select {
            case <-done:
                return errors.New("walk canceled")
            default:
                return nil
            }
        })
        // Walk has returned, so all calls to wg.Add are done.  Start a
        // goroutine to close c once all the sends are done.
        go func() {
            wg.Wait()
            close(c)
        }()
        // No select needed here, since errc is buffered.
        errc <- err
    }()
    return c, errc
}
```

## Bounded parallelism

The [bounded.go](./bounded.go) is a further optimization of the MD5 checksum program. It introduces a three-stage pipeline that restricts the number of files read concurrently, mitigating potential memory issues. It uses the done-channel pattern as well.

The first stage, `walkFiles`, emits the paths of regular files in the tree. The second stage runs a fixed number of `digester` goroutines that receive filenames, read and digest the files, and send the results over a channel. The final stage, implemented in `MD5All`, receives all the results and arranges for the channel to be closed when all digesters are done. This method ensures efficient parallel processing while keeping memory allocation within limits.

The main function also remains.

Main parts of the go code:

> The MD5All controls how many concurrent digestions are in place. It starts a fixed number of goroutines to read and digest files (`wg.Add(20)`); A new function `walkFiles` is in charge of reading all the files

```go
func MD5All(root string) (map[string][md5.Size]byte, error) {
    done := make(chan struct{})
    defer close(done)

    paths, errc := walkFiles(done, root)

    c := make(chan result)

    var wg sync.WaitGroup
    const numDigesters = 20
    wg.Add(numDigesters)
    
    for i := 0; i < numDigesters; i++ {
        go func() {
            digester(done, paths, c)
            wg.Done()
        }()
    }
    go func() {
        wg.Wait()
        close(c)
    }()
   // shortened for brevity...
}
```

> Now the digest retrieval is in a separate function that receives paths from a channel and puts results in the respective channel

```go
func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
    for path := range paths {
        data, err := ioutil.ReadFile(path)
        select {
        case c <- result{path, md5.Sum(data), err}:
        case <-done:
            return
        }
    }
}
```

> The `walkFiles` function defer the digest ingestion by sending the request through the `paths` channel

```go
func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
    paths := make(chan string)
    errc := make(chan error, 1)
    go func() {
        // Close the paths channel after Walk returns.
        defer close(paths)
        // No select needed for this send, since errc is buffered.
        errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
            // shortened for brevity...
            select {
            case paths <- path:
            case <-done:
                return errors.New("walk canceled")
            }
            return nil
        })
    }()
    return paths, errc
}
```

## References

- [Go Concurrency Patterns: Pipelines and cancellation (Go Blog)](https://go.dev/blog/pipelines)
