# frozen_string_literal: true

module LatexYearlyPlanner
  module Pieces
    module Mos
      module Components
        class MonthlyBody < Component
          def generate(month)
            XTeX::CalendarLarge.new(month, **config.large_calendar(section_name)).to_s
          end
        end
      end
    end
  end
end
