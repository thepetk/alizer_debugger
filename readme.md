# Alizer Debugger Script
This repo is a simple golang script which can be used in order to debug the library usage of [Alizer (Application Analyzer)](https://github.com/redhat-developer/alizer/).

## Overview
The script has two parts:

* In the first part it tries to match a devfile with a given path.
* In the second part it detects all components of the given path.

## Usage
In order to use the script you will have to update first the `projectAbsolutePath` variable inside `main.go` with the path of the application you want to check. In case you want to use another registry, you will have to update `devfileRegistryURL` too.

After this update you will have to build and run the binary:
```bash
$ go build -o alizer_debugger main.go
$ ./alizer_debugger
```

## Contributing
This is an open source project open to anyone. This project welcomes contributions and suggestions!

For any suggestion, bug, idea feel free to open an issue.

## License
EPL 2.0, See [LICENSE](LICENSE) for more information.
