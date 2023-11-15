class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.3.8-w4'

  on_linux do
    sha256 '33493136b554818c9b3adc0ffc89fe0bd0ada326ae6030b9a861cd848ab70378'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.8-w4/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '7d6127cedd30a2af656137b3c81c96067e2c5eaef6be337cb06bcc71c2c781ed'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.8-w4/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
