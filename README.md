# Chuck Norris jokes Bot

This is a simple p3p bot that sends Chuck Norris jokes.

# Setup

You can download I2Pd from:
 - https://i2pd.readthedocs.io/en/latest/user-guide/install/
 - https://git.mrcyjanek.net/p3pch4t/flutter_i2p_bins-prebuild/

## Configuration

### I2pd

```conf
$ cat /etc/i2pd/i2pd.conf
tunnelsdir = /var/lib/i2pd/tunnels.d

# [...]
[httpproxy]
enabled = true
address = 127.0.0.1
port = 4567 # Different port to make sure that it doesn't conflict with apps running on default
```

```conf
$ cat /etc/i2pd/tunnels.conf.d/p3pchuck.conf
[p3pchuck]
type = http
host = 127.0.0.1
port = 3895
inport = 3895
gzip = false
signaturetype = 7
enableuniquelocal = true
keys = p3pchuck.dat
ssl = false
```

### .env


First figure your local i2p address (`PRIVATEINFO_ROOT_ENDPOINT` in .env)

```bash
$ printf "i2p://%s.b32.i2p\n" $(sudo head -c 391 /var/lib/i2pd/p3pchuck.dat | sha256sum | xxd -r -p | base32 | sed s/=//g | tr A-Z a-z)
```

Then fill .env file

```
PRIVATEINFO_ROOT_ENDPOINT=i2p://[your address].b32.i2p
I2P_HTTP_PROXY=http://127.0.0.1:4567
LOCAL_SERVER_PORT=3895
```


## Hosted instance

There is a hosted instance at `i2p://rtrnlpzzexbufkhpx3bjwxuim4lp22iz5oso2egy7hmwuz2yvugq.b32.i2p`
