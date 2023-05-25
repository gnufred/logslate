class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.0.19'

  on_linux do
    sha256 '7f0700380f813429199934226d6f0448aa3172432708598fa61265102ad46f81'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.19/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '3ba88f345e744a70cb47b41a3402a109ddfd46d57eb4d754e8827200dc16b41a'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.19/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
