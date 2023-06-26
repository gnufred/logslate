class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.2.7'

  on_linux do
    sha256 'fe6e0674f3c7448b0d23bc06fc6fea56c24a53361e7eb539361a145998745a72'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.7/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'aa2aac8320f23c04b27ad456babc33579cae67a46d652c51301f2cd523f05c95'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.7/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
