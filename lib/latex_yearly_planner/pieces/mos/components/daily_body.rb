# frozen_string_literal: true

module LatexYearlyPlanner
  module Pieces
    module Mos
      module Components
        class DailyBody < Component
          def generate(day)
            schedule_column(day)
          end

          private

          def schedule_column(day)
            TeX::Minipage.new(
              content: "#{schedule}\n#{little_calendar(day)}",
              width: param(:schedule_column_width)
            )
          end

          def schedule
            XTeX::Schedule.new(**schedule_options)
          end

          def little_calendar(day)
            XTeX::CalendarLittle.new(
              day.month,
              **config.little_calendar(section_name).merge({ highlight_day: day })
            )
          end

          def schedule_options
            { compensate_height: config.document.document_class.size }
              .compact
              .merge(param(:schedule).parameters_as_a_hash&.compact)
          end
        end
      end
    end
  end
end
