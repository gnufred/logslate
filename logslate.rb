class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.7.2'

  on_linux do
    sha256 'b5feb2ad74097e81cf1de1b099a356dc7091cf0f69730081e30a6d662b5770a6'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.2/logslate-linux-amd64.xz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'e66a50d1c0c8327156b1e7a16ffcb7939570919badf08b3ae577af0f8c401fb3'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.2/logslate-darwin-amd64.xz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
