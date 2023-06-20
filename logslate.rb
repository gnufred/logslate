class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.2.3'

  on_linux do
    sha256 '9ced70653a5876823d3d1f6525fe3b6d592c19193089d67f3c558a39866b36bc'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.3/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '8e75bc3439b2cd8517b423f8d268f4bec573722410bd137ac1e725f211f61007'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.3/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
