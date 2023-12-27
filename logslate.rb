class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.4.8'

  on_linux do
    sha256 '693113ac2ed6c3312e6a06b6c08f3a28c415c00a62ce8d401c60d172467ee0c7'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.8/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '9c679c01b3abe217f2d2d5dadbef8bb9ac69738d60c754109179675a60152b8a'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.8/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
