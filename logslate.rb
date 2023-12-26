class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.4.7-w6.0'

  on_linux do
    sha256 'c51afd4c38a11e3b5cb2017f94b3b7aba3c0bd88d6d9d35a32b708a0b9bcecba'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.7-w6.0/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '5d92e75c3fe8eb0765c5f70b677b9f4abba41520741a3484e7e4c1c309f700a4'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.7-w6.0/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
