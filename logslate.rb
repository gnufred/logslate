class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.4.1'

  on_linux do
    sha256 '5ef267b39b8cbe8c2fec691ae4e746a6257b045a563f9cc6ac22d49cc26525f7'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.1/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '36d60f15c3c5fde5cc99276efd462fc2530793ba29c3cf40848bb28b11e15e75'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.1/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
