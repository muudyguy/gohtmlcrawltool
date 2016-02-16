package htmlcrawler

//dst is pointer to pointer to avoid copying
func GetFirstTag(tagName string, tag *HtmlTag, dst **HtmlTag) bool {
	if tag.XMLName.Local == tagName {
		*dst = tag
		return true
	}

	for _, div := range tag.ChildTags {
		return GetFirstTag(tagName, &div, dst)
	}

	return false

}
