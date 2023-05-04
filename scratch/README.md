# output

below is output of this program, playing around with some flow, logging, breaking out of a loop, etc...


```
❯ go run .
2023/05/02 20:02:53 Using switch/case
2023/05/02 20:02:53 You really are elite, aren't you? 1337 the type is: int
2023/05/02 20:02:53 ERROR ITS 1337
2023/05/02 20:02:53 if/then block: testing value 42 loop increment 0
2023/05/02 20:02:53 42 in for/if loop
2023/05/02 20:02:53 if/then block: testing value 1337 loop increment 1
2023/05/02 20:02:53 if/then block: testing value 31337 loop increment 2
2023/05/02 20:02:53 Hit that 31337
2023/05/02 20:02:53 printing out updated map now
2023/05/02 20:02:53 string: foo int: 42
2023/05/02 20:02:53 string: bar int: 31337
```

Next up is the `contents.go` file, which is just playing around with files.

```
 go run contents.go
This will look at the contents of /etc/resolv.conf and print the contents to stdout
# This is /run/systemd/resolve/stub-resolv.conf managed by man:systemd-resolved(8).
# Do not edit.
#
# This file might be symlinked as /etc/resolv.conf. If you're looking at
# /etc/resolv.conf and seeing this text, you have followed the symlink.
#
# This is a dynamic resolv.conf file for connecting local clients to the
# internal DNS stub resolver of systemd-resolved. This file lists all
# configured search domains.
#
# Run "resolvectl status" to see details about the uplink DNS servers
# currently in use.
#
# Third party programs should typically not access this file directly, but only
# through the symlink at /etc/resolv.conf. To manage man:resolv.conf(5) in a
# different way, replace this symlink by a static file or a different symlink.
#
# See man:systemd-resolved.service(8) for details about the supported modes of
# operation for /etc/resolv.conf.

nameserver 127.0.0.53
```
