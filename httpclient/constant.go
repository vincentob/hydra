package httpclient

const (
	// Application
	ContentTypeApplicationJson   = "application/json"                                                          // .json
	ContentTypeApplicationXML    = "application/xml"                                                           // .xml
	ContentTypeApplicationBinary = "application/octet-stream"                                                  // .bin
	ContentTypeApplicationWord   = "application/msword"                                                        // .doc
	ContentTypeApplicationWordX  = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"   // .docx
	ContentTypeApplicationPPT    = "application/vnd.ms-powerpoint"                                             // .ppt
	ContentTypeApplicationPPTX   = "application/vnd.openxmlformats-officedocument.presentationml.presentation" // .pptx
	ContentTypeApplicationXLS    = "application/vnd.ms-excel"                                                  // .xls
	ContentTypeApplicationXLSX   = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"         // .xlsx
	ContentTypeApplicationEPUB   = "application/epub+zip"                                                      // .epub
	ContentTypeApplicationRAR    = "application/x-rar-compressed"                                              // .rar
	ContentTypeApplicationSH     = "application/x-sh"                                                          // .sh
	ContentTypeApplicationZIP    = "application/zip"                                                           // .zip

	// Text
	ContentTypeTextPlain = "text/plain" // .txt
	ContentTypeTextXML   = "text/xml"   // .xml
	ContentTypeTextCSS   = "text/css"   // .css
	ContentTypeTextCSV   = "text/csv"   // .csv
	ContentTypeTextHTML  = "text/html"  // .htm .html

	// Image
	ContentTypeImageGIF  = "image/gif"                // .gif
	ContentTypeImageICON = "image/vnd.microsoft.icon" // .ico
	ContentTypeImageJPEG = "image/jpeg"               // .jpg .jpeg
	ContentTypeImagePNG  = "image/png"                // .png
	ContentTypeImageWEBP = "image/webp"               // .webp

	// Audio
	ContentTypeAudioMP3 = "audio/mpeg" // .mp3
)

const (
	HeaderContentType = "Content-Type"
)
