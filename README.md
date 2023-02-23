<img src="https://raw.githubusercontent.com/balena-labs-research/apps-logo/main/logo.png" width="75" />

# Balena Go Developer CLI

This is an experimental community project that provides some basic CLI functions for development on balena IoT devices. It was initially built to allow device scanning inside of a container on a Raspberry Pi. It is made available here for others to use and for additional features to be added.

_Next Steps to be useable:_

- [ ] The `arpscan` command needs to be run as `sudo` to allow it to capture the ARP packets on the primary interface. At the moment there is no warning to the user about this. Need to add a function that checks if it is running as `sudo` and if not, exit with a message.
- [ ] Binaries are not Apple or Windows signed. This will make it difficult to run, and I think impossible to run the arm64 binary on a Mac. In the meantime, you can build the binary yourself from source by cloning the repo and running `make build`, but ultimately should add signing to make the binaries in the releases section more accessible.
- [ ] Update Mutagen to v `0.17`, which although isn't necessary, it has some nice performance improvements for the mounts.

## Basic Usage:

Download the file based on your system type from the releases page. `balena-go --help` will show available options.

There are docker containers published in the `Packages` section of the GitHub repo.

## Mutagen

Mutagen is integrated in to this CLI to allow mounting the contents of a container to the local file system. For the service to work, `mutagen-agents.tar.gz` needs to be present alongside the `balena-go` executable as it contains the agents to be deployed in to the running container.

## Contributing

Contributions are welcome to help grow the number of features. It may be wise to create an issue or discussion topic before starting to help us all coordinate.
