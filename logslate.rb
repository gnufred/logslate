class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.7.4'

  on_linux do
    sha256 '0baa8358af2721be19d753c781a3e4d9774e21dbd1b3644cde0ce2deb9beb640'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.4/logslate-linux-amd64.xz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'cf4abf7cd0864e6b14c9238b1a27a9892ce45cb2155fc12550ce95cbe11e95eb'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.4/logslate-darwin-amd64.xz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
