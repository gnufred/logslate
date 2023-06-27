class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.2.8'

  on_linux do
    sha256 '00e72bcdcccf329cf48fdc53ffd65f111557395127f1deab811fc51dc34b56ea'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.8/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'a89b08374bb52c3fd5230721fc25cf1df3b962c13285cc947175affb91fc2d71'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.8/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
