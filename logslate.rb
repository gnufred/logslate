class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.3.2-w1'

  on_linux do
    sha256 '09fe64f061daba79efc3910ead51c7906ff91bb4eb1838816ebaf437c78ad068'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.2-w1/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'f36ea65e2fcf03576071a91090601d9eaaa57df985f297ba4375c31ee51d196b'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.2-w1/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
