# frozen_string_literal: true

require 'bundler/gem_tasks'
require 'rspec/core/rake_task'

RSpec::Core::RakeTask.new(:spec)

require 'rubocop/rake_task'

RuboCop::RakeTask.new

task default: %i[spec rubocop]

require 'rake/clean'

CLEAN.include(['out'])

require 'latex_yearly_planner'
require 'latex_yearly_planner/app'

desc 'generate kindle'
task :kindle do
  args = ['generate', './config/kindlescribe_left-hand_dotted.yaml']
  LatexYearlyPlanner::App.start(args)
end

desc 'testing'
task :doesthiswork do
  puts 'yo is this thing on?'
end
