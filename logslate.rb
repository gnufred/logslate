class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.3.8'

  on_linux do
    sha256 '0d8e8721dbc717f707712da85e91617b8c4cc8797efa432e3a669d051ed0441f'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.8/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'e9c10583b9fbdd12603a663fb393ca9af6f9e400ffc29686c6f8239b723977db'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.8/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
