# kvatyr [:ukraine:](https://savelife.in.ua/en/donate-en/)

Primitive file sharing tool that also answers the eternal question "How to share files from Linux to Mac?"

[![Go Reference](https://pkg.go.dev/badge/github.com/radikh/kvatyr.svg)](https://pkg.go.dev/github.com/radikh/kvatyr)

Pros:
 - Easy install, easy use, easy remove
 - No configs
 - No garbadge left in the system
 - Primitive as dirt
 
Cons:
 - Security level - N/A
 - Your granny probably will not manage to use it
 - Primitive as dirt

### How to use 

#### Prerequisites

- Browser installed on the target machine
- Not being afraid of terminal
- Optionally a lot of anger on the fact that there is not easy way to send files directly to Mac iOS machines :japanese_goblin:

Check the latest release and find an executable for you.

#### For Mac and Linux
Use terminal.

Give the executable execution permissions:

```
chmod u+x kvatyr-Linux-x86`
```

Then you can run kvatyr with the default params. By default kvatyr runs on port 8080 and shares the directory it's called from.
```
./kvatyr-Linux-x86
```

Or you can specify optional port and folder to share.

```
./kvatyr-Linux-x86 --port=8090 /home/user/folder_to_share
```

When app is started you will see a message with a link to access kvatyr

#### Fow windows

I had no chance to test, so feel free to explore and update this section. Open source innit? :trollface:

# IF YOU DONT READ THIS NOTE - I'LL FIND YOU!!!!
Security is important - so please read this text carefully.

The app allows to connect without authentication. It means that when you're on Wi-Fi everyone can get the access to the shared files, same for local network.

The connection is not encrypted and not private, it can be scanned, identified, hijacked, intercepted etc. 

In short you data is not protected during the transfer. Completely.

### NAQ:

Q: Why to use this instead of SMB or FTP?

A: If you wish - do. I just needed something quick and simple for my home network without configuring something system invasive.


Q: How to use it if it's so insecure?

A: I will not give recommendations but I have certain preconditions to run it. Run only on the network I control (normally home network or phone hotspot). The network should not have my hackerman neighbor to sniff your traffic. Run the tool not longer than needed. Don't leave the tool working on the background. Prepare a separate directory with the files share and run the tool only on that directory. And still it's not secure, you just need to accept some grade of risk.


Q: Will the tool be improved?

A: Who knows, who knows...


Q: What's NAQ?

A: Never asked questions, obviously :) 
