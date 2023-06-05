class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.0.21'

  on_linux do
    sha256 '5b50b09d561c7d9d244ce12fc9c4d5e56a28e1d00b71c1ce3aff2147eedc1ca8'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.21/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '925e548f488c82c5ea1a3193522ff0a6402044fd1c6e42f7577a3fd0131d14b1'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.21/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
