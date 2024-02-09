class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.6.1'

  on_linux do
    sha256 'fc1666f7ff43cbd244d47ebbc37a437a6aef77fd1f9adf6910c6e21ef370bd87'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.6.1/logslate-linux-amd64.tar.gz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '273ebb80a5f1b96614bb7d46b25362b09bb079c9ba5e5596f5e1e2a80ec2fb40'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.6.1/logslate-darwin-amd64.tar.gz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
