class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.1.0'

  on_linux do
    sha256 '073337c2686ec45065d9b56f2b92b0d9aa655d57ad51fd5d0db95de1baa09980'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.1.0/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'aa3414404ab3a11a645887dfd34c950cd5ddfb05b9dca7e9883df186d4925bce'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.1.0/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
