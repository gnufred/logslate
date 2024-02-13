class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.6.3.t1'

  on_linux do
    sha256 '03c1a325f4a10ddf9dd8aeb3b62fc66500324ec436332f8c7cec7bd911a04d08'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.6.3.t1/logslate-linux-amd64.tar.gz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '7b7005e1f5f33911c064967f674a73f8db9da1eb59445755a6da50859a815947'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.6.3.t1/logslate-darwin-amd64.tar.gz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
