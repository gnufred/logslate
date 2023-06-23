class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.2.6'

  on_linux do
    sha256 '3c25877ccba7bfe2c109bf9f4cda0e8998d0c1803362cc14b07ad4e7427e4e96'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.6/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '60f0b04907c05a964a1be5d5b631cb665e690cfb28446876fd9aef99676e77e7'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.6/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
