require 'sprockets'
project_root = File.expand_path(File.dirname(__FILE__))
assets = Sprockets::Environment.new(project_root) do |env|
  env.logger = Logger.new(STDOUT)
end

assets.append_path(File.join(project_root, 'assets'))
assets.append_path(File.join(project_root, 'assets', 'javascripts'))
assets.append_path(File.join(project_root, 'assets', 'stylesheets'))

manifest = Sprockets::Manifest.new(assets, "public/manifest.json")

filenames = [
  File.join(project_root, 'assets', 'javascripts', 'application.js'),
  File.join(project_root, 'assets', 'stylesheets', 'application.css'),
]

map "/assets" do
  run assets
end

map "/" do
  run lambda { |env|
    [
      404,
      {
        'Content-Type'  => 'text/html',
        'Cache-Control' => 'public, max-age=86400'
      },
      "File not found"
    ]
  }
end

map '/manifest.json' do
  run lambda { |env|
    manifest.compile(filenames)
    [
      404,
      {
        'Content-Type'  => 'text/json'
      },
      File.open("public/manifest.json", File::RDONLY)
    ]
  }
end
