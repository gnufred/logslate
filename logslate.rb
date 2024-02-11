class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.6.3'

  on_linux do
    sha256 'e659d49a2a3bcb19f452b32addb10b8dffd6b173b045f10567b193f7e3f7b2d2'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.6.3/logslate-linux-amd64.tar.gz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '729490eb0eab6d50f16241514d6b0a6e205c9dc9978fc98d9576e5135452da58'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.6.3/logslate-darwin-amd64.tar.gz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
