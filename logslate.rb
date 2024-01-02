class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.5.8'

  on_linux do
    sha256 'f3ebca0d64ad29d3931169f3db383a8513e6224c9afcc481fd3c92eefbc0f650'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.8/logslate-linux-amd64.tar.gz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '0a469f185e9b1f1457e248d125b0a8732c2a9904a1bf82011d5ee6f20b96758c'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.8/logslate-darwin-amd64.tar.gz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
