class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.5.9-w6'

  on_linux do
    sha256 '2378a0e605dc46e357db1d137259b8004bad966231b3a5058d754e8e74fdb2c6'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.9-w6/logslate-linux-amd64.tar.gz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '9e3879991a6b920a1f63846f2f512ed1bbf62d1fefeb2075e5f6f7c427ebee12'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.9-w6/logslate-darwin-amd64.tar.gz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
