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
  
  task :has_gsed do
    Rake::Task['command_exists'].invoke('gsed')
  end

  task :has_migrate do
    Rake::Task['command_exists'].invoke('migrate')
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
  task :test => [:has_gsed] do
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
        echo "\n\n${color_white}Tests are passing...${color_off}"
        echo "${color_white}Calculating code coverage${color_off}"
        go test ./... -coverpkg=./internal/... -coverprofile ./coverage.out > /dev/null 2>&1
        code_coverage_ratio=$(go tool cover -func ./coverage.out | grep "total:" | awk '{print $3}')
        echo "${color_white}Total test coverage: ${color_yellow}${code_coverage_ratio}${color_off}"
        code_coverage_ratio_md=${code_coverage_ratio/%/25}
        gsed -i -r "s/coverage-[0-9\.\%]+/coverage-${code_coverage_ratio_md}/" README.md &&
        echo "README updated...\n"
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
  
  desc "run migration (up,down) default: up"
  task :migrate, [:direction] => [:has_migrate] do |_, args|
    args.with_defaults(direction: 'up')
    puts "-> will migrate to: \e[33m#{args.direction}\e[0m"
    system %{
      source .env && 
      migrate -verbose -database "${DATABASE_URL}" -path migrations #{args.direction}
    }
  end
  
  desc "create new migration"
  task :update, [:name] => [:has_migrate] do |_, args|
    abort "please enter the name of migration" if args.name == nil

    system "migrate create -ext sql -dir ./internal/repository/migrations/postgres -seq #{args.name}"
    system "migrate create -ext sql -dir ./internal/repository/migrations/mysql -seq #{args.name}"
    system "migrate create -ext sql -dir ./internal/repository/migrations/mssql -seq #{args.name}"
  end