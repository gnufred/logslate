class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.7.0a'

  on_linux do
    sha256 '8508080acfe208fb40dddd997c78027568f1fd5a70aa61acf33529b71c7df4d9'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.0a/logslate-linux-amd64.tar.gz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '434e04123487086503b9ab71b5fd1454f12b11782df25e2195ffecd95c4414c2'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.0a/logslate-darwin-amd64.tar.gz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
