class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.6.3-t2'

  on_linux do
    sha256 'deb83ea5444254296026ea8f068752c4aac2303b3d625caee299eecf5b468496'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.6.3-t2/logslate-linux-amd64.tar.gz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '72f8ac9cd3f978fd3cdb9fec3f3ea0655dd8f2611868d6e2f1670a37f948533f'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.6.3-t2/logslate-darwin-amd64.tar.gz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
