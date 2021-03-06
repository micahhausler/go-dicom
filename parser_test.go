package dicom

import (
	"os"
	"testing"
)

func TestConstructor(t *testing.T) {
	p, err := NewParser()
	if err != nil {
		t.Error(err)
	}

	if len(p.dictionary) == 0 {
		t.Error("Error constructing parser. Dictionary can not be of 0 length")
	}

}

func TestDictionaryOption(t *testing.T) {

	fh, err := os.Open("dicom.dic")
	defer fh.Close()
	if err != nil {
		t.Error(err)
	}

	p, err := NewParser(Dictionary(fh))
	if err != nil {
		t.Error(err)
	}

	if len(p.dictionary) == 0 {
		t.Error("Error constructing parser. Dictionary can not be of 0 length")
	}
}

func TestGetTag(t *testing.T) {
	elem := &DicomElement{0x7FE0, 0x0010, "PixelData", "ox", 1, nil}

	if tag := elem.getTag(); tag != "(7FE0,0010)" {
		t.Errorf("Error creating tag. Incorrect value %s", tag)
	}

}

func TestReadTag(t *testing.T) {

}

func TestReadDataElement(t *testing.T) {

}
