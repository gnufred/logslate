class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.3.8-w3'

  on_linux do
    sha256 'cbfb9e09b3913c93ffa862f55637277a6f926f9e9d525e3749f134e2818024cd'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.8-w3/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '03f17fc9560cd18a63d9a36606d7cc82a5a19c3d480ba1b77a3e0d1a72b9f84d'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.8-w3/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
