class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.1.0a'

  on_linux do
    sha256 'a28aaedf0d9db303c6562b820a2bfd07a9eb05ed5b89d1365017839b91372ff7'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.1.0a/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '59cda68c354e46c652de938424ef6fafdb9e0c790420c6a84a9b795c5e5da8e8'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.1.0a/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
