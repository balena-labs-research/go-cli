<img src="https://github.com/maggie0002/balena-apps-logo/raw/main/logo.png" width="75" />

# Balena Go CLI

This is an experimental community project that provides some of the basic balena CLI functions in Golang. It was built to allow use of CLI features inside of a container on a Raspberry Pi but works on Windows, Mac and Linux too. It is made available here for others to use and for additional features to be added.

## Basic Usage:

Download the file based on your system type from the releases page. `go-cli -help` will show available options.

You can start with the `-api` flag which starts an API on port 7878 waiting for requests.

There are docker containers published in the `Packages` section of the GitHub repo.

## Supported Features:

Currently the following features are supported both for standalone use and the API:

- Scan the network for balena devices running in development mode

## Contributing

Contributions are welcome to help grow the number of features. It may be wise to create an issue or discussion topic before starting to help us all coordinate.
