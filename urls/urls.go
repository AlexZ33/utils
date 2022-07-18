package urls

import "net/url"

type UrlBuilder struct {
	u     *url.URL
	query url.Values
}

func ParseUrl(rawUrl string) *UrlBuilder {
	urlBuilder := &UrlBuilder{}
	urlBuilder.u, _ = url.Parse(rawUrl)
	urlBuilder.query = urlBuilder.u.Query()
	return urlBuilder
}

func (builder *UrlBuilder) SetQuery(key string, value string) *UrlBuilder {
	builder.query.Set(key, value)
	return builder
}

func (builder *UrlBuilder) AddQuery(key string, value string) *UrlBuilder {
	builder.query.Add(key, value)
	return builder
}

func (builder *UrlBuilder) AddQueries(queries map[string]string) *UrlBuilder {
	for key, value := range queries {
		builder.AddQuery(key, value)
	}
	return builder
}
func (builder *UrlBuilder) GetQuery() url.Values {
	return builder.query
}

func (builder *UrlBuilder) GetUrl() *url.URL {
	return builder.u
}

func (builder *UrlBuilder) Build() *url.URL {
	builder.u.RawQuery = builder.query.Encode()
	return builder.u
}

func (builder *UrlBuilder) BuildString() string {
	return builder.Build().String()
}
