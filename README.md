# idpp [![GoDoc](https://godoc.org/github.com/nzmprlr/idpp?status.svg)](http://godoc.org/github.com/nzmprlr/idpp) [![Go Report Card](https://goreportcard.com/badge/github.com/nzmprlr/idpp)](https://goreportcard.com/report/github.com/nzmprlr/idpp) [![Coverage](http://gocover.io/_badge/github.com/nzmprlr/idpp)](http://gocover.io/github.com/nzmprlr/idpp)

idpp is a time based 8 and 12 byte unique id generator.

8 byte ID is composed of

    42 bits for time in unix nanoseconds
    10 bits for a worker id
    12 bits for a sequence number

12 byte ID is composed of

    64 bits for time in unix nanoseconds
    16 bits for a worker id
    16 bits for a sequence number

## Usage

``` go
id := idpp.NewID8()
fmt.Println(id, id.Hex(), id.Time(), id.WorkerID(), id.Sequence())
// output: 4375014701438501 f8b0d7d2a9625 2021-02-20 18:16:54.701006848 +0300 +03 681 1573

time.Sleep(2 * time.Millisecond)

id = idpp.NewID8()
fmt.Println(id, id.Hex(), id.Time(), id.WorkerID(), id.Sequence())
// output: 4375014705632806 f8b0d7d6a9626 2021-02-20 18:16:54.705201152 +0300 +03 681 1574

id = idpp.NewID12()
fmt.Println(id, id.Hex(), id.Time(), id.WorkerID(), id.Sequence())
// output: f8b0d7d54c11092a95627 f8b0d7d54c11092a95627 2021-02-20 18:16:54.704202 +0300 +03 37545 22055

time.Sleep(1 * time.Nanosecond)

id = idpp.NewID12()
fmt.Println(id, id.Hex(), id.Time(), id.WorkerID(), id.Sequence())
// output: f8b0d7d551ae892a95628 f8b0d7d551ae892a95628 2021-02-20 18:16:54.704225 +0300 +03 37545 22056
```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.