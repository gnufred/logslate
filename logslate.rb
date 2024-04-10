class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.7.8-r2'

  on_linux do
    sha256 'b09a08176715190ca4e63226363e7b8cbe310b579df9c33b89e3aad9a1967cc3'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.8-r2/logslate-linux-amd64.xz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'a8614f7d630d3189efccb5a0bb9a327493f0018470d6fba97ce089fe8563319b'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.8-r2/logslate-darwin-amd64.xz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
