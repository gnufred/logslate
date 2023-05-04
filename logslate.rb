class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.0.11'

  on_linux do
    sha256 '0abde51328464a6d9e967a7e236784742f87e340880cf8ed71f964dcecc5af35'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.11/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '41b9e7cdf6101a8f3153568a73dfc1aa5f787978efcbdc39fc30a1c173296698'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.11/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
