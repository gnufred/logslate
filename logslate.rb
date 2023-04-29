class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.0.10'

  on_linux do
    sha256 'b1c4bbc246fab0ca7f9af2ab96ec688176fad748e20ee4a9e0094c733164fd5f'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.10/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '59a11cd137c5edc026ae6ffd538633bf26ea5d7cae84f3af521e346a38c20bb3'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.10/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
