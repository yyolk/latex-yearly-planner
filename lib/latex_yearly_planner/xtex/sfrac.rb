# frozen_string_literal: true

module LatexYearlyPlanner
  module XTeX
    class Sfrac
      attr_reader :num, :denom, :slash_symbol

      def initialize(num, denom)
        @num = num
        @denom = denom
        @slash_symbol = '\\textfractionsolidus'
      end

      def to_s
        "\\sfrac{#{num}}{#{denom}}"
      end
    end
  end
end
