class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.0.9'

  on_linux do
    sha256 'f678d771564b496511c38ba68f47c44cc2124195de064287037d58e30d71cf50'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.9/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '5ffa6633a4d0e189ae26d7f3a6969114f33f38948139df2e3727ed2bb3e01bfa'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.9/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
