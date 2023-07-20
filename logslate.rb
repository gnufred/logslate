class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.3.6'

  on_linux do
    sha256 'f7979d64a587636b9c4b8b52f4f092f68640a70e5a4f150bc3750dc84be27b9a'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.6/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '8f6a2faf181a34908b998d286362f17ce41c447c611a0333531379b249aedb98'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.6/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
