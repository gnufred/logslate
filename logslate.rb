class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.4.7'

  on_linux do
    sha256 '54eae81eefda1304997fb79f80de5da82f1f72a286616e4077a687daee7d3394'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.7/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '40d3a9ecfb474279ec369d3a4edfcb5eabcdc45aaea70c4fee382937e99930da'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.7/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
