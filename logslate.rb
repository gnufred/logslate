class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.7.1'

  on_linux do
    sha256 '29810c0e364fc973fecd0fc5fa9e93eabe772da447c891689c5cd4b6b711d134'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.1/logslate-linux-amd64.tar.gz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '01039b6b83ab284377f84d8eb9d7f31a480b2ae622179978796f54b238d6556a'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.1/logslate-darwin-amd64.tar.gz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
