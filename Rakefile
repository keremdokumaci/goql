task :command_exists, [:command] do |_, args|
    abort "#{args.command} doesn't exists" if `command -v #{args.command} > /dev/null 2>&1 && echo $?`.chomp.empty?
  end
  
  task :has_bumpversion do
    Rake::Task['command_exists'].invoke('bumpversion')
  end
  
  task :has_golangci_lint do
    Rake::Task['command_exists'].invoke('golangci-lint')
  end
  
  task :has_mockery do
    Rake::Task['command_exists'].invoke('mockery')
  end
  
  AVAILABLE_REVISIONS = %w[major minor patch].freeze
  desc "bump version, default is: patch"
  task :bump, [:revision] => [:has_bumpversion] do |_, args|
    args.with_defaults(revision: 'patch')
    unless AVAILABLE_REVISIONS.include?(args.revision)
      abort "Please provide valid revision: #{AVAILABLE_REVISIONS.join(',')}"
    end
  
    system "bumpversion #{args.revision}"
  end
  
  desc "run tests"
  task :test do
    system %{
      color_red=$'\e[0;31m'
      color_yellow=$'\e[0;33m'
      color_white=$'\e[0;37m'
      color_off=$'\e[0m'
      any_errors="0"
      for s in $(go list ./...); do
        if ! go test -failfast -p 1 -v -race "${s}"; then
          echo "\t\n${color_red}${s}${color_off} ${color_yellow}fails${color_off}...\n"
          any_errors="1"
          break
        fi
      done
      if [[ "${any_errors}" == "0" ]]; then
        echo "\n\n${color_white}Tests are passing...${color_off}\n"
      fi
    }
  end
  
  namespace :mock do
    namespace :mockery do
      desc "upgrade mockery version"
      task :upgrade => [:has_mockery] do
        system %{
          brew update && 
          brew upgrade mockery && 
          brew cleanup mockery
        }
      end
      
      desc "show installed version"
      task :show_version => [:has_mockery] do
        system "mockery --version --log-level error"
      end
    end
  
    namespace :generate do
      desc "generate mock for all"
      task :all => [:has_mockery] do
        rm_rf %w(mocks)
        system %{
          mockery --output ./mocks --case=underscore --dir ./internal/ --all --keeptree --recursive
        }
      end
    end
  end
  
  desc "lint"
  task :lint => [:has_golangci_lint] do
    system "LOG_LEVEL=error golangci-lint run"
  end