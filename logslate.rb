class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.3.5'

  on_linux do
    sha256 '030258adf9aee0176543216223e9ec76697a627ce17111d2a033176ad2952a23'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.5/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '7a0f797fd94ad66f105af57fb1215fd5de5e22f2997127bed387beb0c7d78fbe'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.5/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
