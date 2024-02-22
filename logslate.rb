class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.6.3-t4'

  on_linux do
    sha256 '81f934043fc8203837bc3395ffb50be5235cd5f569f93ab07239048486ee08b0'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.6.3-t4/logslate-linux-amd64.tar.gz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '4037502b3f5087ff26fc2740417419b6b6ccd7d03a6b2314111419b7ba5de4fb'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.6.3-t4/logslate-darwin-amd64.tar.gz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
