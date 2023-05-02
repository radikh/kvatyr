# kvatyr [:ukraine:](https://savelife.in.ua/en/donate-en/)
Primitive file sharing tool that also answers the eternal question "How to share files from Linux to Mac?"

Pros:
 - Easy install, easy remove
 - No configs
 - No garbadge left in the system
 - Primitive as dirt
 
Cons:
 - Security level - N/A
 - Your granny probably will not manage to use it
 - Primitive as dirt

### How to use 
As of now no binaries are distributed, so you need to have go installed :trollface:

When done run this

`$ go install github.com/radikh/kvatyr@latest`

run default to serve current directory and port 8080

`$ kvatyr`

or specify a port and a directory

`$ kvatyr --port=8090 /home/user/folder_to_share`


# IF YOU DONT READ THIS NOTE - I'LL FIND YOU!!!!
Security is important - so please read this text carefully.

The app allows to connect without authentication. It means that when you're on Wi-Fi everyone can get the access to the shared files, same for local network.

The connection is not encrypted and not private, it can be scanned, identified, hijacked, intercepted etc. 

In short you data is not protected during the transfer. Completely.

### NAQ:

Q: Why to use this instead of SMB or FTP?

A: If you wish - do. I just needed something quick and simple for my home network without configuring something system invasive.


Q: How to use it if it's so insecure?

A: I will not give recommendations but I have certain preconditions to run it. Run only on the network I control (normally home network or phone hotspot). The network should not have my hackerman neighbour to sniff your traffic. Run the tool not longer than needed. Don't leave the tool working on the background. Prepare a separate directory with the files share and run the tool only on that directory. And still it's not secure, you just need to accept some grade of risk.


Q: Will the tool be improved?

A: Who knows, who knows...


Q: What's NAQ?

A: Never asked questions, obviously :) 
