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
   wget https://github.com/yourusername/yourappname/releases/download/v1.0.0/yourappname.tar.gz
   tar -xzvf yourappname.tar.gz
   sudo mv yourappname /usr/local/bin/
   ```
