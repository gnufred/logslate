class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.5.1'

  on_linux do
    sha256 '5bc94d441c551e1cf77dacf0a5403c1629a601974d6d25fc75cbb374129d675d'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.1/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'c52d25de70a036e0fb50d349c20c868cbe00a8fafb0fd4fd719162765d1609a0'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.1/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
