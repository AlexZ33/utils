package http

import (
	"encoding/json"
	"github.com/AlexZ33/utils/errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// ContentTypeJSON 表示JSON类型
	ContentTypeJSON = "application/json; charset=utf-8"
	// ContentTypeForm 表示Form类型
	ContentTypeForm = "application/x-www-form-urlencoded; charset=utf-8"
	// ContentTypeFormData 表示FormData类型
	ContentTypeFormData = "multipart/form-data; charset=utf-8"
	// ContentTypeText 表示Text类型
	ContentTypeText = "text/plain; charset=utf-8"
	// ContentTypeHTML 表示HTML类型
	ContentTypeHTML = "text/html; charset=utf-8"
	// ContentTypeXML 表示XML类型
	ContentTypeXML = "application/xml; charset=utf-8"

	// ContentTypeJSONP 表示JSONP类型
	ContentTypeJSONP = "application/javascript; charset=utf-8"
	// ContentTypeXHTML 表示XHTML类型
	ContentTypeXHTML = "application/xhtml+xml; charset=utf-8"

	// ContentTypePDF 表示PDF类型
	ContentTypePDF = "application/pdf; charset=utf-8"
	// ContentTypeZip 表示Zip类型
	ContentTypeZip = "application/zip; charset=utf-8"
	// ContentTypeGZIP 表示GZIP类型
	ContentTypeGZIP = "application/gzip; charset=utf-8"
	// ContentTypeBZIP2 表示BZIP2类型
	ContentTypeBZIP2 = "application/x-bzip2; charset=utf-8"
	// ContentTypeTAR 表示TAR类型
	ContentTypeTAR = "application/x-tar; charset=utf-8"
	// ContentTypeTARGZ 表示TARGZ类型
	ContentTypeTARGZ = "application/x-gzip; charset=utf-8"
	// ContentTypeTARBZ2 表示TARBZ2类型
	ContentTypeTARBZ2 = "application/x-bzip2; charset=utf-8"
	// ContentTypeTARZIP 表示TARZIP类型
	ContentTypeTARZIP = "application/x-compress; charset=utf-8"
	// ContentTypeTAR7Z 表示TAR7Z类型
	ContentTypeTAR7Z = "application/x-7z-compressed; charset=utf-8"
	// ContentTypeTARGZIP 表示TARGZIP类型
	ContentTypeTARGZIP = "application/x-gzip; charset=utf-8"

	// ContentTypeBMP 表示BMP类型
	ContentTypeBMP = "image/bmp; charset=utf-8"
	// ContentTypeICO 表示ICO类型
	ContentTypeICO = "image/vnd.microsoft.icon; charset=utf-8"
	// ContentTypeJPEG 表示JPEG类型
	ContentTypeJPEG = "image/jpeg; charset=utf-8"
	// ContentTypePNG 表示PNG类型
	ContentTypePNG = "image/png; charset=utf-8"
	// ContentTypeGIF 表示GIF类型
	ContentTypeGIF = "image/gif; charset=utf-8"
	// ContentTypeTIFF 表示TIFF类型
	ContentTypeTIFF = "image/tiff; charset=utf-8"
	// ContentTypeSVG 表示SVG类型
	ContentTypeSVG = "image/svg+xml; charset=utf-8"

	// ContentTypeMP3 表示MP3类型
	ContentTypeMP3 = "audio/mpeg; charset=utf-8"
	// ContentTypeWAV 表示WAV类型
	ContentTypeWAV = "audio/x-wav; charset=utf-8"
	// ContentTypeOGG 表示OGG类型
	ContentTypeOGG = "audio/ogg; charset=utf-8"
	// ContentTypeMP4 表示MP4类型
	ContentTypeMP4 = "video/mp4; charset=utf-8"
	// ContentTypeWebM 表示WebM类型
	ContentTypeWebM = "video/webm; charset=utf-8"
	// ContentTypeWMV 表示WMV类型
	ContentTypeWMV = "video/x-ms-wmv; charset=utf-8"
	// ContentTypeFLV 表示FLV类型
	ContentTypeFLV = "video/x-flv; charset=utf-8"
	// ContentTypeMKV 表示MKV类型
	ContentTypeMKV = "video/x-matroska; charset=utf-8"
	// ContentTypeAVI 表示AVI类型
	ContentTypeAVI = "video/x-msvideo; charset=utf-8"
	// ContentTypeMOV 表示MOV类型
	ContentTypeMOV = "video/quicktime; charset=utf-8"
	// ContentTypeWM 表示WM类型
	ContentTypeWM = "video/x-ms-wm; charset=utf-8"
	// ContentTypeASF 表示ASF类型
	ContentTypeASF = "video/x-ms-asf; charset=utf-8"
	// ContentType3GP 表示3GP类型
	ContentType3GP = "video/3gpp; charset=utf-8"
	// ContentType3G2 表示3G2类型
	ContentType3G2 = "video/3gpp2; charset=utf-8"
	// ContentTypeMP2 表示MP2类型
	ContentTypeMP2 = "video/mpeg; charset=utf-8"
	// ContentTypeMPEG 表示MPEG类型
	ContentTypeMPEG = "video/mpeg; charset=utf-8"
	// ContentTypeMPG 表示MPG类型
	ContentTypeMPG = "video/mpeg; charset=utf-8"

	// ContentTypeDOC 表示DOC类型
	ContentTypeDOC = "application/msword; charset=utf-8"
	// ContentTypeDOCX 表示DOCX类型
	ContentTypeDOCX = "application/vnd.openxmlformats-officedocument.wordprocessingml.document; charset=utf-8"
	// ContentTypeXLS 表示XLS类型
	ContentTypeXLS = "application/vnd.ms-excel; charset=utf-8"
	// ContentTypeXLSX 表示XLSX类型
	ContentTypeXLSX = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet; charset=utf-8"
	// ContentTypePPT 表示PPT类型
	ContentTypePPT = "application/vnd.ms-powerpoint; charset=utf-8"
	// ContentTypePPTX 表示PPTX类型
	ContentTypePPTX = "application/vnd.openxmlformats-officedocument.presentationml.presentation; charset=utf-8"
	// ContentTypeTXT 表示TXT类型
	ContentTypeTXT = "text/plain; charset=utf-8"
)

var Cli *http.Client = &http.Client{
	Timeout:   5 * time.Second, // 包括连接时间和响应时间
	Transport: &http.Transport{MaxConnsPerHost: 100},
}

// Request 请求接口
type Request struct {
	Protocol   string // 协议
	DomainName string // 域名
	Port       string // 端口
	API        string // API接口
	Method     string // 请求方法 -> GET, POST, PUT, DELETE

}

/**************************************ReqBody***************************************************/

// ReqBody 请求体
type ReqBody struct {
	inter       interface{}
	reader      io.Reader
	contentType string
}

// Set 设置请求体
func (body *ReqBody) Set(contentType string, bodyInter interface{}) error {
	var bodyReader io.Reader
	switch contentType {
	case ContentTypeJSON:
		bodyByte, err := json.Marshal(bodyInter)
		if err != nil {
			panic(err)
		}
		bodyReader = strings.NewReader(string(bodyByte))
	case ContentTypeFormData:
		if val, ok := bodyInter.(url.Values); ok {
			bodyReader = strings.NewReader(val.Encode())
		} else {
			return errors.New("数据类型[" + ContentTypeFormData + "]错误, 必须是url.Values类型")
		}
	default:
		return errors.New("不支持的数据类型[" + contentType + "]")
	}
	body.inter = bodyInter
	body.reader = bodyReader
	body.contentType = contentType
	return nil
}

// GetBodyReader 获取bodyReader
func (body *ReqBody) GetBodyReader() io.Reader {
	if body == nil {
		return nil
	}
	var bodyReader io.Reader
	switch body.contentType {
	case ContentTypeJSON:
		bodyByte, err := json.Marshal(body.inter)
		if err != nil {
			panic(err)
		}
		bodyReader = strings.NewReader(string(bodyByte))
	case ContentTypeFormData:
		if val, ok := body.inter.(url.Values); ok {
			bodyReader = strings.NewReader(val.Encode())
		} else {
			panic("数据类型[" + ContentTypeFormData + "]错误, 必须是url.Values类型")
		}
	default:
		panic("不支持的数据类型[" + body.contentType + "]")

	}
	return bodyReader
}

// GetBodyInter 获取bodyInter
func (body *ReqBody) GetBodyInter() interface{} {
	if body == nil {
		return nil
	}
	return body.inter
}

//AddParam 为body添加参数, 仅form-data支持使用
func (body *ReqBody) AddParam(key string, value interface{}) error {
	if body.contentType != ContentTypeFormData {
		return errors.New("body不是from-data, 不支持AddParam")
	}
	bodyInter := body.GetBodyInter()
	bodyParams := bodyInter.(url.Values)
	if bodyParams == nil {
		bodyParams = make(url.Values)
	}
	switch value.(type) {
	case string:
		v := value.(string)
		if len(v) > 0 {
			bodyParams.Add(key, v)
		}
	case []string:
		vs := value.([]string)
		for _, v := range vs {
			if len(v) > 0 {
				bodyParams.Add(key, v)
			}
		}
	default:
		panic("param不支持的参数类型")
	}
	body.inter = bodyParams
	return nil
}

// GetContentType 获取contentType
func (body *ReqBody) GetContentType() string {
	if body == nil {
		return ""
	}
	return body.contentType
}

/**************************************ReqInput***************************************************/

// ReqInput 请求输入
type ReqInput struct {
	Header http.Header // 请求头
	Params url.Values  // 请求参数
	Body   *ReqBody    // 请求体
}

func(rInput *ReqInput) AddReqInputParam (key string, value interface) {
	if rInput.Params == nil {
		rInput.Params = make(url.Values)
	}
	switch value.(type) {
	case string:
		v := value.(string)
		if len(v) > 0 {
			rInput.Params.Add(key, v)
		}
	case []string:
		vs := value.([]string)
		for _, v := range vs {
			if len(v) > 0 {
				rInput.Params.Add(key, v)
			}
		}
	default:
		panic("param不支持的参数类型")
	}
}



