# ftpctl

This is an ftp configuration tool written in golang language, which can help you to easily solve the upload, download and other operations

## Usage

You should download the binary installer ftpctl for the corresponding operating system, and execute the following command.
```shell
mv ftpctl /usr/local/bin/
```
You can use this command globally as follows.
```shell
[root@localhost ~]# ftpctl -h
NAME:
   ftpctl - it is a ftp server command line configuration tool

USAGE:
   ftpctl [global options] command [command options] [arguments...]

VERSION:
   0.1

COMMANDS:
   upload    
   download  
   delete    
   mkdir     
   rmdir     
   get       
   rename    
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
[root@localhost ~]# ftpctl upload -h
NAME:
   ftpctl upload

USAGE:
   ftpctl upload [command options] [arguments...]

DESCRIPTION:
   upload file to the remote ftp server of the dirctory, the default directory is /

OPTIONS:
   --host value, -a value      Host ip address for remote connection to ftp server. (default: "127.0.0.1")
   --username value, -u value  The username of login ftp server. (default: "admin")
   --password value, -p value  The password of login ftp server. (default: "admin")
   --srcfile value, -s value   This file is used for upload or download.
   --desfile value, -d value   Rename the file to a new name.
   --path value, -P value      The dirctory is used for upload or download.
   --help, -h                  show help
```
## simple example

* upload file

```shell
[root@localhost ~]# ftpctl  upload -a 192.168.10.179 -u helloftp -p 123456 -s helloworld.yaml 
[root@localhost ~]# ftp 192.168.10.179
Connected to 192.168.10.179 (192.168.10.179).
220 (vsFTPd 3.0.2)
Name (192.168.10.179:root): helloftp
331 Please specify the password.
Password:
230 Login successful.
Remote system type is UNIX.
Using binary mode to transfer files.
ftp> ls
227 Entering Passive Mode (192,168,10,179,24,158).
150 Here comes the directory listing.
-rw-r--r--    1 1001     1001          425 Apr 11 06:18 helloworld.yaml
```

* delete file

```shell
[root@localhost ~]# ftpctl delete  -a 192.168.10.179 -u helloftp -p 123456 -s helloworld.yaml 
[root@localhost ~]# ftp 192.168.10.179
Connected to 192.168.10.179 (192.168.10.179).
220 (vsFTPd 3.0.2)
Name (192.168.10.179:root): helloftp
331 Please specify the password.
Password:
230 Login successful.
Remote system type is UNIX.
Using binary mode to transfer files.
ftp> ls
227 Entering Passive Mode (192,168,10,179,24,158).
150 Here comes the directory listing.
226 Directory send OK.
```

