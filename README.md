## Setup
The script setup.sh takes configurations from setup.config.
The configuration file has these fields:
 - USAGE takes SETUP, RUN, or STOP. 
    - SETUP prepares each VM by cloning or pulling from the git repo, as well as downloading Go. 
    - RUN sets the necessary variables to run Go, retrieves the log files from the CS425 piazza page, and starts the servers. 
    - STOP stops the servers and clears the ports that they were running on.
 -  VM_NODES takes the host names of the VMs, separated by commas.
 - DIRECTORY is the directory holding the git repository.
 - GOPKG is the binary download of Go.
 - HOME is the home directory of the VMs.

To prepare the VMs, run `./setup.sh setup.config` with USAGE set to SETUP.

## Usage
1. Start all the servers. This can be done easily using the "RUN" setting in the [Setup Script](#setup).
2. ssh into a vm `$ ssh tkao4@fa17-cs425-g46-01.cs.illinois.edu`
3. Make sure you can run the go command `$ go`. If not, set the PATH enviorment variable `$  export PATH=$PATH:/usr/local/go/bin`
4. Run the client.
    ```
    $ cd CS425-MP1/src/
    $ go run client/main.go [grep flags] [pattern]
    ```
    Note: You should not enter the "grep" command itself or the name of a file to search through. We automatically grep through files named "machine.i.log" (i corresponding to the machine number).
5. The grep result from each machine is printed out into files named `grep_i_out.txt`, where `i` corresponds to the machine number. The corresponding line count is printed to the terminal. 