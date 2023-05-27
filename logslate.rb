class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.0.20'

  on_linux do
    sha256 'e851a5b2f813dc251b4e8d8a67788829bb3de6c81e66ac66d613e29bc787a8b3'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.20/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'de5aef235225b76c86cf7a0012c2efb5f7538aa78016643a5acf9367f605de0c'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.20/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
