class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.5.3'

  on_linux do
    sha256 '76b886c434255365ec5535b8aa3cbaac838a27246eb12ae2f5d0f077925ae57f'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.3/logslate-linux-amd64.tar.gz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '7236bbb77f9fd2a1a08346b209e55cb441bfc7da03ac24ce2265a3c4a961c632'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.3/logslate-darwin-amd64.tar.gz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
