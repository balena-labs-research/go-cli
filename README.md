<img src="https://raw.githubusercontent.com/balena-labs-research/apps-logo/main/logo.png" width="75" />

# Balena Go CLI

This is an experimental community project that provides some of the basic balena CLI functions in Golang. It was built to allow use of CLI features inside of a container on a Raspberry Pi but works on Windows, Mac and Linux too. It is made available here for others to use and for additional features to be added.

## Basic Usage:

Download the file based on your system type from the releases page. `balena-go -help` will show available options.

There are docker containers published in the `Packages` section of the GitHub repo.

## Supported Features:

Currently the following features are supported both for standalone use and the API:

- Scan the network for balena devices running in development mode
- Connect to local development devices on the network with SSH
- Stream container logs from devices on your network

## Contributing

Contributions are welcome to help grow the number of features. It may be wise to create an issue or discussion topic before starting to help us all coordinate.
