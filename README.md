# merge

automate the next steps:
1. git add -A
2. git commit -m <arg[0]>
3. git push
4. open pr
5. wait for required workflows to run
6. git merge


## Installation

### Linux/Mac

1. Download the latest release from the [releases page](https://github.com/yourusername/yourappname/releases).
2. Extract the tarball and move the binary to `/usr/local/bin`:

```sh
    # Download the latest release of the binary
    curl -OL https://github.com/oshri22004/merge/releases/latest/download/merge

    # Move the binary to /usr/local/bin
    sudo mv merge /usr/local/bin/

    # Make the binary executable
    sudo chmod +x /usr/local/bin/merge

    # Verify the installation
    merge --version
```
