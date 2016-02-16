package htmlcrawler

type HtmlTagNotFoundError struct {

}

func (htmlNotFoundError HtmlTagNotFoundError) Error() string {
	return "No html Tag Found"
}

type BodyTagNotFoundError struct {

}

func (bodyTagNotFoundError BodyTagNotFoundError) Error() string {
	return "No Body Tag Found"
}

func GrabHtmlTag(root *HtmlTag) (error, *HtmlTag) {
	htmlTag := new(HtmlTag)
	bl := GetFirstTag("html", root, &htmlTag)
	if !bl {
		return HtmlTagNotFoundError{}, nil

	}
	return nil, htmlTag
}


func GrabBodyTag(root *HtmlTag) (error, *HtmlTag) {
	bodyTag := new(HtmlTag)
	bl := GetFirstTag("body", root, &bodyTag)
	if !bl {
		return BodyTagNotFoundError{}, nil

	}
	return nil, bodyTag
}

