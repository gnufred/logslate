class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.4.6'

  on_linux do
    sha256 'ad61c277efd4dd5f7d8535226f5428e442fef0c60cfb640e2dd54e729775f97d'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.6/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '0e3d56b328084b2980ef85c3e8c88573571cb9e0c25bd2342c1b07ef99a881ef'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.6/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
