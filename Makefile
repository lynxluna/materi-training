.PHONY: all clean

DISTDIR := dist

SRCFILES := $(wildcard *.adoc)
OBJFILES := $(patsubst %.adoc,$(DISTDIR)/%.html,$(SRCFILES))

$(DISTDIR)/%.html: %.adoc
	bundle exec asciidoctor -r asciidoctor-multipage -b multipage_html5 -D dist $<
	#bundle exec asciidoctor -o $@ $<

all: $(OBJFILES)

clean:
	rm -rf $(DISTDIR)/*.html

