class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.3.7'

  on_linux do
    sha256 'f1b1eb646b0cdeb78a6bd824d5fa377293cb94325c599ddb173d481926eeaa2e'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.7/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '2a192813dd8e13b65a4a3ede7d2497af0d31aca496beda981d0637576c0c4b91'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.7/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
