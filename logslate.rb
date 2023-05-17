class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.0.15'

  on_linux do
    sha256 '88fd178389f1c17e52665ca2af166b9aed580d14811903332a3855228f6b7f9d'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.15/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '585c0df33041247b08dbc780b38c9ed4e0a411bd5a66eba50fe4183619f4cc2b'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.15/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
