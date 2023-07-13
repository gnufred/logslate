class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.3.4'

  on_linux do
    sha256 '0d7fac4a3ca037285c0f2df4c212f89b172b7c5a0159ee75230c00c8b3cf8189'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.4/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'b0a8e492080c3788b999d2c844f97979a2bcd65f57e429f19ab37f3da0ff2052'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.4/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
