# Example for usage of gostream

![builds.sr.ht status](https://builds.sr.ht/~biosmarcel/gostreamexample/arch.yml.svg)](https://builds.sr.ht/~biosmarcel/gostreamexample/arch.yml?)

This repository demonstrates how to use
[gostream](https://github.com/Bios-Marcel/gostream).

Head over to the [build file](.builds/arch.yml) in order to see what has to be
done. You'd just get your dependencies as usual, but in addition to that, you
generate the specific versions of generic files that you want to use, you can
also generate more of them at a later point in time or simply delete the
existing ones to generate a new one with more types.

In this example I ceated a specific package for it, simply because I didn't
want the file to be in the same package as my own code.

In order to execute this example locally, run the following commands:

1. Get and install genny, but make sure your `$GOPATH/bin` is on the path
    ```shell
    go get -v -u github.com/cheekybits/genny
    go install github.com/cheekybits/genny
    ```
2. Get the source
    ```shell
    go get -v -u github.com/Bios-Marcel/gostreamexample
    ```
3. Generate the necessary code
    ```shell
    cd $GOPATH/src/github.com/Bios-Marcel/gostreamexample
    go get -v -u github.com/Bios-Marcel/gostream
    genny -in="$GOPATH/src/github.com/Bios-Marcel/gostream/stream.go" -out="stream/stream.go" -pkg="stream" gen "GenericStreamEntity=string"
    ```
4. Run the test
    ```shell
    go run main.go
    ```
