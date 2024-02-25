# frozen_string_literal: true

module LatexYearlyPlanner
  module TeX
    class TextSize
      attr_reader :text

      def initialize(text)
        @text = text
      end

      def xhuge
        self.size = 'Huge'

        self
      end

      def huge
        self.size = 'huge'

        self
      end

      def xxlarge
        self.size = 'LARGE'

        self
      end

      def xlarge
        self.size = 'Large'

        self
      end

      def large
        self.size = 'large'

        self
      end

      def normal
        self.size = 'normalsize'

        self
      end

      def small
        self.size = 'small'

        self
      end

      def footnote
        self.size = 'footnotesize'

        self
      end

      def script
        self.size = 'scriptsize'

        self
      end

      def tiny
        self.size = 'tiny'

        self
      end

      def to_s
        return text unless size

        "{\\#{size}{}#{text}}"
      end

      private

      attr_accessor :size
    end
  end
end
