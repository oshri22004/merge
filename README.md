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
   wget https://github.com/oshri22004/merge/blob/dev/dist/merge.tar.gz
   tar -xzvf merge.tar.gz
   sudo mv merge /usr/local/bin/
   ```
