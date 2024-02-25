# frozen_string_literal: true

module LatexYearlyPlanner
  module Pieces
    module Mos
      module Components
        class DailyNotesHeader < Header
          def generate(page, day, total_pages)
            make_header(
              top_table(day:),
              total_pages == 1 ? title(day) : title_with_pages_count(day, page, total_pages),
              highlight_quarters: [day.quarter],
              highlight_months: [day.month]
            )
          end

          def title(day)
            link_day(target_daily_notes(TeX::TextSize.new(day.name).huge, day:), day:)
          end

          def title_with_pages_count(day, page, total_pages)
            [title(day), XTeX::Sfrac.new(page, total_pages), '\\usym{1F5CA}'].join(' ')
          end
        end
      end
    end
  end
end
