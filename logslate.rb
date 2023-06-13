class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.1.2'

  on_linux do
    sha256 '5145d6ea83997ffe837d02fa64b5d3a0e82b2b72e5a50e4d258e122d0ccd16b1'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.1.2/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '7e4a5b2a82c9505cab31a0991acaca789640080ed7d10ed6594309bc91240ead'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.1.2/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
