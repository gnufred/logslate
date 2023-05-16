class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.0.14'

  on_linux do
    sha256 'f3a20de1dfe7cd79344f3d3b68b395133bdf49c18ab6032990e98a0da582c994'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.14/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '5d787b4456dd7819ae2968d7bc79aa83ede39d3be9048f728505a107eca771bf'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.14/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
