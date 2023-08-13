package emaildecoder

import (
	"encoding/base64"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"mime/quotedprintable"
	"net/mail"
	"path/filepath"
	"strings"
)

type AttachmentCallback func(attachment Attachment)

type Decoder struct {
	attachmentIdx      int
	attachmentCallback AttachmentCallback
	reader             io.Reader
	plainText          []byte
	html               []byte
}

type Attachment struct {
	Filename    string
	ContentType string
	Reader      io.Reader
}

type EmailContent struct {
	HTML      []byte
	PlainText []byte
	Headers   mail.Header
}

func NewDecoder(reader io.Reader, attachmentCallback AttachmentCallback) *Decoder {
	return &Decoder{
		attachmentCallback: attachmentCallback,
		reader:             reader,
	}
}

func (d *Decoder) Decode() (*EmailContent, error) {
	msg, err := mail.ReadMessage(d.reader)
	if err != nil {
		panic(err)
	}
	mediaType, mediaParams, err := mime.ParseMediaType(msg.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(mediaType, "multipart/") {
		d.findParts(msg.Body, mediaParams["boundary"])
	} else if strings.HasPrefix(msg.Header.Get("Content-Type"), "text/") {
		d.decodeText(msg.Body, mediaType, msg.Header.Get("Content-Transfer-Encoding"), mediaParams)
	}
	return &EmailContent{
		HTML:      d.html,
		PlainText: d.plainText,
		Headers:   msg.Header,
	}, nil
}

func (d *Decoder) findParts(mime_data io.Reader, boundary string) {
	// Instantiate a new io.Reader dedicated to MIME multipart parsing
	// using multipart.NewReader()
	reader := multipart.NewReader(mime_data, boundary)
	if reader == nil {
		return
	}
	// Go through each of the MIME part of the message Body with NextPart(),
	for {
		newPart, err := reader.NextPart()
		if err != nil {
			break
		}
		contentType, params, err := mime.ParseMediaType(newPart.Header.Get("Content-Type"))
		if err == nil && strings.HasPrefix(contentType, "multipart/") {
			d.findParts(newPart, params["boundary"])
		} else if strings.HasPrefix(contentType, "text/") {
			d.decodeText(newPart, contentType, newPart.Header.Get("Content-Transfer-Encoding"), params)
		} else if mediaType, mediaParams, err := mime.ParseMediaType(newPart.Header.Get("Content-Disposition")); err == nil && mediaType == "attachment" && d.attachmentCallback != nil {
			attachment := Attachment{
				Reader:      newPart,
				ContentType: contentType,
			}
			if filename, filenameFound := mediaParams["filename"]; !filenameFound || len(filename) == 0 || filepath.Base(filename) == string(filepath.Separator) {
				d.attachmentIdx++
				attachment.Filename = fmt.Sprintf("attachement-%d.file", d.attachmentIdx)
			} else {
				attachment.Filename = filepath.Base(filename)
			}
			attachment.Reader = d.getDecodeReader(attachment.Reader, newPart.Header.Get("Content-Transfer-Encoding"), params)
			d.attachmentCallback(attachment)
		}
	}
}

func (d *Decoder) getDecodeReader(reader io.Reader, transferEncoding string, contentParams map[string]string) io.Reader {
	if strings.Contains(transferEncoding, "base64") {
		reader = base64.NewDecoder(base64.StdEncoding, reader)
	}
	if strings.Contains(transferEncoding, "quoted-printable") {
		reader = quotedprintable.NewReader(reader)
	}
	if charset, charsetFound := contentParams["charset"]; charsetFound && !strings.Contains(strings.ToLower(charset), "utf-8") {
		if charsetEncoding, charsetEncodingFound := charMapEncoders[strings.ToLower(charset)]; charsetEncodingFound {
			reader = charsetEncoding.NewDecoder().Reader(reader)
		}
	}
	return reader
}

func (d *Decoder) decodeText(reader io.Reader, contentType, transferEncoding string, contentParams map[string]string) {
	reader = d.getDecodeReader(reader, transferEncoding, contentParams)
	if strings.HasPrefix(contentType, "text/plain") {
		d.plainText, _ = io.ReadAll(reader)
	}
	if strings.HasPrefix(contentType, "text/html") {
		d.html, _ = io.ReadAll(reader)
	}
}
