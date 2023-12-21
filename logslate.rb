class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.4.4'

  on_linux do
    sha256 '99cfd2cf63001f4ca53d24d9115c3553f3699a90093574ca65a28cc9fbb6d8d1'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.4/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '00bb7342f0eb67034865f2c196d9607adc5da3bb407f502b9724020f1f2bdf1a'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.4/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
