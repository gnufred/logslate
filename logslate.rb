class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.7.10'

  on_linux do
    sha256 '1b3ef14fd8b5c533c4597f164c0cbd11c2a988a50a788067e4670808ec8eb632'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.10/logslate-linux-amd64.xz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'b0285e8f0d2b3f3ec64b8a621810c01430e16ad5447e2fbd2463aa441534a57e'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.10/logslate-darwin-amd64.xz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
