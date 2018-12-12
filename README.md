# Example for usage of gostream

This repository demonstrates how to use
[gostream](https://github.com/Bios-Marcel/gostream).

Head over to the [build file](.builds/arch.yml) in order to see what has to be
done. You'd just get your dependencies as usual, but in addition to that, you
generate the specific versions of generic files that you want to use, you can
also generate more of them at a later point in time or simply delete the
existing ones to generate a new one with more types.

In this example I ceated a specific package for it, simply because I didn't
want the file to be in the same package as my own code.