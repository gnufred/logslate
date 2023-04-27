class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '18-latest'

  on_linux do
    sha256 '470a7bcd07b0004ec0d4d07039c9ca2a9bbeeb9a73336cf559377ab26c900272'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/18-latest/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '426a4aba35dbd7b5db63798938d9259091c18149d7f97a8148478c117fd91a68'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/18-latest/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
