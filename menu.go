package gopher

import (
	"fmt"
	"io"
)

type Menu struct {
	Links []Link
}

func (m Menu) Encode(out io.Writer) error {
	for _, link := range m.Links {
		if _, err := out.Write([]byte(link.String())); err != nil {
			return err
		}
		if _, err := out.Write([]byte{'\r', '\n'}); err != nil {
			return err
		}
	}
	return nil
}

type LinkType byte

const (
	// Canonical Types (RFC 1436)
	TextFile       LinkType = '0'
	Submenu        LinkType = '1'
	Nameserver     LinkType = '2'
	Error          LinkType = '3'
	BinHexFile     LinkType = '4'
	DOSFile        LinkType = '5'
	UUEncodedFile  LinkType = '6'
	FullTextSearch LinkType = '7'
	Telnet         LinkType = '8'
	BinaryFile     LinkType = '9'
	Mirror         LinkType = '+'
	GIFFile        LinkType = 'g'
	ImageFile      LinkType = 'I'
	Telnet3270     LinkType = 'T'
	// Non-Canonical Types
	HTMLFile             LinkType = 'h'
	InformationalMessage LinkType = 'i'
	SoundFile            LinkType = 's'
)

type Link struct {
	Type     LinkType
	Display  string
	Selector string
	Hostname string
	Port     int
}

func (l Link) String() string {
	return fmt.Sprintf("%s%s\t%s\t%s\t%d", string(l.Type), l.Display, l.Selector, l.Hostname, l.Port)
}
