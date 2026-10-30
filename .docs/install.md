# Installation

## Local

If you want to use `logslate` on your local workstation, it's recommended to use the `brew` package manager.

```
# install
brew tap gnufred/logslate https://github.com/gnufred/logslate
brew install logslate
# update
brew update && brew upgrade logslate
```

For other uses cases, head to the [release page](https://github.com/gnufred/logslate/releases) and download the binary manually.

## Remote

If you're using `logslate` on a server, or need to install it in a built container, it's recommended to use `curl`.

You can use something like:

```
#!/bin/bash
ARCH=linux-amd64
curl --location "https://github.com/gnufred/logslate/releases/latest/download/logslate-$ARCH.xz" --output "logslate-$ARCH.xz" && \
tar -xJf "logslate-$ARCH.xz" && \
rm "logslate-$ARCH.xz" && \
chmod +x "logslate-$ARCH" && \
mv "logslate-$ARCH" logslate
```

