class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.4.3'

  on_linux do
    sha256 '982cddb977c74793cbb848635346f46d70b64aed4baea5b8be8c54571deb0f24'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.3/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '5a2e76c9cb38b1a5af647a774e3e9e89ebc6865a8f95bb0588932f0f94c1be6c'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.3/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
