class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.2.1'

  on_linux do
    sha256 '6673e9b5e5259e3d1d12f410c4239ff16f494df50cb6a5eb4e739eedd5b21941'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.1/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '2e99ad6ab17d662ab2aebb45aa462ef67722cd034562332cedc9ceb0eccb948a'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.1/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
