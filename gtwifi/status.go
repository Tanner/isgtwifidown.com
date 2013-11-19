package gtwifi

import (
	"code.google.com/p/go.net/html"
	"errors"
	"log"
	"net/http"
)

const (
	GREEN int = iota
	YELLOW
	RED
	UKNOWN
)

const STATUS_URL = "http://status.oit.gatech.edu/index.php?action=service&service=lawn"

const STATUS_BLOCK_CLASS = "vip_status"

const (
	GREEN_STATUS_CLASS  = "bg_green"
	YELLOW_STATUS_CLASS = "bg_yellow"
	RED_STATUS_CLASS    = "bg_red"
)

type Status struct {
	Status int
	Reason string
}

func GetStatus() (Status, error) {
	resp, err := http.Get(STATUS_URL)
	if err != nil {
		log.Println(err)

		return Status{}, errors.New("Could not access OIT status page")
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)

	statusNode, err := FindStatusBlock(doc)
	if err != nil {
		log.Println(err)

		return Status{}, err
	}

	status, err := ExtractStatus(statusNode)
	if err != nil {
		log.Println(err)

		return Status{}, err
	}

	reason, err := ExtractReason(statusNode)
	if err != nil {
		log.Println(err)

		return Status{}, err
	}

	return Status{status, reason}, nil
}

func FindStatusBlock(node *html.Node) (*html.Node, error) {
	if node.Type == html.ElementNode {
		for _, attr := range node.Attr {
			if attr.Key == "class" && attr.Val == STATUS_BLOCK_CLASS {
				return node, nil
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		statusNode, err := FindStatusBlock(child)
		if err == nil {
			return statusNode, err
		}
	}

	return nil, errors.New("Could not find status block")
}

func ExtractStatus(node *html.Node) (int, error) {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.ElementNode && child.Data == "td" {
			for _, attr := range child.Attr {
				if attr.Key == "class" {
					class := attr.Val

					switch class {
					case GREEN_STATUS_CLASS:
						return GREEN, nil
					case YELLOW_STATUS_CLASS:
						return YELLOW, nil
					case RED_STATUS_CLASS:
						return RED, nil
					}

					return UKNOWN, errors.New("Status not recognized")
				}
			}
		}
	}

	return UKNOWN, errors.New("Status not found")
}

func ExtractReason(node *html.Node) (string, error) {
	if node.Type == html.ElementNode && node.Data == "p" {
		return node.FirstChild.Data, nil
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		reason, err := ExtractReason(child)
		if err == nil {
			return reason, err
		}
	}

	return "", errors.New("Reason not found")
}
