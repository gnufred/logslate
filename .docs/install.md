# Installation

If you want to use `logslate` on your local workstation, it's recommended to use the `brew` package manager.

```
# install
brew tap gnufred/logslate https://gitlab.com/gnufred/logslate
brew install logslate
# update
brew update && brew upgrade logslate
```

If you're using `logslate` on a server, or need to install it in a built container, it's recommended to use `curl`.

```
curl --location --output ./logslate https://gitlab.com/api/v4/projects/gnufred%2Flogslate/packages/generic/logslate/latest/logslate-linux-amd64
```

For **Windows** and other uses cases, head to the [release page](https://gitlab.com/gnufred/logslate/-/releases) and download the binary manually.
