class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.0.16'

  on_linux do
    sha256 'a6bdeef040c50f5b5131e74720db3dc72210df7ec3d800bd9561ec287555c1db'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.16/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '3e4cea07a35c2f67cd1c8cc9caffbba6a397b530a0b8c7703aa43ba7eaaa8f83'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.16/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
