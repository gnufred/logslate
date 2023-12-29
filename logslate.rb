class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.5.6'

  on_linux do
    sha256 '40255adc8779ffa85a1bdd45e7b5d8398693c24ff9568881111c8198c759b8dd'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.6/logslate-linux-amd64.tar.gz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '9f121c5b63384c50e5bf56bb08df16ac842935aa0732232f34ae9ff892029658'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.6/logslate-darwin-amd64.tar.gz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
