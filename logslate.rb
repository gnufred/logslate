class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.0.18'

  on_linux do
    sha256 '10abcc5a8b7a87308e44c444f327cc321f5a7ad60c3afc9c029af9a168c62854'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.18/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '73661839f183b05ced31dcdac6a85f1377d7ce0f01bfc5412bbe2d04f11da069'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.18/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
