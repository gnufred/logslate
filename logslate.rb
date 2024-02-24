class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.7.0'

  on_linux do
    sha256 '5317240270fc8391a736d09c908030c9161fc2d37d3a4347a311e2d6c8c66c2e'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.0/logslate-linux-amd64.tar.gz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'ab6fd5a14efbefc7ffd0b94383308009f66df57e3ad3f070bbd74901f1e7ca84'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.0/logslate-darwin-amd64.tar.gz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
