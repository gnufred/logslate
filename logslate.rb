class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.3.2'

  on_linux do
    sha256 '1773131468ad9ee9a227fb5941e2b151c2f31b5dc455aae9cdcfc56a637963cb'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.2/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '034efaf4f5b20c8b0e9dfc51d4a8643d4a4ea448a4f960efcdbc11c77c3e3c52'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.2/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
