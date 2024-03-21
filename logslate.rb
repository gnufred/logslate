class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.7.8-r1'

  on_linux do
    sha256 'b72b8339caf6c709be96e6c5aaff75712199786552b8ca575022a277d92571a9'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.8-r1/logslate-linux-amd64.xz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '658b8269b8e44b9ce153117dd9816cae7f7eb53d390a571f03a04c16cafffb54'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.8-r1/logslate-darwin-amd64.xz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
