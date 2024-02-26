# frozen_string_literal: true

module LatexYearlyPlanner
  module Pieces
    module Mos
      module Components
        class IndexNotesHeader < Header
          def generate_index(page)
            make_header(top_table(page:), index_title(page))
          end

          def generate_notes(note, page, total_pages)
            make_header(top_table(page:, note:), notes_title(note, page, total_pages))
          end

          private

          def index_title(page)
            target_reference(TeX::TextSize.new('Index').huge, reference: NOTES_INDEX_REFERENCE, page:)
          end

          def notes_title(note, page, total_pages)
            content = TeX::TextSize.new("Note #{note}").huge
            content = [content, XTeX::Sfrac.new(page, total_pages), '\\usym{1F5CA}'].join(' ') unless total_pages == 1
            target_note(content, note:)
          end
        end
      end
    end
  end
end
