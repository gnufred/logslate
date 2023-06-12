class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.1.1'

  on_linux do
    sha256 'b37a5318f0a671c970717d63598aff931bf20cc94f099c1c0d6cbbafe434fb44'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.1.1/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'aad1673522e2acdad4f474255971cd99359cb8cbd24dcd87c18ff721cb33a1f8'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.1.1/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
