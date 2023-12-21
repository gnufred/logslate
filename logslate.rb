class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.4.2'

  on_linux do
    sha256 'ba1dbe55ba65eeb795461add81fd88e50639c3b3af0933f6f813e762f3893372'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.2/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'faecd53a33bc6bb7b9892e3da567b2d1ebcab1fbfc5034b82ce51779331c6460'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.2/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
